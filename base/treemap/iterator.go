package treemap

import (
	"bytes"
)

type ItemIterator[V any] func(key []byte, value V) bool

func (p *TreeMap[V]) Each(iterator ItemIterator[V]) {
	p.rangeAsc(iterator, p.root)
}

func (p *TreeMap[V]) Range(iterator ItemIterator[V], greaterOrEqual, lessThan []byte) {
	p.rangeAsc(iterator, p.root, greaterOrEqual, lessThan)
}

func (p *TreeMap[V]) GreateOrEqual(iterator ItemIterator[V], greaterOrEqual []byte) {
	p.rangeAsc(iterator, p.root, greaterOrEqual)
}

func (p *TreeMap[V]) LessThan(iterator ItemIterator[V], lessThan []byte) {
	p.rangeAsc(iterator, p.root, nil, lessThan)
}

func (p *TreeMap[V]) Prefix(iterator ItemIterator[V], prefix []byte) {
	p.rangeAsc(iterator, p.root, prefix, BytesInc(prefix))
}

func (p *TreeMap[V]) rangeAsc(iterator ItemIterator[V], elem *node[V], infs ...[]byte) bool {
	if elem == nil {
		return true
	}

	if len(infs) > 1 && bytes.Compare(elem.key, infs[1]) >= 0 {
		return p.rangeAsc(iterator, elem.left, infs...)
	}

	if len(infs) > 0 && bytes.Compare(elem.key, infs[0]) < 0 {
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

func (p *TreeMap[V]) EachDesc(iterator ItemIterator[V]) {
	p.rangeDesc(iterator, p.root)
}

func (p *TreeMap[V]) RangeDesc(iterator ItemIterator[V], greaterOrEqual, lessThan []byte) {
	p.rangeDesc(iterator, p.root, greaterOrEqual, lessThan)
}

func (p *TreeMap[V]) GreateOrEqualDesc(iterator ItemIterator[V], greaterOrEqual []byte) {
	p.rangeDesc(iterator, p.root, greaterOrEqual)
}

func (p *TreeMap[V]) LessThanDesc(iterator ItemIterator[V], lessThan []byte) {
	p.rangeDesc(iterator, p.root, nil, lessThan)
}

func (p *TreeMap[V]) PrefixDesc(iterator ItemIterator[V], prefix []byte) {
	p.rangeDesc(iterator, p.root, prefix, BytesInc(prefix))
}

func (p *TreeMap[V]) rangeDesc(iterator ItemIterator[V], elem *node[V], infs ...[]byte) bool {
	if elem == nil {
		return true
	}

	if len(infs) > 0 && bytes.Compare(elem.key, infs[0]) < 0 {
		return p.rangeDesc(iterator, elem.right, infs...)
	}

	if len(infs) > 1 && bytes.Compare(elem.key, infs[1]) >= 0 {
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
