package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/xuender/oils/base"
)

func Get(engine *gin.Engine, uri string) *http.Response {
	return method(engine, http.MethodGet, uri, "", nil)
}

func GetBySession(engine *gin.Engine, uri, session string) *http.Response {
	return method(engine, http.MethodGet, uri, session, nil)
}

func Post(engine *gin.Engine, uri string, body any) *http.Response {
	data := base.Must1(json.Marshal(body))

	return method(engine, http.MethodPost, uri, "", bytes.NewReader(data))
}

func PostBySession(engine *gin.Engine, uri, session string, body any) *http.Response {
	data := base.Must1(json.Marshal(body))

	return method(engine, http.MethodPost, uri, session, bytes.NewReader(data))
}

func PutBySession(engine *gin.Engine, uri, session string, body any) *http.Response {
	data := base.Must1(json.Marshal(body))

	return method(engine, http.MethodPut, uri, session, bytes.NewReader(data))
}

func Delete(engine *gin.Engine, uri string) *http.Response {
	return method(engine, http.MethodDelete, uri, "", nil)
}

func DeleteBySession(engine *gin.Engine, uri, session string) *http.Response {
	return method(engine, http.MethodDelete, uri, session, nil)
}

func method(engine *gin.Engine, method, uri, session string, body io.Reader) *http.Response {
	req := httptest.NewRequest(method, uri, body)
	req.Header = http.Header{
		"Content-Type": {"application/json; charset=utf-8"},
	}

	if session != "" {
		req.Header.Add("Cookie", session)
	}

	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	return rec.Result()
}
