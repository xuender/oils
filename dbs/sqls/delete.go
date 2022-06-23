package sqls

import (
	"fmt"
	"strings"
)

// BatchDelete 生成批量删除SQL.
func BatchDelete(table, primaryKey string, count int) string {
	where := strings.Repeat(",?", count)

	return fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, primaryKey, where[1:])
}
