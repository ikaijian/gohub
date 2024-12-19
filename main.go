package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

// 加载 config 目录下的配置信息
func init() {
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载的配置文件 env 文件, 默认为 dev")
	flag.Parse()
	config.InitConfig(env)

	// new 一个gin Engine 实例
	router := gin.New()
	// 初始化 DB
	bootstrap.SetupDB()
	// 初始化路由
	bootstrap.SetupRoute(router)
	// 初始化 Logger
	bootstrap.SetupLogger()

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
