package ExchangeRecordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_exchange_record"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, from_cid, to_cid, raw_amount, fee, amount, remark interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":        uid,
		"from_cid":   from_cid,
		"to_cid":     to_cid,
		"raw_amount": raw_amount,
		"fee":        fee,
		"amount":     amount,
		"remark":     remark,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_select(uid interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byCid(uid, cid interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_find(uid, id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"id":  id,
	}
	db.Where(where)
	db.LockForUpdate()
	db.OrderBy("id desc")
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
