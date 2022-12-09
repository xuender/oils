package maps_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs/maps"
)

func TestMaps_Prefix(t *testing.T) {
	t.Parallel()

	tmap := maps.NewString(-1)

	for i := 0; i < 100; i++ {
		tmap.Set(fmt.Sprintf("a%d", i), i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set(fmt.Sprintf("b%d", i), i)
	}

	list := []int{}

	tmap.Prefix(func(key string, value int) bool {
		list = append(list, value)

		return true
	}, "b")

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], i)
	}

	assert.Equal(t, 10, len(list))
}

func TestMaps_PrefixDesc(t *testing.T) {
	t.Parallel()

	tmap := maps.NewString(-1)

	for i := 0; i < 100; i++ {
		tmap.Set(fmt.Sprintf("a%d", i), i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set(fmt.Sprintf("b%d", i), i)
	}

	list := []int{}

	tmap.PrefixDesc(func(key string, value int) bool {
		list = append(list, value)

		return true
	}, "b")

	for i := 0; i < 10; i++ {
		assert.Equal(t, list[i], 9-i)
	}

	assert.Equal(t, 10, len(list))
}
