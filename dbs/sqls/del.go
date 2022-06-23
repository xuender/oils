package sqls

import (
	"fmt"
	"strings"
)

func BatchDel(table, primaryKey string, count int) string {
	where := strings.Repeat(",?", count)

	return fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, primaryKey, where[1:])
}
