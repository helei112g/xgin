package benchmark

import (
	"dayu.com/gindemo/router"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init()  {
	Engine = gin.New()
	router.SetupRouter(Engine)
}
