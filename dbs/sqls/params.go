package sqls

import "strings"

// Params 参数及去除参数的SQL.
func Params(sql string) ([]string, string) {
	sqls := []string{}
	pks := []string{}
	index := 0

	for {
		start := strings.Index(sql[index:], "#")
		if start < 0 {
			break
		}

		start += index

		sqls = append(sqls, sql[index:start])
		sqls = append(sqls, "?")

		start++

		end := strings.IndexAny(sql[start:], "#, '()")
		if end < 0 {
			index = len(sql)
			pks = append(pks, sql[start:])

			break
		}

		end += start
		pks = append(pks, sql[start:end])
		index = end
	}

	sqls = append(sqls, sql[index:])

	return pks, strings.Join(sqls, "")
}
