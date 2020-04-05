package main

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway"
	"github.com/MinterTeam/minter-go-sdk/api"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"os"
)

func main() {
	minterClient := api.NewApiWithClient(os.Getenv("MINTER_HOST"), resty.New())
	ethereumClient, err := ethclient.Dial(os.Getenv("ETH_HOST"))
	if err != nil {
		panic(err)
	}
	minter := gateway.NewMinterAdapter(minterClient)
	eth := gateway.NewEthereumAdapter(ethereumClient)
	hub := gateway.NewCryptoHub(map[string]gateway.Adapter{
		"BIP": minter,
		"ETH": eth,
	})

	bipWallet, err := hub.NewWallet(context.Background(), "BIP")
	ethWallet, err := hub.NewWallet(context.Background(), "ETH")
	balance1, err := hub.GetBalance(context.Background(), "ETH", "0x96c1b7aAC011e930A92F8B14D5C926f26e6eBA46")
	balance2, err := hub.GetBalance(context.Background(), "ETH", "0x024AE5728506Edf423BDce74D431299567e3c3Cf")
	balance3, err := hub.GetBalance(context.Background(), "ETH", "0x24Ae5728506EdF423bdcE74D431299567E3C3CF8")
	balance4, err := hub.GetBalance(context.Background(), "ETH", "0x0de80616a76aF3d92CB36708c5D1F7e2239eA643")
	estimatedAmount, err := hub.EstimateSendFee(context.Background(), "ETH", gateway.Wallet{PrivateKey: os.Getenv("ETH_PRIVATE_KEY")}, "USDT", 0.01, "0x0de80616a76aF3d92CB36708c5D1F7e2239eA643")
	estimatedAmount2, err := hub.EstimateSendFee(context.Background(), "ETH", gateway.Wallet{PrivateKey: os.Getenv("ETH_PRIVATE_KEY2")}, "USDT", 0.01, "0x96c1b7aAC011e930A92F8B14D5C926f26e6eBA46")
	//hash, err := hub.Send(context.Background(), "ETH", gateway.Wallet{PrivateKey: os.Getenv("ETH_PRIVATE_KEY")}, "USDT", 0.01,  "0x0de80616a76aF3d92CB36708c5D1F7e2239eA643")
	//hash, err := hub.Send(context.Background(), "ETH", gateway.Wallet{PrivateKey: os.Getenv("ETH_PRIVATE_KEY")}, "ETH", 0.00008237, "0x0de80616a76aF3d92CB36708c5D1F7e2239eA643")
	//hash, err := hub.Send(context.Background(), "ETH", gateway.Wallet{PrivateKey: os.Getenv("ETH_PRIVATE_KEY2")}, "USDT", 0.01, "0x96c1b7aAC011e930A92F8B14D5C926f26e6eBA46")
	//if err != nil {
	//	panic(err)
	//}
	fmt.Printf("BIP wallet: %s - %s\n", bipWallet.Address, bipWallet.PrivateKey)
	fmt.Printf("ETH wallet: %s - %s\n", ethWallet.Address, ethWallet.PrivateKey)
	fmt.Printf("ETH wallet: %s - %s\n", "0x0de80616a76aF3d92CB36708c5D1F7e2239eA643", os.Getenv("ETH_PRIVATE_KEY2"))
	fmt.Printf("ETH balance1: %v\n", balance1)
	fmt.Printf("ETH balance2: %v\n", balance2)
	fmt.Printf("ETH balance3: %v\n", balance3)
	fmt.Printf("ETH balance4: %v\n", balance4)
	//fmt.Printf("Transaction: https://etherscan.io/tx/%s\n", hash)
	fmt.Printf("Estimated Commission 1: %0.9f ETH\n", estimatedAmount)
	fmt.Printf("Estimated Commission 2: %0.9f ETH\n", estimatedAmount2)
}
