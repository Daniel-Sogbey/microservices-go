package data

type Filter struct {
	Page     int64
	PageSize int64
}

func (f Filter) Skip() int64 {
	return (f.Page - 1) * f.PageSize
}

func (f Filter) Limit() int64 {
	return f.PageSize
}
