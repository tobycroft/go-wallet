package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"main.go/app/v1/balance/action/BalanceAction"
	"main.go/app/v1/coin/action/CoinAction"
	"main.go/app/v1/exchange/model/ExchangeModel"
	"main.go/app/v1/exchange/model/ExchangeRecordModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func PaymentController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("list", payment_list)
	route.Any("select", payment_select)
	route.Any("buy", payment_buy)

}

func payment_list(c *gin.Context) {
	datas := ExchangeModel.Api_select()
	coininfos := CoinAction.App_coin()
	for i, data := range datas {
		data["coin_info"] = coininfos[data["from_cid"].(int64)]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func payment_select(c *gin.Context) {
	from_cid, ok := Input.PostInt("from_cid", c)
	if !ok {
		return
	}
	coininfos := CoinAction.App_coin()
	datas := ExchangeModel.Api_select_byCid(from_cid)
	for i, data := range datas {
		data["coin_info"] = coininfos[data["to_cid"].(int64)]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func payment_buy(c *gin.Context) {
	uid := c.PostForm("uid")
	from_cid, ok := Input.PostInt64("from_cid", c)
	if !ok {
		return
	}
	to_cid, ok := Input.PostInt64("to_cid", c)
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
	if amount.LessThan(decimal.Zero) {
		RET.Fail(c, 406, nil, "闪兑金额应该大于0")
		return
	}
	user := UserModel.Api_find(uid)
	if user["password"] != Calc.Md5(password) {
		RET.Fail(c, 401, nil, "支付密码错误")
		return
	}
	coins := CoinAction.App_coin()
	from_coin := coins[from_cid]
	to_coin := coins[to_cid]
	if len(from_coin) < 1 || len(to_coin) < 1 {
		RET.Fail(c, 404, nil, "币种没有准备好")
		return
	}
	data := ExchangeModel.Api_find_byFull(from_cid, to_cid, amount, amount)
	if len(data) > 0 {
		db := tuuz.Db()
		var bal BalanceAction.Interface
		bal.Db = db
		order_id := Calc.GenerateOrderId()
		err := bal.App_single_balance(uid, from_cid, 102, order_id, amount.Abs().Neg(), "闪兑兑出", "闪兑交易兑出", "")
		if err != nil {
			db.Rollback()
			RET.Fail(c, 300, nil, err.Error())
			return
		}
		fee := Calc.ToDecimal(data["fee"]).Mul(amount).Abs()
		from_price := Calc.ToDecimal(from_coin["price"])
		to_price := Calc.ToDecimal(to_coin["price"])
		//扣除手续费
		from_amount := amount.Sub(fee)
		//领到的钱等于原币*源价格/目标价格=目标币
		get_amount := from_amount.Mul(from_price).Div(to_price)
		err = bal.App_single_balance(uid, to_cid, 101, order_id, get_amount, "闪兑兑入", "闪兑交易兑入,手续费:"+fee.String(), "")
		if err != nil {
			db.Rollback()
			RET.Fail(c, 300, nil, err.Error())
			return
		}

		//插入闪兑记录
		var exc ExchangeRecordModel.Interface
		exc.Db = db
		if !exc.Api_insert(uid, from_cid, to_cid, amount, fee, get_amount, "") {
			db.Rollback()
			RET.Fail(c, 500, nil, "闪兑记录插入失败")
			return
		}
		db.Commit()
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
