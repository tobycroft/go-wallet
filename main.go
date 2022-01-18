package main

import (
	"github.com/gin-gonic/gin"
	"main.go/route"
)

func main() {


	mainroute := gin.Default()
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
