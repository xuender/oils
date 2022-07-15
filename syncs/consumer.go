package syncs

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// ConsumeFunc 消费.
// cacheSize 缓存数量.
// sliceMax 每次消费最大数量.
// routine 消费协程数量.
// minTime 预计的最小消费时间.
// producer 生产者.
// consumer 消费者.
// nolint
func ConsumeFunc[E any](cacheSize, sliceMax, routine int, minTime time.Duration,
	producFunc func(produceChan chan E),
	consumFunc func(num int, elems []E),
) int64 {
	var count int64

	cache := make(chan E, cacheSize)
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := &sync.WaitGroup{}

	for index := 0; index < routine; index++ {
		waitGroup.Add(1)

		go func(num int) {
			var index int64

			datas := make([]E, sliceMax)

			for {
				select {
				case <-ctx.Done():
					if index > 0 {
						consumFunc(num, datas[:index])
						atomic.AddInt64(&count, index)
					}

					waitGroup.Done()

					return
				case data := <-cache:
					datas[index] = data
					index++

					if int(index) >= sliceMax {
						consumFunc(num, datas)
						atomic.AddInt64(&count, index)
						index = 0
					}
				}
			}
		}(index)
	}

	producFunc(cache)
	time.Sleep(minTime)
	cancel()
	waitGroup.Wait()

	return count
}

type ProducerAndConsumer[E any] interface {
	Produc(produceChan chan E)
	Consum(num int, elems []E)
}

// Consume 消费.
// cacheSize 缓存数量.
// sliceMax 每次消费最大数量.
// routine 消费协程数量.
// minTime 预计的最小消费时间.
// producer 生产者.
// consumer 消费者.
func Consume[E any](cacheSize, sliceMax, routine int, minTime time.Duration,
	pac ProducerAndConsumer[E],
) int64 {
	return ConsumeFunc(cacheSize, sliceMax, routine, minTime, pac.Produc, pac.Consum)
}
