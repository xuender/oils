package gins

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StaticHandler fs.
func StaticHandler(urlPrefix string, fsys fs.FS, dir string) gin.HandlerFunc {
	if dir != "" {
		fsys, _ = fs.Sub(fsys, dir)
	}

	handler := http.FileServer(http.FS(fsys))
	if urlPrefix != "" {
		handler = http.StripPrefix(urlPrefix, handler)
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path[1:]
		if _, err := fsys.Open(path); path == "" || err == nil {
			handler.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
