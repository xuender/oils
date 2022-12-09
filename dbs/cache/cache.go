// Package cache 简单的内存缓存.
//
// 如果需要复杂的功能可以使用 github.com/allegro/bigcache 或 https://github.com/coocood/freecache.
package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/oils/dbs"
)

// nolint
var caches = newCleaners()

// Cache support LRU (Least Recently Used).
// nolint
type Cache[K comparable, V any] struct {
	id     uint64
	lock   sync.RWMutex
	data   map[K]V
	access map[K]time.Time

	Expire   time.Duration
	Callback func(key K, value V)
}

// Set value by key.
func (p *Cache[K, V]) Set(key K, value V) error {
	if p.id == 0 {
		return ErrClose
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.data[key] = value
	p.access[key] = time.Now().Add(p.Expire)

	return nil
}

// SetByTime value by key.
func (p *Cache[K, V]) SetByTime(key K, value V, expire time.Time) error {
	if p.id == 0 {
		return ErrClose
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.data[key] = value
	p.access[key] = expire

	return nil
}

// SetByDuration value by key.
func (p *Cache[K, V]) SetByDuration(key K, value V, expire time.Duration) error {
	if p.id == 0 {
		return ErrClose
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.data[key] = value
	p.access[key] = time.Now().Add(expire)

	return nil
}

func (p *Cache[K, V]) get(key K) (V, bool) {
	p.lock.RLock()

	value, has := p.data[key]
	if has {
		if access, ahas := p.access[key]; ahas && access.Before(time.Now()) {
			p.Callback(key, value)

			delete(p.data, key)
			delete(p.access, key)
			p.lock.RUnlock()

			caches.Clean()

			return value, false
		}
	}

	p.lock.RUnlock()

	return value, has
}

// Get 获取值并修改过期时间.
func (p *Cache[K, V]) Get(key K) (V, bool) {
	value, has := p.get(key)

	if has {
		_ = p.Reset(key)
	}

	return value, has
}

// GetOrSet 获取或设置.
func (p *Cache[K, V]) GetOrSet(key K, value V) (V, bool) {
	old, has := p.get(key)
	if has {
		_ = p.Reset(key)

		return old, has
	}

	_ = p.Set(key, value)

	return value, has
}

// GetOrLoad 获取或加载.
func (p *Cache[K, V]) GetOrLoad(key K, loader func(key K) V) (V, bool) {
	old, has := p.get(key)
	if has {
		_ = p.Reset(key)

		return old, has
	}

	value := loader(key)

	_ = p.Set(key, value)

	return value, has
}

// GetOnly 只获取不修改过期时间.
func (p *Cache[K, V]) GetOnly(key K) (V, bool) {
	value, has := p.get(key)

	return value, has
}

// GetString by key.
func (p *Cache[K, V]) GetString(key K) (string, bool) {
	if value, ok := p.Get(key); ok {
		return fmt.Sprintf("%v", value), ok
	}

	return "", false
}

// Reset expire by key.
func (p *Cache[K, V]) Reset(key K) error {
	if p.id == 0 {
		return ErrClose
	}

	p.lock.RLock()
	defer p.lock.RUnlock()

	if a, ok := p.access[key]; ok {
		e := time.Now().Add(p.Expire)
		if e.After(a) {
			p.access[key] = e
		}
	}

	return nil
}

// Keys by map.
func (p *Cache[K, V]) Keys() []K {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return lo.Keys(p.data)
}

// Del key.
func (p *Cache[K, V]) Del(keys ...K) error {
	if p.id == 0 {
		return ErrClose
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	for _, key := range keys {
		if value, ok := p.data[key]; ok {
			p.Callback(key, value)

			delete(p.data, key)
			delete(p.access, key)
		}
	}

	return nil
}

// Clean overdue.
func (p *Cache[K, V]) Clean(overdue time.Time) error {
	return p.Del(p.Overdue(overdue)...)
}

func (p *Cache[K, V]) Close() {
	caches.Del(p.id)

	p.id = 0
}

// Overdue keys.
func (p *Cache[K, V]) Overdue(overdue time.Time) []K {
	p.lock.RLock()
	defer p.lock.RUnlock()

	keys := make([]K, 0, len(p.access))

	for key, v := range p.access {
		if overdue.After(v) {
			keys = append(keys, key)
		}
	}

	return keys
}

// Size by Cache.
func (p *Cache[K, V]) Size() int {
	_ = p.Clean(time.Now())

	return len(p.data)
}

// NewCache new cache.
// nolint: varnamelen
func NewCache[K comparable, V any](expire time.Duration) *Cache[K, V] {
	cache := &Cache[K, V]{
		id:       dbs.ID(),
		data:     map[K]V{},
		access:   map[K]time.Time{},
		Expire:   expire,
		Callback: func(key K, value V) {},
	}

	caches.Set(cache.id, cache)

	return cache
}

// NewCallbackCache new have del callback cache.
func NewCallbackCache[K comparable, V any](
	expire time.Duration,
	callback func(key K, value V),
) *Cache[K, V] {
	cache := NewCache[K, V](expire)
	cache.Callback = callback

	return cache
}
