package base

import (
	"fmt"
	"strings"
)

// Set 通用set.
type Set[K comparable] map[K]struct{}

// NewSet 新建set.
func NewSet[K comparable](elems ...K) Set[K] {
	return Set[K]{}.Add(elems...)
}

// Add 增加.
func (set Set[K]) Add(elems ...K) Set[K] {
	for _, s := range elems {
		set[s] = None
	}

	return set
}

// AddSet 增加set.
func (set Set[K]) AddSet(elems ...Set[K]) Set[K] {
	for _, elem := range elems {
		for i := range elem {
			set[i] = None
		}
	}

	return set
}

// Any 包含任意个.
func (set Set[K]) Any(keys ...K) bool {
	for _, key := range keys {
		if _, has := set[key]; has {
			return true
		}
	}

	return false
}

// All 包含全部.
func (set Set[K]) All(keys ...K) bool {
	for _, key := range keys {
		if _, has := set[key]; !has {
			return false
		}
	}

	return true
}

// Del 删除.
func (set Set[K]) Del(keys ...K) Set[K] {
	for _, key := range keys {
		delete(set, key)
	}

	return set
}

// DelSet 删除set.
func (set Set[K]) DelSet(elems ...Set[K]) Set[K] {
	for _, elem := range elems {
		for i := range elem {
			delete(set, i)
		}
	}

	return set
}

// Has 包含.
func (set Set[K]) Has(key K) bool {
	_, has := set[key]

	return has
}

// Join 集合连接.
func (set Set[K]) Join(sep string) string {
	elems := set.Slice()

	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", elems[0])
	}

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%v", elems[0]))

	for _, elem := range elems[1:] {
		builder.WriteString(sep)
		builder.WriteString(fmt.Sprintf("%v", elem))
	}

	return builder.String()
}

// Slice 转换切片.
func (set Set[K]) Slice() []K {
	keys := make([]K, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	return keys
}

// String 字符串.
func (set Set[K]) String() string {
	return "Set[" + set.Join(" ") + "]"
}
