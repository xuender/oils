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

func Filter[T any](elems []T, condition func(T) bool) []T {
	ret := []T{}

	for _, elem := range elems {
		if condition(elem) {
			ret = append(ret, elem)
		}
	}

	return ret
}

// Append 追加，max最大尺寸,max小于0不限制尺寸.
func Append[S ~[]E, E any](max int, slice S, elems ...E) S {
	size := len(slice) + len(elems)

	if max < 0 || size <= max {
		return append(slice, elems...)
	}

	if max == len(elems) {
		return elems
	}

	if max < len(elems) {
		return elems[len(elems)-max:]
	}

	return append(slice[size-max:], elems...)
}

// SliceMap 切片转换.
func SliceMap[S, T any](elems []S, change func(S) T) []T {
	ret := make([]T, len(elems))

	for index, elem := range elems {
		ret[index] = change(elem)
	}

	return ret
}

// Sub 切片截取.
func Sub[Elem any](slice []Elem, nums ...int) []Elem {
	length := len(slice)

	if length == 0 {
		return []Elem{}
	}

	start, end := 0, length

	if len(nums) > 0 {
		start = nums[0]
	}

	if start < 0 {
		if -start > length {
			start = 0
		} else {
			start += length
		}
	}

	if len(nums) > 1 && nums[1] < length {
		end = nums[1]
	}

	if end < 0 {
		end += length
	}

	if start >= length || start >= end {
		return []Elem{}
	}

	return slice[start:end]
}

// Chunk 切片分块.
func Chunk[Elem any](slice []Elem, size int) [][]Elem {
	if len(slice) == 0 || size < 1 {
		return [][]Elem{}
	}

	length := len(slice)
	retLen := length / size

	if length%size > 0 {
		retLen++
	}

	ret := make([][]Elem, retLen)

	for index := 0; index < retLen; index++ {
		end := index*size + size

		if end > length {
			end = length
		}

		ret[index] = slice[index*size : end]
	}

	return ret
}
