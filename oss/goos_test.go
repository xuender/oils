package oss_test

import (
	"runtime"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/oss"
)

func TestIsWindows(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		assert.True(t, oss.IsWindows())
	} else {
		assert.False(t, oss.IsWindows())
	}
}
