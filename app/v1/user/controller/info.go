package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/action/UserInfoAction"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Any("get", info_get)
}

func info_get(c *gin.Context) {
	uid := c.PostForm("uid")
	user := UserInfoAction.App_userinfo(uid)
	RET.Success(c, 0, user, nil)
}
