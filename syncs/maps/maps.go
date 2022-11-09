package maps

import (
	"bytes"
	"sync"

	"github.com/xuender/oils/base/treemap"
)

const _defaultGroup = 32

// Maps 线程安全的map.
type Maps[K, V any] struct {
	maps     []*treemap.TreeMap[V]
	locks    []sync.RWMutex
	group    int
	notFound V
}

// New 新建map，默认分组32个.
func New[K, V any](notFound V) *Maps[K, V] {
	return NewByGroup[K](_defaultGroup, notFound)
}

// NewByGroup 新建map，自定义分组数量.
func NewByGroup[K, V any](group int, notFound V) *Maps[K, V] {
	maps := make([]*treemap.TreeMap[V], group)
	locks := make([]sync.RWMutex, group)

	for i := 0; i < group; i++ {
		maps[i] = treemap.New(notFound)
		locks[i] = sync.RWMutex{}
	}

	return &Maps[K, V]{
		maps:     maps,
		locks:    locks,
		group:    group,
		notFound: notFound,
	}
}

// Len map存储数量.
func (p *Maps[K, V]) Len() int {
	sum := 0

	for i := 0; i < p.group; i++ {
		p.locks[i].RLock()
		sum += p.maps[i].Len()
		p.locks[i].RUnlock()
	}

	return sum
}

// Get 取值.
func (p *Maps[K, V]) Get(key K) (V, bool) {
	return p.GetByBytes(Encode(key))
}

// GetByBytes 根据[]byte键取值.
func (p *Maps[K, V]) GetByBytes(key []byte) (V, bool) {
	group := Group(key, p.group)

	p.locks[group].RLock()
	defer p.locks[group].RUnlock()

	return p.maps[group].Get(key)
}

// Set 设值.
func (p *Maps[K, V]) Set(key K, value V) bool {
	return p.SetByBytes(Encode(key), value)
}

// SetByBytes 根据[]byte键设值.
func (p *Maps[K, V]) SetByBytes(key []byte, value V) bool {
	group := Group(key, p.group)

	p.locks[group].Lock()
	defer p.locks[group].Unlock()

	return p.maps[group].Set(key, value)
}

// Add 增加，如果有则忽略.
func (p *Maps[K, V]) Add(key K, value V) bool {
	return p.AddByBytes(Encode(key), value)
}

// AddByBytes 根据[]byte键增加，如果有则忽略.
func (p *Maps[K, V]) AddByBytes(key []byte, value V) bool {
	group := Group(key, p.group)

	p.locks[group].Lock()
	defer p.locks[group].Unlock()

	return p.maps[group].Add(key, value)
}

// Clear 清空.
func (p *Maps[K, V]) Clear() {
	for i := 0; i < p.group; i++ {
		p.locks[i].Lock()
		p.maps[i].Clear()
		p.locks[i].Unlock()
	}
}

// Min 最小键值.
func (p *Maps[K, V]) Min() (V, K) {
	value, key := p.MinByBytes()

	return value, Decode[K](key)
}

// Min 最小键值,[]byte.
func (p *Maps[K, V]) MinByBytes() (V, []byte) {
	var key []byte

	value := p.notFound

	for index := 0; index < p.group; index++ {
		p.locks[index].RLock()
		valueTmp, keyTmp := p.maps[index].Min()

		if keyTmp != nil && (key == nil || bytes.Compare(key, keyTmp) >= 0) {
			key = keyTmp
			value = valueTmp
		}

		p.locks[index].RUnlock()
	}

	return value, key
}

// Max 最大键值.
func (p *Maps[K, V]) Max() (V, K) {
	value, key := p.MaxByBytes()

	return value, Decode[K](key)
}

// Min 最小键值,[]byte.
func (p *Maps[K, V]) MaxByBytes() (V, []byte) {
	var key []byte

	value := p.notFound

	for index := 0; index < p.group; index++ {
		p.locks[index].RLock()
		valueTmp, keyTmp := p.maps[index].Max()

		if keyTmp != nil && bytes.Compare(key, keyTmp) <= 0 {
			key = keyTmp
			value = valueTmp
		}

		p.locks[index].RUnlock()
	}

	return value, key
}

// Del 删除.
func (p *Maps[K, V]) Del(key K) bool {
	return p.DelByBytes(Encode(key))
}

// DelByBytes 根据[]byte键删除.
func (p *Maps[K, V]) DelByBytes(key []byte) bool {
	group := Group(key, p.group)

	p.locks[group].Lock()
	defer p.locks[group].Unlock()

	return p.maps[group].Del(key)
}

// DelMin 删除最小键.
func (p *Maps[K, V]) DelMin() bool {
	_, key := p.MinByBytes()

	if key == nil {
		return false
	}

	return p.DelByBytes(key)
}

// DelMax 删除最大键.
func (p *Maps[K, V]) DelMax() bool {
	_, key := p.MaxByBytes()

	if key == nil {
		return false
	}

	return p.DelByBytes(key)
}
