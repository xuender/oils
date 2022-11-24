package tests_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/gins/tests"
)

func TestSession(t *testing.T) {
	t.Parallel()

	res := &http.Response{Header: http.Header{}}

	res.Header.Add("Set-Cookie", "aa;f=2")
	assert.Equal(t, "aa", tests.Session(res))

	res = &http.Response{Header: http.Header{}}
	res.Header.Add("Set-Cookie", "aa")
	assert.Equal(t, "aa", tests.Session(res))
}
