package base

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Slice[T constraints.Ordered] []T

func (s Slice[T]) Len() int { return len(s) }
func (s Slice[T]) Less(i, j int) bool {
	// nolint
	return s[i] < s[j]
}
func (s Slice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func NewSlice[T constraints.Ordered](elems ...T) Slice[T] {
	newSlice := Slice[T]{}
	newSlice.Add(elems...)

	return newSlice
}

func (s *Slice[T]) Add(elems ...T) Slice[T] {
	*s = append(*s, elems...)

	return *s
}

// All 全部包含.
func (s Slice[T]) All(elems ...T) bool {
	for _, e := range elems {
		if !s.Has(e) {
			return false
		}
	}

	return true
}

// Any 任意包含.
func (s Slice[T]) Any(elems ...T) bool {
	for _, e := range elems {
		if s.Has(e) {
			return true
		}
	}

	return false
}

// Del 删除.
func (s *Slice[T]) Del(elems ...T) Slice[T] {
	for _, elem := range elems {
		if i := s.Index(elem); i >= 0 {
			*s = append((*s)[0:i], (*s)[i+1:]...)
		}
	}

	return *s
}

// DelAll 删除全部.
func (s *Slice[T]) DelAll(elems ...T) Slice[T] {
	for _, elem := range elems {
		for i := s.Index(elem); i >= 0; i = s.Index(elem) {
			*s = append((*s)[0:i], (*s)[i+1:]...)
		}
	}

	return *s
}

// Indexs 包含.
func (s Slice[T]) Indexs(elems []T) int {
	if len(s) < len(elems) {
		return -1
	}

	for index := 0; index <= len(s)-len(elems); index++ {
		has := true

		for elemIndex, elem := range elems {
			// nolint
			if s[index+elemIndex] != elem {
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

// Equal 比较.
func (s Slice[T]) Equal(target Slice[T]) bool {
	if len(s) != len(target) {
		return false
	}

	for i, elem := range target {
		// nolint
		if s[i] != elem {
			return false
		}
	}

	return true
}

// Count 包含元素数量.
func (s Slice[T]) Count(elem T) (count int) {
	for _, value := range s {
		// nolint
		if value == elem {
			count++
		}
	}

	return
}

// Count 包含切片数量.
func (s Slice[T]) Counts(elems []T) (count int) {
	if len(s) < len(elems) {
		return
	}

	for index := 0; index <= len(s)-len(elems); index++ {
		has := true

		for elemIndex, elem := range elems {
			// nolint
			if s[index+elemIndex] != elem {
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

// Replace 替换.
func (s *Slice[T]) Replace(oldSlice, newSlice []T, num int) Slice[T] {
	if num == 0 {
		return *s
	}

	if count := s.Counts(oldSlice); count == 0 {
		return *s
	} else if num < 0 || count < num {
		num = count
	}

	start := 0
	for i := 0; i < num; i++ {
		index := (*s)[start:].Indexs(oldSlice)
		*s = append((*s)[0:start+index], append(newSlice, (*s)[start+index+len(oldSlice):]...)...)
	}

	return *s
}

// ReplaceAll 全部替换.
func (s Slice[T]) ReplaceAll(oldSlice, newSlice []T) Slice[T] {
	return s.Replace(oldSlice, newSlice, -1)
}

// Has 包含.
func (s Slice[T]) Has(elem T) bool {
	for _, value := range s {
		// nolint
		if value == elem {
			return true
		}
	}

	return false
}

// Index 位置.
func (s Slice[T]) Index(elem T) int {
	for index, value := range s {
		// nolint
		if value == elem {
			return index
		}
	}

	return -1
}

// String 转换成字符串.
func (s Slice[T]) String() string {
	return "Slice[" + s.Join(" ") + "]"
}

// Join 集合连接.
func (s Slice[T]) Join(sep string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", s[0])
	}

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%v", s[0]))

	for _, elem := range s[1:] {
		builder.WriteString(sep)
		builder.WriteString(fmt.Sprintf("%v", elem))
	}

	return builder.String()
}
