package logs_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

// nolint: paralleltest
func TestLogName(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.IsWindows, false)

	assert.Equal(t, "/var/tmp/logs.log", logs.LogName())
	patches.Reset()

	patches = gomonkey.ApplyFuncReturn(oss.IsWindows, true)

	assert.NotEqual(t, "/var/tmp/logs.log", logs.LogName())
	patches.Reset()
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
