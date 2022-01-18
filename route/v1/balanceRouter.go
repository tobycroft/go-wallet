package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/balance/controller"
	"main.go/common/BaseController"
)

func BalanceRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())

	controller.TransferController(route.Group("transfer"))
	controller.ReceiveController(route.Group("receive"))
	controller.RecordController(route.Group("record"))
	controller.RefundController(route.Group("refund"))
}
