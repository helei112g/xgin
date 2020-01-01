package router

import (
	v1 "dayu.com/gindemo/api/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) {
	apiv1 := e.Group("/api/v1")
	apiv1.GET("/ping", v1.Ping)
}
