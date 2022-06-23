package sqls

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/xuender/oils/base"
)

// TableName sql语句中获取表名.
func TableName(sql string) string {
	sql = strings.ToLower(sql)
	sql = strings.ReplaceAll(sql, "`", "")

	next := false

	for _, str := range base.SplitFunc(sql, unicode.IsSpace) {
		if next {
			return str
		}

		if str == "from" || str == "into" {
			next = true
		}
	}

	return ""
}

// TableNames 数据库名及表名.
func TableNames(sql string) (schema string, table string) {
	name := TableName(sql)
	names := strings.Split(name, ".")

	if len(names) == 1 {
		table = names[0]
	}

	if len(names) > 1 {
		schema = names[0]
		table = names[1]
	}

	return
}

const regSpace = "\\s+"

var RegSpace = regexp.MustCompile(regSpace)

// Count 计数SQL.
func Count(sql string) string {
	sql = RegSpace.ReplaceAllString(strings.TrimSpace(sql), " ")
	lower := strings.ToLower(sql)
	start := strings.IndexFunc(lower, unicode.IsSpace)
	from := strings.Index(lower, "from")
	end := len(lower)

	if order := strings.Index(lower, "order "); order > 0 && order < end {
		end = order
	}

	if group := strings.Index(lower, "group "); group > 0 && group < end {
		end = group
	}

	if limit := strings.Index(lower, "limit "); limit > 0 && limit < end {
		end = limit
	}

	return sql[:start] + " count(*) " + strings.TrimSpace(sql[from:end])
}

// Offset 倒叙生成offset，防止丢数据.
func Offset(count, limit int) []int {
	size := count / limit
	if count%limit > 0 {
		size++
	}

	ret := make([]int, size)

	for i := 0; i < size; i++ {
		ret[i] = (size - i - 1) * limit
	}

	return ret
}
