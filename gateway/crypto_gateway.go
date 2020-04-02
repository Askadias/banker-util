package gateway

import "fmt"

type Adapter interface {
	NewWallet() (Wallet, error)
	GetBalance(w Wallet) (map[string]float64, error)
	Send(w Wallet, coin string, amount float64, address string) (string, error)
	MultiSend(w Wallet, coin string, amounts map[string]float64) (string, error)
}

type CryptoHub struct {
	blockChains map[string]Adapter
}

func NewCryptoHub(blockChains map[string]Adapter) *CryptoHub {
	return &CryptoHub{blockChains}
}

func (ch *CryptoHub) NewWallet(baseCoin string) (Wallet, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.NewWallet()
	} else {
		return Wallet{}, fmt.Errorf("network adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) GetBalance(baseCoin string, w Wallet) (map[string]float64, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.GetBalance(w)
	} else {
		return map[string]float64{}, fmt.Errorf("network adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) Send(baseCoin string, w Wallet, coin string, amount float64, address string) (string, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.Send(w, coin, amount, address)
	} else {
		return "", fmt.Errorf("network adapter for coin %s not found", baseCoin)
	}
}

func (ch *CryptoHub) MultiSend(baseCoin string, w Wallet, coin string, amounts map[string]float64) (string, error) {
	if adapter, ok := ch.blockChains[baseCoin]; ok {
		return adapter.MultiSend(w, coin, amounts)
	} else {
		return "", fmt.Errorf("network adapter for coin %s not found", baseCoin)
	}
}



