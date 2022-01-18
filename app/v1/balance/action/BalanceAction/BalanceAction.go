package BalanceAction

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/balance/model/BalanceModel"
	"main.go/app/v1/balance/model/BalanceRecordModel"
	"main.go/app/v1/balance/model/TransferRecordModel"
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Date"
)

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) App_check_balance(uid, cid interface{}) decimal.Decimal {
	var balmodel BalanceModel.Interface
	balmodel.Db = self.Db
	userbalance := balmodel.Api_find(uid, cid)
	if len(userbalance) > 0 {
		return Calc.ToDecimal(userbalance["balance"])
	} else {
		balmodel.Api_insert(uid, cid, 0)
		return decimal.Zero
	}
}

func (self *Interface) App_single_balance(uid, cid, Type, order_id interface{}, amount decimal.Decimal, extra, remark1, remark2 interface{}) error {
	db := self.Db
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	db.Begin()
	balance_left := self.App_check_balance(uid, cid)
	if balance_left.Add(amount).LessThan(decimal.Zero) {
		db.Rollback()
		return errors.New("余额不足")
	}
	var balance BalanceModel.Interface
	balance.Db = db
	balance.Api_incr(uid, cid, amount)

	//插入变动数据
	var balancerecord BalanceRecordModel.Interface
	balancerecord.Db = db
	last_record := balancerecord.Api_find_last(uid, cid)
	after := decimal.Zero
	before := "0"
	if len(last_record) > 0 {
		after = Calc.Bc_add(last_record["after"], amount)
		before = last_record["after"].(string)
	} else {
		after = amount
	}
	if after.LessThan(decimal.Zero) {
		db.Rollback()
		return errors.New("余额记录不足")
	}
	if !balancerecord.Api_insert(uid, cid, Type, order_id, before, amount, after, extra, remark1, remark2) {
		db.Rollback()
		return errors.New("balance_record添加失败")
	}
	db.Commit()
	return nil
}

func (self *Interface) App_single_transfer(uid, cid, to_uid, order_id interface{}, amount decimal.Decimal, extra, remark1, remark2 interface{}) (err error) {
	db := self.Db
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	user := UserModel.Api_find(to_uid)
	if len(user) < 1 {
		err = errors.New("没有找到接收人")
		return
	}
	db.Begin()
	err = self.App_single_balance(uid, cid, 22, order_id, amount.Abs().Neg(), extra, "用户转账转出", remark2)
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
	transfer_fee := coin["transfer_fee"].(float64)
	transfer_time_limit := coin["transfer_time_limit"].(int64)
	transfer_limit := coin["transfer_limit"].(float64)
	transfer_per_limit := coin["transfer_per_limit"].(float64)
	var trm TransferRecordModel.Interface
	trm.Db = db
	todays_time := trm.Api_count_today(uid, Date.Today())
	if todays_time > transfer_time_limit {
		db.Rollback()
		err = errors.New("超过当日最大转出数量")
		return
	}
	todays_amount := trm.Api_sum_today(uid, Date.Today())
	if todays_amount.GreaterThan(Calc.ToDecimal(transfer_limit)) {
		db.Rollback()
		err = errors.New("超过当日最大转出数量")
		return
	}
	if decimal.NewFromFloat(transfer_per_limit).LessThan(amount) {
		db.Rollback()
		err = errors.New("超过单笔最大限制")
		return
	}
	fee := Calc.Bc_mul(transfer_fee, amount)
	after_amount := Calc.Bc_min(amount, fee)
	err = self.App_single_balance(to_uid, cid, 21, order_id, after_amount.Abs(), extra, "用户收款转入", remark2)
	if err != nil {
		db.Rollback()
		return
	}
	db.Commit()
	return
}
