package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/app/v1/invest/model/InvestUserModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("get", user_get)
}

func user_get(c *gin.Context) {
	uid := c.PostForm("uid")
	Type, ok := Input.PostIn("type", c, []string{"eth", "trc"})
	if !ok {
		return
	}
	contract, ok := Input.Post("contract", c, false)
	if !ok {
		return
	}
	coin := CoinModel.Api_find_byTypeAndContract(Type, contract)
	if len(coin) < 1 {
		RET.Fail(c, 404, nil, "未找到coin")
		return
	}
	var iu InvestUserModel.Interface
	iu.Db = tuuz.Db()
	investuser := iu.Api_find(uid, coin["id"])
	if len(investuser) > 0 {
		RET.Success(c, 0, 1, nil)
	} else {
		RET.Success(c, 0, 0, nil)
	}
}
