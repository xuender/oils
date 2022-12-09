package oss_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/oss"
)

func TestAbs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "parent", filepath.Base(base.Must1(oss.Abs("parent/c/.."))))
	assert.NotEqual(t, "~", filepath.Base(base.Must1(oss.Abs("~"))))

	home, _ := os.UserHomeDir()
	path, _ := oss.Abs("~")
	assert.Equal(t, home, path)

	path, _ = oss.Abs("~/ff")
	assert.True(t, strings.HasPrefix(path, home))

	path, _ = oss.Abs("~ff")
	assert.NotEqual(t, len(home)+3, len(path))
}
