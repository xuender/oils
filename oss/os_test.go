package oss_test

import (
	"os/exec"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/oss"
)

// nolint: paralleltest
func TestOpen(t *testing.T) {
	cmd := &exec.Cmd{}

	patches1 := gomonkey.ApplyMethodReturn(cmd, "Start", nil)
	defer patches1.Reset()

	patches := gomonkey.ApplyFuncReturn(exec.Command, cmd)
	defer patches.Reset()

	assert.Nil(t, oss.Open("file"))
	assert.Nil(t, oss.Show("file"))

	pat := gomonkey.ApplyFuncReturn(oss.IsWindows, true)

	assert.Nil(t, oss.Open("file"))
	assert.Nil(t, oss.Show("."))
	assert.NotNil(t, oss.Show("file"))
	assert.Nil(t, oss.Show("os.go"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsLinux, true)

	assert.Nil(t, oss.Open("file"))
	assert.Nil(t, oss.Show("file"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsDarwin, true)

	assert.Nil(t, oss.Open("file"))
	assert.Nil(t, oss.Show("file"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsLinux, false)

	assert.NotNil(t, oss.Open("file"))
	assert.NotNil(t, oss.Show("file"))
	pat.Reset()
}
