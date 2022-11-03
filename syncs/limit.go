package syncs

import (
	"sync"
	"time"
)

// Limit 频率限制.
type Limit struct {
	interval time.Duration
	last     time.Time
	mutex    sync.Mutex
}

// NewLimit 频率限制，qps每秒限制次数.
func NewLimit(qps uint) *Limit {
	return &Limit{
		interval: time.Second / time.Duration(qps),
		mutex:    sync.Mutex{},
		last:     time.Now(),
	}
}

// Wait 等待执行.
func (p *Limit) Wait() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if wait := p.waiting(); wait > 0 {
		time.Sleep(wait)
		p.last = time.Now()
	}
}

func (p *Limit) waiting() time.Duration {
	now := time.Now()
	sleep := p.interval - now.Sub(p.last)

	if sleep <= 0 {
		p.last = now

		return 0
	}

	return sleep
}

// Try 尝试执行.
func (p *Limit) Try() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.waiting() > 0 {
		return ErrLimit
	}

	return nil
}
