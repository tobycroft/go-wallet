package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumModel"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/app/v1/forum/model/ForumThreadReplyModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func ReplyController(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", reply_add)
}

func reply_add(c *gin.Context) {
	uid := c.PostForm("uid")
	tid, ok := Input.PostInt("tid", c)
	if !ok {
		return
	}
	content, ok := Input.Post("content", c, true)
	if !ok {
		return
	}
	img, ok := Input.Post("img", c, true)
	if !ok {
		return
	}
	thread := ForumThreadModel.Api_find(tid)
	if len(thread) < 1 {
		RET.Fail(c, 403, nil, "帖子不存在")
		return
	}
	if thread["can_reply"].(int64) != 1 {
		RET.Fail(c, 403, nil, "帖子不允许回复")
		return
	}
	forum := ForumModel.Api_find(thread["fid"])
	if len(forum) < 1 {
		RET.Fail(c, 403, nil, "板块不存在")
		return
	}
	if forum["is_private"].(int64) != 1 {
		RET.Fail(c, 403, nil, "私有板块不允许回复")
		return
	}
	if ForumThreadReplyModel.Api_insert(tid, uid, content, img) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
