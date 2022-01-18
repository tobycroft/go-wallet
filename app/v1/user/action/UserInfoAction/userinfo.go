package UserInfoAction

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/user/model/UserModel"
)

func App_userinfo(uid interface{}) gorose.Data {
	user := UserModel.Api_find(uid)
	delete(user, "password")
	return user
}
