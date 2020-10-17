package gateway

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway/listener"
)

type Adapter interface {
	IsValidAddress(ctx context.Context, address string) bool
	NewWallet(ctx context.Context) (Wallet, error)
	FindWallet(ctx context.Context, privateKey string) (Wallet, error)
	GetBalance(ctx context.Context, address string) (map[string]float64, error)
	EstimateSendFee(ctx context.Context, w Wallet, coin string, amount float64, address string) (float64, float64, error)
	Send(ctx context.Context, w Wallet, coin string, amount float64, address string) (string, error)
	EstimateMultiSendFee(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, float64, error)
	MultiSend(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error)
	Subscribe(ctx context.Context, consumer listener.EventConsumer) error
	IsTransactionComplete(ctx context.Context, hash string) bool
	Unsubscribe()
}

type CryptoHub struct {
	BlockChains map[string]Adapter
}

func NewCryptoHub(blockChains map[string]Adapter) *CryptoHub {
	return &CryptoHub{blockChains}
}

func (ch *CryptoHub) IsValidAddress(ctx context.Context, baseCoin string, address string) bool {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.IsValidAddress(ctx, address)
	} else {
		return false
	}
}

func (ch *CryptoHub) MustNewWallet(ctx context.Context, baseCoin string) Wallet {
	wallet, err := ch.NewWallet(ctx, baseCoin)
	if err != nil {
		panic(err)
	}
	return wallet
}

func (ch *CryptoHub) NewWallet(ctx context.Context, baseCoin string) (Wallet, error) {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.NewWallet(ctx)
	} else {
		return Wallet{}, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustFindWallet(ctx context.Context, baseCoin string, privateKey string) Wallet {
	wallet, err := ch.FindWallet(ctx, baseCoin, privateKey)
	if err != nil {
		panic(err)
	}
	return wallet
}

func (ch *CryptoHub) FindWallet(ctx context.Context, baseCoin string, privateKey string) (Wallet, error) {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.FindWallet(ctx, privateKey)
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
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
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
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.Send(ctx, w, coin, amount, address)
	} else {
		return "", fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustEstimateSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) (float64, float64) {
	estimation, gasPrice, err := ch.EstimateSendFee(ctx, baseCoin, w, coin, amount, address)
	if err != nil {
		panic(err)
	}
	return estimation, gasPrice
}

func (ch *CryptoHub) EstimateSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, amount float64, address string) (float64, float64, error) {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.EstimateSendFee(ctx, w, coin, amount, address)
	} else {
		return 0, 0, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
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
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		if len(addresses) != len(amounts) {
			return "", fmt.Errorf("number of addresses should be equal to numberr of amounts")
		}
		return adapter.MultiSend(ctx, w, coin, addresses, amounts)
	} else {
		return "", fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MustEstimateMultiSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) (float64, float64) {
	estimation, gasPrice, err := ch.EstimateMultiSendFee(ctx, baseCoin, w, coin, addresses, amounts)
	if err != nil {
		panic(err)
	}
	return estimation, gasPrice
}

func (ch *CryptoHub) EstimateMultiSendFee(ctx context.Context, baseCoin string, w Wallet, coin string, addresses []string, amounts []float64) (float64, float64, error) {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		if len(addresses) != len(amounts) {
			return 0, 0, fmt.Errorf("number of addresses should be equal to numberr of amounts")
		}
		return adapter.EstimateMultiSendFee(ctx, w, coin, addresses, amounts)
	} else {
		return 0, 0, fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) IsTransactionComplete(ctx context.Context, baseCoin string, hash string) bool {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		return adapter.IsTransactionComplete(ctx, hash)
	} else {
		return false
	}
}

func (ch *CryptoHub) SubscribeAll(ctx context.Context, consumer listener.EventConsumer) error {
	for baseCoin, adapter := range ch.BlockChains {
		if err := adapter.Subscribe(ctx, consumer); err != nil {
			return fmt.Errorf("unable to subscribe to %s blockchain", baseCoin)
		}
	}
	return nil
}

func (ch *CryptoHub) UnsubscribeAll() {
	for _, adapter := range ch.BlockChains {
		adapter.Unsubscribe()
	}
}

func (ch *CryptoHub) Subscribe(ctx context.Context, baseCoin string, consumer listener.EventConsumer) error {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		if err := adapter.Subscribe(ctx, consumer); err != nil {
			return fmt.Errorf("unable to subscribe to %s blockchain", baseCoin)
		}
	} else {
		return fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
	return nil
}

func (ch *CryptoHub) Unsubscribe(_ context.Context, baseCoin string) error {
	if adapter, ok := ch.BlockChains[baseCoin]; ok {
		adapter.Unsubscribe()
	} else {
		return fmt.Errorf("blockchain adapter for coin %s not found", baseCoin)
	}
	return nil
}
