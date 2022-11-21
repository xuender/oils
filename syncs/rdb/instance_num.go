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
	client  redis.Cmdable
	num     uint64
	old     uint64
	subject *modes.Subject[uint64, int]
	key     string
}

// NewInstanceNum 新建实例数量.
func NewInstanceNum(client redis.Cmdable) *InstanceNum {
	return NewInstanceNumByKey(client, "ins_num")
}

func NewInstanceNumByKey(client redis.Cmdable, key string) *InstanceNum {
	return &InstanceNum{
		client:  client,
		subject: &modes.Subject[uint64, int]{},
		key:     key,
		num:     1,
	}
}

// Register 订阅主题.
func (p *InstanceNum) Register(observer modes.Observer[uint64, int]) {
	p.subject.Register(observer)
}

// Run 运行.
//
// 使用 redis INCR 命令计算一秒内增加的数量得出实例数量.
func (p *InstanceNum) Run() {
	ctx := context.Background()
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		num, err := p.client.Incr(ctx, p.key).Result()
		if err != nil {
			p.client.Del(ctx, p.key)

			num = base.Must1(p.client.Incr(ctx, p.key).Result())
		}

		old := p.num

		if p.old > 0 {
			p.num = uint64(num) - p.old

			if p.num <= 0 {
				p.num = 1
			}
		}

		if old != p.num || p.old == 0 {
			p.subject.NotifyAll(p.num)
		}

		p.old = uint64(num)
	}
}

// Num 实例数量.
func (p *InstanceNum) Num() int {
	return int(p.num)
}
