package InvestTransfer

import (
	"main.go/app/v1/coin/model/CoinModel"
	"main.go/app/v1/invest/model/InvestOrderModel"
	"main.go/app/v1/wallet/model/UserAddressModel"
	"main.go/common/BaseModel/SystemParamModel"
	Erc20_Usdt "main.go/extend/Erc20_Usdt"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"time"
)

func Refresh_eth() {
	for {
		InvestTransfer_eth()
		time.Sleep(60 * time.Second)
	}
}

func InvestTransfer_eth() {
	coin := CoinModel.Api_find_byTypeAndName("eth", "usdt")
	eth_address := SystemParamModel.Api_find_val("eth_address").(string)
	db := tuuz.Db()
	var io InvestOrderModel.Interface
	io.Db = db
	datas := io.Api_select_byProgress(0)
	for _, data := range datas {
		t := Erc20_Usdt.InitTranns(coin["contract"].(string))
		var us UserAddressModel.Interface
		us.Db = db
		useraddr := us.Api_find(data["uid"], "eth")
		if len(useraddr) < 1 {
			continue
		}
		err, txs := t.TransferFrom("c2e34562e0478a3e4e8f1f79f0d9f156c81249da3df00013531191888a18d7cf", useraddr["address"].(string), eth_address, Calc.ToDecimal(data["amount"]))
		//fmt.Println("err",err)
		if err != nil {
			db.Begin()
			if !io.Api_update_progress(data["id"], -1) {
				db.Rollback()
				continue
			}
			if !io.Api_update_remark(data["id"], err.Error()) {
				db.Rollback()
				continue
			}
			db.Commit()
		} else {
			db.Begin()
			if !io.Api_update_progress(data["id"], 1) {
				db.Rollback()
				continue
			}
			if !io.Api_update_txId(data["id"], txs.Hash().String()) {
				db.Rollback()
				continue
			}
			db.Commit()
		}
	}
}
