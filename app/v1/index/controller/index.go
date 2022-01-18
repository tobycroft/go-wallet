package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/LogMailModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/Mail"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("mail", index_mail)
}

func index_mail(c *gin.Context) {
	mailaddr, ok := Input.Post("mail", c, false)
	if !ok {
		return
	}
	count1 := LogMailModel.Api_count_60(c.ClientIP())
	if count1 > 10 {
		RET.Fail(c, 403, nil, "邮件需要间隔一分钟")
		return
	}
	count := LogMailModel.Api_count_today(c.ClientIP())
	if count > 10 {
		RET.Fail(c, 403, nil, "一天内只能注册10次")
		return
	}
	mail_host := SystemParamModel.Api_find_val("mail_host")
	mail_user := SystemParamModel.Api_find_val("mail_user")
	mail_password := SystemParamModel.Api_find_val("mail_password")
	mail := Mail.SendStruct{
		Host:     mail_host.(string),
		Port:     "25",
		User:     mail_user.(string),
		Password: mail_password.(string),
		Title:    "[GoWallet]Your Verify Code",
	}
	mail.To = mailaddr
	rand := Calc.Rand(100000, 999999)
	mail.Content = "Your verify code is:[" + Calc.Int2String(rand) + "], this code will be avail in 24H"

	err := mail.SendMail()
	if err != nil {
		LogMailModel.Api_insert(c.ClientIP(), 0, mailaddr, err.Error())
	} else {
		LogMailModel.Api_insert(c.ClientIP(), 1, mailaddr, "")
	}
}

func index(c *gin.Context) {
	c.String(0, "index")
}
