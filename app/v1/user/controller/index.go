package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func IndexController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Use(BaseController.LoginedController(), gin.Recovery())
}
