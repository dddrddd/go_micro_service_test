package router

import (
	"github.com/gin-gonic/gin"
	"test-api/goods-web/api/category"
	"test-api/goods-web/middlewares"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategorysRouter := Router.Group("categorys")
	{
		CategorysRouter.GET("/list", category.List)
		CategorysRouter.DELETE("/delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.Delete)
		CategorysRouter.GET("/detail", category.Detail)
		CategorysRouter.PUT("/update/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.Update)
		CategorysRouter.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.New)
	}

}
