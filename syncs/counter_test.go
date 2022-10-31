package syncs_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func ExampleCounter() {
	counter := syncs.NewCounter[string]()

	counter.Inc("test1")
	counter.Inc("test2")
	counter.Inc("test1")
	counter.Add("test3", 5)
	counter.Dec("test3")

	fmt.Println(counter.Get("test1"))
	fmt.Println(counter.Get("test2"))
	fmt.Println(counter.Get("test3"))
	counter.Clean()
	fmt.Println(counter.Get("test1"))

	// Output:
	// 2
	// 1
	// 4
	// 0
}

func TestNewCounter(t *testing.T) {
	t.Parallel()

	count := syncs.NewCounter[int]()
	for i := 0; i < 1000; i++ {
		count.Inc(1)
	}
	for i := 0; i < 2000; i++ {
		count.Inc(2)
	}

	assert.Equal(t, 1000, count.Get(1))
	assert.Equal(t, 2000, count.Get(2))
}

func BenchmarkCounter(b *testing.B) {
	count := syncs.NewCounter[int]()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		count.Inc(1)
	}
}

// func BenchmarkCounterb(b *testing.B) {
// 	count := syncs.NewCounterb[int]()

// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
//		count.Inc(1)
// 	}
// }
