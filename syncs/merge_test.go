package syncs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	out := syncs.Merge(func(int1, int2 int) bool {
		return int1 < int2
	}, ch1, ch2, ch3)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				ch2 <- i
			}
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%3 == 0 {
				ch1 <- i
			}
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%5 == 0 {
				ch3 <- i
			}
		}
		close(ch3)
	}()

	res := []int{}
	for value := range out {
		res = append(res, value)
	}

	assert.Equal(t, 11, len(res))
	assert.Equal(t, 0, res[0])
	assert.Equal(t, 9, res[10])
}
