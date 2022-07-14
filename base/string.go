package base

import (
	"unicode"

	"golang.org/x/exp/slices"
)

// SepInitialisms 首字母缩写分割.
const SepInitialisms = rune(-3)

// nolint
var CommonInitialisms = NewSet(
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
)

// Split 根据分割符拆分字符串.
// SepInitialisms 缩写子母拆分.
func Split(str string, seps ...rune) []string {
	if len(seps) == 0 {
		return []string{str}
	}

	ret := []string{}
	slice := NewSlice[rune]()
	hasSepInitialisms := Has(seps, SepInitialisms)
	isLower := false

	for index, elem := range str {
		if hasSepInitialisms && index > 0 && isLower && unicode.IsUpper(elem) {
			ret = append(ret, string(slice))
			slice.Clean()
		}

		if slices.Contains(seps, elem) {
			ret = append(ret, string(slice))
			slice.Clean()

			continue
		}

		slice.Add(elem)
		isLower = unicode.IsLower(elem)

		if !isLower && CommonInitialisms.Has(string(slice)) {
			ret = append(ret, string(slice))
			slice.Clean()
		}
	}

	ret = append(ret, string(slice))

	return ret
}

// SplitFunc 根据函数拆分字符串.
func SplitFunc(str string, splitFunc func(rune) bool) []string {
	ret := []string{}
	start := 0

	for index, value := range str {
		if splitFunc(value) {
			if start < 0 {
				continue
			}

			if start == index {
				start++

				continue
			}

			ret = append(ret, str[start:index])
			start = -1
		} else if start < 0 {
			start = index
		}
	}

	if start > -1 {
		ret = append(ret, str[start:])
	}

	return ret
}

// LevenshteinDistance 编辑距离.
func LevenshteinDistance(str1, str2 string) int {
	if str1 == str2 {
		return 0
	}

	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)
	runeStr1Len := len(runeStr1)
	runeStr2Len := len(runeStr2)

	if runeStr1Len == 0 {
		return runeStr2Len
	}

	if runeStr2Len == 0 {
		return runeStr1Len
	}

	column := make([]int, runeStr1Len+1)

	for index := 1; index <= runeStr1Len; index++ {
		column[index] = index
	}

	for index2 := 1; index2 <= runeStr2Len; index2++ {
		column[0] = index2
		lastkey := index2 - 1

		for index1 := 1; index1 <= runeStr1Len; index1++ {
			oldkey := column[index1]
			sub := 0

			if runeStr1[index1-1] != runeStr2[index2-1] {
				sub = 1
			}

			column[index1] = Min(
				column[index1]+1,   // 插入
				column[index1-1]+1, // 删除
				lastkey+sub,        // 修改
			)
			lastkey = oldkey
		}
	}

	return column[runeStr1Len]
}

// StringMatch 字符串匹配.
// pattern 支持 * 和 ?.
func StringMatch(str string, pattern string) bool {
	strLen, patternLen := len(str), len(pattern)
	data := make([][]bool, strLen+1)

	for i := 0; i <= strLen; i++ {
		data[i] = make([]bool, patternLen+1)
	}

	data[0][0] = true

	for index := 1; index <= patternLen; index++ {
		if pattern[index-1] == '*' {
			data[0][index] = true
		} else {
			break
		}
	}

	for i := 1; i <= strLen; i++ {
		for j := 1; j <= patternLen; j++ {
			if pattern[j-1] == '*' {
				data[i][j] = data[i][j-1] || data[i-1][j]
			} else if pattern[j-1] == '?' || str[i-1] == pattern[j-1] {
				data[i][j] = data[i-1][j-1]
			}
		}
	}

	return data[strLen][patternLen]
}
