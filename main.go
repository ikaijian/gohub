package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello word!")
	// gin Engine 实例
	r := gin.New()
	// 设置中间件
	r.Use(gin.Logger(), gin.Recovery())
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		// 响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// html
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
	// 运行服务
	r.Run(":8000")
}
