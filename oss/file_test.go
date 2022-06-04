package oss_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/oss"
)

func TestExist(t *testing.T) {
	t.Parallel()

	assert.True(t, oss.Exist("doc.go"))
	assert.False(t, oss.Exist("unknown"))
}

func TestIsFile(t *testing.T) {
	t.Parallel()

	assert.False(t, oss.IsFile(".."))
	assert.False(t, oss.IsFile("~"))
	assert.True(t, oss.IsFile("file.go"))
	assert.False(t, oss.IsFile("fdfdfd"))
}
