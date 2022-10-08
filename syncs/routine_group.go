package syncs

import "sync"

// RoutineGroup 协程组，是sync.WaitGroup的增强版本.
type RoutineGroup struct {
	ch chan struct{}
	wg sync.WaitGroup
}

// NewRoutineGroup 协程组，控制协程总数量.
func NewRoutineGroup(size uint) *RoutineGroup {
	if size < 1 {
		panic(ErrSizeLessZero)
	}

	return &RoutineGroup{
		ch: make(chan struct{}, size),
		wg: sync.WaitGroup{},
	}
}

// Add 增加协程.
func (p *RoutineGroup) Add(delta uint) {
	for i := uint(0); i < delta; i++ {
		p.ch <- struct{}{}
	}

	p.wg.Add(int(delta))
}

// Done 协程完成.
func (p *RoutineGroup) Done() {
	<-p.ch
	p.wg.Done()
}

// Wait 等待全部完成.
func (p *RoutineGroup) Wait() {
	p.wg.Wait()
}
