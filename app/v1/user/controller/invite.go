package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/RET"
)

func InviteController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("get", invite_share)
}

func invite_share(c *gin.Context) {
	uid := c.PostForm("uid")
	user := UserModel.Api_find(uid)
	share_url := SystemParamModel.Api_find_val("share_url")
	RET.Success(c, 0, share_url.(string)+"/login.html?code="+user["username"].(string), nil)
}
