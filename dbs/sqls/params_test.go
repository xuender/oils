package sqls_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs/sqls"
)

func TestParams(t *testing.T) {
	t.Parallel()

	str, sql := sqls.Params("aa #id bb")

	assert.Equals(t, []string{"id"}, str)
	assert.Equal(t, "aa ? bb", sql)

	str, sql = sqls.Params("aa #id,#ff #dd bb")

	assert.Equals(t, []string{"id", "ff", "dd"}, str)
	assert.Equal(t, "aa ?,? ? bb", sql)

	str, sql = sqls.Params("aa #id")

	assert.Equals(t, []string{"id"}, str)
	assert.Equal(t, "aa ?", sql)

	str, sql = sqls.Params("#id aa")

	assert.Equals(t, []string{"id"}, str)
	assert.Equal(t, "? aa", sql)

	str, sql = sqls.Params("'#id' aa")

	assert.Equals(t, []string{"id"}, str)
	assert.Equal(t, "'?' aa", sql)
}
