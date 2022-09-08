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
	return &Limit{interval: time.Second / time.Duration(qps)}
}

// Try 尝试执行.
func (p *Limit) Try() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := time.Now()
	sleep := p.interval - now.Sub(p.last)

	if sleep <= 0 {
		p.last = now

		return
	}

	time.Sleep(sleep)
	p.last = time.Now()
}
