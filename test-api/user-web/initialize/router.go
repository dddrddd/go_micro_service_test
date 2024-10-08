package initialize

import (
	"github.com/gin-gonic/gin"
	"test-api/user-web/middlewares"
	router2 "test-api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	return Router
}
