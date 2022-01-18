package UserModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_user"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(pid, username, password, pass_notify, info, language, share interface{}) int64 {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"pid":         pid,
		"username":    username,
		"password":    password,
		"pass_notify": pass_notify,
		"info":        info,
		"language":    language,
		"share":       share,
	}
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return ret
	} else {
		return ret
	}
}

func Api_find_byUsername(username interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"username": username,
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

func Api_find(id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"id": id,
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

func Api_find_active(id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"id":     id,
		"active": true,
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

func Api_find_byPid(pid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"pid": pid,
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

func Api_find_byUsernameandPassword(username, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"username": username,
		"password": password,
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

func (self *Interface) Api_update_password(id, password interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"password": password,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
func (self *Interface) Api_update_username(id, username interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"username": username,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select_byPid(pid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"pid": pid,
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

func Api_select_id_inPids(pids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("id")
	db.WhereIn("pid", pids)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_idPid() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("id,pid")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count_byPid(pid interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"pid": pid,
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
