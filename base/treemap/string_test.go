package treemap_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/treemap"
)

func TestTreeMap_Prefix(t *testing.T) {
	t.Parallel()

	tmap := treemap.NewString(-1)

	for i := 0; i < 100; i++ {
		tmap.Set(fmt.Sprintf("a%d", i), i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set(fmt.Sprintf("b%d", i), i*3)
	}

	list := []int{}

	tmap.Prefix(func(_ string, i int) bool {
		list = append(list, i)

		return true
	}, "b")

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], i*3)
	}

	assert.Equal(t, 10, len(list))
}

func TestTreeMap_PrefixDesc(t *testing.T) {
	t.Parallel()

	tmap := treemap.NewString(-1)

	for i := 0; i < 100; i++ {
		tmap.Set(fmt.Sprintf("a%d", i), i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set(fmt.Sprintf("b%d", i), i)
	}

	list := []int{}

	tmap.PrefixDesc(func(_ string, i int) bool {
		list = append(list, i)

		return true
	}, "b")

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], 9-i)
	}

	assert.Equal(t, 10, len(list))
}

func TestPrefixNext(t *testing.T) {
	t.Parallel()

	assert.True(t, "" < treemap.PrefixNext(""))
	assert.True(t, "a" < treemap.PrefixNext("a"))
	assert.True(t, "aa" < treemap.PrefixNext("a"))
	assert.True(t, "af" < treemap.PrefixNext("a"))
}
