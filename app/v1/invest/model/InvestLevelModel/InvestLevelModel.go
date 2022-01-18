package InvestLevelModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "invest_level"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.OrderBy("id asc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
	}
	return ret
}
