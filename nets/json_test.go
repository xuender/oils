package nets_test

import (
	"net/http"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/nets"
)

// nolint: exhaustruct
func TestIsJSON(t *testing.T) {
	t.Parallel()

	header := map[string][]string{}
	header[nets.ContentType] = []string{"application/json"}

	req := &http.Request{Header: header}
	assert.True(t, nets.IsJSON(req))

	header[nets.ContentType] = []string{"application/javascript"}

	assert.False(t, nets.IsJSON(req))
}
