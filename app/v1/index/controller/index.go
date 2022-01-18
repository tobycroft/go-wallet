package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"log"
	"net/smtp"
	"strings"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("login", loginss)
	route.Any("register")
}

func index_register(c *gin.Context) {
	// Setup an unencrypted connection to a local mail server.
	c, err := smtp.Dial("localhost:25")
	if err != nil {
		return err
	}
	defer c.Close()

	// Set the sender and recipient, and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := c.SendMail("sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
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
