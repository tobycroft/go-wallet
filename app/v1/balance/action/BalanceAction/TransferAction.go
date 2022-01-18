package BalanceAction

import (
	"errors"
	"github.com/shopspring/decimal"
	"main.go/app/v1/balance/model/TransferRecordModel"
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Date"
)

func (self *Interface) App_transfer_out(uid, cid, to_address, order_id interface{}, amount decimal.Decimal, extra, remark1, remark2 interface{}) (err error) {
	db := self.Db
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	db.Begin()
	err = self.App_single_balance(uid, cid, 11, order_id, amount.Abs().Neg(), extra, remark1, remark2)
	if err != nil {
		db.Rollback()
		return
	}
	coin := CoinModel.Api_find(cid)
	if len(coin) < 1 {
		db.Rollback()
		err = errors.New("未找到币种")
		return
	}
	transfer_out_fee := coin["transfer_out_fee"].(float64)
	transfer_out_time_limit := coin["transfer_out_time_limit"].(int64)
	transfer_out_limit := coin["transfer_out_limit"].(float64)
	transfer_out_per_limit := coin["transfer_out_per_limit"].(float64)
	var trm TransferRecordModel.Interface
	trm.Db = db
	todays_time := trm.Api_count_today(uid, Date.Today())
	if todays_time > transfer_out_time_limit {
		db.Rollback()
		err = errors.New("超过当日最大转出数量")
		return
	}
	todays_amount := trm.Api_sum_today(uid, Date.Today())
	if todays_amount.GreaterThan(Calc.ToDecimal(transfer_out_limit)) {
		db.Rollback()
		err = errors.New("超过当日最大转出数量")
		return
	}
	if decimal.NewFromFloat(transfer_out_per_limit).LessThan(amount) {
		db.Rollback()
		err = errors.New("超过单笔最大限制")
		return
	}
	fee := Calc.Bc_mul(transfer_out_fee, amount)
	after_amount := Calc.Bc_min(amount, fee)
	sign := Transfer_sign(uid, to_address, amount)
	if !trm.Api_insert(uid, cid, to_address, order_id, amount, fee, after_amount, 0, 0, sign) {
		db.Rollback()
		err = errors.New("转出记录插入失败")
		return
	}
	db.Commit()
	return
}

func Transfer_sign(uid, to_address, amount interface{}) string {
	sign := Calc.Md5(Calc.Any2String(uid) + "_" + Calc.Any2String(to_address) + "_" + Calc.Any2String(amount))
	return sign
}
