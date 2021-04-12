package doc

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("project")
	router.GET("branch")
}
