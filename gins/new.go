package gins

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	ret := gin.New()
	ret.Use(GinLogHandler, GinRecoveryHandler)

	return ret
}
