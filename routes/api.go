package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

// 注册网页相关路由
func RegisterApiRoutes(r *gin.Engine) {
	//V1版本的路由
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			//判断手机是否已经注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			//判断Email是否已经注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码，需要加限流
			authGroup.POST("verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}
