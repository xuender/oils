package maps_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/maps"
)

func TestGroup(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, maps.Group(1, 10))
	assert.Equal(t, 1, maps.Group(-1, 10))
	assert.Equal(t, 1, maps.Group(11, 10))

	assert.Equal(t, 1, maps.Group(1.0, 10))
	assert.Equal(t, 1, maps.Group(-1.0, 10))
	assert.Equal(t, 1, maps.Group(11.0, 10))

	assert.Equal(t, 1, maps.Group("aaaa", 10))
}
