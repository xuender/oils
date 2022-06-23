package sqls_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs/sqls"
)

func TestTableName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "table", sqls.TableName("Select * From table Limit 1"))
	assert.Equal(t, "", sqls.TableName("table Limit 1"))
}

func TestTableNames(t *testing.T) {
	t.Parallel()

	schema, table := sqls.TableNames("Select * From table Limit 1")

	assert.Equal(t, "table", table)
	assert.Equal(t, "", schema)

	schema, table = sqls.TableNames("Select * From schema.table Limit 1")

	assert.Equal(t, "table", table)
	assert.Equal(t, "schema", schema)
}

func TestCount(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "select count(*) from a where b=?", sqls.Count("select a, b\tfrom a\t\twhere b=? order by a limit 10"))
	assert.Equal(t, "select count(*) from a where b=?", sqls.Count("select a, b from a where b=? limit 10"))
	assert.Equal(t, "select count(*) from a where b=?", sqls.Count("select a, b from a where b=? group by b"))
}

func TestOffset(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []int{}, sqls.Offset(0, 100))
	assert.Equals(t, []int{0}, sqls.Offset(10, 100))
	assert.Equals(t, []int{8, 4, 0}, sqls.Offset(10, 4))
	assert.Equals(t, []int{5, 0}, sqls.Offset(10, 5))
	assert.Equals(t, []int{9, 6, 3, 0}, sqls.Offset(10, 3))
}
