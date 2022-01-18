package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/exchange/controller"
	"main.go/common/BaseController"
)

func ExchangeRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())

	controller.PaymentController(route.Group("payment"))
	controller.RecordController(route.Group("record"))
}
