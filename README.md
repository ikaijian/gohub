## 项目结构设计

~~~text
.├── app                            // 程序具体逻辑代码
│   ├── cmd                         // 命令
│   │   ├── cache.go                
│   │   ├── cmd.go
│   │   ├── key.go
│   │   ├── make                    // make 命令及子命令
│   │   │   ├── make.go
│   │   │   ├── make_apicontroller.go
│   │   │   ├── make_cmd.go
│   │   │   ├── make_factory.go
│   │   │   ├── make_migration.go
│   │   │   ├── make_model.go
│   │   │   ├── make_policy.go
│   │   │   ├── make_request.go
│   │   │   ├── make_seeder.go
│   │   │   └── stubs               // make 命令的模板
│   │   │       ├── apicontroller.stub
│   │   │       ├── cmd.stub
│   │   │       ├── factory.stub
│   │   │       ├── migration.stub
│   │   │       ├── model
│   │   │       │   ├── model.stub
│   │   │       │   ├── model_hooks.stub
│   │   │       │   └── model_util.stub
│   │   │       ├── policy.stub
│   │   │       ├── request.stub
│   │   │       └── seeder.stub
│   │   ├── migrate.go
│   │   ├── play.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── http                        // http 请求处理逻辑
│   │   ├── controllers             // 控制器，存放 API 和视图控制器
│   │   │   ├── api                 // API 控制器，支持多版本的 API 控制器
│   │   │   │   └── v1              // v1 版本的 API 控制器
│   │   │   │       ├── users_controller.go
│   │   │   │       └── ...
│   │   └── middlewares             // 中间件
│   │       ├── auth_jwt.go
│   │       ├── guest_jwt.go
│   │       ├── limit.go
│   │       ├── logger.go
│   │       └── recovery.go
│   ├── models                      // 数据模型
│   │   ├── user                    // 单独的模型目录
│   │   │   ├── user_hooks.go       // 模型钩子文件
│   │   │   ├── user_model.go       // 模型主文件
│   │   │   └── user_util.go        // 模型辅助方法
│   │   └── ...
│   ├── policies                    // 授权策略目录
│   │   ├── category_policy.go
│   │   └── ...
│   └── requests                    // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
│       ├── validators              // 自定的验证规则
│       │   ├── custom_rules.go
│       │   └── custom_validators.go
│       ├── user_request.go
│       └── ...
├── bootstrap                       // 程序模块初始化目录
│   ├── app.go  
│   ├── cache.go
│   ├── database.go
│   ├── logger.go
│   ├── redis.go
│   └── route.go
├── config                          // 配置信息目录
│   ├── app.go
│   ├── captcha.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── mail.go
│   ├── pagination.go
│   ├── redis.go
│   ├── sms.go
│   └── verifycode.go
├── database                        // 数据库相关目录
│   ├── database.db                 // sqlite 数据文件（加入到 .gitignore 中）
│   ├── factories                   // 模型工厂目录
│   │   ├── user_factory.go
│   │   └── ...
│   ├── migrations                  // 数据库迁移目录
│   │   ├── 2021_12_21_102259_create_users_table.go
│   │   ├── 2021_12_21_102340_create_categories_table.go
│   │   └── ...
│   └── seeders                     // 数据库填充目录
│       ├── users_seeder.go
│       ├── ...
├── pkg                             // 内置辅助包
│   ├── app
│   ├── auth
│   ├── cache
│   ├── captcha
│   ├── config
│   └── ...
├── public                          // 静态文件存放目录
│   ├── css
│   ├── js
│   └── uploads                     // 用户上传文件目录
│       └── avatars                 // 用户上传头像目录
├── routes                          // 路由
│   ├── api.go
│   └── web.go
├── storage                         // 内部存储目录
│   ├── app
│   └── logs                        // 日志存储目录
│       ├── 2021-12-28.log
│       ├── 2021-12-29.log
│       ├── 2021-12-30.log
│       └── logs.log
└── tmp                             // air 的工作目录
├── .env                            // 环境变量文件
├── .env.example                    // 环境变量示例文件
├── .gitignore                      // git 配置文件
├── .air.toml                       // air 配置文件
├── .editorconfig                   // editorconfig 配置文件
├── go.mod                          // Go Module 依赖配置文件
├── go.sum                          // Go Module 模块版本锁定文件
├── main.go                         // Gohub 程序主入口
├── Makefile                        // 自动化命令文件
├── readme.md                       // 项目 readme
~~~

driver_aliyun.go 文件，其中结果日志我是写到 mysql 表里了

~~~go
package sms

import (
	"encoding/json"
	"gohub/pkg/config"
	"gohub/pkg/logger"

	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Aliyun struct{}

func (a *Aliyun) Send(phone string, message Message) bool {
	// 从配置中获取阿里云短信配置
	aliyunConfig := config.GetStringMapString("sms.aliyun")

	// 打印配置信息用于调试
	logger.DebugJSON("短信[阿里云]", "配置信息", aliyunConfig)

	if len(aliyunConfig) == 0 {
		logger.ErrorString("短信[阿里云]", "配置错误", "找不到阿里云短信配置信息")
		return false
	}

	// 检查必要的配置项
	requiredConfigs := []string{"access_key_id", "access_key_secret", "sign_name"}
	for _, c := range requiredConfigs {
		if aliyunConfig[c] == "" {
			logger.ErrorString("短信[阿里云]", "配置错误", c+" 未配置")
			return false
		}
	}

	clientConfig := &openapi.Config{
		AccessKeyId:     tea.String(aliyunConfig["access_key_id"]),
		AccessKeySecret: tea.String(aliyunConfig["access_key_secret"]),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		RegionId:        tea.String("cn-hangzhou"), // 添加区域配置
	}

	// 打印客户端配置用于调试
	logger.DebugJSON("短信[阿里云]", "客户端配置", map[string]string{
		"AccessKeyId": *clientConfig.AccessKeyId,
		"Endpoint":    *clientConfig.Endpoint,
		"RegionId":    *clientConfig.RegionId,
	})

	smsClient, err := dysmsapi.NewClient(clientConfig)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "创建短信客户端失败", err.Error())
		return false
	}

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "模板参数序列化失败", err.Error())
		return false
	}

	// 打印请求参数用于调试
	logger.DebugJSON("短信[阿里云]", "请求参数", map[string]string{
		"PhoneNumbers":  phone,
		"SignName":      aliyunConfig["sign_name"],
		"TemplateCode":  message.Template,
		"TemplateParam": string(templateParam),
	})

	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(aliyunConfig["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		TemplateParam: tea.String(string(templateParam)),
	}

	// 添加更多的运行时选项
	runtime := &util.RuntimeOptions{
		ReadTimeout:    tea.Int(5000),
		ConnectTimeout: tea.Int(5000),
	}

	response, err := smsClient.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "发送短信失败", err.Error())
		return false
	}

	if response.Body == nil || *response.Body.Code != "OK" {
		if response.Body != nil {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", *response.Body.Message)
		} else {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", "response body is nil")
		}
		return false
	}

	logger.DebugString("短信[阿里云]", "发送成功", "")
	return true
}
~~~
