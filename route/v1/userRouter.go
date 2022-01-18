package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/controller"
	"main.go/common/BaseController"
)

func UserRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())

	controller.IndexController(route.Group("index"))
	controller.InfoController(route.Group("info"))
	controller.InviteController(route.Group("invite"))
}
