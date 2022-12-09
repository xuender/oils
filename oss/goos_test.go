package oss_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIsLinux(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "linux" {
		assert.True(t, oss.IsLinux())
	} else {
		assert.False(t, oss.IsLinux())
	}
}

func TestIsDarwin(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "darwin" {
		assert.True(t, oss.IsDarwin())
	} else {
		assert.False(t, oss.IsDarwin())
	}
}
