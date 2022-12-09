package base

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

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

// Sub 切片截取.
func Sub[Elem any](slice []Elem, startAndEnd ...int) []Elem {
	length := len(slice)

	if length == 0 {
		return []Elem{}
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
		return []Elem{}
	}

	return slice[start:end]
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

func Del[E comparable](slice []E, elems ...E) []E { // nolint
	if len(slice) == 0 {
		return []E{}
	}

	if len(elems) == 0 {
		return slice
	}

	ret := append([]E{}, slice...)

	for _, elem := range elems {
		if i := slices.Index(ret, elem); i >= 0 {
			ret = append((ret)[0:i], (ret)[i+1:]...)
		}
	}

	return ret
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

func Counts[E comparable](slice, sub []E) int {
	count := 0

	if len(slice) < len(sub) {
		return count
	}

	if len(sub) < 1 {
		return len(slice)
	}

	for index := 0; index <= len(slice)-len(sub); index++ {
		has := true

		for elemIndex, elem := range sub {
			if slice[index+elemIndex] != elem {
				has = false

				break
			}
		}

		if has {
			count++
		}
	}

	return count
}

// All 全部包含.
func All[E comparable](slice, elems []E) bool {
	for _, elem := range elems {
		if slices.Index(slice, elem) < 0 {
			return false
		}
	}

	return true
}

// Any 任意包含.
func Any[E comparable](slice, elems []E) bool {
	for _, elem := range elems {
		if slices.Index(slice, elem) > -1 {
			return true
		}
	}

	return false
}

// Replace 替换.
func Replace[E comparable](slice, sub, newSub []E, num int) []E { // nolint
	if num == 0 {
		return slice
	}

	if count := Counts(slice, sub); count == 0 {
		return slice
	} else if num < 0 || count < num {
		num = count
	}

	ret := append([]E{}, slice...)

	start := 0
	for i := 0; i < num; i++ {
		index := Index(ret[start:], sub)
		ret = append(ret[:start+index], append(newSub, ret[start+index+len(sub):]...)...)
	}

	return ret
}

// ReplaceAll 全部替换.
func ReplaceAll[E comparable](slice, sub, newSub []E) []E {
	return Replace(slice, sub, newSub, -1)
}
