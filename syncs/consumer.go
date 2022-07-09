package syncs

import (
	"context"
	"sync"
	"time"
)

type Consumer[E any] struct {
	cache chan E
	max   int
	min   time.Duration
}

type ProducerAndConsumer[E any] interface {
	Produc(produceChan chan E)
	Consum(num int, elems []E)
}

// NewConsumer 新建消费器.
// cacheSize 缓存数量.
// sliceMax 每次消费最大数量.
// minTime 预计的最小消费时间.
func NewConsumer[E any](cacheSize, sliceMax int, minTime time.Duration) *Consumer[E] {
	return &Consumer[E]{
		cache: make(chan E, cacheSize),
		max:   sliceMax,
		min:   minTime,
	}
}

func (p *Consumer[E]) ConsumeFunc(
	routine int,
	producFunc func(produceChan chan E),
	consumFunc func(num int, elems []E),
) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := &sync.WaitGroup{}

	for index := 0; index < routine; index++ {
		waitGroup.Add(1)

		go func(num int) {
			p.consume(ctx, num, consumFunc)
			waitGroup.Done()
		}(index)
	}

	producFunc(p.cache)
	time.Sleep(p.min)
	cancel()
	waitGroup.Wait()
}

// Consume 消费.
// routine 消费协程数量.
// producer 生产者.
// consumer 消费者.
func (p *Consumer[E]) Consume(routine int, pac ProducerAndConsumer[E]) {
	p.ConsumeFunc(routine, pac.Produc, pac.Consum)
}

func (p *Consumer[E]) consume(ctx context.Context, num int, consumFunc func(num int, elems []E)) {
	datas := make([]E, p.max)
	index := 0

	for {
		select {
		case <-ctx.Done():
			if index > 0 {
				consumFunc(num, datas[:index])
			}

			return
		case data := <-p.cache:
			datas[index] = data
			index++

			if index >= p.max {
				index = 0

				consumFunc(num, datas)
			}
		}
	}
}
