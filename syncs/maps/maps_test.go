package maps_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs"
	"github.com/xuender/oils/syncs/maps"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tmap := maps.New[int](0)

	for i := 0; i < 10000; i++ {
		tmap.Set(i, i)
	}

	assert.Equal(t, 10000, tmap.Len())
	value, _ := tmap.Get(1000)
	assert.Equal(t, 1000, value)
	tmap.Clear()
	assert.Equal(t, 0, tmap.Len())
}

func TestMaps_Set(t *testing.T) {
	t.Parallel()

	tmap := maps.New[int](0)

	tmap.Set(7, 7)

	for i := 0; i < 10000; i++ {
		tmap.Set(i, i)
	}

	for i := 0; i < 10000; i++ {
		num, has := tmap.Get(i)

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	_, has := tmap.Get(-3)
	assert.False(t, has)
}

func TestMaps_Del(t *testing.T) {
	t.Parallel()

	tmap := maps.New[int](0)

	tmap.Set(7, 7)

	for i := 0; i < 10000; i++ {
		tmap.Set(i, i)
	}

	for i := 500; i < 2000; i++ {
		tmap.Del(i)
	}

	tmap.Del(0)

	for i := 1; i < 500; i++ {
		num, has := tmap.Get(i)

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	for i := 2000; i < 10000; i++ {
		num, has := tmap.Get(i)

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	for i := 500; i < 2000; i++ {
		_, has := tmap.Get(i)
		assert.False(t, has)
	}

	_, has := tmap.Get(0)
	assert.False(t, has)
}

func TestMaps_Del1(t *testing.T) {
	t.Parallel()

	tmap := maps.New[uint64](0)

	for index := 0; index < 100000; index++ {
		if dbs.Rand()%7 > 0 {
			key := dbs.Rand() / 100
			tmap.Del(key)
		} else {
			key := dbs.Rand() / 100
			tmap.Set(key, index)
		}
	}
}

func TestMaps_DelMin(t *testing.T) {
	t.Parallel()

	tmap := maps.New[int](-1)

	min, _ := tmap.Min()
	assert.Equal(t, -1, min)
	assert.Equal(t, 0, tmap.Len())
	assert.False(t, tmap.DelMin())
	assert.Equal(t, 0, tmap.Len())

	tmap.Add(80, 80)
	assert.Equal(t, 1, tmap.Len())
	tmap.Add(80, 80)
	assert.Equal(t, 1, tmap.Len())

	for i := 0; i < 100; i++ {
		tmap.Add(i, i)
	}

	assert.Equal(t, 100, tmap.Len())

	assert.True(t, tmap.DelMin())
	assert.True(t, tmap.DelMin())

	assert.Equal(t, 98, tmap.Len())
	min, _ = tmap.Min()
	assert.Equal(t, 2, min)
}

func TestMaps_DelMax(t *testing.T) {
	t.Parallel()

	tmap := maps.New[int](-1)

	max, _ := tmap.Max()
	assert.Equal(t, -1, max)
	assert.False(t, tmap.DelMax())

	for i := 0; i < 100; i++ {
		tmap.Add(i, i)
	}

	assert.True(t, tmap.DelMax())
	assert.True(t, tmap.DelMax())

	max, _ = tmap.Max()
	assert.Equal(t, 97, max)
}
