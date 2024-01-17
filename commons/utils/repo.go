package utils

import "reflect"

func ReturnFirstIfNotNil[T any](first T, second T) T {
	if reflect.ValueOf(first).IsZero() {
		return second
	}
	return first
}
