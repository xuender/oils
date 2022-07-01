package tests

import (
	"net/http"
	"strings"
)

func Session(res *http.Response) string {
	str := res.Header.Get("Set-Cookie")

	return str[:strings.Index(str, ";")]
}
