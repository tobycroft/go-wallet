package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"log"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Mail"
	"net/smtp"
	"strings"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("login", loginss)
	route.Any("register", index_register)
}

func index_register(c *gin.Context) {
	mail_host := SystemParamModel.Api_find_val("mail_host")
	mail_user := SystemParamModel.Api_find_val("mail_user")
	mail_password := SystemParamModel.Api_find_val("mail_password")
	mail := Mail.SendStruct{
		Host:     mail_host.(string),
		User:     mail_user.(string),
		Password: mail_password.(string),
		Title:    "[GoWallet]Your Verify Code",
		Content:  "Your code is:123456",
	}

	err := mail.SendMail()
	fmt.Println(err)
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func loginss(c *gin.Context) {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(0, json)
}
