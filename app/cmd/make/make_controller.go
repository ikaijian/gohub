package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"os"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，exmaple: make controller users v1/users",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(2), // 只允许且必须传 1 个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	// 格式化模型名称对象
	modelModule := makeModelFromString(args[0])
	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[1], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.Model 实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	dir := fmt.Sprintf("app/http/controllers/api/%s/%s", apiVersion, modelModule.PackageName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.LogIf(err)
	}

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s/%s_controller.go", apiVersion, modelModule.PackageName, model.TableName)
	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model)
}
