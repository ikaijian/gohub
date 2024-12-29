package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/app"
	"gohub/pkg/console"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table users",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(2), // 只允许且必须传 1 个参数
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	// 日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	modelName := makeModelFromString(args[1])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName, "{{ModelName}}": modelName.StructName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}
