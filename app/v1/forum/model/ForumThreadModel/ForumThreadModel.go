package ForumThreadModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "xrc_forum_thread"

func Api_find(id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byFidandId(fid, id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"id":  id,
		"fid": fid,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(fid interface{}, limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"fid": fid,
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

func Api_select_hot(limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
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

func Api_select_public(fid interface{}, limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"fid":       fid,
		"is_public": true,
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

func Api_select_byTypeandUid(Type interface{}, limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_select_byUidandInFids(uid interface{}, fids []interface{}, limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.WhereIn("fid", fids)
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

func Api_like(search_text string) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("title", "like", "%"+search_text+"%")
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_like_in(search_text string, fids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.WhereIn("fid", fids)
	db.Where("title", "like", "%"+search_text+"%")
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(Type, fid, uid, title, tag, content, img, imgs, extra, is_public, can_reply interface{}) int64 {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"type":      Type,
		"fid":       fid,
		"uid":       uid,
		"title":     title,
		"tag":       tag,
		"content":   content,
		"img":       img,
		"imgs":      imgs,
		"extra":     extra,
		"is_public": is_public,
		"can_reply": can_reply,
	}
	db.Data(data)
	ret, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return ret
	} else {
		return ret
	}
}

func (self *Interface) Api_update(uid, id, title, tag, content, img, imgs, extra, is_public, can_reply interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"id":  id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"title":     title,
		"tag":       tag,
		"content":   content,
		"img":       img,
		"imgs":      imgs,
		"extra":     extra,
		"is_public": is_public,
		"can_reply": can_reply,
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

func (self *Interface) Api_update_canUpdate(uid, id, can_reply interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"id":  id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"can_reply": can_reply,
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

func (self *Interface) Api_update_isPublic(uid, id, is_public interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
		"id":  id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"is_public": is_public,
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
