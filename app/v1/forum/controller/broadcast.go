package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumModel"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BroadcastController(route *gin.RouterGroup) {
	route.Use(BaseController.CorsController())

	route.Any("list", broadcast_list)
	route.Any("get", broadcast_get)
}

func broadcast_list(c *gin.Context) {
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	broadcast_fid := SystemParamModel.Api_find_val("broadcast_fid")
	if broadcast_fid != nil {
		forum := ForumModel.Api_find(broadcast_fid)
		if len(forum) > 0 {
			datas := ForumThreadModel.Api_select(broadcast_fid, limit, page)
			RET.Success(c, 0, datas, nil)
		} else {
			RET.Fail(c, 404, nil, "未找到板块")
		}
	} else {
		RET.Fail(c, 400, nil, "broadcast_fid")
	}
}

func broadcast_get(c *gin.Context) {
	update_log_fid := SystemParamModel.Api_find_val("update_log_fid")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	threads := ForumThreadModel.Api_find_byFidandId(update_log_fid, id)
	RET.Success(c, 0, threads, nil)
}
