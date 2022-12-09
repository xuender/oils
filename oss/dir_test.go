package oss_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Parallel()

	path := "/tmp/a/b/c/d"

	os.Remove(path)
	os.Remove(filepath.Dir(path))
	oss.MkdirParent(path)
	assert.True(t, oss.Exist("/tmp/a/b/c"))
	os.Remove(path)
}
