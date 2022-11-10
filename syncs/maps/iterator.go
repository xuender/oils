package maps

import (
	"github.com/xuender/oils/base/treemap"
	"github.com/xuender/oils/syncs"
	"golang.org/x/exp/constraints"
)

type (
	// Iterator 迭代器, 返回 false 终止迭代.
	Iterator[K constraints.Ordered, V any] func(key K, value V) bool
	item[K constraints.Ordered, V any]     struct {
		key   K
		value V
		index int
	}
)

func aes[K constraints.Ordered, V any](item1, item2 *item[K, V]) bool  { return item1.key < item2.key }
func desc[K constraints.Ordered, V any](item1, item2 *item[K, V]) bool { return item2.key < item1.key }

// Each 遍历.
func (p *Maps[K, V]) Each(iterator Iterator[K, V]) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.Each(ii)
		},
		aes[K, V],
	)
}

func (p *Maps[K, V]) each(
	iterator Iterator[K, V],
	run func(*treemap.TreeMap[K, V], treemap.ItemIterator[K, V]),
	less func(*item[K, V], *item[K, V]) bool,
) {
	ins := make([]chan *item[K, V], p.group)
	outs := make([]chan bool, p.group)

	for i := 0; i < p.group; i++ {
		ins[i] = make(chan *item[K, V])
		outs[i] = make(chan bool)
	}

	for i := 0; i < p.group; i++ {
		p.locks[i].RLock()
	}

	for group := 0; group < p.group; group++ {
		go func(index int) {
			run(p.maps[index], func(key K, value V) bool {
				ins[index] <- &item[K, V]{key: key, value: value, index: index}

				return <-outs[index]
			})
			close(ins[index])
			close(outs[index])
			outs[index] = nil
		}(group)
	}

	for elem := range syncs.Merge(less, ins...) {
		next := iterator(elem.key, elem.value)

		if next {
			outs[elem.index] <- true
		} else {
			for _, out := range outs {
				if out != nil {
					out <- false
				}
			}

			break
		}
	}

	for i := 0; i < p.group; i++ {
		p.locks[i].RUnlock()
	}
}

// Range 范围.
func (p *Maps[K, V]) Range(iterator Iterator[K, V], greaterOrEqual, lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.Range(ii, greaterOrEqual, lessThan)
		},
		aes[K, V],
	)
}

// GreateOrEqual 大于或等于.
func (p *Maps[K, V]) GreateOrEqual(iterator Iterator[K, V], greaterOrEqual K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.GreateOrEqual(ii, greaterOrEqual)
		},
		aes[K, V],
	)
}

// LessThan 小于.
func (p *Maps[K, V]) LessThan(iterator Iterator[K, V], lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.LessThan(ii, lessThan)
		},
		aes[K, V],
	)
}

// Prefix 字符串前缀.
// func (p *Maps[K, V]) Prefix(iterator Iterator[K, V], prefix string) {
// 	p.each(
// 		iterator,
// 		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
// 			t.Prefix(ii, prefix)
// 		},
// 		aes[K, V],
// 	)
// }

// EachDesc 倒叙遍历.
func (p *Maps[K, V]) EachDesc(iterator Iterator[K, V]) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.EachDesc(ii)
		},
		desc[K, V],
	)
}

// RangeDesc 倒叙范围.
func (p *Maps[K, V]) RangeDesc(iterator Iterator[K, V], greaterOrEqual, lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.RangeDesc(ii, greaterOrEqual, lessThan)
		},
		desc[K, V],
	)
}

// GreateOrEqualDesc 倒叙大于等于.
func (p *Maps[K, V]) GreateOrEqualDesc(iterator Iterator[K, V], greaterOrEqual K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.GreateOrEqualDesc(ii, greaterOrEqual)
		},
		desc[K, V],
	)
}

// LessThanDesc 倒叙小于.
func (p *Maps[K, V]) LessThanDesc(iterator Iterator[K, V], lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
			t.LessThanDesc(ii, lessThan)
		},
		desc[K, V],
	)
}

// PrefixDesc 倒叙字符串前缀.
// func (p *Maps[K, V]) PrefixDesc(iterator Iterator[K, V], prefix string) {
// 	p.each(
// 		iterator,
// 		func(t *treemap.TreeMap[K, V], ii treemap.ItemIterator[K, V]) {
// 			t.PrefixDesc(ii, prefix)
// 		},
// 		desc[K, V],
// 	)
// }
