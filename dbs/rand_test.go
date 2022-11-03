package dbs_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs"
)

func TestRand(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		assert.NotEqual(t, dbs.Rand(), dbs.Rand())
		assert.Greater(t, uint64(1<<54), dbs.Rand())
	}
}
