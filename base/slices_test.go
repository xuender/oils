package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	slices := []int{1, 2, 3}
	filter := base.Filter(slices, func(elem int) bool {
		return elem > 1
	})

	assert.Equals(t, []int{2, 3}, filter)
}

func TestHas(t *testing.T) {
	t.Parallel()

	slices := []int{1, 2, 3}

	assert.True(t, base.Has(slices, 2))
	assert.False(t, base.Has(slices, 4))
}

func TestAppend(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []int{2, 3, 4}, base.Append(3, []int{1, 2}, 3, 4))
	assert.Equals(t, []int{4}, base.Append(1, []int{1, 2}, 3, 4))
	assert.Equals(t, []int{3, 4}, base.Append(2, []int{0, 1, 2}, 3, 4))
	assert.Equals(t, []int{}, base.Append(0, []int{1, 2}, 3, 4))
	assert.Equals(t, []int{2}, base.Append(1, []int{1, 2}))
	assert.Equals(t, []int{1, 2, 3, 4}, base.Append(-1, []int{1, 2}, 3, 4))
}

func FuzzAppend(f *testing.F) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testcases := []int{-1, 0, 1, 3, 4, 5}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		t.Helper()

		rev := base.Append(orig, slice, slice...)

		if orig < 0 {
			assert.Equal(t, 20, len(rev))
			assert.Equal(t, 0, rev[0])
			assert.Equal(t, 9, rev[19])
		}

		if orig == 0 {
			assert.Equal(t, 0, len(rev))
		}

		if orig > 0 {
			if orig >= 20 {
				assert.Equal(t, 20, len(rev))
				assert.Equal(t, 0, rev[0])
				assert.Equal(t, 9, rev[19])
			} else {
				assert.Equal(t, orig, len(rev))
			}

			assert.Equal(t, 9, rev[len(rev)-1])
		}
	})
}

func TestCount(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 1}
	assert.Equal(t, 2, base.Count(slice, 1))
}

func TestIndex(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 1}
	assert.Equal(t, 1, base.Index(slice, []int{2, 1}))
}
