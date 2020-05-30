package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/api"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

const (
	BIPDecimal = 18.0
)

type MinterAdapter struct {
	baseCoin        string
	client          *api.Api
	pollingClient   *api.Api
	ticker          *time.Ticker
	pollingDuration time.Duration
	lastBlockHeight int
	blockListeners  []func(lastBlockHeight int, transactionsCount int)
}

func NewMinterAdapter(minterClient *api.Api, minterPollingClient *api.Api, pollingDuration time.Duration) *MinterAdapter {
	return &MinterAdapter{baseCoin: "BIP", client: minterClient, pollingClient: minterPollingClient, pollingDuration: pollingDuration}
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

func (ma *MinterAdapter) IsValidAddress(_ context.Context, address string) bool {
	return wallet.IsValidAddress(address)
}

func (ma *MinterAdapter) FindWallet(ctx context.Context, privateKey string) (Wallet, error) {
	emptyWallet := Wallet{"", "", ""}
	mntWallet, err := ma.getWallet(privateKey)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, mntWallet.Address(), privateKey}, nil
}

func (ma *MinterAdapter) NewWallet(ctx context.Context) (Wallet, error) {
	mnemonic, err := wallet.NewMnemonic()
	emptyWallet := Wallet{"", "", ""}
	if err != nil {
		return emptyWallet, fmt.Errorf("unable to generate wallet mnemonic: %v", err)
	}
	mntWallet, err := ma.getWallet(mnemonic)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, mntWallet.Address(), mnemonic}, nil
}

func (ma *MinterAdapter) GetBalance(ctx context.Context, address string) (map[string]float64, error) {
	balance := map[string]float64{}
	result, err := ma.client.Balance(address)
	if err != nil {
		return balance, fmt.Errorf("unable to get balance: %v", err)
	}
	for coin, balanceAmount := range result {
		balanceVal, _, err := big.ParseFloat(balanceAmount, 10, 0, big.ToNegativeInf)
		if err != nil {
			return balance, fmt.Errorf("unable to parse %s balance: %v", coin, err)
		}
		quotient := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(BIPDecimal), nil))
		balanceAmount, _ := new(big.Float).Quo(balanceVal, quotient).Float64()
		balance[coin] = balanceAmount
	}
	return balance, nil
}

func (ma *MinterAdapter) SellAll(ctx context.Context, w Wallet, coin string) (string, error) {
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

	nonce, err := ma.client.Nonce(mntWallet.Address())
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

	res, err := ma.client.SendTransaction(signedTransaction)
	if err != nil {
		return "", fmt.Errorf("unable to sell coin %s: %v", coin, err)
	}

	return res.Hash, nil
}

func (ma *MinterAdapter) EstimateSendFee(ctx context.Context, w Wallet, coin string, amount float64, address string) (float64, error) {
	signedTransaction, _, err := ma.prepareSendTx(w, coin, amount, address)
	if err != nil {
		return 0, fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
	}
	res, err := ma.client.EstimateTxCommission(signedTransaction)
	if err != nil {
		return 0, fmt.Errorf("unable to estimate send transatcion of %s to %s: %v", coin, address, err)
	}
	fee, _ := pipToBIP(res.Commission).Float64()
	return fee, nil
}

func (ma *MinterAdapter) Send(ctx context.Context, w Wallet, coin string, amount float64, address string) (string, error) {
	var (
		signedTransaction transaction.SignedTransaction
		err               error
		nonce             = ^uint64(0)
		prevNonce         = ^uint64(0)
	)
	for {
		signedTransaction, nonce, err = ma.prepareSendTx(w, coin, amount, address)
		if err != nil {
			return "", fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
		}
		if prevNonce == nonce {
			continue
		}
		prevNonce = nonce

		res, err := ma.client.SendTransaction(signedTransaction)
		if err != nil {
			err = tryParseError(err)
			if isTransactionInMempool(err) {
				time.Sleep(1 * time.Second)
				continue
			}
			return "", fmt.Errorf("unable to send coin %s to %s: %v", coin, address, err)
		}
		return "Mt" + strings.ToLower(res.Hash), nil
	}
}

func (ma *MinterAdapter) prepareSendTx(w Wallet, coin string, amount float64, address string) (transaction.SignedTransaction, uint64, error) {
	data, _ := transaction.NewSendData().
		SetCoin(coin).
		SetValue(bipToCoin(amount)).
		SetTo(address)

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return nil, 0, fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return nil, 0, err
	}

	nonce, err := ma.client.Nonce(mntWallet.Address())
	if err != nil {
		return nil, 0, fmt.Errorf("unable to generate nonce: %v", err)
	}

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin("BIP").
		SetGasPrice(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey())
	return signedTransaction, nonce, nil
}

func (ma *MinterAdapter) EstimateMultiSendFee(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, error) {
	signedTransaction, _, err := ma.prepareMultiSendTx(w, coin, addresses, amounts)
	if err != nil {
		return 0, fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
	}
	res, err := ma.client.EstimateTxCommission(signedTransaction)
	if err != nil {
		return 0, fmt.Errorf("unable to estimate multisend transaction of %s to %v: %v", coin, amounts, err)
	}
	fee, _ := pipToBIP(res.Commission).Float64()
	return fee, nil
}

func (ma *MinterAdapter) MultiSend(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	var (
		signedTransaction transaction.SignedTransaction
		err               error
		nonce             = ^uint64(0)
		prevNonce         = ^uint64(0)
	)
	for {
		signedTransaction, nonce, err = ma.prepareMultiSendTx(w, coin, addresses, amounts)
		if err != nil {
			return "", fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
		}
		if prevNonce == nonce {
			continue
		}
		prevNonce = nonce

		res, err := ma.client.SendTransaction(signedTransaction)
		if err != nil {
			err = tryParseError(err)
			if isTransactionInMempool(err) {
				time.Sleep(1 * time.Second)
				continue
			}
			return "", fmt.Errorf("unable to multisendsend coin %s to %v: %v", coin, addresses, err)
		}
		return "Mt" + strings.ToLower(res.Hash), nil
	}
}

