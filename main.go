package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	//加载config目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	//配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载.env文件，如--env=testing加载的是.env.testing文件")
	flag.Parse()
	config.InitConfig(env)

	//初始化Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	//new一个Gin Engine实例
	router := gin.New()

	//初始化DB
	bootstrap.SetupDB()

	//初始化Redis
	bootstrap.SetupRedis()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		//错误处理，端口被占用了或其他错误
		fmt.Println(err.Error())
	}
}
