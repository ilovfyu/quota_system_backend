package main

import (
	"github.com/gin-gonic/gin"
	"quota_system/app"
	"quota_system/common"
)

func init() {
	// 初始化配置
	common.InitViper()
	// 初始化日志
	common.InitLogger()
	// 初始化MYSQL
	common.InitMySQL()
	//// 初始化REDIS
	common.InitRedis()

}

func main() {

	engine := gin.Default()
	app.InitService(engine)

	engine.Run(":8088")
}
