package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/model/ForumModel"
	"main.go/app/v1/forum/model/ForumThreadModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Use(BaseController.CorsController())

	route.Any("hot", index_hot)
	route.Any("list", index_list)
}

func index_hot(c *gin.Context) {
	datas := ForumThreadModel.Api_select_hot(5, 1)
	RET.Success(c, 0, datas, nil)
}

func index_list(c *gin.Context) {
	datas := ForumModel.Api_select_isShow()
	RET.Success(c, 0, datas, nil)
}
