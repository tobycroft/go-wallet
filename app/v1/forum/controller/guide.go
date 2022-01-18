package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumModel"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GuideController(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	route.Any("upper", guide_upper)
	route.Any("downer", guide_downer)
	route.Any("get", thread_get)
}

func guide_upper(c *gin.Context) {
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	help_hot_fid := SystemParamModel.Api_find_val("guide_upper")
	if help_hot_fid != nil {
		forum := ForumModel.Api_find(help_hot_fid)
		if len(forum) > 0 {
			datas := ForumThreadModel.Api_select(help_hot_fid, limit, page)
			RET.Success(c, 0, datas, nil)
		} else {
			RET.Fail(c, 404, nil, "未找到板块")
		}
	} else {
		RET.Fail(c, 400, nil, "未设定热门帮助板块")
	}
}

func guide_downer(c *gin.Context) {
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	help_hot_fid := SystemParamModel.Api_find_val("guide_downer")
	if help_hot_fid != nil {
		forum := ForumModel.Api_find(help_hot_fid)
		if len(forum) > 0 {
			datas := ForumThreadModel.Api_select(help_hot_fid, limit, page)
			RET.Success(c, 0, datas, nil)
		} else {
			RET.Fail(c, 404, nil, "未找到板块")
		}
	} else {
		RET.Fail(c, 400, nil, "未设定热门帮助板块")
	}
}
