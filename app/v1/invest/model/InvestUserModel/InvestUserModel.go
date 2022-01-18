package InvestUserModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "invest_user"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, cid, auth interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":  uid,
		"cid":  cid,
		"auth": auth,
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

func (self *Interface) Api_select() []gorose.Data {
	db := self.Db.Table(table)
	db.Order("uid desc")
	db.LockForUpdate()
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_incr_amount(uid, amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Increment("amount", amount)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_incr_levelAmount(uid, level_amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Increment("level_amount", level_amount)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update(uid, cid, auth interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"cid":  cid,
		"auth": auth,
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

func (self *Interface) Api_update_levelAmount(uid, level_amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"level_amount": level_amount,
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

func (self *Interface) Api_update_level(uid, level interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"level": level,
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

func (self *Interface) Api_find_maxLevelAmount(uids []interface{}) interface{} {
	db := self.Db.Table(table)
	db.WhereIn("uid", uids)
	db.LockForUpdate()
	ret, err := db.Max("level_amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
