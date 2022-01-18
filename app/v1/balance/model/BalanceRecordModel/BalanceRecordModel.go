package BalanceRecordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_balance_record"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, cid, Type, order_id, before, amount, after, extra, remark1, remark2 interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":      uid,
		"cid":      cid,
		"type":     Type,
		"order_id": order_id,
		"before":   before,
		"amount":   amount,
		"after":    after,
		"extra":    extra,
		"remark1":  remark1,
		"remark2":  remark2,
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

func (self *Interface) Api_find_last(uid, cid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.OrderBy("id desc")
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
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

func (self *Interface) Api_select_byType(uid, Type interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid":  uid,
		"type": Type,
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

func (self *Interface) Api_select_inType(uid interface{}, Type []interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.WhereIn("type", Type)
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

func (self *Interface) Api_select_inTypeAndCid(uid interface{}, cid interface{}, Type []interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.WhereIn("type", Type)
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
