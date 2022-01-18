package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/invest/model/InvestOrderModel"
	"main.go/app/v1/invest/model/InvestUserModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GroupController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("amount", group_amount)
	route.Any("list", group_list)
	route.Any("down", group_down)
	route.Any("get", group_get)
	route.Any("my", group_my)
}


func group_amount(c *gin.Context) {
	//uid := c.PostForm("uid")
	var in InvestUserModel.Interface
	in.Db = tuuz.Db()
	//ud := in.Api_find(uid)
	//RET.Success(c, 0, ud["level_amount"], nil)
}

func group_list(c *gin.Context) {
	uid := c.PostForm("uid")
	users := UserModel.Api_select_byPid(uid)
	var inv InvestOrderModel.Interface
	inv.Db = tuuz.Db()
	var in InvestUserModel.Interface
	in.Db = tuuz.Db()
	for i, user := range users {
		//ud := in.Api_find(user["id"])
		//level := ud["level"]
		//if level == nil {
		//	level = 0
		//}
		//delete(user, "password")
		//delete(user, "pass_notify")
		//user["team_num"] = ud["downer"]
		//user["team_avail"] = inv.Api_count_byUid(user["id"]) > 0
		//user["level"] = level
		//user["amount"] = ud["level_amount"]
		users[i] = user
	}
	RET.Success(c, 0, users, nil)
}

func group_down(c *gin.Context) {
	uid, ok := Input.PostInt("user_id", c)
	if !ok {
		return
	}
	users := UserModel.Api_select_byPid(uid)
	var inv InvestOrderModel.Interface
	inv.Db = tuuz.Db()
	var in InvestUserModel.Interface
	in.Db = tuuz.Db()
	for i, user := range users {
		//ud := in.Api_find(user["id"])
		//level := ud["level"]
		//if level == nil {
		//	level = 0
		//}
		//delete(user, "password")
		//delete(user, "pass_notify")
		//user["team_num"] = ud["downer"]
		//user["team_avail"] = inv.Api_count_byUid(user["id"]) > 0
		//user["level"] = level
		//user["amount"] = ud["level_amount"]
		users[i] = user
	}
	RET.Success(c, 0, users, nil)
}

func group_get(c *gin.Context) {
	user_id, ok := Input.PostInt("user_id", c)
	if !ok {
		return
	}
	user := UserModel.Api_find(user_id)
	if len(user) > 0 {
		puser := UserModel.Api_find(user["pid"])
		if len(puser) < 1 {
			RET.Fail(c, 404, nil, nil)
			return
		}
		//num := UserModel.Api_count_byPid(user["pid"])
		var inv InvestOrderModel.Interface
		inv.Db = tuuz.Db()
		var in InvestUserModel.Interface
		in.Db = tuuz.Db()
		//ud := in.Api_find(user["id"])
		//level := ud["level"]
		//if level == nil {
		//	level = 0
		//}
		//data := map[string]interface{}{
		//	"id":         puser["pid"],
		//	"name":       puser["username"],
		//	"username":   puser["username"],
		//	"head_img":   puser["head_img"],
		//	"team_num":   num,
		//	"team_avail": inv.Api_count_byUid(user["id"]) > 0,
		//	"amount":     ud["level_amount"],
		//	"level":      level,
		//}
		//RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func group_my(c *gin.Context) {
	user_id, ok := Input.PostInt("user_id", c)
	if !ok {
		return
	}
	user := UserModel.Api_find(user_id)
	if len(user) > 0 {
		//num := UserModel.Api_count_byPid(user_id)
		//var inv InvestOrderModel.Interface
		//inv.Db = tuuz.Db()
		//var in InvestUserModel.Interface
		//in.Db = tuuz.Db()
		//ud := in.Api_find(user["id"])
		//level := ud["level"]
		//if level == nil {
		//	level = 0
		//}
		//data := map[string]interface{}{
		//	"id":         user["id"],
		//	"name":       user["username"],
		//	"username":   user["username"],
		//	"head_img":   user["head_img"],
		//	"team_num":   num,
		//	"team_avail": inv.Api_count_byUid(user["id"]) > 0,
		//	"amount":     ud["level_amount"],
		//	"level":      level,
		//}
		//RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
