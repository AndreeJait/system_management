// Package pagination provides support for pagination requests and responses.
package pagination

import (
	"net/http"
	"strconv"
)

var (
	lastID   = "last_id"
	typeSort = "type_sort"
	perPage  = "per_page"
)

type Page[T any] struct {
	PaginationInfo Pagination `json:"pagination_info"`
	Items          []T        `json:"items"`
}

type Pagination struct {
	HasNext   bool  `json:"has_next"`
	PerPage   int   `json:"per_page"`
	LastId    int64 `json:"last_id"`
	OrderType int   `json:"order_type"`
}

func (p *Page[T]) SetItem(items []T, hasNext bool, lastID int64) {
	p.Items = items
	p.PaginationInfo.HasNext = hasNext
	p.PaginationInfo.LastId = lastID
}

func (p *Page[T]) GetLastID() int64 {
	return p.PaginationInfo.LastId
}

func (p *Page[T]) GetOrder() int {
	return p.PaginationInfo.OrderType
}

func NewFromRequest[T any](req *http.Request) Page[any] {
	lastID := parseInt64(req.URL.Query().Get(lastID), 0)
	perPage := parseInt64(req.URL.Query().Get(perPage), 10)
	typeSort := int(parseInt64(req.URL.Query().Get(typeSort), 0))
	return Page[any]{
		PaginationInfo: Pagination{
			LastId:    lastID,
			PerPage:   int(perPage),
			OrderType: typeSort,
		},
	}
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt64(value string, defaultValue int64) int64 {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.ParseInt(value, 10, 64); err == nil {
		return result
	}
	return defaultValue
}
