package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestSorts(t *testing.T) {
	t.Parallel()

	keys := []int{3, 1, 2}
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"1", "2", "3"}

	base.Sorts(keys, slice1, slice2)

	assert.Equal(t, []int{1, 2, 3}, keys)
	assert.Equal(t, []string{"b", "c", "a"}, slice1)
	assert.Equal(t, []string{"2", "3", "1"}, slice2)
}

func TestSorts_error(t *testing.T) {
	t.Parallel()

	keys := []int{3, 1, 2}
	slice1 := []string{"a", "b"}
	slice2 := []string{"1", "2"}

	base.Sorts(keys, slice1, slice2)

	assert.Equal(t, []int{1, 2, 3}, keys)
	assert.Equal(t, []string{"b", "a"}, slice1)
	assert.Equal(t, []string{"2", "1"}, slice2)
}
