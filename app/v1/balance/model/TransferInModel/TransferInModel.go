package TransferInModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_transfer_in"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, cid, order_id, from_address, amount, transaction_id, block interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":            uid,
		"cid":            cid,
		"order_id":       order_id,
		"from_address":   from_address,
		"amount":         amount,
		"transaction_id": transaction_id,
		"block":          block,
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

func (self *Interface) Api_find(transaction_id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"transaction_id": transaction_id,
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

func (self *Interface) Api_find_max_uid(uid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
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

func (self *Interface) Api_find_today_uid(uid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.Where("date > CURRENT_DATE")
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_find_max_block(uid, cid interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Max("block")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		if ret == nil {
			return 0
		} else {
			return ret.(int64)
		}
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

func (self *Interface) Api_select_byUidAndCid(uid, cid interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"cid": cid,
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
