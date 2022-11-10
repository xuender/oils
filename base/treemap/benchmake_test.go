package treemap_test

import (
	"testing"

	"github.com/xuender/oils/base/treemap"
)

func BenchmarkTreeMap(b *testing.B) {
	obj := treemap.New(-1, -1)

	for i := 0; i < 10_000_000; i++ {
		obj.Set(i, i)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		obj.Get(n)
	}
}

func BenchmarkMap(b *testing.B) {
	obj := map[int]int{}

	for i := 0; i < 10_000_000; i++ {
		obj[i] = i
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = obj[n]
	}
}
