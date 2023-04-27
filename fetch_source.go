package novelsvc

type FetchSource interface {
	GetSourceByKey(key string) (BookSource, bool)
	GetAllSource() []BookSource
}

func NewFetchSource() FetchSource {
	return &fetchSource{
		NewFetchAction(BookSources),
	}
}

type fetchSource struct {
	rep FetchAction
}

func (s *fetchSource) GetSourceByKey(key string) (BookSource, bool) {
	return s.rep.Select(func(s BookSource) bool {
		return s.SourceKey == key
	})
}

func (s *fetchSource) GetAllSource() []BookSource {
	return s.rep.SelectMany(func(_ BookSource) bool {
		return true
	}, -1)
}