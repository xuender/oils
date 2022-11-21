package syncs

import (
	"time"
)

// Queue 队列.
type Queue[T any] struct {
	elems chan T
	size  uint
	call  func(T)
}

// NewQueue 新建队列.
func NewQueue[T any](size uint, call func(T)) *Queue[T] {
	if size < 1 {
		size = 1
	}

	return &Queue[T]{
		elems: make(chan T, size),
		size:  size,
		call:  call,
	}
}

// Consume 队列消费.
func (p *Queue[T]) Consume() {
	elems := p.elems
	for elem := range elems {
		p.call(elem)
	}
}

// Size 队列尺寸.
func (p *Queue[T]) Size() uint {
	return p.size
}

// SetSize 修改队列尺寸.
func (p *Queue[T]) SetSize(size uint) {
	if size < 1 {
		size = 1
	}

	if p.size != size {
		p.size = size

		close(p.elems)

		p.elems = make(chan T, size)
		go p.Consume()
	}
}

// Add 队列增加.
func (p *Queue[T]) Add(elem T) error {
	select {
	case p.elems <- elem:
		return nil
	case <-time.After(time.Millisecond):
		return ErrFull
	}
}
