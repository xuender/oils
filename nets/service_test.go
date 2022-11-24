package nets_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/nets"
	"google.golang.org/protobuf/proto"
)

type testHandler struct {
	writer http.ResponseWriter
	req    *http.Request
}

func (p *testHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	p.writer = writer
	p.req = req
}

func (p *testHandler) GET() []byte {
	return []byte("123")
}

func (p *testHandler) POST(data []byte, isJSON bool) proto.Message {
	panic("error")
}

// nolint: exhaustruct
func (p *testHandler) Recover(any) proto.Message {
	return &Data{Str: "error"}
}

func TestNewService(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, nets.NewService(nil))
}

// nolint: exhaustruct
func TestService_Handler(t *testing.T) {
	t.Parallel()

	handler := &testHandler{}
	service := nets.NewService(handler)
	hand := service.Handler("api")
	req := httptest.NewRequest(http.MethodGet, "/api", bytes.NewReader([]byte("test")))
	res := &httptest.ResponseRecorder{Body: &bytes.Buffer{}}

	hand.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

// nolint: exhaustruct
func TestWithNotFound(t *testing.T) {
	t.Parallel()

	handler := &testHandler{}
	notHandler := nets.WithNotFound("/api", handler)
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte("test")))
	res := &httptest.ResponseRecorder{}

	notHandler.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)

	req = httptest.NewRequest(http.MethodGet, "/api/", bytes.NewReader([]byte("test")))
	res = &httptest.ResponseRecorder{Body: &bytes.Buffer{}}

	notHandler.ServeHTTP(res, req)
	assert.Equal(t, handler.req, req)
}

// nolint: exhaustruct
func TestWithLogging(t *testing.T) {
	t.Parallel()

	handler := &testHandler{}
	logHander := nets.WithLogging(handler)
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte("test")))
	res := &httptest.ResponseRecorder{}

	logHander.ServeHTTP(res, req)
	assert.True(t, handler.writer == res)
	assert.True(t, handler.req == req)
}

// nolint: exhaustruct
func TestWithRecover(t *testing.T) {
	t.Parallel()

	handler := &testHandler{}
	recoverHandler := nets.WithRecover(handler)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("test")))
	res := &httptest.ResponseRecorder{Body: &bytes.Buffer{}}

	recoverHandler.ServeHTTP(res, req)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, 7, len(res.Body.Bytes()))

	req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("test")))
	// req.Header = map[string][]string{}
	req.Header = http.Header(base.NewMapSameValue([]string{"json"}, nets.ContentType))
	res = &httptest.ResponseRecorder{Body: &bytes.Buffer{}}

	recoverHandler.ServeHTTP(res, req)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, 15, len(res.Body.Bytes()))

	req = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte("test")))
	res = &httptest.ResponseRecorder{Body: &bytes.Buffer{}}

	recoverHandler.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, 3, len(res.Body.Bytes()))
}

// nolint: exhaustruct
func TestWrite(t *testing.T) {
	t.Parallel()

	write := &bytes.Buffer{}
	data := &Data{Str: "xx"}

	nets.Write(write, data, true)
	assert.Equal(t, 12, len(write.Bytes()))

	write = &bytes.Buffer{}
	nets.Write(write, data, false)
	assert.Equal(t, 4, len(write.Bytes()))
}
