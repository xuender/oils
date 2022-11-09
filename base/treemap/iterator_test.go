package treemap_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/treemap"
)

func TestTreeMap_Each(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.Each(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	})

	for i := 0; i < 100; i++ {
		assert.Equal(t, list[i], i)
	}

	assert.Equal(t, 100, len(list))

	list = []int{}

	tmap.Each(func(b []byte, i int) bool {
		list = append(list, i)

		return i < 10
	})

	assert.Equal(t, 11, len(list))
}

func TestTreeMap_Range(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.Range(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{byte(20)}, []byte{byte(30)})

	for i := 20; i < 30; i++ {
		assert.Equal(t, list[i-20], i)
	}

	assert.Equal(t, 10, len(list))

	list = []int{}

	tmap.Range(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{byte(70)}, []byte{byte(80)})

	for i := 70; i < 80; i++ {
		assert.Equal(t, list[i-70], i)
	}
}

func TestTreeMap_GreateOrEqual(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.GreateOrEqual(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{byte(20)})

	for i := 20; i < 100; i++ {
		assert.Equal(t, list[i-20], i)
	}

	assert.Equal(t, 80, len(list))
}

func TestTreeMap_LessThan(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.LessThan(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{byte(20)})

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], i)
	}

	assert.Equal(t, 20, len(list))
}

func TestTreeMap_Prefix(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{'a', byte(i)}, i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set([]byte{'b', byte(i)}, i*3)
	}

	list := []int{}

	tmap.Prefix(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{'b'})

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], i*3)
	}

	assert.Equal(t, 10, len(list))
}

func TestTreeMap_Iterator_EachDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)
	num := 100

	for i := 0; i < num; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.EachDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	})

	for i := 0; i < num; i++ {
		assert.Equal(t, list[i], 99-i)
	}

	list = []int{}

	tmap.EachDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return i > 10
	})

	assert.Equal(t, 90, len(list))
}

func TestTreeMap_Iterator_EachDesc2(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)
	num := 3

	for i := 0; i < num; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.EachDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	})

	for i := 0; i < num; i++ {
		assert.Equal(t, list[i], 2-i)
	}
}

func TestTreeMap_RangeDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.RangeDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{20}, []byte{30})

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], 29-i)
	}
	assert.Equal(t, 10, len(list))

	list = []int{}

	tmap.RangeDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{70}, []byte{90})

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 89-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestTreeMap_GreateOrEqualDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.GreateOrEqualDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{80})

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 99-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestTreeMap_LessThanDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 30; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	list := []int{}

	tmap.LessThanDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{20})

	for i := 0; i < 20; i++ {
		assert.Equal(t, list[i], 19-i)
	}

	assert.Equal(t, 20, len(list))
}

func TestTreeMap_PrefixDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.New(-1)

	for i := 0; i < 100; i++ {
		tmap.Set([]byte{'a', byte(i)}, i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set([]byte{'b', byte(i)}, i)
	}

	list := []int{}

	tmap.PrefixDesc(func(b []byte, i int) bool {
		list = append(list, i)

		return true
	}, []byte{'b'})

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], 9-i)
	}

	assert.Equal(t, 10, len(list))
}
