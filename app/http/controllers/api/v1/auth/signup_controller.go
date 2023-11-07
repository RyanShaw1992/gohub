package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"
)

// SignupController注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	//请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

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

	//检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}