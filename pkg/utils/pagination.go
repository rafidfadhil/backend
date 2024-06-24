package utils

import "go.mongodb.org/mongo-driver/mongo/options"

type MongoPaginate struct {
	limit int64
	page  int64
	total int64
}

type PaginationData struct {
	CurrentPage int64 `json:"current_page"`
	MaxPage     int64 `json:"max_page"`
	Total       int64 `json:"total"`
	Limit       int64 `json:"limit"`
}

func NewMongoPaginate(limit int, page int, total int) *MongoPaginate {
	return &MongoPaginate{
		limit: int64(limit),
		page:  int64(page),
		total: int64(total),
	}
}

func (mp *MongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}

func (mp *MongoPaginate) GetPaginationData() *PaginationData {
	maxPage := int64(0)
	if mp.limit != 0 {
		maxPage = mp.total / mp.limit
		if mp.total%mp.limit > 0 {
			maxPage++
		}
	}

	return &PaginationData{
		CurrentPage: mp.page,
		MaxPage:     maxPage,
		Total:       mp.total,
		Limit:       mp.limit,
	}
}
