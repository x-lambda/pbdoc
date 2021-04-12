package pbgen

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 路由注册
func RegisterRouter(router *gin.RouterGroup) {

	// 新建/修改分支
	router.POST("create", CreateDoc)

	// 删除分支
	router.POST("delete", DeleteDoc)
}
