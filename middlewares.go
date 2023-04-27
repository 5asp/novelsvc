package novelsvc

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
			w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
			w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
			if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusNoContent)
					return
			}
			next.ServeHTTP(w, r)
	})
}

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

func (mw loggingMiddleware) GetListByKeyword(ctx context.Context, keyword string) (list []*BookInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetListByKeyword", "keyword=", keyword, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetListByKeyword(ctx, keyword)
}

func (mw loggingMiddleware) GetInfo(ctx context.Context,url string, key string) ( info BookInfo) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetInfo", "url:", url,"source:", key,  "took", time.Since(begin))
	}(time.Now())
	return mw.next.GetInfo(ctx, url,key)
}

func (mw loggingMiddleware) GetChapterList(ctx context.Context,url string, key string) (chapterList []BookChapter) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetChapterList", "url:", url,"source:", key,  "took", time.Since(begin))

	}(time.Now())
	return mw.next.GetChapterList(ctx, url,key)
}
func (mw loggingMiddleware) GetContent(ctx context.Context,detailURL string,chapterURL string, key string) ( content  BookContent) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetContent", "detail_url:", detailURL,"chapter_url",chapterURL,"source:", key,  "took", time.Since(begin))
	}(time.Now())
	return mw.next.GetContent(ctx, detailURL,chapterURL, key)
}
