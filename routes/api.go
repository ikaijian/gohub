package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	categoryCtl "gohub/app/http/controllers/api/v1/category"
	linkCtl "gohub/app/http/controllers/api/v1/link"
	topicCtl "gohub/app/http/controllers/api/v1/topic"
	userCtl "gohub/app/http/controllers/api/v1/user"
	"gohub/app/middlewares"
	"net/http"
)

// RegisterAPIRoutes 注册 API 路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("gohub/api/v1")
	// ping
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点
	v1.Use(middlewares.LimitIp("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 用户授权
		authGroup.Use(middlewares.LimitIp("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断手机是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password/reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password/reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}

		// 发送验证码
		vcc := new(auth.VerifyCodeController)
		codeGroup := v1.Group("/captcha")
		{
			// 图片验证码，需要加限流
			codeGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			codeGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			codeGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("50-H"), vcc.SendUsingEmail)
		}

		// 用户
		uc := new(userCtl.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		userGroup := v1.Group("/users")
		{
			userGroup.GET("/list", middlewares.AuthJWT(), uc.Index)
			userGroup.PUT("/update", middlewares.AuthJWT(), uc.UpdateProfile)
			userGroup.PUT("/update/email", middlewares.AuthJWT(), uc.UpdateEmail)
			userGroup.PUT("/update/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			userGroup.PUT("/update/password", middlewares.AuthJWT(), uc.UpdatePassword)
			userGroup.PUT("/update/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
		}

		// 分类
		cgc := new(categoryCtl.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("/list", cgc.Index)
			cgcGroup.POST("/add", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/update/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/delete/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		// topic话题
		tpc := new(topicCtl.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("/list", tpc.Index)
			tpcGroup.GET("/show/:id", tpc.Show)
			tpcGroup.POST("/add", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/update/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/delete/:id", middlewares.AuthJWT(), tpc.Delete)
		}

		// link友情链接
		lkc := new(linkCtl.LinksController)
		lcGroup := v1.Group("/topics")
		{
			lcGroup.GET("/list", lkc.Index)

		}
	}

}
