package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestSlice_String(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("a", "c", "b")

	assert.Equal(t, "Slice[a c b]", s.String())
	assert.Equal(t, "Slice[]", base.NewSlice[string]().String())
	assert.Equal(t, "Slice[1]", base.NewSlice(1).String())
}

func TestSlice_Clean(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 4)
	assert.Equal(t, 3, len(slice))
	slice.Clean()
	assert.Equal(t, 0, len(slice))
}

func TestSlice_Add(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 4)
	assert.Equal(t, 3, len(slice))
	slice.Add(4)
	assert.Equal(t, 4, len(slice))
}

func TestSlice_Push(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	slice.Push(5, 5, 5)
	assert.EqualValues(t, []int{5, 5, 5, 1, 2, 3}, slice)
}
