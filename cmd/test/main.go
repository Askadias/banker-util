package main

import (
	"fmt"
	"github.com/Askadias/banker-util/gateway"
	"github.com/MinterTeam/minter-go-sdk/api"
	"github.com/go-resty/resty/v2"
)

func main() {
	minterClient := api.NewApiWithClient("http://api.minter.one", resty.New())
	minter := gateway.NewMinterAdapter(minterClient)
	hub := gateway.NewCryptoHub(map[string]gateway.Adapter{"BIP": minter})
	wallet, err := hub.NewWallet("BIP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("New wallet: %s", wallet.Address)
}
