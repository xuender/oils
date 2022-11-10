package maps_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/xuender/oils/syncs/maps"
)

func BenchmarkMaps(b *testing.B) {
	obj := maps.New("", -1)

	for i := 0; i < 1_000_000; i++ {
		obj.Set(strconv.Itoa(i), i)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		group := sync.WaitGroup{}

		for index := 0; index < 10; index++ {
			group.Add(1)

			go func(num int) {
				for i := 0; i < 10_000; i++ {
					obj.Get(strconv.Itoa(i * num))
				}

				group.Done()
			}(index)
		}

		group.Wait()
	}
}

func BenchmarkMap(b *testing.B) {
	obj := map[string]int{}

	for i := 0; i < 1_000_000; i++ {
		obj[strconv.Itoa(i)] = i
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		group := sync.WaitGroup{}

		for index := 0; index < 10; index++ {
			group.Add(1)

			go func(num int) {
				for i := 0; i < 10_000; i++ {
					_ = obj[strconv.Itoa(i*num)]
				}

				group.Done()
			}(index)
		}

		group.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	obj := sync.Map{}

	for i := 0; i < 1_000_000; i++ {
		obj.Store(strconv.Itoa(i), i)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		group := sync.WaitGroup{}

		for index := 0; index < 10; index++ {
			group.Add(1)

			go func(num int) {
				for i := 0; i < 10_000; i++ {
					obj.Load(strconv.Itoa(i * num))
				}

				group.Done()
			}(index)
		}

		group.Wait()
	}
}