func (ma *MinterAdapter) prepareMultiSendTx(w Wallet, coin string, addresses []string, amounts []float64) (transaction.SignedTransaction, uint64, error) {
	data := transaction.NewMultisendData()
	for i := 0; i < len(addresses); i++ {
		data = data.AddItem(*transaction.NewMultisendDataItem().
			SetCoin(coin).SetValue(bipToCoin(amounts[i])).
			MustSetTo(addresses[i]))
	}

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return nil, 0, fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return nil, 0, err
	}

	nonce, err := ma.client.Nonce(mntWallet.Address())
	if err != nil {
		return nil, 0, fmt.Errorf("unable to generate nonce: %v", err)
	}

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin("BIP").
		SetGasPrice(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey())
	return signedTransaction, nonce, nil
}

func (ma *MinterAdapter) Subscribe(ctx context.Context, consumer EventConsumer) error {
	ma.Unsubscribe()
	go func() {
		ma.ticker = time.NewTicker(ma.pollingDuration)
		for range ma.ticker.C {
			st, _ := ma.pollingClient.Status()
			newLastBlockHeight, _ := strconv.ParseInt(st.LatestBlockHeight, 10, 32)
			startBlock := ma.lastBlockHeight + 1
			if ma.lastBlockHeight == 0 {
				startBlock = int(newLastBlockHeight)
			}
			if int(newLastBlockHeight) > ma.lastBlockHeight {
				for blockHeight := startBlock; blockHeight <= int(newLastBlockHeight); blockHeight++ {
					block, _ := ma.pollingClient.Block(blockHeight)
					if block != nil && block.Transactions != nil {
						for _, tx := range block.Transactions {
							bigFee := new(big.Int)
							bigFee, _ = bigFee.SetString(tx.Gas, 10)
							fee, _ := pipToBIP(big.NewInt(0).Mul(bigFee, big.NewInt(1000000000000000)).String()).Float64()
							fee = fee * float64(tx.GasPrice)
							if tx.Type == int(transaction.TypeMultisend) {
								msend := &api.MultisendData{}
								err := tx.Data.FillStruct(msend)
								if err != nil {
									consumer.Consume(Event{Error: err})
								}
								for _, item := range msend.List {
									amount, _ := pipToBIP(item.Value).Float64()
									if err != nil {
										consumer.Consume(Event{Error: err})
									}
									consumer.Consume(Event{
										Hash:    tx.Hash,
										From:    tx.From,
										To:      item.To,
										Coin:    item.Coin,
										Amount:  amount,
										Type:    TypeMultisend,
										FeeCoin: tx.GasCoin,
										Fee:     fee,
									})
								}
							} else if tx.Type == int(transaction.TypeSend) {
								send := &api.SendData{}
								err := tx.Data.FillStruct(send)
								if err != nil {
									consumer.Consume(Event{Error: err})
								}
								amount, _ := pipToBIP(send.Value).Float64()
								if err != nil {
									consumer.Consume(Event{Error: err})
								}
								consumer.Consume(Event{
									Hash:    tx.Hash,
									From:    tx.From,
									To:      send.To,
									Coin:    send.Coin,
									Amount:  amount,
									Type:    TypeSend,
									FeeCoin: tx.GasCoin,
									Fee:     fee,
								})
							} else {
								continue
							}
						}
						lastBlockHeight, _ := strconv.ParseInt(block.Height, 10, 32)
						ma.lastBlockHeight = int(lastBlockHeight)
						for _, listener := range ma.blockListeners {
							listener(ma.lastBlockHeight, len(block.Transactions))
						}
					}
				}
			}
		}
	}()
	return nil
}

func (ma *MinterAdapter) Unsubscribe() {
	if ma.ticker != nil {
		ma.ticker.Stop()
		ma.ticker = nil
	}
}

func (ma *MinterAdapter) GetLastBlockHeight() int {
	return ma.lastBlockHeight
}

func (ma *MinterAdapter) SetLastBlockHeight(lastBlockHeight int) {
	ma.lastBlockHeight = lastBlockHeight
}

func (ma *MinterAdapter) AddBlockListener(listener func(lastBlockHeight int, transactionsCount int)) {
	ma.blockListeners = append(ma.blockListeners, listener)
}

func pipToBIP(pip string) *big.Float {
	weiFloat := new(big.Float)
	weiFloat.SetString(pip)
	return new(big.Float).Quo(weiFloat, big.NewFloat(math.Pow10(BIPDecimal)))
}

func bipToCoin(bip float64) *big.Int {
	value := big.NewFloat(0).Copy(big.NewFloat(bip))
	multiplier := big.NewFloat(math.Pow10(BIPDecimal))
	value.Mul(value, multiplier)
	// value = amount * 10^18
	// amount = 0.0000000123
	// value = 1230000000000
	val, _ := value.Int(big.NewInt(0))
	return val
}

func tryParseError(err error) error {
	if txErr, ok := err.(*api.ResponseError); ok { //Tx already exists in mempool
		response := new(api.SendTransactionResponse)
		err = json.Unmarshal(txErr.Body(), response)
		if err == nil {
			return response.Error
		}
	}
	return err
}

func isTransactionInMempool(err error) bool {
	if txErr, ok := err.(*api.TxError); ok { //Tx already exists in mempool
		if txErr.TxResult.Code == 113 {
			return true
		}
	}
	return false
}
