package syncs

import (
	"sync"
)

// Pool Goroutine 池.
type Pool[I, O any] struct {
	input chan *job[I, O]
	yield func(I, int) O
}

type job[I, O any] struct {
	wgp    *sync.WaitGroup
	input  I
	output O
	index  int
}

// NewPool 新建 Goroutine 池.
func NewPool[I, O any](size int, yield func(I, int) O) *Pool[I, O] {
	pool := &Pool[I, O]{
		input: make(chan *job[I, O], size),
		yield: yield,
	}

	for i := 0; i < size; i++ {
		go pool.run(i)
	}

	return pool
}

func (p *Pool[I, O]) run(num int) {
	for input := range p.input {
		input.output = p.yield(input.input, num)
		input.wgp.Done()
	}
}

func (p *Pool[I, O]) Post(inputs []I) []O {
	jobs := make([]*job[I, O], len(inputs))
	wgp := sync.WaitGroup{}

	wgp.Add(len(inputs))

	for index, input := range inputs {
		jobs[index] = &job[I, O]{
			wgp:   &wgp,
			input: input,
			index: index,
		}

		p.input <- jobs[index]
	}

	wgp.Wait()

	res := make([]O, len(inputs))

	for _, job := range jobs {
		res[job.index] = job.output
	}

	return res
}
