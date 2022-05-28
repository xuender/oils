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
