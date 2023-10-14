package utils

import (
	"reflect"
)

func Includes[T any](slice []T, value T) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}
