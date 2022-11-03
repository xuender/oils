package rdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/syncs"
)

const (
	_windowNum  = 10
	_expiration = time.Second / _windowNum
)

// Limit 滑动窗口限流.
type Limit struct {
	key      string
	keyTimer string
	qps      int64
	client   *redis.Client
}

// NewLimit 频率限制, client redis客户端, key 键，qps 每秒限制次数.
func NewLimit(client *redis.Client, key string, qps uint) *Limit {
	ret := &Limit{
		client:   client,
		key:      key,
		keyTimer: fmt.Sprintf("_%s_", key),
		qps:      int64(qps / _windowNum),
	}

	return ret
}

// Wait 等待执行.
func (p *Limit) Wait() {
	for dur := p.waiting(); dur > 0; dur = p.waiting() {
		time.Sleep(dur)
	}
}

func (p *Limit) waiting() time.Duration {
	ctx := context.Background()
	// 设置计时器
	if base.Must1(p.client.SetNX(ctx, p.keyTimer, "1", _expiration).Result()) {
		// 计时器设置成功后删除计数器
		base.Must1(p.client.Del(ctx, p.key).Result())
	}
	// 计数器大与QPS
	if base.Must1(p.client.Incr(ctx, p.key).Result()) > p.qps {
		// 获取计时器剩余时间
		if dur := base.Must1(p.client.PTTL(ctx, p.keyTimer).Result()); dur > 0 {
			return dur
		}
	}

	return 0
}

// Try 尝试执行.
func (p *Limit) Try() error {
	if p.waiting() > 0 {
		return syncs.ErrLimit
	}

	return nil
}
