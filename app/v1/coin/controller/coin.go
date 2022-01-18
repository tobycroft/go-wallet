package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func CoinController(route *gin.RouterGroup) {
	route.Use(BaseController.CorsController())

	route.Any("list", coin_list)
	route.Any("get", coin_get)
}

func coin_list(c *gin.Context) {
	coins := CoinModel.Api_select()
	for i, coin := range coins {
		coins[i] = coin
	}
	RET.Success(c, 0, coins, nil)
}

func coin_get(c *gin.Context) {
	cid, ok := Input.PostInt("cid", c)
	if !ok {
		return
	}
	coin := CoinModel.Api_find(cid)
	if len(coin) > 1 {
		RET.Success(c, 0, coin, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
