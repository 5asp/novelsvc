package novelsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SearchEndpoint   endpoint.Endpoint
	ReadEndpoint     endpoint.Endpoint
	ChaptersEndpoint endpoint.Endpoint
	InfoEndpoint     endpoint.Endpoint
}

type searchRequest struct {
	Keyword string
}
type infoRequest struct {
	Source string
	DetailURL string
}
type readRequest struct {
	Source string
	DetailURL string
	ChapterURL string
}
type searchResponse struct {
	Err  error `json:"err,omitempty"`
	List []*BookInfo
}

func (r searchResponse) error() error { return r.Err }

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		SearchEndpoint:   MakeSearchEndpoint(s),
		ReadEndpoint:     MakeReadEndpoint(s),
		ChaptersEndpoint: MakeChaptersEndpoint(s),
		InfoEndpoint:     MakeInfoEndpoint(s),
	}
}

func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchRequest)
		list, e := s.GetListByKeyword(ctx, req.Keyword)
		return searchResponse{Err: e, List: list}, nil
	}
}
func MakeReadEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(readRequest)
		content := s.GetContent(ctx, req.DetailURL,req.ChapterURL,req.Source)
		return content,nil
	}
}

func MakeChaptersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(infoRequest)
		list := s.GetChapterList(ctx, req.DetailURL,req.Source)
		return list, nil
	}
}

func MakeInfoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(infoRequest)
		item := s.GetInfo(ctx, req.DetailURL,req.Source)
		return item,nil
	}
}
