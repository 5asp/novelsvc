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
		req := request.(searchRequest)
		list, e := s.GetListByKeyword(ctx, req.Keyword)
		return searchResponse{Err: e, List: list}, nil
	}
}

func MakeChaptersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchRequest)
		list, e := s.GetListByKeyword(ctx, req.Keyword)
		return searchResponse{Err: e, List: list}, nil
	}
}

func MakeInfoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchRequest)
		list, e := s.GetListByKeyword(ctx, req.Keyword)
		return searchResponse{Err: e, List: list}, nil
	}
}
