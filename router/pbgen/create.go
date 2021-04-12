package pbgen

import (
	"fmt"
	"net/http"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
)

type createReq struct {
	AccessKey   string `json:"access_key" form:"access_key"`
	ProjectName string `json:"project_name" form:"project_name"`
	BranchName  string `json:"branch_name" form:"branch_name"`
}

func CreateDoc(ctx *gin.Context) {
	req := createReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "bad request",
		})
	}

	// 这里的文件都视为安全文件
	fd, err := ctx.FormFile("pb")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "Get Upload File Failed",
		})

		return
	}

	dst := path.Join("/tmp", fd.Filename)
	err = ctx.SaveUploadedFile(fd, dst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "Save File Failed",
		})

		return
	}

	// TODO protoc-gen-markdown
	// 修改 github.com/lvht/protoc-gen-markdown
	// 直接生成临时文件
	cmd := fmt.Sprintf("protoc -I /tmp --markdown_out=. %s", fd.Filename)
	fmt.Println(cmd)
	c := exec.Command(cmd)
	_, err = c.CombinedOutput()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "protoc markdown error",
		})
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})

	// TODO pandoc 生成静态资源
	return
}
