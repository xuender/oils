package sqls_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs/sqls"
)

func TestBatchInsert(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "a (?,?),(?,?) b", sqls.BatchInsert("a (?,?) b", 2))
}
