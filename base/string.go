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
