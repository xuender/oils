package cache

import (
	"sync"
	"time"
)

type cleaner interface {
	Clean(overdue time.Time) error
}

type cleaners struct {
	data     map[uint64]cleaner
	lock     sync.Mutex
	previous time.Time
}

func newCleaners() *cleaners {
	return &cleaners{
		data:     map[uint64]cleaner{},
		lock:     sync.Mutex{},
		previous: time.Now(),
	}
}

func (p *cleaners) Set(id uint64, cleaner cleaner) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.data[id] = cleaner
}

func (p *cleaners) Del(id uint64) {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.data, id)
}

func (p *cleaners) Clean() {
	now := time.Now()
	if now.Sub(p.previous) < time.Second {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	now = time.Now()
	if now.Sub(p.previous) < time.Second {
		return
	}

	for id, close := range p.data {
		if err := close.Clean(now); err != nil {
			delete(p.data, id)
		}
	}

	p.previous = time.Now()
}
