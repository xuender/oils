package treemap_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/base/treemap"
	"github.com/xuender/oils/dbs"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(0)

	for i := 0; i < 10000; i++ {
		tmap.Set([]byte(base.Itoa(i)), i)
	}

	assert.Equal(t, 10000, tmap.Len())
	value, _ := tmap.Get([]byte(base.Itoa(1000)))
	assert.Equal(t, 1000, value)
	tmap.Clear()
	assert.Equal(t, 0, tmap.Len())
}

func TestTreeMap_Set(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(0)

	tmap.Set(base.Number2Bytes(7), 7)

	for i := 0; i < 10000; i++ {
		tmap.Set(base.Number2Bytes(i), i)
	}

	for i := 0; i < 10000; i++ {
		num, has := tmap.Get(base.Number2Bytes(i))

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	_, has := tmap.Get(base.Number2Bytes(-3))
	assert.False(t, has)
}

func TestTreeMap_Del(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(0)

	tmap.Set(base.Number2Bytes(7), 7)

	for i := 0; i < 10000; i++ {
		tmap.Set(base.Number2Bytes(i), i)
	}

	for i := 500; i < 2000; i++ {
		tmap.Del(base.Number2Bytes(i))
	}

	tmap.Del(base.Number2Bytes(0))

	for i := 1; i < 500; i++ {
		num, has := tmap.Get(base.Number2Bytes(i))

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	for i := 2000; i < 10000; i++ {
		num, has := tmap.Get(base.Number2Bytes(i))

		assert.True(t, has)
		assert.Equal(t, i, num)
	}

	for i := 500; i < 2000; i++ {
		_, has := tmap.Get(base.Number2Bytes(i))
		assert.False(t, has)
	}

	_, has := tmap.Get(base.Number2Bytes(0))
	assert.False(t, has)
}

func TestTreeMap_Del1(t *testing.T) {
	t.Parallel()

	tmap := treemap.New[uint64](0)

	for i := 0; i < 100000; i++ {
		if dbs.Rand()%7 > 0 {
			key := dbs.Rand() / 100
			tmap.Del(base.Number2Bytes(key))
		} else {
			key := dbs.Rand() / 100
			tmap.Set(base.Number2Bytes(key), key)
		}
	}
}

func TestTreeMap_DelMin(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	min, _ := tmap.Min()
	assert.Equal(t, -1, min)
	assert.Equal(t, 0, tmap.Len())
	assert.False(t, tmap.DelMin())
	assert.Equal(t, 0, tmap.Len())

	tmap.Add(base.Number2Bytes(80), 80)
	assert.Equal(t, 1, tmap.Len())
	tmap.Add(base.Number2Bytes(80), 80)
	assert.Equal(t, 1, tmap.Len())

	for i := 0; i < 100; i++ {
		tmap.Add(base.Number2Bytes(i), i)
	}

	assert.Equal(t, 100, tmap.Len())

	assert.True(t, tmap.DelMin())
	assert.True(t, tmap.DelMin())

	assert.Equal(t, 98, tmap.Len())
	min, _ = tmap.Min()
	assert.Equal(t, 2, min)
}

func TestTreeMap_DelMax(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	max, _ := tmap.Max()
	assert.Equal(t, -1, max)
	assert.False(t, tmap.DelMax())

	for i := 0; i < 100; i++ {
		tmap.Add(base.Number2Bytes(i), i)
	}

	assert.True(t, tmap.DelMax())
	assert.True(t, tmap.DelMax())

	max, _ = tmap.Max()
	assert.Equal(t, 97, max)
}
