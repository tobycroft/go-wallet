package InvestRecordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "invest_record"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(Type, uid, oid, before, amount, after interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"type":   Type,
		"uid":    uid,
		"oid":    oid,
		"before": before,
		"amount": amount,
		"after":  after,
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

func (self *Interface) Api_select(uid interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_find_last(Type, uid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"type": Type,
		"uid":  uid,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
