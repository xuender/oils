package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.True(&errorfer{}, true))
	assert.False(t, assert.True(&errorfer{}, false))
}

func TestFalse(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.False(&errorfer{}, false))
	assert.False(t, assert.False(&errorfer{}, true))
}
