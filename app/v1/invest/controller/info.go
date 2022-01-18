package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/invest/model/InvestOrderModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("get", info_get)
}

func info_get(c *gin.Context) {
	invest_total_num := SystemParamModel.Api_find_val("invest_total_num")
	var iv InvestOrderModel.Interface
	iv.Db = tuuz.Db()
	amount := iv.Api_sum_amount()
	RET.Success(c, 0, map[string]interface{}{
		"total": invest_total_num,
		"has":   amount,
	}, nil)
}
