package doc

import (
	"context"
	"fmt"
	"net/http"

	"pbdoc/dao/repo"

	"github.com/gin-gonic/gin"
)

func ListProject(ctx *gin.Context) {
	projects, err := repo.ListAllProject(context.TODO())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  fmt.Printf("%+v", err),
		})
		return
	}

}
