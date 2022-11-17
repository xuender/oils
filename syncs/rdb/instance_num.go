package rdb

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/modes"
)

// InstanceNum 实例数量.
type InstanceNum struct {
	client redis.Cmdable
	count  uint64
	old    uint64
	sub    *modes.Subject[uint64, int]
	key    string
}

// NewInstanceNum 新建实例数量.
func NewInstanceNum(client redis.Cmdable) *InstanceNum {
	return NewInstanceNumByKey(client, "ins_num")
}

func NewInstanceNumByKey(client redis.Cmdable, key string) *InstanceNum {
	return &InstanceNum{
		client: client,
		sub:    &modes.Subject[uint64, int]{},
		key:    key,
		count:  1,
	}
}

// Register 订阅主题.
func (p *InstanceNum) Register(observer modes.Observer[uint64, int]) {
	p.sub.Register(observer)
}

// Run 运行.
//
// 使用 redis INCR 命令计算一秒钟内增加的数量得出实例数量.
func (p *InstanceNum) Run() {
	ctx := context.Background()
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		num, err := p.client.Incr(ctx, p.key).Result()
		if err != nil {
			p.client.Del(ctx, p.key)

			num = base.Must1(p.client.Incr(ctx, p.key).Result())
		}

		old := p.count

		if p.old > 0 {
			p.count = uint64(num) - p.old

			if p.count <= 0 {
				p.count = 1
			}
		}

		if old != p.count || p.old == 0 {
			p.sub.NotifyAll(p.count)
		}

		p.old = uint64(num)
	}
}

// Count 数量统计.
func (p *InstanceNum) Count() int {
	return int(p.count)
}
