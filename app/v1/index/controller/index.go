package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"log"
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
	// Setup an unencrypted connection to a local mail server.
	// Set the sender and recipient, and send the email all in one step.

	mail := Mail.SendStruct{
		Host:     "localhost",
		User:     "verify@tuuz.cc",
		Password: "qwerty123",
		Title:    "[GoWallet]Your Verify Code",
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
