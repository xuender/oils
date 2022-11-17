package nets_test

import (
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/nets"
)

// nolint
func TestListenAndServe(t *testing.T) {
	patches := gomonkey.ApplyMethodReturn(&http.Server{}, "ListenAndServe", nil)
	defer patches.Reset()

	assert.Nil(t, nets.ListenAndServe(":8080", nil))
}
