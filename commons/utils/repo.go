package utils

import (
	"github.com/uptrace/bun"
	"reflect"
)

func ReturnFirstIfNotNil[T any](first T, second T) T {
	if reflect.ValueOf(first).IsZero() {
		return second
	}
	return first
}

func BuildPagination(lastID int64, limit int, orderType int, query *bun.SelectQuery) {
	if orderType == 0 {
		query.Where("id > ?", lastID)
	} else {
		query.Where("id < ?", lastID)
	}
	query.Limit(limit)
}
