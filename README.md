## 章节总结

### 第一章 项目介绍

#### 功能

+ 登录注册
+ 用户
+ 话题

#### 第三方依赖

使用到的开源库：

+ [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
+ [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
+ [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
+ [cobra](https://github.com/spf13/cobra) —— 命令行结构
+ [viper](https://github.com/spf13/viper) —— 配置信息
+ [cast](https://github.com/spf13/cast) —— 类型转换
+ [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
+ [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
+ [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
+ [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
+ [limiter](https://github.com/ulule/limiter/v3) —— 限流器
+ [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
+ [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
+ [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
+ [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
+ [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理
+ [faker](https://learnku.com/courses/go-api/1.19/finish-up/github.com/bxcodec/faker) —— 假数据填充
+ [imaging](https://learnku.com/courses/go-api/1.19/finish-up/github.com/disintegration/imaging) —— 图片裁切

### 第二章 环境搭建

#### 本章知识点

请参考以下这几个视频教程来配置环境：

+ [Go 开发环境配置（Windows 10）](https://learnku.com/courses/go-video/2022/go-development-environment-configuration-windows-10/11304)
+ [Go 开发环境配置（Mac）](https://learnku.com/courses/go-video/2022/001-go-development-environment-configuration-mac/11303)
+ [配置 VSCode Go 开发编辑器](https://learnku.com/courses/go-video/2022/vscode-as-go-development-editor/11306)
+ [VSCode Go 插件详解](https://learnku.com/courses/go-video/2022/vscode-go-plug-in-details/11314)
+ [Go Modules 详解 - 相关命令](https://learnku.com/courses/go-video/2022/006-go-modules-related-commands/11454)
  另外国内网络下载第三方包不稳定，请按照以下教程配置 Goproxy：

> Wiki：[Go 文档和加速：Go 国内加速镜像](https://learnku.com/go/wikis/38122)

#### 其他软件

Redis 管理工具
推荐使用 [AnotherRedisDesktopManager](https://github.com/qishibo/AnotherRedisDesktopManager/releases)，跨平台，支持
Mac、Windows 和 Linux。

数据库管理工具
推荐使用 [TablePlus](https://tableplus.com/) ，跨平台，支持 Mac、Windows 和 Linux。

### 第三章 开始编码

#### 本章知识点

- git init 初始化版本控制
- go mod init 初始化项目
- 通用的 Go 项目 .gitignore 文件
- air 自动重载工具
- 集成 gin 包
- 自定义 404 处理器
- 养成『关注错误提示』的习惯
- go mod tidy 的使用
- API 路由使用单独的 api.go 文件
- 路由组实现 API 版本（v1）
- 配置方案 .env 和 config 的设计
- viper 和 cast 包的应用
- 应用 Go 的 init() 函数
- 利用 reflect 标准库构建 Empty() 方法

### 第四章 手机或邮箱是否注册

#### 本章知识点

- 登录流程
- 注册流程
- 找回密码流程
- 安全知识 —— 轰炸机
- 安全知识 —— 暴力破解
- Gorm 连接 MySQL 服务器
- Gorm 连接 Sqlite 数据库
- Gorm 模型映射
- 验证手机合法性
- 验证邮箱合法性
- 解析用户请求

#### 推荐学习资料

+ [Gin框架在中文文档](https://learnku.com/docs/gin-gonic/1.7)
+ [绑定表单数据至自定义结构体](https://learnku.com/docs/gin-gonic/1.7/examples-bind-form-data-request-with-custom-struct/11393)
+ [绑定查询字符串或表单数据](https://learnku.com/docs/gin-gonic/1.7/examples-bind-query-or-post/11392)
+ [Multipart/Urlencoded 绑定](https://learnku.com/docs/gin-gonic/1.7/examples-multipart-urlencoded-binding/11366)
+ [只绑定 url 查询字符串](https://learnku.com/docs/gin-gonic/1.7/examples-only-bind-query-string/11380)
+ [模型定义《GORM 中文文档](https://learnku.com/docs/gorm/v2/models/9729)
+ [查询《GORM 中文文档》](https://learnku.com/docs/gorm/v2/query/9733)
+ [连接数据库《GORM 中文文档](https://learnku.com/docs/gorm/v2/connecting_to_the_database/9731)
+ [使用中间件《Gin 框架中文文档](https://learnku.com/docs/gin-gonic/1.7/examples-using-middleware/11379)
+ [记录日志《GORM 中文文档》](https://learnku.com/docs/gorm/v2/logger/9761)

### 第五章 日志和错误处理

- 日志系统的等级；
- 日志的使用场景；
- 如何分割日志文件；
- zap 高性能日志的使用；
- 使用 zap 记录 HTTP 访问日志；
- Panic Recovery 方案；
- 自定义 Gorm Logger

### 第六章 图片验证码

+ 图片验证码与数字验证码的区别；
+ 图片验证码的流程；
+ 使用 Redis；
+ 使用 logger.Dump() 方法调试；
+ 使用 Postman 环境变量；
+ response 库统一响应格式；
+ 使用『多变参数』实现函数『参数默认值』；
+ 功能模块与底层库之间的错误处理实践

#### base64图片在线转换工具

[base64图片在线转换](https://tool.chinaz.com/tools/imgtobase)

##### 安装

~~~
go install github.com/mailhog/MailHog@latest
~~~

### 第七章 数字验证码

+ 使用阿里短信平台的测试账号；
+ 发短信的四个要素；
+ 使用 interface 来抽象 sms 包的驱动；
+ 短信验证码发送和验证流程；
+ 发送短信；
+ 验证码使用 interface 来抽象存储驱动；
+ 发送短信验证码；
+ 使用 sync.Once 实现单例模式；
+ Mailhog 邮件测试服务；
+ 使用 interface 来抽象 mail 包发邮件驱动；
+ 发送 SMTP 邮件；
+ 发送邮件验证码

#### Mailhog工具

##### 安装

~~~
go install github.com/mailhog/MailHog@latest
~~~

##### 启动

~~~
MailHog
~~~

以上输出有两个比较重要的信息：
0.0.0.0:1025 是 SMTP 端口
127.0.0.1:8025/ 是网页界面

### 第八章 用户注册

+ 用户注册信息的验证；
+ 自定义 not_exists 规则；
+ 自定义验证器；
+ 使用测试手机号和测试邮件进行注册；
+ 密码加密和匹对；
+ JWT 创建 token ；
+ 使用通用时区；

#### 推荐学习资料：

+ [govalidator 的验证规则](https://github.com/thedevsaddam/govalidator)
+ [JWT](https://jwt.io/introduction/)

### 第九章 用户登录

+ 用来管理身份认证的的 auth 包；
+ 使用手机号 + 验证短信登录用户；
+ 使用密码登录，支持手机号、用户名、邮箱；
+ 使用 GORM 发送 Where Or 请求；
+ 登录成功后签发 JWT Token；
+ JWT Token 刷新；
+ 解码 JWT Token，分析里面的三段信息以及安全性；
+ 使用 Postman 发起 Authorization:Bearer xxxxx 标头的请求；
+ Auth 中间件，授权才能操作；
+ Guest 中间件，只允许游客操作；

#### 推荐学习资料：

+ [自定义中间件《Gin 框架中文文档》](https://learnku.com/docs/gin-gonic/1.7/examples-custom-middleware/11395)
+ [在中间件中使用 Goroutine](https://learnku.com/docs/gin-gonic/1.7/examples-goroutines-inside-a-middleware/11381)
+ [使用 BasicAuth 中间件](https://learnku.com/docs/gin-gonic/1.7/examples-using-basicauth-middleware/11377)
+ [使用中间件《Gin 框架中文文档》](https://learnku.com/docs/gin-gonic/1.7/examples-using-middleware/11379)
+ [不使用默认的中间件《Gin 框架中文文档》](https://learnku.com/docs/gin-gonic/1.7/examples-without-middleware/11374)

### 第十章 找回密码

+ 通过手机号找回密码；
+ 通过邮箱找回密码；
+ 全局限流中间件；
+ 针对单个路由的限流；
+ 使用 Postman 查看 Header 信息；
+ AuthJWT 和 GuestJWT 的使用

### 第十一章 命令行模式

#### 本章知识点：

+ 命令组成的四个部分；
+ 参数（args） 与选项（flag）的区别；
+ 长选项与短选项的区别（–env -e）
+ 最佳实践：终端英文输出
+ Cobra 与 Cli 的选择；
+ cobra.Command 所有选项的学习；
+ cobra 的命令钩子；
+ 终端打印高亮信息；
+ 创建 cobra 主命令；
+ 生成随机字符串；
+ key 命令；
+ cobra 的参数校验器；
+ 类似于 go.dev/play/ 的 play 命令；

#### 推荐学习资料：

+ [cobra 官方指南](github.com/spf13/cobra/blob/master)
+ [pflag](github.com/spf13/pflag) 项目也需要了解一下
+ 另外一个[知名项目](cli github.com/urfave/cli)
+ 代码高亮的两个库了解一下
+ [ansi](github.com/mgutz/ansi)
+ [color](github.com/fatih/color)

### 第十二章 Make命令

#### 本章知识点：

+ 代码生成的逻辑（模板、变量替换）；
+ str 包，大小写处理，单复数处理；
+ 文件读写；
+ 创建目录；
+ 使用 Go 的 embed 功能；
+ make cmd 命令；
+ make model 命令；
+ make apicontroller 命令；
+ make request 命令；

#### 推荐学习

钩子[《GORM 中文文档》](https://learnku.com/docs/gorm/v2/hooks/9744)



