package syncs_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	que := syncs.NewQueue(2, func(i int) {})

	assert.Equal(t, uint(2), que.Size())
	assert.Nil(t, que.Add(1))
	assert.Nil(t, que.Add(2))
	assert.NotNil(t, que.Add(3))
}

func TestQueueSize2(t *testing.T) {
	t.Parallel()

	que := syncs.NewQueue(0, func(i int) {})

	assert.Equal(t, uint(1), que.Size())
	que.SetSize(0)
	assert.Equal(t, uint(1), que.Size())
}

func TestQueueConsum(t *testing.T) {
	t.Parallel()

	que := syncs.NewQueue(2, func(i int) {})

	go que.Consume()

	time.Sleep(time.Millisecond)
	assert.Nil(t, que.Add(1))
	assert.Nil(t, que.Add(2))
	assert.Nil(t, que.Add(3))
}

func TestQueuePush(t *testing.T) {
	t.Parallel()

	que := syncs.NewQueue(2, func(i int) { time.Sleep(time.Millisecond) })

	go que.Consume()

	time.Sleep(time.Millisecond)
	assert.Nil(t, que.Add(1))
	assert.Nil(t, que.Add(2))
	assert.Nil(t, que.Add(3))
	assert.NotNil(t, que.Add(4))
}

func TestQueueSize(t *testing.T) {
	t.Parallel()

	set := 0
	que := syncs.NewQueue(2, func(i int) {
		func() {
			time.Sleep(time.Millisecond * 5)
			set += i
		}()
	})

	go que.Consume()

	time.Sleep(time.Millisecond)
	assert.Nil(t, que.Add(1))
	assert.Nil(t, que.Add(2))
	assert.Nil(t, que.Add(3))
	que.SetSize(10)

	for i := 10; i < 20; i++ {
		assert.Nil(t, que.Add(i))
	}

	time.Sleep(time.Millisecond * 5 * 13)
	assert.Equal(t, set, 151)
}

func TestQueueLen(t *testing.T) {
	t.Parallel()

	sum := 0
	que := syncs.NewQueue(2, func(i int) { sum += i })

	assert.Nil(t, que.Add(1))
	assert.Nil(t, que.Add(2))
	assert.NotNil(t, que.Add(3))

	go que.Consume()

	time.Sleep(time.Millisecond * 10)
	assert.Equal(t, 3, sum)
}
