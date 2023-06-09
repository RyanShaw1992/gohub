package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	//new一个Gin Engine实例
	router := gin.New()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":3000")
	if err != nil {
		//错误处理，端口被占用或其他错误
		fmt.Println(err.Error())
	}
}
