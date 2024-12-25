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
