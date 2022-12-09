package syncs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs"
)

func TestNewCounter(t *testing.T) {
	t.Parallel()

	count := syncs.NewCounter[int]()
	for i := 0; i < 1000; i++ {
		count.Inc(1)
	}

	for i := 0; i < 2000; i++ {
		count.Inc(2)
	}

	assert.Equal(t, int64(1000), count.Get(1))
	assert.Equal(t, int64(2000), count.Get(2))
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
