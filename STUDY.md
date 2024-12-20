## 推荐学习资料【第四章】

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

### 第五章 日志和错误处理

- 日志系统的等级；
- 日志的使用场景；
- 如何分割日志文件；
- zap 高性能日志的使用；
- 使用 zap 记录 HTTP 访问日志；
- Panic Recovery 方案；
- 自定义 Gorm Logger。

