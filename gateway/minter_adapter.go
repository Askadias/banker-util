package gateway

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/api"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"math"
	"math/big"
)

type MinterAdapter struct {
	Client *api.Api
}

func NewMinterAdapter(minterClient *api.Api) *MinterAdapter {
	return &MinterAdapter{Client: minterClient}
}

func (ma *MinterAdapter) getWallet(mnemonic string) (*wallet.Wallet, error) {
	seed, err := wallet.Seed(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare wallet seed: %v", err)
	}
	mntWallet, err := wallet.NewWallet(seed)
	if err != nil {
		return nil, fmt.Errorf("unable create wallet: %v", err)
	}
	return mntWallet, nil
}

func (ma *MinterAdapter) NewWallet() (Wallet, error) {
	mnemonic, err := wallet.NewMnemonic()
	emptyWallet := Wallet{"", ""}
	if err != nil {
		return emptyWallet, fmt.Errorf("unable to generate wallet mnemonic: %v", err)
	}
	mntWallet, err := ma.getWallet(mnemonic)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{mntWallet.Address(), mnemonic}, nil
}

func (ma *MinterAdapter) GetBalance(w Wallet) (map[string]float64, error) {
	balance := map[string]float64{}
	wlt, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return balance, err
	}
	result, err := ma.Client.Balance(wlt.Address())
	if err != nil {
		return balance, fmt.Errorf("unable to get balance: %v", err)
	}
	for coin, balanceAmount := range result {
		balanceVal, _, err := big.ParseFloat(balanceAmount, 10, 0, big.ToNegativeInf)
		if err != nil {
			return balance, fmt.Errorf("unable to parse %s balance: %v", coin, err)
		}
		quotient := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
		balanceAmount, _ := new(big.Float).Quo(balanceVal, quotient).Float64()
		balance[coin] = balanceAmount
	}
	return balance, nil
}

func (ma *MinterAdapter) SellAll(w Wallet, coin string) (string, error) {
	data := transaction.NewSellAllCoinData().
		SetCoinToSell(coin).
		SetCoinToBuy("BIP")

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Client.Nonce(mntWallet.Address())
	if err != nil {
		return "", err
	}
	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(coin).
		SetGasPrice(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey())
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to sell coin %s: %v", coin, err)
	}

	res, err := ma.Client.SendTransaction(signedTransaction)
	if err != nil {
		return "", fmt.Errorf("unable to sell coin %s: %v", coin, err)
	}

	return res.Hash, nil
}

func (ma *MinterAdapter) Send(w Wallet, coin string, amount float64, address string) (string, error) {
	value := big.NewFloat(0).Copy(big.NewFloat(amount))
	multiplier := big.NewFloat(math.Pow10(18))
	value.Mul(value, multiplier)
	// value = amount * 10^18
	// amount = 0.0000000123
	// value = 1230000000000
	val, _ := value.Int(big.NewInt(0))
	data, _ := transaction.NewSendData().
		SetCoin(coin).
		SetValue(val).
		SetTo(address)

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Client.Nonce(mntWallet.Address())
	if err != nil {
		return "", fmt.Errorf("unable to generate nonce: %v", err)
	}

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(coin).
		SetGasPrice(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey())
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
	}

	res, err := ma.Client.SendTransaction(signedTransaction)
	if err != nil {
		return "", fmt.Errorf("unable to send coin %s to %s: %v", coin, address, err)
	}
	return res.Hash, nil
}

func (ma *MinterAdapter) MultiSend(w Wallet, coin string, amounts map[string]float64) (string, error) {
	data := transaction.NewMultisendData()
	for address, amount := range amounts {
		value := big.NewFloat(0).Copy(big.NewFloat(amount))
		multiplier := big.NewFloat(math.Pow10(18))
		value.Mul(value, multiplier)
		val, _ := value.Int(big.NewInt(0))
		data.AddItem(*transaction.NewMultisendDataItem().
			SetCoin(coin).SetValue(val).
			MustSetTo(address))
	}

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Client.Nonce(mntWallet.Address())
	if err != nil {
		return "", fmt.Errorf("unable to generate nonce: %v", err)
	}

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(coin).
		SetGasPrice(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey())
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
	}

	res, err := ma.Client.SendTransaction(signedTransaction)
	if err != nil {
		return "", fmt.Errorf("unable to send coin %s to %v: %v", coin, amounts, err)
	}
	return res.Hash, nil
}
