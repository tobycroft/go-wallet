package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/balance/action/BalanceAction"
	"main.go/app/v1/balance/model/TransferRecordModel"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func RefundController(route *gin.RouterGroup) {

	route.Any("remote", refund_remote)
}

func refund_remote(c *gin.Context) {
	remote, ok := Input.Post("remote", c, false)
	if !ok {
		return
	}
	if remote != app_conf.RemoteKey {
		RET.Fail(c, 401, nil, nil)
		return
	}
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	var tr TransferRecordModel.Interface
	tr.Db = tuuz.Db()
	data := tr.Api_find(id)
	if len(data) > 0 {
		sign := BalanceAction.Transfer_sign(data["uid"], data["to_address"], data["amount"])
		if sign == Calc.Any2String(data["sign"]) {
			db := tuuz.Db()
			db.Begin()
			var bal BalanceAction.Interface
			bal.Db = db
			err := bal.App_single_balance(data["uid"], data["cid"], 13, data["order_id"], Calc.ToDecimal(data["amount"]).Abs(), data["to_address"], "提现退款", "")
			if err != nil {
				db.Rollback()
				RET.Fail(c, 500, nil, err.Error())
				return
			}
			if !tr.Api_update_status(data["id"], -2) {
				db.Rollback()
				RET.Fail(c, 500, nil, "无法修改状态")
				return
			}
			RET.Success(c, 0, nil, nil)
		} else {
			RET.Fail(c, 403, nil, "验证不通过")
		}
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
