package ExchangeModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_exchange"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.GroupBy("from_cid")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_byCid(from_cid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("from_cid", "=", from_cid)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byFull(from_cid, to_cid, max_amount, min_amount interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("from_cid", "=", from_cid)
	db.Where("to_cid", "=", to_cid)
	db.Where("max_amount", ">=", max_amount)
	db.Where("min_amount", "<=", min_amount)
	db.Where("is_open", "=", true)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
