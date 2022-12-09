package base

import (
	"unicode"

	"github.com/samber/lo"
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
	slice := []rune{}
	hasSepInitialisms := slices.Contains(seps, SepInitialisms)
	isLower := false

	for index, elem := range str {
		if hasSepInitialisms && index > 0 && isLower && unicode.IsUpper(elem) {
			ret = append(ret, string(slice))
			slice = []rune{}
		}

		if slices.Contains(seps, elem) {
			ret = append(ret, string(slice))
			slice = []rune{}

			continue
		}

		slice = append(slice, elem)
		isLower = unicode.IsLower(elem)

		if !isLower && CommonInitialisms.Has(string(slice)) {
			ret = append(ret, string(slice))
			slice = []rune{}
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

			column[index1] = lo.Min([]int{
				column[index1] + 1,   // 插入
				column[index1-1] + 1, // 删除
				lastkey + sub,        // 修改
			})
			lastkey = oldkey
		}
	}

	return column[runeStr1Len]
}

// Wildcard 字符串通配符匹配.
// pattern 支持 * 和 ?.
func Wildcard(str string, pattern string) bool {
	strIndex, patternIndex := 0, 0
	match, startIndex := 0, -1
	isStar := func() bool {
		return patternIndex < len(pattern) && pattern[patternIndex] == '*'
	}

	for strIndex < len(str) {
		switch {
		case patternIndex < len(pattern) && (pattern[patternIndex] == '?' || pattern[patternIndex] == str[strIndex]):
			patternIndex++
			strIndex++
		case isStar():
			startIndex = patternIndex
			match = strIndex
			patternIndex++
		case startIndex != -1:
			patternIndex = startIndex + 1
			match++
			strIndex = match
		default:
			return false
		}
	}

	for isStar() {
		patternIndex++
	}

	return patternIndex == len(pattern)
}
