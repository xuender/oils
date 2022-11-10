package maps_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/maps"
)

func TestMaps_Each(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 1_000; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.Each(func(key, value int) bool {
		list = append(list, value)

		return true
	})

	for i := 0; i < 1_000; i++ {
		assert.Equal(t, list[i], i)
	}

	assert.Equal(t, 1_000, len(list))

	list = []int{}

	tmap.Each(func(key, value int) bool {
		list = append(list, value)

		return value < 10
	})

	assert.Equal(t, 11, len(list))
}

func TestMaps_Range(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 100; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.Range(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 20, 30)

	for i := 20; i < 30; i++ {
		assert.Equal(t, list[i-20], i)
	}

	assert.Equal(t, 10, len(list))

	list = []int{}

	tmap.Range(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 70, 80)

	for i := 70; i < 80; i++ {
		assert.Equal(t, list[i-70], i)
	}
}

func TestMaps_GreateOrEqual(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 100; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.GreateOrEqual(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 20)

	for i := 20; i < 100; i++ {
		assert.Equal(t, list[i-20], i)
	}

	assert.Equal(t, 80, len(list))
}

func TestMaps_LessThan(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 100; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.LessThan(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 20)

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], i)
	}

	assert.Equal(t, 20, len(list))
}

func TestMaps_Iterator_EachDesc(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)
	num := 100

	for i := 0; i < num; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.EachDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	})

	for i := 0; i < num; i++ {
		assert.Equal(t, list[i], 99-i)
	}

	list = []int{}

	tmap.EachDesc(func(key, value int) bool {
		list = append(list, value)

		return value > 10
	})

	assert.Equal(t, 90, len(list))
}

func TestMaps_Iterator_EachDesc2(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)
	num := 3

	for i := 0; i < num; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.EachDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	})

	for i := 0; i < num; i++ {
		assert.Equal(t, list[i], 2-i)
	}
}

func TestMaps_RangeDesc(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 100; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.RangeDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 20, 30)

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], 29-i)
	}
	assert.Equal(t, 10, len(list))

	list = []int{}

	tmap.RangeDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 70, 90)

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 89-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestMaps_GreateOrEqualDesc(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 100; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.GreateOrEqualDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 80)

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 99-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestMaps_LessThanDesc(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-1, -1)

	for i := 0; i < 30; i++ {
		tmap.Set(i, i)
	}

	list := []int{}

	tmap.LessThanDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 20)

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 19-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestMaps_Each_float32(t *testing.T) {
	t.Parallel()

	tmap := maps.New(-0.1, -1)

	tmap.Set(1.0, 1)
	tmap.Set(1.1, 2)
	tmap.Set(1.1111, 3)
	tmap.Set(1.4, 4)
	tmap.Set(2.0, 5)

	list := []int{}

	tmap.Each(func(key float64, value int) bool {
		list = append(list, value)

		return true
	})

	assert.Equals(t, []int{1, 2, 3, 4, 5}, list)
}

func TestTreeMap_Each_sync(t *testing.T) {
	t.Parallel()

	num := -1
	tmap := maps.New(num, num)

	for i := 0; i < 10_000; i++ {
		tmap.Set(i, i)
	}

	group := sync.WaitGroup{}

	group.Add(10)
	rand.Seed(time.Now().UnixMilli())

	for f := 0; f < 10; f++ {
		go func() {
			for i := 0; i < 1_000; i++ {
				// nolint
				tmap.Del(rand.Intn(10_000))
			}

			group.Done()
		}()
	}

	group.Wait()
	tmap.Each(func(key, value int) bool {
		assert.Greater(t, key, num)
		num = key

		return true
	})
}
