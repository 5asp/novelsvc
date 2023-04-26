package novelsvc

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

// func (mw loggingMiddleware) GetListByKeyword(ctx context.Context, keyword string) (err error) {
// 	defer func(begin time.Time) {
// 		mw.logger.Log("method", "GetListByKeyword", "keyword=", keyword, "took", time.Since(begin), "err", err)
// 	}(time.Now())
// 	return mw.next.GetListByKeyword(ctx, keyword)
// }

func (mw loggingMiddleware) GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetListByKeyword", "keyword=", keyword, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetListByKeyword(ctx, keyword)
}
