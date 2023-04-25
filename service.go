package novelsvc

import (
	"context"
	"errors"
	"sync"
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
	GetListByKeyword(ctx context.Context, keyword string) error
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

func (s *inmemService) GetListByKeyword(ctx context.Context, keyword string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	// if _, ok := s.m[p.ID]; ok {
	// 	return ErrAlreadyExists // POST = create, don't overwrite
	// }
	// s.m[p.ID] = p
	return nil
}
