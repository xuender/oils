package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestNil(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Nil(&errorfer{}, nil))
	assert.False(t, assert.Nil(&errorfer{}, 1))
}

func TestNotNil(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.NotNil(&errorfer{}, 1))
	assert.False(t, assert.NotNil(&errorfer{}, nil))
}
