package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/wallet/model/UserAddressModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/RET"
)

func ReceiveController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("get", receive_get)
}

func receive_get(c *gin.Context) {
	uid := c.PostForm("uid")
	var ur UserAddressModel.Interface
	ur.Db = tuuz.Db()
	datas := ur.Api_select_InType(uid, []interface{}{
		"eth",
		"trc",
	})
	temp := map[string]interface{}{}
	for _, data := range datas {
		temp[data["type"].(string)] = data["address"]
	}
	RET.Success(c, 0, temp, nil)
}
