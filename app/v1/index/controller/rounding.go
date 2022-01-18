package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/RoundingModel"
	"main.go/tuuz/RET"
)

func RoundingController(route *gin.RouterGroup) {
	route.Any("index", rounding_index)
}

func rounding_index(c *gin.Context) {
	data := RoundingModel.Api_select()
	RET.Success(c, 0, data, nil)
}
