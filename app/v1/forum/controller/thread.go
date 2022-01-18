package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/app/v1/forum/model/ForumThreadReplyModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func ThreadController(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	route.Any("list", thread_list)
	route.Any("get", thread_get)
}

func thread_list(c *gin.Context) {
	fid, ok := Input.PostInt("fid", c)
	if !ok {
		return
	}
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := ForumThreadModel.Api_select(fid, limit, page)
	RET.Success(c, 0, datas, nil)
}

func thread_get(c *gin.Context) {
	tid, ok := Input.PostInt("tid", c)
	if !ok {
		return
	}
	fid, ok := Input.PostInt("fid", c)
	if !ok {
		return
	}
	thread := ForumThreadModel.Api_find_byFidandId(fid, tid)
	reply := ForumThreadReplyModel.Api_select_byTid(tid)
	if len(thread) > 0 {
		RET.Success(c, 0, map[string]interface{}{
			"thread": thread,
			"reply":  reply,
		}, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
