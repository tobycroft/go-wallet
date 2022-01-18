package VerifyModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "verify"

func Api_insert(username, code interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"username": username,
		"code":     code,
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

func Api_find(username interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	whehe := map[string]interface{}{
		"username": username,
	}
	db.Where(whehe)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_today(username, code interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	whehe := map[string]interface{}{
		"username": username,
		"code":     code,
	}
	db.Where(whehe)
	db.Where("date > FROM_UNIXTIME((UNIX_TIMESTAMP()-86400))")
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_delete(username interface{}) bool {
	db := tuuz.Db().Table(table)
	whehe := map[string]interface{}{
		"username": username,
	}
	db.Where(whehe)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
