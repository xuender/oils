package treemap

import (
	"golang.org/x/exp/constraints"
)

type ItemIterator[K constraints.Ordered, V any] func(key K, value V) bool

// Each 遍历.
func (p *TreeMap[K, V]) Each(iterator ItemIterator[K, V]) {
	p.rangeAsc(iterator, p.root)
}

// Range 范围.
func (p *TreeMap[K, V]) Range(iterator ItemIterator[K, V], greaterOrEqual, lessThan K) {
	p.rangeAsc(iterator, p.root, greaterOrEqual, lessThan)
}

// GreateOrEqual 大于等于.
func (p *TreeMap[K, V]) GreateOrEqual(iterator ItemIterator[K, V], greaterOrEqual K) {
	p.rangeAsc(iterator, p.root, greaterOrEqual)
}

// LessThan 小于.
func (p *TreeMap[K, V]) LessThan(iterator ItemIterator[K, V], lessThan K) {
	p.rangeAsc(iterator, p.root, p.notFoundKey, lessThan)
}

func (p *TreeMap[K, V]) rangeAsc(iterator ItemIterator[K, V], elem *node[K, V], infs ...K) bool {
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
func (p *TreeMap[K, V]) EachDesc(iterator ItemIterator[K, V]) {
	p.rangeDesc(iterator, p.root)
}

// RangeDesc 倒叙范围.
func (p *TreeMap[K, V]) RangeDesc(iterator ItemIterator[K, V], greaterOrEqual, lessThan K) {
	p.rangeDesc(iterator, p.root, greaterOrEqual, lessThan)
}

// GreateOrEqualDesc 倒叙大于等于.
func (p *TreeMap[K, V]) GreateOrEqualDesc(iterator ItemIterator[K, V], greaterOrEqual K) {
	p.rangeDesc(iterator, p.root, greaterOrEqual)
}

// LessThanDesc 倒叙小于.
func (p *TreeMap[K, V]) LessThanDesc(iterator ItemIterator[K, V], lessThan K) {
	p.rangeDesc(iterator, p.root, p.notFoundKey, lessThan)
}

func (p *TreeMap[K, V]) rangeDesc(iterator ItemIterator[K, V], elem *node[K, V], infs ...K) bool {
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
