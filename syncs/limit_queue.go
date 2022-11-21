package syncs

import (
	"time"

	"github.com/xuender/oils/oss"
)

// LimitQueue 限频队列.
type LimitQueue[T any] struct {
	limit   *Limit
	queue   *Queue[T]
	call    func(T)
	timeOut time.Duration
	id      int
}

// NewLimitQueue 新建限频队列.
func NewLimitQueue[T any](qps uint, timeOut time.Duration, call func(T)) *LimitQueue[T] {
	limit := NewLimit(qps)
	res := &LimitQueue[T]{
		limit:   limit,
		call:    call,
		timeOut: timeOut,
		id:      oss.UniqueID(),
	}

	res.queue = NewQueue(qps*uint(timeOut/time.Second), res.consume)

	go res.queue.Consume()

	return res
}

func (p *LimitQueue[T]) consume(elem T) {
	p.limit.Wait()
	p.call(elem)
}

func (p *LimitQueue[T]) Add(elem T) error {
	return p.queue.Add(elem)
}

// Update 修改QPS.
func (p *LimitQueue[T]) Update(qps uint64) {
	p.limit.QPS(uint(qps))
	p.queue.SetSize(uint(qps) * uint(p.timeOut/time.Second))
}

// ID 观察者主键.
func (p *LimitQueue[T]) ID() int {
	return p.id
}
