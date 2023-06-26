package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册网页相关路由
func RegisterApiRoutes(r *gin.Engine) {
	//V1版本的路由
	v1 := r.Group("/v1")
	{
		//注册一个路由
		v1.GET("/", func(c *gin.Context) {
			//以json格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
