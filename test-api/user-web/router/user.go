package router

import (
	"github.com/gin-gonic/gin"
	"test-api/user-web/api"
	"test-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("/pwd_login", api.PassWordLogin)
		UserRouter.POST("/register", api.Register)
	}

}
