package bootstrap

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/redis"
)

func SetupRedis() {
	//建立redis连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
