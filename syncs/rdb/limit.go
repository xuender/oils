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
	_windowsNum = 10
	_expiration = time.Second / _windowsNum
)

// Limit 滑动窗口限流.
type Limit struct {
	key      string
	keyTimer string
	qps      int64
	client   *redis.Client
}

// NewLimit 频率限制，qps每秒限制次数.
func NewLimit(client *redis.Client, key string, qps uint) *Limit {
	ret := &Limit{
		client:   client,
		key:      key,
		keyTimer: fmt.Sprintf("_%s_", key),
		qps:      int64(qps / _windowsNum),
	}

	return ret
}

// Wait 等待执行.
func (p *Limit) Wait() {
	for {
		if wait := p.wait(); wait > 0 {
			time.Sleep(wait)
		} else {
			return
		}
	}
}

func (p *Limit) wait() time.Duration {
	ctx := context.Background()
	// 判断计时器是否存在
	if base.Must1(p.client.Exists(ctx, p.keyTimer).Result()) == 0 {
		// 不存在计时器: 创建计时器，删除计数器
		base.Must1(p.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			if _, err := pipe.Set(ctx, p.keyTimer, "1", _expiration).Result(); err != nil {
				return err
			}

			_, err := pipe.Del(ctx, p.key).Result()

			return err
		}))
	}
	// 计数器大与QPS
	if base.Must1(p.client.Incr(ctx, p.key).Result()) > p.qps {
		wait := base.Must1(p.client.PTTL(ctx, p.keyTimer).Result())

		if wait > 0 {
			return wait
		}
	}

	return 0
}

// Try 尝试执行.
func (p *Limit) Try() error {
	if p.wait() > 0 {
		return syncs.ErrLimit
	}

	return nil
}
