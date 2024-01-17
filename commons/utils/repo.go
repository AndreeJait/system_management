package utils

import "reflect"

func ReturnFirstIfNotNil[T interface{}](first T, second T) T {
	if reflect.ValueOf(first).IsNil() {
		return second
	}
	return first
}
