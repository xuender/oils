package base_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestSet_String(t *testing.T) {
	t.Parallel()

	set := base.NewSet(2, 3, 1)

	assert.True(t, strings.Contains(set.String(), "Set["))
}

func TestSet_Join(t *testing.T) {
	t.Parallel()

	set := base.NewSet(2, 3, 1)

	assert.True(t, strings.Contains(set.Join("-"), "1"))

	set.Del(1, 2)
	assert.Equal(t, "3", set.Join("-"))

	set.Del(3)
	assert.Equal(t, "", set.Join("-"))
}

func TestSet_All(t *testing.T) {
	t.Parallel()

	set := base.NewSet(2, 3, 1)

	assert.True(t, set.All(2, 1, 3))
	assert.False(t, set.All(4, 1, 3))
}

func TestSet_Has(t *testing.T) {
	t.Parallel()

	set := base.NewSet(2, 3, 1)

	assert.True(t, set.Has(2))
	assert.False(t, set.Has(4))
}

func TestSet_AddUint64s(t *testing.T) {
	t.Parallel()

	set := base.NewSet[int]()
	set.AddSet(base.NewSet(1, 2), base.NewSet(2, 3))

	assert.Equal(t, 3, len(set))
}

func TestSet_DelUint64s(t *testing.T) {
	t.Parallel()

	set := base.NewSet(1, 2, 3, 4)
	set.DelSet(base.NewSet(1, 2), base.NewSet(2, 3))

	assert.Equal(t, 1, len(set))
}

func TestSet_Add(t *testing.T) {
	t.Parallel()

	data := uint64(1)
	set := base.NewSet[uint64]()

	assert.False(t, set.Any(data))

	set.Add(data)
	assert.Equal(t, 1, len(set))
	assert.True(t, set.Any(data))

	set.Del(data)
	assert.Equal(t, 0, len(set))
	assert.False(t, set.Any(data))
}

func TestNewSet(t *testing.T) {
	t.Parallel()

	set := base.NewSet(1, 2, 2)
	assert.Equal(t, 2, len(set))
}
