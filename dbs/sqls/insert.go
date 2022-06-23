package sqls

import "strings"

// BatchInsert 生成批量插入SQL.
func BatchInsert(sql string, num int) string {
	start := strings.LastIndex(sql, "(")
	end := strings.LastIndex(sql, ")") + 1

	ret := make([]string, num)
	for i := 0; i < num; i++ {
		ret[i] = sql[start:end]
	}

	return sql[0:start] + strings.Join(ret, ",") + sql[end:]
}
