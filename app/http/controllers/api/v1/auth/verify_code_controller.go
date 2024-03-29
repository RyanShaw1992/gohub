package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

// VerifyCodeController用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, _, err := captcha.NewCaptcha().GenerateCaptcha()
	//记录错误日志，因为验证码是用户的入口，出错时应该记录error等级的日志
	logger.LogIf(err)
	//返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
