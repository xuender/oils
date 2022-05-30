package logs_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/logs"
)

func TestLogName(t *testing.T) {
	t.Parallel()

	switch runtime.GOOS {
	case "window":
		assert.Equal(t, filepath.Join(filepath.Base(os.Args[0]), "logs.log"), logs.LogName())
	default:
		assert.Equal(t, "/var/tmp/logs.log", logs.LogName())
	}
}

func TestRotateLog(t *testing.T) {
	t.Parallel()

	old := logs.Desugar()

	logs.RotateLog("/tmp", "test")

	newLog := logs.Desugar()

	assert.NotEqual(t, old, newLog)
}

func TestNewDebug(t *testing.T) {
	t.Parallel()

	old := logs.Desugar()

	logs.NewDebug()

	newLog := logs.Desugar()

	assert.NotEqual(t, old, newLog)
}

func TestNewInfo(t *testing.T) {
	t.Parallel()

	old := logs.Desugar()

	logs.NewInfo()

	newLog := logs.Desugar()

	assert.NotEqual(t, old, newLog)
}
