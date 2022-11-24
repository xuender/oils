package base_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestNewMap(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[3] = 1

	assert.Equal(t, 1, len(map1))
}

func TestMapHas(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[3] = 1

	assert.True(t, map1.Has(3))
	assert.False(t, map1.Has(2))
}

func TestMapAny(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[1] = 1
	map1[2] = 1

	assert.True(t, map1.Any(1, 3))
	assert.False(t, map1.Any(3, 4))
}

func TestMapAll(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[1] = 1
	map1[2] = 1

	assert.True(t, map1.All(1, 2))
	assert.False(t, map1.All(1, 3))
}

func TestMapDel(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[1] = 1
	map1[2] = 1
	map1[3] = 1
	map1.Del(1, 2)

	assert.Equal(t, 1, len(map1))
	assert.True(t, map1.Has(3))
}

func TestMapPut(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()
	map1[1] = 1
	map1[2] = 1
	map1[3] = 1

	map2 := base.NewMap[int, int]()

	map2[2] = 2
	map2[3] = 2
	map2[4] = 2
	map1.Put(map2)

	assert.Equal(t, 4, len(map1))
	assert.Equal(t, 2, map1[2])
	assert.Equal(t, 1, map1[1])
}

func TestMapDelMap(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()
	map1[1] = 1
	map1[2] = 1
	map1[3] = 1

	map2 := base.NewMap[int, int]()

	map2[2] = 2
	map2[3] = 2
	map2[4] = 2
	map1.DelMap(map2)

	assert.Equal(t, 1, len(map1))
	assert.Equal(t, 1, map1[1])
}

func TestMap_String(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()
	map1[1] = 1
	map1[2] = 1

	assert.True(t, strings.Contains(map1.String(), "Map["))
}

func TestMap_Has(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	assert.True(t, map1.Has(1))
	assert.False(t, map1.Has(4))
}

func TestMap_Any(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	assert.True(t, map1.Any(1, 5, 4))
	assert.False(t, map1.Any(4, 4))
}

func TestMap_All(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	assert.True(t, map1.All(1, 3, 2))
	assert.False(t, map1.All(4, 1))
}

func TestMap_Del(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	map1.Del(2)
	assert.Equal(t, 2, len(map1))
}

func TestMap_DelMap(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	map2 := base.NewMapSameValue(false, 1, 5)
	map1.DelMap(map2)
	assert.Equal(t, 2, len(map1))
}

func TestMap_Put(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	map2 := base.NewMapSameValue(false, 4, 5)

	map1.Put(map2)
	assert.Equal(t, 5, len(map1))
}

func TestNewMapSameValue(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	assert.True(t, map1.Has(1))
	assert.False(t, map1.Has(4))
}

func TestMap_ValuesByKeys(t *testing.T) {
	t.Parallel()

	testMap := base.NewMap[int, int]()
	testMap[3] = 1
	testMap[4] = 6
	assert.Equal(t, []int{6, 1}, testMap.ValuesByKeys([]int{4, 3}))
	assert.Equal(t, []int{6, 0}, testMap.ValuesByKeys([]int{4, 2}))
}
