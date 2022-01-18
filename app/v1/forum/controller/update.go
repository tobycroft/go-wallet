package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func UpdateController(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	route.Any("current_version", update_current_version)
	route.Any("list", update_list)
	route.Any("get", update_get)

}

func update_list(c *gin.Context) {
	update_log_fid := SystemParamModel.Api_find_val("update_log_fid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	threads := ForumThreadModel.Api_select(update_log_fid, limit, page)
	RET.Success(c, 0, threads, nil)
}

func update_get(c *gin.Context) {
	update_log_fid := SystemParamModel.Api_find_val("update_log_fid")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	threads := ForumThreadModel.Api_find_byFidandId(update_log_fid, id)
	RET.Success(c, 0, threads, nil)
}

func update_current_version(c *gin.Context) {
	current_version := SystemParamModel.Api_find_val("current_version")
	RET.Success(c, 0, current_version, nil)
}
