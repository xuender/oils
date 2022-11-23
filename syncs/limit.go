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

// QPS 设置qps.
func (p *Limit) QPS(qps uint) {
	p.interval = time.Second / time.Duration(qps)
}

// Wait 等待执行.
func (p *Limit) Wait() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if wait := p.waiting(); wait > 0 {
		time.Sleep(wait)

		p.last = p.last.Add(p.interval)
	}
}

func (p *Limit) waiting() time.Duration {
	dru := time.Since(p.last)

	if sleep := p.interval - dru; sleep > 0 {
		return sleep
	}

	p.last = p.last.Add(dru)

	return 0
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
