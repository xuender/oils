package syncs_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func TestNewRoutineGroup(t *testing.T) {
	t.Parallel()

	count := 0
	group := syncs.NewRoutineGroup(3)

	for i := 0; i < 10; i++ {
		group.Add(1)

		go func() {
			defer group.Done()
			count++
		}()
	}

	group.Wait()

	assert.Equal(t, count, 10)
}

func TestRoutineGroupInc(t *testing.T) {
	t.Parallel()

	count := 0
	group := syncs.NewRoutineGroup(3)

	for i := 0; i < 10; i++ {
		group.Inc()

		go func() {
			defer group.Done()
			count++
		}()
	}

	group.Wait()

	assert.Equal(t, count, 10)
}

func TestNewRoutineGroup_panic(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, syncs.ErrSizeLessZero.Error(), func() {
		syncs.NewRoutineGroup(0)
	})
}
