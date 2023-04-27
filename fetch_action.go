package novelsvc

import "sync"

type Query func(source BookSource) bool

type FetchAction interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie BookSource, found bool)
	SelectMany(query Query, limit int) (results []BookSource)
}

func NewFetchAction(source map[int64]BookSource) FetchAction {
	return &fetchAction{
		source: source,
	}
}

type fetchAction struct {
	source map[int64]BookSource
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *fetchAction) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

func (r *fetchAction) Select(query Query) (movie BookSource, found bool) {
	found = r.Exec(query, func(m BookSource) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)
	//设置一个空的datamodels.Movie，如果根本找不到的话。
	if !found {
		movie = BookSource{}
	}
	return
}

func (r *fetchAction) SelectMany(query Query, limit int) (results []BookSource) {
	r.Exec(query, func(m BookSource) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}