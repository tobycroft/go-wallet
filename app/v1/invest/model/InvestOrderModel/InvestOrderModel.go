package InvestOrderModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

const table = "invest_order"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, pid, cid, mid, order_id, amount, from, to, tx_id, tx_compelete interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":          uid,
		"pid":          pid,
		"cid":          cid,
		"mid":          mid,
		"order_id":     order_id,
		"amount":       amount,
		"from":         from,
		"to":           to,
		"tx_id":        tx_id,
		"tx_compelete": tx_compelete,
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

func (self *Interface) Api_find(id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
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

func (self *Interface) Api_select_txCompelete(uid, tx_compelete interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid":          uid,
		"tx_compelete": tx_compelete,
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

func (self *Interface) Api_find_first(uid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.OrderBy("id asc")
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byUid(uid interface{}) []gorose.Data {
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

func (self *Interface) Api_select_byPid(pid interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"pid": pid,
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

func (self *Interface) Api_select() []gorose.Data {
	db := self.Db.Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byProgress(progress interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"progress": progress,
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

func (self *Interface) Api_sum_amount() interface{} {
	db := self.Db.Table(table)
	ret, err := db.Sum("amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return ret
	} else {
		return ret
	}
}

func (self *Interface) Api_count_byPid(pid interface{}) int64 {
	db := self.Db.Table(table)
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

func (self *Interface) Api_count_byUid(uid interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"uid": uid,
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

func (self *Interface) Api_select_outOfTime() []gorose.Data {
	db := self.Db.Table(table)
	db.Where("last_exec", "<", time.Now().Unix())
	db.Where("generation", "<", 12)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_update_generationAndLastExec(id, generation, last_exec interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"generation": generation,
		"last_exec":  last_exec,
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

func (self *Interface) Api_update_leftAmountAndDeployAmount(id, left_amount, deploy_amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"left_amount":   left_amount,
		"deploy_amount": deploy_amount,
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

func (self *Interface) Api_update_progress(id, progress interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"progress": progress,
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

func (self *Interface) Api_update_txId(id, tx_id interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"tx_id": tx_id,
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

func (self *Interface) Api_update_remark(id, remark interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"remark": remark,
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
