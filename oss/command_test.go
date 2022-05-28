package oss_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/oss"
)

func TestExec(t *testing.T) {
	t.Parallel()

	assert.Nil(t, oss.Exec("ls"))
}

func TestExecOut(t *testing.T) {
	t.Parallel()

	o, err := oss.ExecOut("ls")

	assert.Nil(t, err)
	assert.NotNil(t, o)
}
