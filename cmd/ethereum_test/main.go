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
	ethereum := gateway.NewEthereumAdapter(ethereumClient)
	hub := gateway.NewCryptoHub(map[string]gateway.Adapter{
		"ETH": ethereum,
	})

	sourceWallet := gateway.Wallet{Address: os.Getenv("ETH_ADDRESS"), PrivateKey: os.Getenv("ETH_PRIVATE_KEY")}
	targetWallet := gateway.Wallet{Address: os.Getenv("ETH_ADDRESS2"), PrivateKey: os.Getenv("ETH_PRIVATE_KEY2")}

	// ==============================================================================================
	// CREATE MULTI_SEND CONTRACT
	//address, contractHash, err := ethereum.DeployMultiSendContract(context.Background(), sourceWallet)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Contract address: %s\n", address)
	//fmt.Printf("Transaction: https://etherscan.io/tx/%s\n", contractHash)


	// ==============================================================================================
	// CREATE WALLET
	//targetWallet := hub.MustNewWallet(context.Background(), "ETH")
	fmt.Printf("New Wallet: %s - %s\n", targetWallet.Address, targetWallet.PrivateKey)
	sourceBalance := hub.MustGetBalance(context.Background(), "ETH", sourceWallet.Address)
	targetBalance := hub.MustGetBalance(context.Background(), "ETH", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// MULTI_SEND
	multiETHWallets := []string{targetWallet.Address, targetWallet.Address, targetWallet.Address}
	multiETHAmounts := []float64{0.01, 0.01, 0.01}
	multiUSDTWallets := []string{targetWallet.Address, targetWallet.Address, targetWallet.Address}
	multiUSDTAmounts := []float64{0.1, 0.1, 0.1}
	estimatedETH := hub.MustEstimateMultiSendFee(context.Background(), "ETH", sourceWallet, "ETH", multiETHWallets, multiETHAmounts)
	estimatedUSDT := hub.MustEstimateMultiSendFee(context.Background(), "ETH", sourceWallet, "USDT", multiUSDTWallets, multiUSDTAmounts)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 0.01 ETH\n", sourceWallet.Address, targetWallet.Address, estimatedETH)
	fmt.Printf("Multisend Estimation: %s -> %s %0.9f 0.1 USDT\n", sourceWallet.Address, targetWallet.Address, estimatedUSDT)

	multisendHashETH := hub.MustMultiSend(context.Background(), "ETH", sourceWallet, "ETH", multiETHWallets, multiETHAmounts)
	//time.Sleep(5 * time.Second)
	multisendHashUSDT := hub.MustMultiSend(context.Background(), "ETH", sourceWallet, "USDT", multiUSDTWallets, multiUSDTAmounts)
	fmt.Printf("Transaction MultiSend ETH: https://etherscan.io/tx/%s\n", multisendHashETH)
	fmt.Printf("Transaction MultiSend USDT: https://etherscan.io/tx/%s\n", multisendHashUSDT)
	time.Sleep(5 * time.Second)
	sourceBalance = hub.MustGetBalance(context.Background(), "ETH", sourceWallet.Address)
	targetBalance = hub.MustGetBalance(context.Background(), "ETH", targetWallet.Address)
	fmt.Printf("Balance: %s = %v\n", sourceWallet.Address, sourceBalance)
	fmt.Printf("Balance: %s = %v\n", targetWallet.Address, targetBalance)

	// ==============================================================================================
	// SEND USDT
	estimatedUSDT = hub.MustEstimateSendFee(context.Background(), "ETH", targetWallet, "USDT", targetBalance["USDT"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f %0.9f USDT\n", targetWallet.Address, sourceWallet.Address, estimatedUSDT, targetBalance["USDT"])
	sendHashUSDT := hub.MustSend(context.Background(), "ETH", targetWallet, "USDT", targetBalance["USDT"], sourceWallet.Address)
	fmt.Printf("Transaction Send USDT: https://etherscan.io/tx/%s\n", sendHashUSDT)

	// ==============================================================================================
	// SEND ETH
	estimatedETH = hub.MustEstimateSendFee(context.Background(), "ETH", targetWallet, "ETH", targetBalance["ETH"], sourceWallet.Address)
	fmt.Printf("Estimation: %s -> %s %0.9f %0.9f ETH\n", targetWallet.Address, sourceWallet.Address, estimatedETH, targetBalance["ETH"]-estimatedUSDT)
	sendHashETH := hub.MustSend(context.Background(), "ETH", targetWallet, "ETH", targetBalance["ETH"]-estimatedETH-estimatedUSDT, sourceWallet.Address)
	fmt.Printf("Transaction Send ETH: https://etherscan.io/tx/%s\n", sendHashETH)
}
