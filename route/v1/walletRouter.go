package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/wallet/controller"
	"main.go/common/BaseController"
)

func WalletRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	route.Use(BaseController.CorsController())
	controller.AddressController(route.Group("address"))
}
