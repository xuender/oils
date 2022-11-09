package maps

import (
	"bytes"

	"github.com/xuender/oils/base/treemap"
	"github.com/xuender/oils/syncs"
)

type (
	// Iterator 迭代器, 返回 false 终止迭代.
	Iterator[K, V any] func(key K, value V) bool
	item[V any]        struct {
		key   []byte
		value V
		index int
	}
)

func aes[V any](item1, item2 *item[V]) bool  { return bytes.Compare(item1.key, item2.key) < 0 }
func desc[V any](item1, item2 *item[V]) bool { return bytes.Compare(item2.key, item1.key) < 0 }

// Each 遍历.
func (p *Maps[K, V]) Each(iterator Iterator[K, V]) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.Each(ii)
		},
		aes[V],
	)
}

func (p *Maps[K, V]) each(
	iterator Iterator[K, V],
	run func(*treemap.TreeMap[V], treemap.ItemIterator[V]),
	less func(*item[V], *item[V]) bool,
) {
	ins := make([]chan *item[V], p.group)
	outs := make([]chan bool, p.group)

	for i := 0; i < p.group; i++ {
		ins[i] = make(chan *item[V])
		outs[i] = make(chan bool)
	}

	for i := 0; i < p.group; i++ {
		p.locks[i].RLock()
	}

	for group := 0; group < p.group; group++ {
		go func(index int) {
			run(p.maps[index], func(key []byte, value V) bool {
				ins[index] <- &item[V]{key: key, value: value, index: index}

				return <-outs[index]
			})
			close(ins[index])
			close(outs[index])
			outs[index] = nil
		}(group)
	}

	for elem := range syncs.Merge(less, ins...) {
		key := Decode[K](elem.key)

		next := iterator(key, elem.value)

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
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.Range(ii, Encode(greaterOrEqual), Encode(lessThan))
		},
		aes[V],
	)
}

// GreateOrEqual 大于或等于.
func (p *Maps[K, V]) GreateOrEqual(iterator Iterator[K, V], greaterOrEqual K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.GreateOrEqual(ii, Encode(greaterOrEqual))
		},
		aes[V],
	)
}

// LessThan 小于.
func (p *Maps[K, V]) LessThan(iterator Iterator[K, V], lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.LessThan(ii, Encode(lessThan))
		},
		aes[V],
	)
}

// Prefix 字符串前缀.
func (p *Maps[K, V]) Prefix(iterator Iterator[K, V], prefix string) {
	p.PrefixBytes(iterator, []byte(prefix))
}

// PrefixBytes 字节码前缀.
func (p *Maps[K, V]) PrefixBytes(iterator Iterator[K, V], prefix []byte) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.Prefix(ii, prefix)
		},
		aes[V],
	)
}

// EachDesc 倒叙遍历.
func (p *Maps[K, V]) EachDesc(iterator Iterator[K, V]) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.EachDesc(ii)
		},
		desc[V],
	)
}

// RangeDesc 倒叙范围.
func (p *Maps[K, V]) RangeDesc(iterator Iterator[K, V], greaterOrEqual, lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.RangeDesc(ii, Encode(greaterOrEqual), Encode(lessThan))
		},
		desc[V],
	)
}

// GreateOrEqualDesc 倒叙大于等于.
func (p *Maps[K, V]) GreateOrEqualDesc(iterator Iterator[K, V], greaterOrEqual K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.GreateOrEqualDesc(ii, Encode(greaterOrEqual))
		},
		desc[V],
	)
}

// LessThanDesc 倒叙小于.
func (p *Maps[K, V]) LessThanDesc(iterator Iterator[K, V], lessThan K) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.LessThanDesc(ii, Encode(lessThan))
		},
		desc[V],
	)
}

// PrefixDesc 倒叙字符串前缀.
func (p *Maps[K, V]) PrefixDesc(iterator Iterator[K, V], prefix string) {
	p.PrefixBytesDesc(iterator, []byte(prefix))
}

// PrefixBytesDesc 倒叙字节码前缀.
func (p *Maps[K, V]) PrefixBytesDesc(iterator Iterator[K, V], prefix []byte) {
	p.each(
		iterator,
		func(t *treemap.TreeMap[V], ii treemap.ItemIterator[V]) {
			t.PrefixDesc(ii, prefix)
		},
		desc[V],
	)
}
