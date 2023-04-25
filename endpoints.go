package novelsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SearchEndpoint endpoint.Endpoint
}

type searchRequest struct {
	Keyword string
}
type searchResponse struct {
	Err error `json:"err,omitempty"`
}

func (r searchResponse) error() error { return r.Err }

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		SearchEndpoint: MakeSearchEndpoint(s),
	}
}

func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchRequest)
		e := s.GetListByKeyword(ctx, req.Keyword)
		return searchResponse{Err: e}, nil
	}
}
