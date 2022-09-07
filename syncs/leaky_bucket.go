package syncs

import (
	"sync"
	"time"

	"github.com/xuender/oils/base"
)

// LeakyBucket 漏桶算法.
type LeakyBucket struct {
	capacity     int
	leakInterval time.Duration
	used         int
	lastLeak     time.Time
	mutex        sync.Mutex
}

// NewLeakyBucket 新建漏桶.
func NewLeakyBucket(capacity uint, leakInterval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:     int(capacity),
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

	if p.used+int(drop) > p.capacity {
		return false
	}

	p.used += int(drop)

	return true
}

// Consume 消费，如果无法消费，则 Sleep 直到可以消费.
func (p *LeakyBucket) Consume(drop uint) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for {
		sleep := p.leak()

		if p.used+int(drop) > p.capacity {
			time.Sleep(sleep)

			continue
		}

		p.used += int(drop)

		return
	}
}

func (p *LeakyBucket) leak() time.Duration {
	now := time.Now()
	sub := now.Sub(p.lastLeak)
	// 不到漏水时间，返回时间差
	if p.leakInterval > sub {
		return p.leakInterval - sub
	}
	// 泄漏一个桶的容量或者根据时间间隔泄漏
	p.used -= base.Max(p.capacity, int(sub/p.leakInterval)*p.capacity)
	p.lastLeak = now
	// 如果桶内余额大于桶的容量，返回足够的间隔时间
	if p.used > p.capacity {
		return p.leakInterval * time.Duration(p.used/p.capacity)
	}

	return 0
}
