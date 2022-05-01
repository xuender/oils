package base

import (
	"golang.org/x/exp/slices"
)

// Has 包含.
func Has[E comparable](elems []E, elem E) bool {
	return slices.Index(elems, elem) > -1
}

// Count 统计元素数量.
func Count[E comparable](elems []E, elem E) int {
	count := 0

	for _, value := range elems {
		if value == elem {
			count++
		}
	}

	return count
}

// Index 索引.
func Index[T comparable](slices, sub []T) int {
	if len(slices) < len(sub) {
		return -1
	}

	for index := 0; index <= len(slices)-len(sub); index++ {
		has := true

		for elemIndex, elem := range sub {
			if slices[index+elemIndex] != elem {
				has = false

				break
			}
		}

		if has {
			return index
		}
	}

	return -1
}
