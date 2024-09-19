package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api/goods-web/models"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 { //设置这个为2的是管理员
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
