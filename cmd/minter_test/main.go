package main

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway"
	"github.com/Askadias/banker-util/gateway/listener"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client"
	"os"
	"time"
)

func main() {
	minterClient, _ := http_client.New(os.Getenv("MINTER_HOST"))
	minterPollingClient, _ := http_client.New(os.Getenv("MINTER_POLLING_HOST"))
	minter := gateway.NewMinterAdapter(minterClient, minterPollingClient, 2*time.Second)
	hub := gateway.NewCryptoHub(map[string]gateway.Adapter{
		"BIP": minter,
	})
	ctx := context.Background()

	sourceWallet := hub.MustFindWallet(ctx, "BIP", os.Getenv("BIP_PRIVATE_KEY"))
	targetWallet := hub.MustFindWallet(ctx, "BIP", os.Getenv("BIP_PRIVATE_KEY2"))

	_ = minter.Subscribe(ctx, listener.EventConsumerFunc(func(event listener.Event) {
		if event.From == sourceWallet.Address ||
			event.To == sourceWallet.Address ||
			event.From == targetWallet.Address ||
			event.To == targetWallet.Address || len(event.Items) > 0 {
			fmt.Println(event)
			fmt.Println("Transaction complete:", hub.IsTransactionComplete(ctx, "BIP", event.Hash))
		}
	}))
	//select {}

	//hash := hub.MustSend(ctx, "BIP", sourceWallet, "BIP", 1, targetWallet.Address)
	//fmt.Printf("Transaction MultiSend BIP: https://minterscan.net/tx/%s\n", hash)
	//fmt.Println("Transaction complete:", hub.IsTransactionComplete(ctx, "ETH", hash))
	//time.Sleep(20 * time.Minute)

	// ==============================================================================================
	// CREATE WALLET
	//targetWallet := hub.MustNewWallet(ctx, "BIP")
	//fmt.Printf("New Wallet: %s - %s\n", targetWallet.Address, targetWallet.PrivateKey)
	sourceBalance := hub.MustGetBalance(ctx, "BIP", sourceWallet.Address)
	targetBalance := hub.MustGetBalance(ctx, "BIP", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	//cost, fee, err := minter.EstimateBuy(ctx, "MUSD", 1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Estimate buy 1 MUSD: %v fee: %v", cost, fee)
	//buy, err := minter.BuySwapPool(ctx, sourceWallet, "MUSD", 1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Transaction BuySwapPool MUSD: https://minterscan.net/tx/%s\n", buy)
	//sourceBalance = hub.MustGetBalance(ctx, "BIP", sourceWallet.Address)
	//fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)

	//sellAll, err := minter.SellAllSwapPool(ctx, sourceWallet, "HUB")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Transaction SellAll HUB: https://minterscan.net/tx/%s\n", sellAll)
	//sourceBalance = hub.MustGetBalance(ctx, "BIP", sourceWallet.Address)
	//targetBalance = hub.MustGetBalance(ctx, "BIP", targetWallet.Address)
	//fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	//fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// MULTI_SEND
	multiBIPWallets := []string{targetWallet.Address, targetWallet.Address}
	multiBIPAmounts := []float64{2, 2}
	multiONLY1Wallets := []string{targetWallet.Address, targetWallet.Address}
	multiONLY1Amounts := []float64{10, 10}
	estimatedBIP, _ := hub.MustEstimateMultiSendFee(ctx, "BIP", sourceWallet, "BIP", multiBIPWallets, multiBIPAmounts)
	estimatedONLY1, _ := hub.MustEstimateMultiSendFee(ctx, "BIP", sourceWallet, "ONLY1", multiONLY1Wallets, multiONLY1Amounts)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 0.1x3 BIP\n", sourceWallet.Address, targetWallet.Address, estimatedBIP)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 10x3 ONLY1\n", sourceWallet.Address, targetWallet.Address, estimatedONLY1)

	multisendHashBIP := hub.MustMultiSend(ctx, "BIP", sourceWallet, "BIP", multiBIPWallets, multiBIPAmounts)
	//time.Sleep(5 * time.Second)
	multisendHashONLY1 := hub.MustMultiSend(ctx, "BIP", sourceWallet, "ONLY1", multiONLY1Wallets, multiONLY1Amounts)
	fmt.Printf("Transaction MultiSend BIP: https://minterscan.net/tx/%s\n", multisendHashBIP)
	fmt.Printf("Transaction MultiSend ONLY1: https://minterscan.net/tx/%s\n", multisendHashONLY1)
	time.Sleep(5 * time.Second)
	sourceBalance = hub.MustGetBalance(ctx, "BIP", sourceWallet.Address)
	targetBalance = hub.MustGetBalance(ctx, "BIP", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// SEND ONLY1
	estimatedONLY1, _ = hub.MustEstimateSendFee(ctx, "BIP", targetWallet, "ONLY1", targetBalance["ONLY1"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f %0.9f ONLY1\n", targetWallet.Address, sourceWallet.Address, estimatedONLY1, targetBalance["ONLY1"])
	sendHashONLY1 := hub.MustSend(ctx, "BIP", targetWallet, "ONLY1", targetBalance["ONLY1"], sourceWallet.Address)
	fmt.Printf("Transaction Send ONLY1: https://minterscan.net/tx/%s\n", sendHashONLY1)

	// ==============================================================================================
	// SEND BIP
	estimatedBIP, _ = hub.MustEstimateSendFee(ctx, "BIP", targetWallet, "BIP", targetBalance["BIP"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f %0.9f BIP\n", targetWallet.Address, sourceWallet.Address, estimatedBIP, targetBalance["BIP"]-estimatedONLY1)
	sendHashBIP := hub.MustSend(ctx, "BIP", targetWallet, "BIP", targetBalance["BIP"]-estimatedBIP-estimatedONLY1, sourceWallet.Address)
	fmt.Printf("Transaction Send BIP: https://minterscan.net/tx/%s\n", sendHashBIP)
	time.Sleep(3 * time.Second)
	minter.Unsubscribe()
}
