package oss_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/oss"
)

// nolint: paralleltest
func TestExist(t *testing.T) {
	assert.True(t, oss.Exist("doc.go"))
	assert.False(t, oss.Exist("unknown"))

	patches := gomonkey.ApplyFuncReturn(filepath.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	assert.False(t, oss.Exist(""))
}

func TestIsFile(t *testing.T) {
	t.Parallel()

	assert.False(t, oss.IsFile(".."))
	assert.False(t, oss.IsFile("~"))
	assert.True(t, oss.IsFile("file.go"))
	assert.False(t, oss.IsFile("fdfdfd"))
}
