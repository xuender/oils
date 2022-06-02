package base_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
	"golang.org/x/exp/slices"
)

func ExampleMap() {
	map1 := base.NewMap[int, int]()
	map1[1] = 2
	map1[3] = 5
	fmt.Println(len(map1))

	// Output:
	// 2
}

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

func TestMapKeys(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()

	map1[1] = 1
	map1[2] = 1
	map1[3] = 1

	assert.Equal(t, 3, len(map1.Keys()))
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

func TestMap_Clear(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	assert.Equal(t, 3, len(map1))
	map1.Clear()
	assert.Equal(t, 0, len(map1))
}

func TestMap_Clone(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	map2 := map1.Clone()
	keys1 := map1.Keys()
	keys2 := map2.Keys()

	slices.Sort(keys1)
	slices.Sort(keys2)
	assert.Equals(t, keys1, keys2)
}

func TestMap_Copy(t *testing.T) {
	t.Parallel()

	map1 := base.NewMapSameValue(true, 1, 2, 3)
	map2 := base.NewMapSameValue(false, 4, 5)

	map1.Copy(map2)
	assert.Equal(t, 5, len(map2))
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

func TestMap_Keys(t *testing.T) {
	t.Parallel()

	keys := base.NewMapSameValue(true, 1, 2, 3).Keys()
	slices.Sort(keys)
	assert.Equals(t, []int{1, 2, 3}, keys)
}

func TestMap_Values(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []bool{true, true, true}, base.NewMapSameValue(true, 1, 2, 3).Values())
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
