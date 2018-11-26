package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeSerializer struct {
	c *gin.Context
}

func HomeService(c *gin.Context) {
	serializer := HomeSerializer{c}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})

}

type HomeResponse struct {
	Msg string `json:"Msg"`
}

func (h *HomeSerializer) Response() HomeResponse {
	home := HomeResponse{
		Msg: "这是首页",
	}
	return home
}