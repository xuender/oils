package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestMessage(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", assert.Message())
	assert.Equal(t, "1", assert.Message(1))
	assert.Equal(t, "1(2)", assert.Message(1, 2))
	assert.Equal(t, "1(2, 3)", assert.Message(1, 2, 3))
	assert.Equal(t, "1(2, 3, \"x\")", assert.Message(1, 2, 3, "x"))
}

func TestFail(t *testing.T) {
	t.Parallel()

	assert.False(t, assert.Fail(&errorfer{}, "fail"))
	assert.False(t, assert.Fail(&errorfer{}, "fail", 1, 2, 3))
}
