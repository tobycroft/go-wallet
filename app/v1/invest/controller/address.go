package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/RET"
)

func AddressController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("get", address_get)
}

func address_get(c *gin.Context) {
	trx_address := SystemParamModel.Api_find_val("trx_address")
	eth_address := SystemParamModel.Api_find_val("eth_address")
	RET.Success(c, 0, map[string]interface{}{
		"trx_address": trx_address,
		"eth_address": eth_address,
	}, nil)
}
