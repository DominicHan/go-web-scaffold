package controllers

import (
	"go-web-scaffold/services"
	"github.com/gin-gonic/gin"
)

func UserController(router *gin.RouterGroup) {
	router.GET("/", services.UserService)
	router.GET("/info", services.UserInfoService)
	router.POST("/info", services.UserInfoService)
	router.GET("/follow/:username", services.UserFollowService)  // http://127.0.0.1:8005/user/follow/wangye
	router.GET("/id", services.UserIdService)  // http://127.0.0.1:8005/user/id?id=123

}