package treemap_test

import (
	"testing"

	"github.com/xuender/oils/base"
	"github.com/xuender/oils/base/treemap"
)

func BenchmarkTreeMap(b *testing.B) {
	obj := treemap.New(0)

	for i := 0; i < 1_000_000; i++ {
		key := base.Number2Bytes(i)
		obj.Set(key, i)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := base.Number2Bytes(n)
		obj.Get(key)
	}
}

func BenchmarkMap(b *testing.B) {
	obj := map[string]int{}

	for i := 0; i < 1_000_000; i++ {
		obj[base.Itoa(i)] = i
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		// base.Number2Bytes(n)
		_ = obj[base.Itoa(n)]
	}
}
