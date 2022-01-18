package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/controller"
	"main.go/common/BaseController"
)

func IndexRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())

	controller.IndexController(route.Group("index"))
	controller.LoginController(route.Group("login"))
	controller.RoundingController(route.Group("rounding"))
}
