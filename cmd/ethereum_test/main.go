package main

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"time"
)

func main() {
	ethereumClient, err := ethclient.Dial(os.Getenv("ETH_HOST"))
	if err != nil {
		panic(err)
	}
	ethereum := gateway.NewEthereumAdapter(ethereumClient, 0)
	hub := gateway.NewCryptoHub(map[string]gateway.Adapter{
		"ETH": ethereum,
	})
	ctx := context.Background()

	sourceWallet := hub.MustFindWallet(ctx, "ETH", os.Getenv("ETH_PRIVATE_KEY"))
	targetWallet := hub.MustFindWallet(ctx, "ETH", os.Getenv("ETH_PRIVATE_KEY2"))

	_ = hub.SubscribeAll(ctx, gateway.EventConsumerFunc(func(event gateway.Event) {
		if event.From == sourceWallet.Address ||
			event.To == sourceWallet.Address ||
			event.From == targetWallet.Address ||
			event.To == targetWallet.Address || len(event.Items) > 0 {
			fmt.Println(event)
			fmt.Println("Transaction complete:", hub.IsTransactionComplete(ctx, "ETH", event.Hash))
		}
	}))

	hash := hub.MustSend(ctx, "ETH", sourceWallet, "USDT", 0.1, targetWallet.Address)
	fmt.Printf("Transaction Send USDT: https://etherscan.io/tx/%s\n", hash)
	fmt.Println("Transaction complete:", hub.IsTransactionComplete(ctx, "ETH", hash))
	time.Sleep(20 * time.Minute)
	// ==============================================================================================
	// CREATE MULTI_SEND CONTRACT
	//address, contractHash, err := ethereum.DeployMultiSendContract(ctx, sourceWallet)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Contract address: %s\n", address)
	//fmt.Printf("Transaction: https://etherscan.io/tx/%s\n", contractHash)

	// ==============================================================================================
	// CREATE WALLET
	//targetWallet := hub.MustNewWallet(ctx, "ETH")
	//fmt.Printf("New Wallet: %s - %s\n", targetWallet.Address, targetWallet.PrivateKey)
	sourceBalance := hub.MustGetBalance(ctx, "ETH", sourceWallet.Address)
	targetBalance := hub.MustGetBalance(ctx, "ETH", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// MULTI_SEND
	//multisendApproveHashUSDT, err := ethereum.ApproveTokenMultisend(ctx, sourceWallet, "USDT")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Transaction MultiSend USDT: https://etherscan.io/tx/%s\n", multisendApproveHashUSDT)
	multiETHWallets := []string{targetWallet.Address, targetWallet.Address, targetWallet.Address}
	multiETHAmounts := []float64{0.01, 0.01, 0.01}
	multiUSDTWallets := []string{targetWallet.Address, targetWallet.Address, targetWallet.Address}
	multiUSDTAmounts := []float64{0.01, 0.01, 0.01}
	estimatedETH := hub.MustEstimateMultiSendFee(ctx, "ETH", sourceWallet, "ETH", multiETHWallets, multiETHAmounts)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 0.01x3 ETH\n", sourceWallet.Address, targetWallet.Address, estimatedETH)
	estimatedUSDT := hub.MustEstimateMultiSendFee(ctx, "ETH", sourceWallet, "USDT", multiUSDTWallets, multiUSDTAmounts)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 0.1x3 USDT\n", sourceWallet.Address, targetWallet.Address, estimatedUSDT)

	multisendHashETH := hub.MustMultiSend(ctx, "ETH", sourceWallet, "ETH", multiETHWallets, multiETHAmounts)
	fmt.Printf("Transaction MultiSend ETH: https://etherscan.io/tx/%s\n", multisendHashETH)
	//time.Sleep(1 * time.Minute)
	multisendHashUSDT := hub.MustMultiSend(ctx, "ETH", sourceWallet, "USDT", multiUSDTWallets, multiUSDTAmounts)
	fmt.Printf("Transaction MultiSend USDT: https://etherscan.io/tx/%s\n", multisendHashUSDT)
	time.Sleep(2 * time.Minute)
	sourceBalance = hub.MustGetBalance(ctx, "ETH", sourceWallet.Address)
	targetBalance = hub.MustGetBalance(ctx, "ETH", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// SEND USDT
	estimatedUSDTSend := hub.MustEstimateSendFee(ctx, "ETH", targetWallet, "USDT", targetBalance["USDT"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f USDT + %0.9f ETH\n", targetWallet.Address, sourceWallet.Address, targetBalance["USDT"], estimatedUSDTSend)
	sendHashUSDT := hub.MustSend(ctx, "ETH", targetWallet, "USDT", targetBalance["USDT"], sourceWallet.Address)
	fmt.Printf("Transaction Send USDT: https://etherscan.io/tx/%s\n", sendHashUSDT)
	time.Sleep(1 * time.Minute)

	// ==============================================================================================
	// SEND ETH
	estimatedETHSend := hub.MustEstimateSendFee(ctx, "ETH", targetWallet, "ETH", targetBalance["ETH"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f ETH + %0.9f ETH\n", targetWallet.Address, sourceWallet.Address, targetBalance["ETH"]-estimatedETHSend, estimatedETHSend)
	sendHashETH := hub.MustSend(ctx, "ETH", targetWallet, "ETH", targetBalance["ETH"]-estimatedETHSend, sourceWallet.Address)
	fmt.Printf("Transaction Send ETH: https://etherscan.io/tx/%s\n", sendHashETH)
	time.Sleep(2 * time.Minute)
	ethereum.Unsubscribe()
}
