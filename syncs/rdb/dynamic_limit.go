package rdb

import (
	"github.com/xuender/oils/oss"
	"github.com/xuender/oils/syncs"
)

// DynamicLimit 动态可变QPS限流.
type DynamicLimit struct {
	syncs.Limit
	id int
}

// NewDynamicLimit 新建动态限流.
func NewDynamicLimit(qps uint) *DynamicLimit {
	return &DynamicLimit{
		Limit: *syncs.NewLimit(qps),
		id:    oss.UniqueID(),
	}
}

// Update 修改QPS.
func (p *DynamicLimit) Update(qps uint64) {
	p.QPS(uint(qps))
}

// ID 观察者主键.
func (p *DynamicLimit) ID() int {
	return p.id
}
