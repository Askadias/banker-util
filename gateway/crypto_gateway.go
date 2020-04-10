package gateway

import (
	"context"
	"fmt"
)

type Adapter interface {
	NewWallet(ctx context.Context) (Wallet, error)
	GetBalance(ctx context.Context, address string) (map[string]float64, error)
	EstimateSendFee(ctx context.Context, w Wallet, coin string, amount float64, address string) (float64, error)
	Send(ctx context.Context, w Wallet, coin string, amount float64, address string) (string, error)
	EstimateMultiSendFee(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, error)
	MultiSend(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error)
}

type CryptoHub struct {
	blockChains map[string]Adapter
}

func NewCryptoHub(blockChains map[string]Adapter) *CryptoHub {
	return &CryptoHub{blockChains}
}

func (ch *CryptoHub) MustNewWallet(ctx context.Context, baseCoin string) Wallet {
	wallet, err := ch.NewWallet(ctx, baseCoin)
	if err != nil {
		panic(err)
	}
	return wallet
}

func (ch *CryptoHub) NewWallet(ctx context.Context, baseCoin string) (Wallet, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.NewWallet(ctx)
	} else {
		return Wallet{}, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustGetBalance(ctx context.Context, baseCoin string, address string) map[string]float64 {
	balance, err := ch.GetBalance(ctx, baseCoin, address)
	if err != nil {
		panic(err)
	}
	return balance
}

func (ch *CryptoHub) GetBalance(ctx context.Context, baseCoin string, address string) (map[string]float64, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.GetBalance(ctx, address)
	} else {
		return map[string]float64{}, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustSend(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) string {
	hash, err := ch.Send(ctx, baseCoin, w, coin, amount, address)
	if err != nil {
		panic(err)
	}
	return hash
}

func (ch *CryptoHub) Send(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) (string, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.Send(ctx, w, coin, amount, address)
	} else {
		return "", fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustEstimateSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) float64 {
	estimation, err := ch.EstimateSendFee(ctx, baseCoin, w, coin, amount, address)
	if err != nil {
		panic(err)
	}
	return estimation
}

func (ch *CryptoHub) EstimateSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) (float64, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.EstimateSendFee(ctx, w, coin, amount, address)
	} else {
		return 0, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustMultiSend(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) string {
	hash, err := ch.MultiSend(ctx, baseCoin, w, coin, addresses, amounts)
	if err != nil {
		panic(err)
	}
	return hash
}

func (ch *CryptoHub) MultiSend(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		if len(addresses) != len(amounts) {
			return "", fmt.Errorf("number of addresses should be equal to numberr of amounts")
		}
		return adapter.MultiSend(ctx, w, coin, addresses, amounts)
	} else {
		return "", fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustEstimateMultiSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) float64 {
	estimation, err := ch.EstimateMultiSendFee(ctx, baseCoin, w, coin,  addresses, amounts)
	if err != nil {
		panic(err)
	}
	return estimation
}

func (ch *CryptoHub) EstimateMultiSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) (float64, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		if len(addresses) != len(amounts) {
			return 0, fmt.Errorf("number of addresses should be equal to numberr of amounts")
		}
		return adapter.EstimateMultiSendFee(ctx, w, coin, addresses, amounts)
	} else {
		return 0, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}
