package syncs

import (
	"sync"
	"time"
)

// LeakyBucket 漏桶算法.
type LeakyBucket struct {
	capacity     uint
	leakInterval time.Duration
	used         uint
	lastLeak     time.Time
	mutex        sync.Mutex
}

// NewLeakyBucket 新建漏桶.
func NewLeakyBucket(capacity uint, leakInterval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		leakInterval: leakInterval,
		used:         0,
		lastLeak:     time.Now(),
	}
}

// TryConsume 尝试消费，消费失败返回 false.
func (p *LeakyBucket) TryConsume(drop uint) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.leak()

	if p.used+drop > p.capacity {
		return false
	}

	p.used += drop

	return true
}

// Consume 消费，无法消费 Sleep 直到可以消费.
func (p *LeakyBucket) Consume(drop uint) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for {
		sleep := p.leak()

		if p.used+drop > p.capacity {
			time.Sleep(sleep)

			continue
		}

		p.used += drop

		return
	}
}

func (p *LeakyBucket) leak() time.Duration {
	now := time.Now()

	if sub := now.Sub(p.lastLeak); p.leakInterval > sub {
		return p.leakInterval - sub
	}

	p.lastLeak = now
	p.used = 0

	return 0
}
