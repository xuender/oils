package nets

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	handler Handler
}

func NewService(handler Handler) *Service {
	return &Service{handler: handler}
}

func (p *Service) Handler(path string) http.Handler {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return WithNotFound(path, WithLogging(WithRecover(p.handler)))
}

func WithNotFound(path string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.RequestURI, path) {
			handler.ServeHTTP(writer, req)

			return
		}

		http.DefaultServeMux.ServeHTTP(writer, req)
		logs.Debugw("404", "ip", req.RemoteAddr, "uri", req.RequestURI, "method", req.Method)
	})
}

func WithLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		start := time.Now()
		handler.ServeHTTP(writer, req)

		logs.Infow("access", "ip", req.RemoteAddr, "uri", req.RequestURI, "method", req.Method, "time", time.Since(start))
	})
}

func WithRecover(handler Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write(handler.GET())

			return
		}

		isJSON := IsJSON(req)

		defer func() {
			if err := recover(); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				Write(writer, handler.Recover(err), isJSON)
				logs.Error(err)
			}
		}()

		defer req.Body.Close()

		if isJSON {
			writer.Header().Set(ContentType, "application/json; charset=utf-8")
		} else {
			writer.Header().Set(ContentType, "application/x-protobuf")
		}

		data := base.Must1(io.ReadAll(req.Body))
		msg := handler.POST(data, isJSON)

		writer.WriteHeader(http.StatusOK)
		Write(writer, msg, isJSON)
	})
}

func Write(writer http.ResponseWriter, msg proto.Message, isJSON bool) {
	var data []byte
	if isJSON {
		// nolint
		data = base.Must1(json.Marshal(msg))
	} else {
		data = base.Must1(proto.Marshal(msg))
	}

	base.Must1(writer.Write(data))
}
