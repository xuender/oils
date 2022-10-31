package syncs

import (
	"sync"
	"sync/atomic"

	"golang.org/x/exp/slices"
)

// Counter 计数器.
type Counter[T comparable] struct {
	keys   []T
	values []int64
	mutex  sync.RWMutex
	size   int
}

// NewCounter 新建计数器.
func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{
		keys:   []T{},
		values: []int64{},
		mutex:  sync.RWMutex{},
		size:   0,
	}
}

// Inc 自增key.
func (p *Counter[T]) Inc(key T) int64 {
	return p.Add(key, 1)
}

// Add 增加key的值num.
func (p *Counter[T]) Add(key T, num int64) int64 {
	return atomic.AddInt64(&p.values[p.index(key)], num)
}

// index 获取key的位置.
func (p *Counter[T]) index(key T) int {
	p.mutex.RLock()

	if index := slices.Index(p.keys[:p.size], key); index >= 0 {
		p.mutex.RUnlock()

		return index
	}

	p.mutex.RUnlock()

	return p.addKey(key)
}

// addKey 增加key.
func (p *Counter[T]) addKey(key T) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	num := int(p.size)

	if index := slices.Index(p.keys[:num], key); index >= 0 {
		return index
	}

	if len(p.keys) > num {
		p.keys[num] = key
		p.values[num] = 0
	} else {
		p.keys = append(p.keys, key)
		p.values = append(p.values, 0)
	}

	p.size++

	return num
}

// Dec 自减key.
func (p *Counter[T]) Dec(key T) int64 {
	return p.Add(key, -1)
}

// Size 键值数量.
func (p *Counter[T]) Size() int {
	return int(p.size)
}

// Sum 计数总量.
func (p *Counter[T]) Sum() int64 {
	var ret int64

	for index := 0; index < p.size; index++ {
		ret += atomic.LoadInt64(&p.values[index])
	}

	return ret
}

// Get 获取key的计数.
func (p *Counter[T]) Get(key T) int64 {
	return atomic.LoadInt64(&p.values[p.index(key)])
}

// Clean 清空.
func (p *Counter[T]) Clean() {
	for i := 0; i < p.size; i++ {
		atomic.StoreInt64(&p.values[i], 0)
	}

	p.size = 0
}
