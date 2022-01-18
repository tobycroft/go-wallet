package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/balance/action/TypeAction"
	"main.go/app/v1/balance/model/BalanceRecordModel"
	"main.go/app/v1/wallet/model/UserAddressModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func RecordController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("all", record_all)
	route.Any("list", record_list)
	route.Any("in", record_in)
	route.Any("out", record_out)
	route.Any("pay", record_pay)

}

func record_all(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}
	Type, ok := Input.PostIn("type", c, []string{"all", "in", "out", "exchange"})
	if !ok {
		return
	}
	cid, ok := Input.PostInt("cid", c)
	if !ok {
		return
	}
	if !ok {
		return
	}
	var bal BalanceRecordModel.Interface
	bal.Db = tuuz.Db()
	in, out, exchange := TypeAction.App_select_in_out_exchange()

	datas := []gorose.Data{}
	switch Type {
	case "in":
		datas = bal.Api_select_inTypeAndCid(uid, cid, in, limit, page)
		break

	case "out":
		datas = bal.Api_select_inTypeAndCid(uid, cid, out, limit, page)
		break

	case "exchange":
		datas = bal.Api_select_inTypeAndCid(uid, cid, exchange, limit, page)
		break

	default:
		datas = bal.Api_select_byCid(uid, cid, limit, page)
		break
	}
	types, direction := TypeAction.App_map_type()
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	useraddr := ur.Api_find_byType(uid, "eth")
	for i, data := range datas {
		data["type_name"] = types[data["type"].(int64)]
		data["type_icon"] = direction[data["type"].(int64)]
		data["extra"] = useraddr["address"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func record_list(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}
	var bal BalanceRecordModel.Interface
	bal.Db = tuuz.Db()
	datas := bal.Api_select(uid, limit, page)
	types, direction := TypeAction.App_map_type()
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	useraddr := ur.Api_find_byType(uid, "eth")
	for i, data := range datas {
		data["type_name"] = types[data["type"].(int64)]
		data["type_icon"] = direction[data["type"].(int64)]
		data["extra"] = useraddr["address"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func record_in(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}

	var bal BalanceRecordModel.Interface
	bal.Db = tuuz.Db()
	in, _, _ := TypeAction.App_select_in_out_exchange()
	datas := bal.Api_select_inType(uid, in, limit, page)
	types, direction := TypeAction.App_map_type()
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	useraddr := ur.Api_find_byType(uid, "eth")
	for i, data := range datas {
		data["type_name"] = types[data["type"].(int64)]
		data["type_icon"] = direction[data["type"].(int64)]
		data["extra"] = useraddr["address"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func record_out(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}

	var bal BalanceRecordModel.Interface
	bal.Db = tuuz.Db()
	_, out, _ := TypeAction.App_select_in_out_exchange()
	datas := bal.Api_select_inType(uid, out, limit, page)
	types, direction := TypeAction.App_map_type()
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	useraddr := ur.Api_find_byType(uid, "eth")
	for i, data := range datas {
		data["type_name"] = types[data["type"].(int64)]
		data["type_icon"] = direction[data["type"].(int64)]
		data["extra"] = useraddr["address"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}

func record_pay(c *gin.Context) {
	uid := c.PostForm("uid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}

	var bal BalanceRecordModel.Interface
	bal.Db = tuuz.Db()
	_, _, exchange := TypeAction.App_select_in_out_exchange()
	datas := bal.Api_select_inType(uid, exchange, limit, page)
	types, direction := TypeAction.App_map_type()
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	useraddr := ur.Api_find_byType(uid, "eth")
	for i, data := range datas {
		data["type_name"] = types[data["type"].(int64)]
		data["type_icon"] = direction[data["type"].(int64)]
		data["extra"] = useraddr["address"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
