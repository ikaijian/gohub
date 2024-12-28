package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, example: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	createFileFromStub(filePath, "factory", model)
}
