package BalanceModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_balance"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_find(uid, cid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
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

func (self *Interface) Api_value(uid, cid interface{}) interface{} {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_insert(uid, cid, balance interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":     uid,
		"cid":     cid,
		"balance": balance,
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

func (self *Interface) Api_update(uid, cid, balance interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"balance": balance,
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

func (self *Interface) Api_incr(uid, cid, incr_balance interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Increment("balance", incr_balance)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
