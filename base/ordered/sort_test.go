package ordered_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/ordered"
)

func TestSort(t *testing.T) {
	t.Parallel()

	slice := []int{2, 3, 1}
	ordered.Sort(slice)
	assert.Equals(t, []int{1, 2, 3}, slice)
}
