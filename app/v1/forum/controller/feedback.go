package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/app/v1/forum/model/ForumThreadReplyModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FeedbackController(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("list", feedback_list)
	route.Any("add", feedback_add)
	route.Any("get", feedback_get)

}

func feedback_list(c *gin.Context) {
	uid := c.PostForm("uid")
	fid := SystemParamModel.Api_find_val("default_feedback_fid")
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := ForumThreadModel.Api_select_byUidandInFids(uid, []interface{}{fid}, limit, page)
	RET.Success(c, 0, datas, nil)
}

func feedback_add(c *gin.Context) {
	uid := c.PostForm("uid")
	fid := SystemParamModel.Api_find_val("default_feedback_fid")
	title, ok := Input.Post("title", c, true)
	if !ok {
		return
	}
	tag, ok := Input.Post("tag", c, true)
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
	imgs, ok := Input.Post("imgs", c, true)
	if !ok {
		return
	}
	extra, ok := Input.Post("extra", c, true)
	if !ok {
		return
	}
	var thread ForumThreadModel.Interface
	thread.Db = tuuz.Db()
	tid := thread.Api_insert("feedback", fid, uid, title, tag, content, img, imgs, extra, false, true)
	if tid != 0 {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func feedback_get(c *gin.Context) {
	tid, ok := Input.PostInt("tid", c)
	if !ok {
		return
	}
	fid := SystemParamModel.Api_find_val("default_feedback_fid")
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
