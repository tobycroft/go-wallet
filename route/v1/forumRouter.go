package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/forum/controller"
	"main.go/common/BaseController"
)

func ForumRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())

	controller.FeedbackController(route.Group("feedback"))
	controller.GuideController(route.Group("guide"))
	controller.UpdateController(route.Group("update"))

	controller.IndexController(route.Group("index"))
	controller.ThreadController(route.Group("thread"))
	controller.ReplyController(route.Group("reply"))
	controller.BroadcastController(route.Group("broadcast"))
}
