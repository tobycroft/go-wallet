package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/coin/action/CoinAction"
	"main.go/app/v1/exchange/model/ExchangeRecordModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func RecordController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("list", record_list)
}

func record_list(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	var exc ExchangeRecordModel.Interface
	exc.Db = tuuz.Db()
	datas := exc.Api_select(uid, limit, page)
	coins := CoinAction.App_coin()
	for i, data := range datas {
		data["from_coin_info"] = coins[data["from_cid"].(int64)]
		data["to_coin_info"] = coins[data["to_cid"].(int64)]
		data["raw_amount"], _ = Calc.String2Float64(data["raw_amount"].(string))
		data["fee"], _ = Calc.String2Float64(data["fee"].(string))
		data["amount"], _ = Calc.String2Float64(data["amount"].(string))
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
