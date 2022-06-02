package base

import (
	"unicode"
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

		if Has(seps, elem) {
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
