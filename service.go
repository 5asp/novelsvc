package novelsvc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type BookInfo struct {
	//Id_	string		`json:"id" bson:"_id" comment:"小说ID"`
	Name        string `json:"name" bson:"name" comment:"小说名称"`
	Author      string `json:"author" bson:"author" comment:"小说作者"`
	Cover       string `json:"cover" bson:"cover" comment:"小说封面"`
	Category    string `json:"category" bson:"category" comment:"小说分类"`
	Description string `json:"description" bson:"description" comment:"小说描述"`
	NewChapter  string `json:"new_chapter" bson:"new_chapter" comment:"搜索结果最新章节"`
	URL         string `json:"url" bson:"url" comment:"搜索结果链接"`
	Source      string `json:"source" bson:"source" comment:"搜索结果来源"`
}
type Service interface {
	GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error)
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

func (s *inmemService) GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[keyword]; ok {
		return nil, ErrAlreadyExists // POST = create, don't overwrite
	}

	bookSources := sourceService.GetAllSource()

	for i := range bookSources {
		source := &bookSources[i]
		f := fetcher.NewFetcher()
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
	// return

	return list, ErrNotFound
}
