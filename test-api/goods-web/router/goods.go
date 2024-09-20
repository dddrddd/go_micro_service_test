package router

import (
	"github.com/gin-gonic/gin"
	"test-api/goods-web/api/goods"
	"test-api/goods-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("/list", goods.List)
		GoodsRouter.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
		GoodsRouter.GET("/detail", goods.Detail)
		GoodsRouter.PATCH("/update_status/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
		GoodsRouter.PUT("/update/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.DELETE("/delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)
		GoodsRouter.GET("/stocks", goods.Stocks)

	}

}
