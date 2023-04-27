package novelsvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	cors := cors.New(cors.Options{
        AllowedOrigins: []string{"https://moutfire.com"},
        AllowedMethods: []string{
            http.MethodGet,
        },
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
    })
	
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/search").Handler(httptransport.NewServer(
		e.SearchEndpoint,
		decodeSearchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/info").Handler(httptransport.NewServer(
		e.InfoEndpoint,
		decodeInfoRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/chapters").Handler(httptransport.NewServer(
		e.ChaptersEndpoint,
		decodeChaptersRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/read").Handler(httptransport.NewServer(
		e.ReadEndpoint,
		decodeReadRequest,
		encodeResponse,
		options...,
	))
	return cors.Handler(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeSearchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	keyword := r.URL.Query().Get("k")

	if keyword == "" {
		return nil, ErrBadRouting
	}

	return searchRequest{Keyword: keyword}, nil
}

func decodeInfoRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	detailURL := r.URL.Query().Get("detail_url")
	source := r.URL.Query().Get("source")
	if detailURL == "" &&  source == ""{
		return nil, ErrBadRouting
	}

	return infoRequest{DetailURL: detailURL,Source: source}, nil
}

func decodeReadRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	detailURL := r.URL.Query().Get("detail_url")
	chapterURL := r.URL.Query().Get("chapter_url")
	source := r.URL.Query().Get("source")
	if detailURL == "" &&  source == "" && chapterURL == ""{
		return nil, ErrBadRouting
	}

	return readRequest{DetailURL: detailURL,Source: source,ChapterURL: chapterURL}, nil
}

func decodeChaptersRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	detailURL := r.URL.Query().Get("detail_url")
	source := r.URL.Query().Get("source")
	if detailURL == "" &&  source == ""{
		return nil, ErrBadRouting
	}

	return infoRequest{DetailURL: detailURL,Source: source}, nil
}

type errorer interface {
	error() error
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
