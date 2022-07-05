package tests

import (
	"net/http"
	"strings"
)

func Session(res *http.Response) string {
	str := res.Header.Get("Set-Cookie")
	if index := strings.Index(str, ";"); index > 0 {
		return str[:index]
	}

	return str
}
