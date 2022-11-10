package treemap

import (
	"golang.org/x/exp/constraints"
)

// Iterator 迭代器，返回 false 终止迭代..
type Iterator[K constraints.Ordered, V any] func(key K, value V) bool

// Each 遍历.
func (p *TreeMap[K, V]) Each(iterator Iterator[K, V]) {
	p.rangeAsc(iterator, p.root)
}

// Range 范围.
func (p *TreeMap[K, V]) Range(iterator Iterator[K, V], greaterOrEqual, lessThan K) {
	p.rangeAsc(iterator, p.root, greaterOrEqual, lessThan)
}

// GreateOrEqual 大于等于.
func (p *TreeMap[K, V]) GreateOrEqual(iterator Iterator[K, V], greaterOrEqual K) {
	p.rangeAsc(iterator, p.root, greaterOrEqual)
}

// LessThan 小于.
func (p *TreeMap[K, V]) LessThan(iterator Iterator[K, V], lessThan K) {
	p.rangeAsc(iterator, p.root, p.notFoundKey, lessThan)
}

func (p *TreeMap[K, V]) rangeAsc(iterator Iterator[K, V], elem *node[K, V], infs ...K) bool {
	if elem == nil {
		return true
	}

	if len(infs) > 1 && elem.key >= infs[1] {
		return p.rangeAsc(iterator, elem.left, infs...)
	}

	if len(infs) > 0 && elem.key < infs[0] {
		return p.rangeAsc(iterator, elem.right, infs...)
	}

	if !p.rangeAsc(iterator, elem.left, infs...) {
		return false
	}

	if !iterator(elem.key, elem.value) {
		return false
	}

	return p.rangeAsc(iterator, elem.right, infs...)
}

// EachDesc 倒叙遍历.
func (p *TreeMap[K, V]) EachDesc(iterator Iterator[K, V]) {
	p.rangeDesc(iterator, p.root)
}

// RangeDesc 倒叙范围.
func (p *TreeMap[K, V]) RangeDesc(iterator Iterator[K, V], greaterOrEqual, lessThan K) {
	p.rangeDesc(iterator, p.root, greaterOrEqual, lessThan)
}

// GreateOrEqualDesc 倒叙大于等于.
func (p *TreeMap[K, V]) GreateOrEqualDesc(iterator Iterator[K, V], greaterOrEqual K) {
	p.rangeDesc(iterator, p.root, greaterOrEqual)
}

// LessThanDesc 倒叙小于.
func (p *TreeMap[K, V]) LessThanDesc(iterator Iterator[K, V], lessThan K) {
	p.rangeDesc(iterator, p.root, p.notFoundKey, lessThan)
}

func (p *TreeMap[K, V]) rangeDesc(iterator Iterator[K, V], elem *node[K, V], infs ...K) bool {
	if elem == nil {
		return true
	}

	if len(infs) > 0 && elem.key < infs[0] {
		return p.rangeDesc(iterator, elem.right, infs...)
	}

	if len(infs) > 1 && elem.key >= infs[1] {
		return p.rangeDesc(iterator, elem.left, infs...)
	}

	if !p.rangeDesc(iterator, elem.right, infs...) {
		return false
	}

	if !iterator(elem.key, elem.value) {
		return false
	}

	return p.rangeDesc(iterator, elem.left, infs...)
}
