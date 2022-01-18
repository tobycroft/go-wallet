package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/invest/model/InvestModeModel"
	"main.go/tuuz/RET"
)

func ModeController(route *gin.RouterGroup) {

	route.Any("get", mode_get)
}

func mode_get(c *gin.Context) {
	mode := InvestModeModel.Api_select()
	RET.Success(c, 0, mode, nil)
}
