package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func IndexController(route *gin.RouterGroup) {

	route.Use(BaseController.CorsController())
	//route.Use(BaseController.LoginedController(), gin.Recovery())

}
