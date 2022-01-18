package Trc20_Usdt

import (
	"bytes"
	"fmt"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/shopspring/decimal"
	"main.go/config/app_conf"
	"main.go/extend/trx-sign-go-1.0.3/grpcs"
	"main.go/extend/trx-sign-go-1.0.3/sign"
)

type TokenTransaction struct {
	client          *grpcs.Client
	contractAddress string
}

func InitTranns(contractAddress string) *TokenTransaction {
	//EthRPC_API := SystemParamModel.Api_find_val("EthRPC_API").(string)
	TrcRPC_API := app_conf.TrcRPC_API
	c, err := grpcs.NewClient(TrcRPC_API)
	if err != nil {
		panic(err)
	}
	return &TokenTransaction{client: c, contractAddress: contractAddress}
}

func (c *TokenTransaction) TransferFrom(privatekey, from, to string, amount decimal.Decimal) (error, *api.TransactionExtention) {
	a, err := ethabi.JSON(bytes.NewReader([]byte(abiJson)))
	if err != nil {
		fmt.Println("JSON", err)
		return err, nil
	}
	//method:=a.Methods["transferFrom"]
	fromaddress, err := addr.Base58ToAddress(from)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	toaddress, err := addr.Base58ToAddress(to)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	amount = amount.Mul(decimal.NewFromInt(1000000))
	bz, err := a.Pack("transferFrom", common.BytesToAddress(fromaddress.Bytes()), common.BytesToAddress(toaddress.Bytes()), amount.BigInt())
	//bz, err := abi2.Pack("transferFrom", ab)
	if err != nil {
		fmt.Println("Pack", err)
		return err, nil
	}
	s := common.Bytes2Hex(bz)

	tx, err := c.client.GRPC.TRC20Call(from, c.contractAddress, s, false, 2000000)
	if err != nil {
		fmt.Println("TRC20Call", err)
		return err, nil
	}
	signTx, err := sign.SignTransaction(tx.Transaction, privatekey)
	if err != nil {
		fmt.Println("SignTransaction", err)
		return err, nil
	}
	//fmt.Println("signTx", signTx.String())
	err = c.client.BroadcastTransaction(signTx)
	if err != nil {
		fmt.Println("BroadcastTransaction", err)
		return err, nil
	}
	fmt.Println(common.Bytes2Hex(tx.GetTxid()))
	return nil, tx
}
