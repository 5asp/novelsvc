package novelsvc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type Service interface {
	GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error)
	GetInfo(ctx context.Context,url string, source string) BookInfo
	GetChapterList(ctx context.Context,url string, source string) []BookChapter
	GetContent(ctx context.Context,detailURL string, chapterURL string, source string) (Content BookContent)
}

type inmemService struct {
	mtx sync.RWMutex
	m   map[string]BookInfo
}

func NewInmemService() Service {
	return &inmemService{
		m: map[string]BookInfo{},
	}
}
var fetchSourceObject = NewFetchSource()

func (s *inmemService) parseItemSearch(source *BookSource, doc *colly.XMLElement) (item *BookInfo) {
	var ele = NewXMLElement(doc)
	item = &BookInfo{
		Name:        ele.ChildText(source.SearchItemNameRule),
		Author:      ele.ChildText(source.SearchItemAuthorRule),
		Cover:       ele.ChildAttr(source.SearchItemCoverRule, "src"),
		NewChapter:  ele.ChildText(source.SearchItemNewChapterRule),
		URL:         ele.ChildUrl(source.SearchItemURLRule, "href"),
		Source:      source.SourceKey,
	}
	return
}

func (s *inmemService) GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[keyword]; ok {
		return nil, ErrAlreadyExists // POST = create, don't overwrite
	}

	bookSources := fetchSourceObject.GetAllSource()
	for i := range bookSources {
		source := &bookSources[i]
		f := NewColly()
		q, _ := queue.New(
			2, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
		)

		f.OnXML(source.SearchItemRule, func(e *colly.XMLElement) {
			list = append(list, s.parseItemSearch(source, e))
		})
		q.AddURL(fmt.Sprintf(source.SearchURL, keyword))
		q.Run(f)
	}
	return
}

func (s *inmemService) parseItemInfo(source *BookSource, doc *colly.XMLElement) (info BookInfo) {
	var ele = NewXMLElement(doc)
	info = BookInfo{
		Name:	ele.ChildText(source.DetailBookNameRule),
		Author:	ele.ChildText(source.DetailBookAuthorRule),
		Cover:	ele.ChildAttr(source.DetailBookCoverRule, "src"),
		Category:	ele.ChildText(source.DetailBookCategoryRule),
		Description:	ele.ChildHtml(source.DetailBookDescriptionRule),
	}
	return
}

func (s *inmemService) GetInfo(ctx context.Context,url string, key string) (info BookInfo) {
	source, ok := fetchSourceObject.GetSourceByKey(key)
	if !ok {
		return
	}

	f := NewColly()
	f.OnXML(source.DetailBookItemRule, func(e *colly.XMLElement) {
		info = s.parseItemInfo(&source, e)
	})

	f.Visit(url)
	return
}

func (s *inmemService) parseChapterList(source *BookSource, doc *colly.XMLElement, url string) (chapter BookChapter) {
	var ele  = NewXMLElement(doc)
	chapter = BookChapter{
		Title: 	ele.ChildText(source.DetailChapterTitleRule),
		DetailURL:	url,
		ChapterURL:   ele.ChildUrl(source.DetailChapterURLRule, "href"),
		Source:	source.SourceKey,
	}

	return
}

func (s *inmemService) GetChapterList(ctx context.Context,url string, key string) (chapterList []BookChapter) {
	source, ok := fetchSourceObject.GetSourceByKey(key)
	if !ok {
		return
	}

	var chapterListURL string
	f := NewColly()

	if source.DetailChapterListURLRule != "" {
		f.OnXML(source.DetailChapterURLRule, func(e *colly.XMLElement) {
			var ele  = NewXMLElement(e)
			chapterListURL = ele.ChildUrl(source.DetailChapterListURLRule, "href")
			fc := NewColly()
			fc.OnXML(source.DetailChapterRule, func(e *colly.XMLElement) {
				chapterList = append(chapterList, s.parseChapterList(&source, e, chapterListURL))
				return
			})
			fc.Visit(chapterListURL)
			return
		})
	}else {
		f.OnXML(source.DetailChapterRule, func(e *colly.XMLElement) {
			chapterList = append(chapterList, s.parseChapterList(&source, e, url))
			return
		})
	}
	f.Visit(url)
	return
}

func (s *inmemService) parseContent(source *BookSource, doc *colly.XMLElement, url string) (content BookContent) {
	var ele  = NewXMLElement(doc)
	content = BookContent{
		Title:	ele.ChildText(source.ContentTitleRule),
		Text:	ele.ChildHtml(source.ContentTextRule),
		DetailURL:	url,
		PreviousURL:	ele.ChildUrl(source.ContentPreviousURLRule, "href"),
		NextURL:	ele.ChildUrl(source.ContentNextURLRule, "href"),
		Source:	source.SourceKey,
	}
	return content
}

func (s *inmemService) GetContent(ctx context.Context,detailURL string, chapterURL string, key string) (content BookContent) {

	source, ok := fetchSourceObject.GetSourceByKey(key)

	if !ok {
		return
	}

	f := NewColly()
	f.OnXML("//body", func(e *colly.XMLElement) {
		content = s.parseContent(&source, e, detailURL)
	})
	f.Visit(chapterURL)
	return
}