package gins

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuender/oils/logs"
)

func GinLogHandler(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery

	ctx.Next()

	logs.Infow(path,
		"status", ctx.Writer.Status(),
		"method", ctx.Request.Method,
		"path", path,
		"query", query,
		"ip", ctx.ClientIP(),
		"user-agent", ctx.Request.UserAgent(),
		"errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
		"cost", time.Since(start),
	)
}
