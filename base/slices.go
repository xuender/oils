package base

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/exp/slices"
)

// Has 包含.
//
// Deprecated: 使用 slices.Contains.
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

// Reverse 反转切片.
func Reverse[E any](slice []E) {
	if len(slice) <= 1 {
		return
	}

	for start, end := 0, len(slice)-1; start < end; start, end = start+1, end-1 {
		slice[start], slice[end] = slice[end], slice[start]
	}
}

func Counts[E comparable](slice, sub []E) (count int) {
	if len(slice) < len(sub) {
		return
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

	return
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

// Sample 获得一个随机元素.
func Sample[E any](slice []E) E {
	if len(slice) == 0 {
		panic(ErrEmpty)
	}
	// nolint
	return slice[rand.Intn(len(slice))]
}

// Shuffle 打乱切片.
func Shuffle[E any](slice []E) []E { // nolint
	length := len(slice)

	if length <= 1 {
		return slice
	}

	ret := append([]E{}, slice...)

	length--

	for i := 0; i < length; i++ {
		// nolint
		num := rand.Intn(length - i)
		ret[i], ret[length-num] = ret[length-num], ret[i]
	}

	return ret
}
