package rdb_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs/rdb"
)

func TestLimits(t *testing.T) {
	t.Parallel()

	limits := rdb.NewLimits()

	limits.QPS("test", 100)
	assert.Equal(t, uint64(100), limits.Get("test"))

	limit := limits.Limit("test")
	start := time.Now()

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 50; i++ {
				limit.Wait()
			}
		}()
	}

	for i := 0; i < 200; i++ {
		limit.Wait()
	}

	assert.GreaterOrEqual(t, time.Since(start), time.Second*4)
}

func TestLimits_Update(t *testing.T) {
	t.Parallel()

	limits := rdb.NewLimits()

	limits.QPS("test", 100)
	assert.Equal(t, uint64(100), limits.Get("test"))

	limit := limits.Limit("test")
	start := time.Now()
	group := sync.WaitGroup{}

	group.Add(4)

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 50; i++ {
				limit.Wait()
			}
			group.Done()
		}()
	}

	group.Wait()

	assert.GreaterOrEqual(t, time.Since(start), time.Second*2)
	assert.Less(t, time.Since(start), time.Second*3)

	limits.Update(2)

	assert.Equal(t, uint(50), limits.GetAll()["test"])

	for i := 0; i < 100; i++ {
		limit.Wait()
	}

	assert.GreaterOrEqual(t, time.Since(start), time.Second*4)
	assert.Less(t, time.Since(start), time.Second*5)
}

func TestLimits_Keys(t *testing.T) {
	t.Parallel()

	limits := rdb.NewLimits()

	limits.QPS("a", 1000)
	assert.Equal(t, []string{"a"}, limits.Keys())
	assert.Equal(t, 1, limits.ID())
}
