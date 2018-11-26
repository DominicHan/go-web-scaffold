package main

import (
	"go-web-scaffold/common"
	"go-web-scaffold/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	tools.ConfigInit() // 配置文件初始化

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例
	api := router.Group("/")

	db := tools.DBInit()       // 数据库初始化
	common.DBMigrate(db, true) // 数据库同步 true开启
	defer db.Close()           // 无引用 关闭db连接

	common.RouterInit(api) // 路由初始化

	router.Run(":8005") // 0.0.0.0:8005
}
