package dbs_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/dbs"
)

func ExampleNewID() {
	fmt.Println(dbs.NewID(3))
	id := dbs.ID()
	fmt.Println(id)
	fmt.Println(dbs.DecodeID(id))
}

func ExampleID() {
	fmt.Println(dbs.ID())
	fmt.Println(dbs.ID())
}

func TestNewID(t *testing.T) {
	t.Parallel()

	set := base.NewSet[uint64]()

	for i := 0; i < 100_000; i++ {
		set.Add(dbs.NewID(3))
	}

	assert.Equal(t, 100_000, len(set))
}

func TestID(t *testing.T) {
	t.Parallel()

	set := base.NewSet[uint64]()

	for i := 0; i < 100_000; i++ {
		set.Add(dbs.ID())
	}

	assert.Equal(t, 100_000, len(set))
}

func BenchmarkID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dbs.ID()
	}
}

func BenchmarkNewID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dbs.NewID(3)
	}
}

// func BenchmarkID2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		u64.ID2()
// 	}
// }

func TestDecodeID(t *testing.T) {
	t.Parallel()

	id := dbs.ID()
	ts, serial, machine := dbs.DecodeID(id)

	assert.Greater(t, ts, 100, "ts")
	assert.GreaterOrEqual(t, serial, 0, "serial")
	assert.GreaterOrEqual(t, machine, 0, "machine")
}
