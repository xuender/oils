package oss_test

import (
	"os"
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

func TestMkdirParent(t *testing.T) {
	path := "/tmp/a/b/c/d"

	os.Remove(path)
	oss.MkdirParent(path)
	assert.True(t, oss.Exist("/tmp/a/b/c"))
}
