package syncs

import (
	"time"
)

// LimitQueue 限频队列.
type LimitQueue[T any] struct {
	limit   *Limit
	queue   *Queue[T]
	call    func(T)
	timeOut time.Duration
	qps     uint
}

// NewLimitQueue 新建限频队列.
func NewLimitQueue[T any](qps uint, timeOut time.Duration, call func(T)) *LimitQueue[T] {
	limit := NewLimit(qps)
	res := &LimitQueue[T]{
		limit:   limit,
		call:    call,
		timeOut: timeOut,
		qps:     qps,
	}

	res.queue = NewQueue(qps*uint(timeOut/time.Second), res.consume)

	go res.queue.Consume()

	return res
}

func (p *LimitQueue[T]) consume(elem T) {
	p.call(elem)
	p.limit.Wait()
}

func (p *LimitQueue[T]) Add(elem T) error {
	return p.queue.Add(elem)
}

// SetQPS 修改QPS.
func (p *LimitQueue[T]) SetQPS(qps uint64) {
	p.limit.QPS(uint(qps))
	p.queue.SetSize(uint(qps) * uint(p.timeOut/time.Second))
}

// QPS 当前.
func (p *LimitQueue[T]) QPS() uint {
	return p.qps
}

func (p *LimitQueue[T]) SetTimeOut(timeOut time.Duration) {
	p.timeOut = timeOut
	p.queue.SetSize(p.qps * uint(p.timeOut/time.Second))
}

func (p *LimitQueue[T]) TimeOut() time.Duration {
	return p.timeOut
}
