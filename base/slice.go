package base

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

// Slice 切片.
type Slice[T constraints.Ordered] []T

// NewSlice 新建切片.
func NewSlice[T constraints.Ordered](elems ...T) Slice[T] {
	newSlice := Slice[T]{}
	newSlice.Add(elems...)

	return newSlice
}

// Add 增加.
func (p *Slice[T]) Add(elems ...T) Slice[T] {
	*p = append(*p, elems...)

	return *p
}

// Del 删除.
func (p *Slice[T]) Del(elems ...T) Slice[T] {
	for _, elem := range elems {
		if i := p.Index(elem); i >= 0 {
			*p = append((*p)[0:i], (*p)[i+1:]...)
		}
	}

	return *p
}

// DelAll 删除全部.
func (p *Slice[T]) DelAll(elems ...T) Slice[T] {
	for _, elem := range elems {
		for i := p.Index(elem); i >= 0; i = p.Index(elem) {
			*p = append((*p)[0:i], (*p)[i+1:]...)
		}
	}

	return *p
}

// Replace 替换.
func (p *Slice[T]) Replace(oldSlice, newSlice []T, num int) Slice[T] {
	if num == 0 {
		return *p
	}

	if count := p.Counts(oldSlice); count == 0 {
		return *p
	} else if num < 0 || count < num {
		num = count
	}

	start := 0
	for i := 0; i < num; i++ {
		index := (*p)[start:].Indexs(oldSlice)
		*p = append((*p)[0:start+index], append(newSlice, (*p)[start+index+len(oldSlice):]...)...)
	}

	return *p
}

// Unique 唯一.
func (p *Slice[T]) Unique() Slice[T] {
	unique := NewSlice[T]()

	for _, elem := range *p {
		if unique.Has(elem) {
			continue
		}

		unique.Add(elem)
	}

	*p = unique

	return *p
}

// All 全部包含.
func (p Slice[T]) All(elems ...T) bool {
	for _, e := range elems {
		if !p.Has(e) {
			return false
		}
	}

	return true
}

// Any 任意包含.
func (p Slice[T]) Any(elems ...T) bool {
	for _, e := range elems {
		if p.Has(e) {
			return true
		}
	}

	return false
}

// Count 包含元素数量.
func (p Slice[T]) Count(elem T) (count int) {
	for _, value := range p {
		// nolint
		if value == elem {
			count++
		}
	}

	return
}

// Count 包含切片数量.
func (p Slice[T]) Counts(elems []T) (count int) {
	if len(p) < len(elems) {
		return
	}

	for index := 0; index <= len(p)-len(elems); index++ {
		has := true

		for elemIndex, elem := range elems {
			// nolint
			if p[index+elemIndex] != elem {
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

// Equal 比较.
func (p Slice[T]) Equal(target Slice[T]) bool {
	if len(p) != len(target) {
		return false
	}

	for i, elem := range target {
		// nolint
		if p[i] != elem {
			return false
		}
	}

	return true
}

// Has 包含.
func (p Slice[T]) Has(elem T) bool {
	for _, value := range p {
		// nolint
		if value == elem {
			return true
		}
	}

	return false
}

// Index 位置.
func (p Slice[T]) Index(elem T) int {
	for index, value := range p {
		// nolint
		if value == elem {
			return index
		}
	}

	return -1
}

// Indexs 包含.
func (p Slice[T]) Indexs(elems []T) int {
	if len(p) < len(elems) {
		return -1
	}

	for index := 0; index <= len(p)-len(elems); index++ {
		has := true

		for elemIndex, elem := range elems {
			// nolint
			if p[index+elemIndex] != elem {
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

// Join 集合连接.
func (p Slice[T]) Join(sep string) string {
	switch len(p) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", p[0])
	}

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%v", p[0]))

	for _, elem := range p[1:] {
		builder.WriteString(sep)
		builder.WriteString(fmt.Sprintf("%v", elem))
	}

	return builder.String()
}

// Len 长度.
func (p Slice[T]) Len() int { return len(p) }

// Less 小于.
func (p Slice[T]) Less(i, j int) bool {
	// nolint
	return p[i] < p[j]
}

// ReplaceAll 全部替换.
func (p Slice[T]) ReplaceAll(oldSlice, newSlice []T) Slice[T] {
	return p.Replace(oldSlice, newSlice, -1)
}

// String 转换成字符串.
func (p Slice[T]) String() string {
	return "Slice[" + p.Join(" ") + "]"
}

// Swap 交换.
func (p Slice[T]) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
