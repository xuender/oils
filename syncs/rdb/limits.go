package rdb

import (
	"github.com/xuender/oils/syncs"
)

// Limits 批量限流.
type Limits struct {
	limits map[string]*DynamicLimit
	qps    map[string]uint64
	num    uint64
}

// NewLimits 新建批量限流.
func NewLimits() *Limits {
	return &Limits{
		limits: map[string]*DynamicLimit{},
		qps:    map[string]uint64{},
		num:    1,
	}
}

func (p *Limits) Update(num uint64) {
	p.num = num

	for key, limit := range p.limits {
		limit.Update(p.qps[key] / num)
	}
}

func (p *Limits) ID() int { return 1 }

// QPS 设置QPS.
func (p *Limits) QPS(key string, qps uint) syncs.Limiter {
	limit := NewDynamicLimit(qps / uint(p.num))

	p.limits[key] = limit
	p.qps[key] = uint64(qps)

	return limit
}

// Limit 限制流.
func (p *Limits) Limit(key string) syncs.Limiter {
	return p.limits[key]
}

// GetQPS 获取所有限速的QPS.
func (p *Limits) GetQPS() map[string]uint {
	res := map[string]uint{}

	for key, qps := range p.qps {
		res[key] = uint(qps / p.num)
	}

	return res
}
