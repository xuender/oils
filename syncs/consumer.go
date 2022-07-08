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

// Consume 消费.
// routine 消费协程数量.
// producer 生产者.
// consumer 消费者.
func (p *Consumer[E]) Consume(routine int, producer func(produceChan chan E), consumer func(num int, elems []E)) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := &sync.WaitGroup{}

	for index := 0; index < routine; index++ {
		waitGroup.Add(1)

		go func(num int) {
			p.consume(ctx, num, consumer)
			waitGroup.Done()
		}(index)
	}

	producer(p.cache)
	time.Sleep(p.min)
	cancel()
	waitGroup.Wait()
}

func (p *Consumer[E]) consume(ctx context.Context, num int, consumer func(int, []E)) {
	datas := make([]E, p.max)
	index := 0

	for {
		select {
		case <-ctx.Done():
			if index > 0 {
				consumer(num, datas[:index])
			}

			return
		case data := <-p.cache:
			datas[index] = data
			index++

			if index >= p.max {
				index = 0

				consumer(num, datas)
			}
		}
	}
}
