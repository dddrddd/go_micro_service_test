package router

import (
	"github.com/gin-gonic/gin"
	"test-api/goods-web/api/banners"
	"test-api/goods-web/middlewares"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	bannerRouter := Router.Group("banner")
	{
		bannerRouter.GET("/list", banners.BannerList)
		bannerRouter.DELETE("/delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.DeleteBanner)
		bannerRouter.PUT("/update/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.UpdateBanner)
		bannerRouter.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.NewBanner)
	}

}
