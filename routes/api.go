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
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
