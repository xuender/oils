package oss_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/oss"
)

func TestIsDir(t *testing.T) {
	t.Parallel()

	assert.True(t, oss.IsDir(".."))
	assert.True(t, oss.IsDir("~"))
	assert.False(t, oss.IsDir("dir.go"))
	assert.False(t, oss.IsDir("fdfdfd"))
}
