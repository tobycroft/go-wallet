package CoinAction

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/coin/model/CoinModel"
)

func App_coin() map[int64]gorose.Data {
	coins := CoinModel.Api_select()
	arr := map[int64]gorose.Data{}
	for _, coin := range coins {
		arr[coin["id"].(int64)] = coin
	}
	return arr
}
