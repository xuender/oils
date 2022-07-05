package gins

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xuender/oils/logs"
)

// Check for a broken connection, as it is not really a
// condition that warrants a panic stack trace.
func brokenPipe(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		// nolint
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				return true
			}
		}
	}

	return false
}

func GinRecoveryHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
			if brokenPipe(err) {
				logs.Errorw(ctx.Request.URL.Path,
					"error", err,
					"request", string(httpRequest),
				)
				// If the connection is dead, we can't write a status to it.
				// nolint
				ctx.Error(err.(error))
				ctx.Abort()

				return
			}

			logs.Errorw("[Recovery from panic]",
				"error", err,
				"request", string(httpRequest),
			)

			switch err.(type) {
			case validator.ValidationErrors:
				ctx.JSON(http.StatusBadRequest, err)
				ctx.About()
			default:
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}
	}()

	ctx.Next()
}
