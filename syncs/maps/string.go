package maps

import "github.com/xuender/oils/base/treemap"

// String 键值 string.
type String[V any] struct {
	Maps[string, V]
}

func NewString[V any](notFoundValue V) *String[V] {
	return NewStringByGroup(_defaultGroup, notFoundValue)
}

func NewStringByGroup[V any](group int, notFoundValue V) *String[V] {
	return &String[V]{Maps: *NewByGroup(group, "", notFoundValue)}
}

// Prefix 前缀.
func (p *String[V]) Prefix(iterator Iterator[string, V], prefix string) {
	p.Range(iterator, prefix, treemap.PrefixNext(prefix))
}

// PrefixDesc 倒叙前缀.
func (p *String[V]) PrefixDesc(iterator Iterator[string, V], prefix string) {
	p.RangeDesc(iterator, prefix, treemap.PrefixNext(prefix))
}
