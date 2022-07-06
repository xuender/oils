package base

import (
	"fmt"
	"strings"

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
func Sub[Elem any](slice []Elem, startAndEnd ...int) (sub []Elem) {
	length := len(slice)

	if length == 0 {
		return
	}

	start, end := 0, length

	if len(startAndEnd) > 0 {
		start = startAndEnd[0]

		if start < 0 {
			start += length

			if start < 0 {
				start = 0
			}
		}
	}

	if len(startAndEnd) > 1 && startAndEnd[1] < length {
		end = startAndEnd[1]
	}

	if end < 0 {
		end += length
	}

	if start >= length || start >= end {
		return
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

// Unique 去重.
func Unique[E comparable](slice []E) []E {
	unique := []E{}

	for _, elem := range slice {
		if slices.Index(unique, elem) > -1 {
			continue
		}

		unique = append(unique, elem)
	}

	return unique
}

// Join 集合连接.
func Join[E any](slice []E, sep string) string {
	switch len(slice) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", slice[0])
	}

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%v", slice[0]))

	for _, elem := range slice[1:] {
		builder.WriteString(sep)
		builder.WriteString(fmt.Sprintf("%v", elem))
	}

	return builder.String()
}

func Del[E comparable](slice []E, elems ...E) []E {
	if len(slice) == 0 {
		return []E{}
	}

	if len(elems) == 0 {
		return slice
	}

	for _, elem := range elems {
		if i := slices.Index(slice, elem); i >= 0 {
			slice = append((slice)[0:i], (slice)[i+1:]...)
		}
	}

	return slice
}

func DelAll[E comparable](slice []E, elems ...E) []E {
	if len(slice) == 0 {
		return []E{}
	}

	if len(elems) == 0 {
		return slice
	}

	for _, elem := range elems {
		for i := slices.Index(slice, elem); i >= 0; i = slices.Index(slice, elem) {
			slice = append(slice[0:i], slice[i+1:]...)
		}
	}

	return slice
}
