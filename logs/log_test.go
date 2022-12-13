package logs_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
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

	logs.RotateLog("/tmp", "test")
}
