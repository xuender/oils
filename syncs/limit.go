package syncs

import (
	"sync"
	"time"
)

// Limiter 等待器.
type Limiter interface {
	Wait()
	Try() error
}

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
	if wait := p.wait(); wait > 0 {
		time.Sleep(wait)
	}
}

func (p *Limit) wait() time.Duration {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := time.Now()
	sleep := p.interval - now.Sub(p.last)

	if sleep <= 0 {
		p.last = now

		return 0
	}

	p.last = now.Add(sleep)

	return sleep
}

// Try 尝试执行.
func (p *Limit) Try() error {
	if p.wait() > 0 {
		return ErrLimit
	}

	return nil
}
