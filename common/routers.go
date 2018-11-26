package common

import (
	"go-web-scaffold/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit(api *gin.RouterGroup) {
	controllers.HomeController(api.Group("/"))
	controllers.UserController(api.Group("/user"))
}
