package router

import (
	"github.com/gin-gonic/gin"
	"test-api/goods-web/api/brands"
	"test-api/goods-web/middlewares"
)

func InitBrandsRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brand")
	{
		BrandRouter.GET("/list", brands.BrandList)
		BrandRouter.DELETE("/delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.DeleteBrand)
		BrandRouter.PUT("/update/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.UpdateBrand)
		BrandRouter.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.NewBrand)
	}
	CategoryBrandRouter := Router.Group("categorybrand")
	{
		CategoryBrandRouter.GET("/list", brands.CategoryBrandList)
		CategoryBrandRouter.DELETE("/delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.DeleteCategoryBrand)
		CategoryBrandRouter.PUT("/update/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.UpdateCategoryBrand)
		CategoryBrandRouter.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), brands.NewCategoryBrand)
		CategoryBrandRouter.GET("/getBycategoryId", brands.GetCategoryBrandList)
	}
}
