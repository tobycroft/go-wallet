package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/balance/action/BalanceAction"
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/app/v1/wallet/model/UserAddressModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"main.go/tuuz/Vali"
)

func TransferController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("eth", transfer_eth)
	route.Any("out", transfer_out)
	route.Any("auto", transfer_auto)
}

func transfer_eth(c *gin.Context) {
	uid := c.PostForm("uid")
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	amount, ok := Input.PostDecimal("amount", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	cid, ok := Input.PostInt("cid", c)
	if !ok {
		return
	}
	user := UserModel.Api_find(uid)
	if Calc.Md5(password) != user["password"].(string) {
		RET.Fail(c, 403, nil, "密码不正确")
		return
	}
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	to_user := ur.Api_find_address(address)
	if len(to_user) < 1 {
		RET.Fail(c, 404, nil, "未能找到用户")
		return
	} else {
		coin := CoinModel.Api_find(cid)
		if len(coin) < 1 {
			RET.Fail(c, 404, nil, "未找到币种")
			return
		}
		var balanceaction BalanceAction.Interface
		db := tuuz.Db()
		balanceaction.Db = db
		//todo:在这里加入一个判断币种是否可用的程序
		err := balanceaction.App_single_transfer(uid, coin["id"], to_user["uid"], nil, amount, "", "ETH用户转账", "")
		if err != nil {
			RET.Fail(c, 300, err.Error(), err.Error())
		} else {
			RET.Success(c, 0, nil, nil)
		}
	}
}

func transfer_out(c *gin.Context) {
	uid := c.PostForm("uid")
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	amount, ok := Input.PostDecimal("amount", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	cid, ok := Input.PostInt("cid", c)
	if !ok {
		return
	}
	user := UserModel.Api_find(uid)
	if Calc.Md5(password) != user["password"].(string) {
		RET.Fail(c, 403, nil, "密码不正确")
		return
	}
	err := Vali.Length(address, 34, 34)
	if err != nil {
		RET.Fail(c, 400, nil, "地址不正确")
		return
	}
	coin := CoinModel.Api_find(cid)
	if len(coin) < 1 {
		RET.Fail(c, 404, nil, "未找到币种")
		return
	}
	var balanceaction BalanceAction.Interface
	db := tuuz.Db()
	balanceaction.Db = db
	//todo:在这里加入一个判断币种是否可用的程序
	err = balanceaction.App_transfer_out(uid, coin["id"], address, nil, amount, nil, "用户提现", nil)
	if err != nil {
		RET.Fail(c, 300, err.Error(), err.Error())
	} else {
		RET.Success(c, 0, nil, nil)
	}
}

func transfer_auto(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	to_user := ur.Api_find_address(address)
	if len(to_user) < 1 {
		transfer_out(c)
	} else {
		transfer_eth(c)
	}
}
