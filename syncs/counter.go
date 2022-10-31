package syncs

import (
	"sync"
	"sync/atomic"
)

// Counter 计数器.
type Counter[T comparable] struct {
	keys  map[T]int
	value []int64
	mutex sync.Mutex
}

// NewCounter 新建计数器.
func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{
		keys:  map[T]int{},
		value: []int64{},
		mutex: sync.Mutex{},
	}
}

// Get 获取key的计数.
func (p *Counter[T]) Get(key T) int64 {
	return atomic.LoadInt64(&p.value[p.index(key)])
}

// index 获取key的位置.
func (p *Counter[T]) index(key T) int {
	if index, has := p.keys[key]; has {
		return index
	}

	return p.addKey(key)
}

// addKey 增加key.
func (p *Counter[T]) addKey(key T) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if index, has := p.keys[key]; has {
		return index
	}

	p.keys[key] = len(p.keys)
	p.value = append(p.value, 0)

	return len(p.keys) - 1
}

// Inc 自增key.
func (p *Counter[T]) Inc(key T) {
	p.Add(key, 1)
}

// Dec 自减key.
func (p *Counter[T]) Dec(key T) {
	p.Add(key, -1)
}

// Add 增加key的值num.
func (p *Counter[T]) Add(key T, num int64) {
	atomic.AddInt64(&p.value[p.index(key)], num)
}

// Clean 清空.
func (p *Counter[T]) Clean() {
	for i := range p.value {
		atomic.StoreInt64(&p.value[i], 0)
	}
}
