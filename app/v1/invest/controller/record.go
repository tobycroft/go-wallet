package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func RecordController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

}
