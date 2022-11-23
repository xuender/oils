package syncs

import (
	"time"
)

// LimitQueue 限频队列.
type LimitQueue[T any] struct {
	interval time.Duration
	last     time.Time
	queue    *Queue[T]
	call     func(T)
	timeOut  time.Duration
	qps      uint
}

// NewLimitQueue 新建限频队列.
func NewLimitQueue[T any](qps uint, timeOut time.Duration, call func(T)) *LimitQueue[T] {
	res := &LimitQueue[T]{
		call:     call,
		timeOut:  timeOut,
		qps:      qps,
		interval: time.Second / time.Duration(qps),
		last:     time.Now(),
	}

	res.queue = NewQueue(qps*uint(timeOut/time.Second), res.consume)

	go res.queue.Consume()

	return res
}

func (p *LimitQueue[T]) Len() int {
	return len(p.queue.elems)
}

func (p *LimitQueue[T]) consume(elem T) {
	dur := time.Since(p.last)

	if sleep := p.interval - dur; sleep > 0 {
		time.Sleep(sleep)

		p.last = p.last.Add(p.interval)
	} else {
		p.last = p.last.Add(dur)
	}

	p.call(elem)
}

func (p *LimitQueue[T]) Add(elem T) error {
	return p.queue.Add(elem)
}

// SetQPS 修改QPS.
func (p *LimitQueue[T]) SetQPS(qps uint64) {
	p.interval = time.Second / time.Duration(qps)
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
