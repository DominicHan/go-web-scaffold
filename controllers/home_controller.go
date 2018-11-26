package controllers

import (
	"github.com/gin-gonic/gin"
	"go-web-scaffold/services"
)

func HomeController(router *gin.RouterGroup) {
	router.GET("/", services.HomeService)
}
