package rdb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/base"
)

// Locker 分布式锁.
type Locker struct {
	client     redis.Cmdable
	keys       sync.Map
	lockTime   time.Duration
	expireTime time.Duration
}

// NewLocker 新建分布式锁.
// times[0] 锁时间, times[1] 续期间隔.
func NewLocker(client redis.Cmdable, times ...time.Duration) *Locker {
	locker := &Locker{
		client:     client,
		keys:       sync.Map{},
		lockTime:   time.Second * base.Three,
		expireTime: time.Second,
	}

	if len(times) > 0 {
		locker.lockTime = times[0]
	}

	if len(times) > 1 {
		locker.expireTime = times[1]
	}

	if locker.expireTime >= locker.lockTime {
		panic(ErrExpire)
	}

	go locker.autoExpire()

	return locker
}

// Lock 锁成功则执行，执行完毕解锁.
func (p *Locker) Lock(key string, yield func() error) error {
	key = fmt.Sprintf("oils-%s-LOCK", key)
	ctx := context.Background()
	res, err := p.client.SetNX(ctx, key, true, p.lockTime).Result()

	if err == nil {
		if res {
			p.keys.Store(key, true)
			defer p.keys.Delete(key)

			return base.Errors(yield(), p.client.Del(ctx, key).Err())
		}

		return ErrLock
	}

	return err
}

// autoExpire 自动续期.
func (p *Locker) autoExpire() {
	ticker := time.NewTicker(p.expireTime)
	ctx := context.Background()

	for range ticker.C {
		p.keys.Range(func(key, value any) bool {
			base.Must(p.client.Expire(ctx, base.Conversion[string](key, true), p.lockTime).Err())

			return true
		})
	}
}
