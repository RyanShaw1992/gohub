package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
)

// SignupController注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	//初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	//解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {
		//解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		//打印错误信息
		fmt.Println(err.Error())
		//出错了，中断请求
		return
	}

	//表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	//errs返回长度等于0即验证通过，大于0即有错误发生
	if len(errs) > 0 {
		//验证失败，返回422状态码以及错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	//检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
