package base_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
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

func TestMapSting(t *testing.T) {
	t.Parallel()

	map1 := base.NewMap[int, int]()
	map1[1] = 1
	map1[2] = 1

	assert.True(t, strings.Contains(map1.String(), "Map["))
}
