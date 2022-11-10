package maps

import (
	"sync"

	"github.com/xuender/oils/base/treemap"
	"golang.org/x/exp/constraints"
)

const _defaultGroup = 32

// Maps 线程安全的map.
type Maps[K constraints.Ordered, V any] struct {
	maps          []*treemap.TreeMap[K, V]
	locks         []sync.RWMutex
	group         int
	notFoundKey   K
	notFoundValue V
}

// New 新建map，默认分组32个.
func New[K constraints.Ordered, V any](notFoundKey K, notFoundValue V) *Maps[K, V] {
	return NewByGroup(_defaultGroup, notFoundKey, notFoundValue)
}

// NewByGroup 新建map，自定义分组数量.
func NewByGroup[K constraints.Ordered, V any](group int, notFoundKey K, notFoundValue V) *Maps[K, V] {
	maps := make([]*treemap.TreeMap[K, V], group)
	locks := make([]sync.RWMutex, group)

	for i := 0; i < group; i++ {
		maps[i] = treemap.New(notFoundKey, notFoundValue)
		locks[i] = sync.RWMutex{}
	}

	return &Maps[K, V]{
		maps:          maps,
		locks:         locks,
		group:         group,
		notFoundKey:   notFoundKey,
		notFoundValue: notFoundValue,
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
	group := Group(key, p.group)

	p.locks[group].RLock()
	defer p.locks[group].RUnlock()

	return p.maps[group].Get(key)
}

// Set 设值.
func (p *Maps[K, V]) Set(key K, value V) bool {
	group := Group(key, p.group)

	p.locks[group].Lock()
	defer p.locks[group].Unlock()

	return p.maps[group].Set(key, value)
}

// Add 增加，如果有则忽略.
func (p *Maps[K, V]) Add(key K, value V) bool {
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
	key := p.notFoundKey
	value := p.notFoundValue

	for group := 0; group < p.group; group++ {
		p.locks[group].RLock()
		valueTmp, keyTmp := p.maps[group].Min()

		if keyTmp != p.notFoundKey && (key == p.notFoundKey || key >= keyTmp) {
			key = keyTmp
			value = valueTmp
		}

		p.locks[group].RUnlock()
	}

	return value, key
}

// Max 最大键值.
func (p *Maps[K, V]) Max() (V, K) {
	key := p.notFoundKey
	value := p.notFoundValue

	for group := 0; group < p.group; group++ {
		p.locks[group].RLock()
		valueTmp, keyTmp := p.maps[group].Max()

		if keyTmp != p.notFoundKey && key <= keyTmp {
			key = keyTmp
			value = valueTmp
		}

		p.locks[group].RUnlock()
	}

	return value, key
}

// Del 删除.
func (p *Maps[K, V]) Del(key K) bool {
	group := Group(key, p.group)

	p.locks[group].Lock()
	defer p.locks[group].Unlock()

	return p.maps[group].Del(key)
}

// DelMin 删除最小键.
func (p *Maps[K, V]) DelMin() bool {
	_, key := p.Min()

	if key == p.notFoundKey {
		return false
	}

	return p.Del(key)
}

// DelMax 删除最大键.
func (p *Maps[K, V]) DelMax() bool {
	_, key := p.Max()

	if key == p.notFoundKey {
		return false
	}

	return p.Del(key)
}
