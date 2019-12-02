package router

import (
	"dayu.com/gindemo/app/controller/site"
	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) {
	e.GET("/ping", site.Ping)
}
