package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controlls/api/v1/auth"
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
		// 用户授权
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断手机是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
		}

		// 发送验证码
		vcc := new(auth.VerifyCodeController)
		codeGroup := v1.Group("/captcha")
		{
			// 图片验证码，需要加限流
			codeGroup.POST("/verify/captcha", vcc.ShowCaptcha)
			codeGroup.POST("/verify/phone", vcc.SendUsingPhone)
			codeGroup.POST("/verify/email", vcc.SendUsingEmail)
		}

	}

}
