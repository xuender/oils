package nets

import (
	"net/http"
	"strings"
)

func IsJSON(req *http.Request) bool {
	return strings.Contains(req.Header.Get(ContentType), "json")
}
