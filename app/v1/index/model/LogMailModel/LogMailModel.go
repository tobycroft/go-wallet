package LogMailModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_mail"

func Api_insert(ip, success, to, log interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"ip":      ip,
		"success": success,
		"to":      to,
		"log":     log,
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

func Api_count(ip interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"ip": ip,
	}
	db.Where(where)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count_today(ip interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"ip": ip,
	}
	db.Where(where)
	db.Where("date > current_date")
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count_60(ip interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"ip": ip,
	}
	db.Where(where)
	db.Where("date > FROM_UNIXTIME((UNIX_TIMESTAMP()-60))")
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}
