package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/portal")
	{
		UserRouter.POST("/column/create_column", v1.CreateSpecialColumn)
		UserRouter.POST("/column/add_to_column", v1.AddToColumn)
		UserRouter.POST("/column/list_all_from_column", v1.ListAllFromAColumn)
		UserRouter.POST("/column/remove_from_column", v1.RemovePaperFromColumn)
		UserRouter.POST("/author", v1.SearchAuthor)
		UserRouter.POST("/column/searchcol", v1.SearchSpecialColumn)

		UserRouter.POST("/recommend/create_recommend", v1.CreateRecommend)
		UserRouter.POST("/recommend/remove_recommend", v1.RemoveRecommend)
		UserRouter.POST("/recommend/recommends_from_one_author", v1.ListRecommendsFromOneAuthor)
		UserRouter.POST("/recommend/recommends_from_one_paper", v1.ListRecommendsFromOnePaper)
		UserRouter.POST("/recommend/top", v1.ListTopSevenPapers)
	}
}
