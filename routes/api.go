package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册 API 路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("gohub/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
