package sqls_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs/sqls"
)

func TestBatchDelete(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "DELETE FROM t WHERE i IN (?,?)", sqls.BatchDelete("t", "i", 2))
}
