package route

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}
		balance := version1.Group("balance")
		{
			v1.BalanceRouter(balance)
		}
		coin := version1.Group("coin")
		{
			v1.CoinRouter(coin)
		}
		exchange := version1.Group("exchange")
		{
			v1.ExchangeRouter(exchange)
		}
		invest := version1.Group("invest")
		{
			v1.InvestRouter(invest)
		}
		wallet := version1.Group("wallet")
		{
			v1.WalletRouter(wallet)
		}
		user := version1.Group("user")
		{
			v1.UserRouter(user)
		}
		forum := version1.Group("forum")
		{
			v1.ForumRouter(forum)
		}
	}
}
