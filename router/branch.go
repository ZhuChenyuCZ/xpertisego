package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBranchRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/branch")
	{
		UserRouter.POST("/comment/create", v1.CreateAComment) // 注意不用加逗号
		UserRouter.POST("/comment/operate", v1.OperateComment)
	}
}
