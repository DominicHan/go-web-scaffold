package services

import (
	"go-web-scaffold/models"
	"go-web-scaffold/tools"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Msg  string `json:"Msg"`
	Info string `json:"Info"`
}

func UserService(c *gin.Context) {
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"data": serializer.UserGet()})

}

func (h *UserSerializer) UserGet() UserResponse {
	user := UserResponse{
		Msg: "这是用户页",
	}
	return user
}

func UserInfoService(c *gin.Context) {
	serializer := UserSerializer{c}
	method := c.Request.Method
	switch {
	case method == "GET":
		c.JSON(http.StatusOK, gin.H{"data": serializer.UserInfoGet()})
	case method == "POST":
		c.JSON(http.StatusOK, gin.H{"data": serializer.UserInfoPost()})
	}
}

func (h *UserSerializer) UserInfoGet() UserResponse {
	userInfo := UserResponse{
		Msg:  "这是用户详情页 get",
		Info: "这是详情..",
	}
	return userInfo
}

func (h *UserSerializer) UserInfoPost() UserResponse {
	projectId := h.c.PostForm("project_id")
	projectName := h.c.DefaultPostForm("project_name", "pro")
	userInfo := UserResponse{
		Msg:  "这是用户详情页 post",
		Info: "这是详情.. project_id = " + projectId + " project_name = " + projectName,
	}
	return userInfo
}

func UserFollowService(c *gin.Context) {
	serializer := UserSerializer{c}
	method := c.Request.Method
	switch {
	case method == "GET":
		c.JSON(http.StatusOK, gin.H{"data": serializer.UserFollowGet()})
	}
}

func (h *UserSerializer) UserFollowGet() UserResponse {
	username := h.c.Param("username")
	userInfo := UserResponse{
		Msg:  "这是用户详情页 get",
		Info: "这是详情.. username = " + username,
	}
	return userInfo
}

func UserIdService(c *gin.Context) {
	serializer := UserSerializer{c}
	method := c.Request.Method
	switch {
	case method == "GET":
		c.JSON(http.StatusOK, gin.H{"data": serializer.UserIdGet()})
	}
}

func (h *UserSerializer) UserIdGet() UserResponse {
	id := h.c.Query("id")
	userInfo := UserResponse{
		Msg:  "这是用户详情页 get",
		Info: "这是详情.. id = " + id,
	}

	db := tools.GetDB()
	tx := db.Begin()
	userA := models.UserModel{
		Username: "AAAAAA",
		Num:      id,
		Hash:     rand.Uint64(),
	}
	//models.SaveOne()
	tx.Save(&userA)
	tx.Commit()

	return userInfo
}
