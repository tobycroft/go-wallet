package TransferRecordModel

import (
	"github.com/shopspring/decimal"
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

const table = "xrc_transfer_record"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, cid, to_address, order_id, raw_amount, fee, amount, approved, status, sign interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":        uid,
		"cid":        cid,
		"to_address": to_address,
		"order_id":   order_id,
		"raw_amount": raw_amount,
		"fee":        fee,
		"amount":     amount,
		"approved":   approved,
		"status":     status,
		"sign":       sign,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_select(uid interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byStatus(status interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"status": status,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_find(id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_update_status(id, status interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"status": status,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_sign(id, sign interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"sign": sign,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_remark(id, remark interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"remark": remark,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_count_today(uid, gt_date interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.Where("date", ">", gt_date)
	db.LockForUpdate()
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func (self *Interface) Api_sum_today(uid, gt_date interface{}) decimal.Decimal {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.Where("date", ">", gt_date)
	db.LockForUpdate()
	ret, err := db.Sum("amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return decimal.Zero
	} else {
		if ret == nil {
			return decimal.Zero
		} else {
			return Calc.ToDecimal(ret)
		}
	}
}
