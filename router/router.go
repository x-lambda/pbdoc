package router

import (
	"pbdoc/router/doc"
	"pbdoc/router/pbgen"

	"github.com/gin-gonic/gin"
)

var g *gin.Engine

func init() {
	g = gin.Default()

	// 统一注册路由的地方
	pbgen.RegisterRouter(g.Group("pbgen"))
	doc.RegisterRouter(g.Group("api-doc"))
}

func GetRouter() *gin.Engine {
	return g
}
