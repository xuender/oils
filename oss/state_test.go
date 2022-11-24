package oss_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/oss"
)

func TestStateFile(t *testing.T) {
	t.Parallel()

	assert.Equal(t, oss.Play, oss.StateByFile("/tmpfs/play"))
	assert.Equal(t, oss.Build, oss.StateByFile(filepath.Join("aa", "bb")))
	assert.Equal(t, oss.Run, oss.StateByFile(filepath.Join(os.TempDir(), "go-build123", "exe", "main")))
	assert.Equal(t, oss.Test, oss.StateByFile(filepath.Join(os.TempDir(), "go-build123", "main.test")))
}

func TestState(t *testing.T) {
	t.Parallel()

	assert.Equal(t, oss.Test, oss.State())
}
