package rdb_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/rdb"
)

func TestDynamicLimit_ID(t *testing.T) {
	t.Parallel()

	limit := rdb.NewDynamicLimit(1000)
	assert.GreaterOrEqual(t, limit.ID(), 0)
}
