package gateway

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway/listener"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"math"
	"math/big"
	"time"
)

const (
	ValuableDecimals = 6
	BIPDecimal = 18.0
)

type MinterAdapter struct {
	baseCoin        string
	client          *http_client.Client
	pollingClient   *http_client.Client
	ticker          *time.Ticker
	pollingDuration time.Duration
	lastBlockHeight uint64
	blockListeners  []func(lastBlockHeight uint64, transactionsCount int)
}

func NewMinterAdapter(minterClient *http_client.Client, minterPollingClient *http_client.Client, pollingDuration time.Duration) *MinterAdapter {
	return &MinterAdapter{baseCoin: "BIP", client: minterClient, pollingClient: minterPollingClient, pollingDuration: pollingDuration}
}

func (ma *MinterAdapter) getWallet(mnemonic string) (*wallet.Wallet, error) {
	mntWallet, err := wallet.Create(mnemonic, "")
	if err != nil {
		return nil, fmt.Errorf("unable create wallet: %v", err)
	}
	return mntWallet, nil
}

func (ma *MinterAdapter) IsValidAddress(_ context.Context, address string) bool {
	return wallet.IsValidAddress(address)
}

func (ma *MinterAdapter) FindWallet(_ context.Context, privateKey string) (Wallet, error) {
	emptyWallet := Wallet{"", "", ""}
	mntWallet, err := ma.getWallet(privateKey)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, mntWallet.Address, privateKey}, nil
}

func (ma *MinterAdapter) NewWallet(_ context.Context) (Wallet, error) {
	mnemonic, err := wallet.NewMnemonic()
	emptyWallet := Wallet{"", "", ""}
	if err != nil {
		return emptyWallet, fmt.Errorf("unable to generate wallet mnemonic: %v", err)
	}
	mntWallet, err := ma.getWallet(mnemonic)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, mntWallet.Address, mnemonic}, nil
}

func (ma *MinterAdapter) GetBalance(c context.Context, address string) (map[string]float64, error) {
	balance := map[string]float64{}
	addr, err := ma.client.Address(api_service.NewAddressParamsWithContext(c).WithAddress(address).Address)
	if err != nil {
		return balance, fmt.Errorf("unable to get balance: %v", err)
	}
	for _, b := range addr.Balance {
		balanceVal, _, err := big.ParseFloat(b.Value, 10, 0, big.ToNegativeInf)
		if err != nil {
			return balance, fmt.Errorf("unable to parse %s balance: %v", b.Coin.Symbol, err)
		}
		quotient := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(BIPDecimal), nil))
		balanceAmount, _ := new(big.Float).Quo(balanceVal, quotient).Float64()
		balance[b.Coin.Symbol] = balanceAmount
	}
	return balance, nil
}

func pstr(val string) *string {
	return &val
}

func (ma *MinterAdapter) EstimateBuy(c context.Context, coin string, amount float64) (float64, float64, error) {
	res, err := ma.client.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithContext(c).
		WithCoinToSell(pstr("BIP")).WithCoinToBuy(&coin).WithValueToBuy(bipToCoin(amount).String()))
	if err != nil {
		return 0, 0, fmt.Errorf("unable to buy coin %s: %v", coin, err)
	}
	fee, _ := pipToBIP(res.Payload.Commission).Float64()
	cost, _ := pipToBIP(res.Payload.WillPay).Float64()

	return cost, fee, nil
}

func (ma *MinterAdapter) Nonce(c context.Context, mntWallet *wallet.Wallet) (uint64, error) {
	addr, err := ma.client.Address(api_service.NewAddressParamsWithContext(c).WithAddress(mntWallet.Address).Address)
	if err != nil {
		return 0, fmt.Errorf("unable to get balance: %v", err)
	}
	return addr.TransactionCount + 1, nil
}

func (ma *MinterAdapter) CoinID(coin string) uint64 {
	if coin == "BIP" {
		return 0
	} else {
		id, _ := ma.client.CoinID(coin)
		return id
	}
}

func (ma *MinterAdapter) Buy(c context.Context, w Wallet, coin string, amount float64) (string, error) {
	cost, _, err := ma.EstimateBuy(c, coin, amount)
	data := transaction.NewBuyCoinData().
		SetCoinToSell(ma.CoinID("BIP")).
		SetCoinToBuy(ma.CoinID(coin)).
		SetMaximumValueToSell(bipToCoin(cost*2)).
		SetValueToBuy(bipToCoin(amount))

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return "", err
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return "", err
	}
	minGas := uint8(gas.MinGasPrice)
	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID("BIP")).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to buy coin coin %s: %v", coin, err)
	}

	encode, err := signedTransaction.Encode()
	if err != nil {
		return "", fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.SendTransaction(encode)
	if err != nil {
		return "", fmt.Errorf("unable to buy coin %s: %v", coin, err)
	}

	return res.Hash, nil
}

func (ma *MinterAdapter) BuySwapPool(c context.Context, w Wallet, coin string, amount float64) (string, error) {
	cost, _, err := ma.EstimateBuy(c, coin, amount)
	data := transaction.NewBuySwapPoolData().
		AddCoin(ma.CoinID("BIP")).
		AddCoin(ma.CoinID(coin)).
		SetMaximumValueToSell(bipToCoin(cost*2)).
		SetValueToBuy(bipToCoin(amount))

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return "", err
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return "", err
	}
	minGas := uint8(gas.MinGasPrice)
	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID("BIP")).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to buy coin coin %s: %v", coin, err)
	}

	encode, err := signedTransaction.Encode()
	if err != nil {
		return "", fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.SendTransaction(encode)
	if err != nil {
		return "", fmt.Errorf("unable to buy coin %s: %v", coin, err)
	}

	return res.Hash, nil
}

func (ma *MinterAdapter) SellAll(c context.Context, w Wallet, coin string) (string, error) {
	data := transaction.NewSellAllCoinData().
		SetCoinToSell(ma.CoinID(coin)).
		SetCoinToBuy(ma.CoinID("BIP"))

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return "", err
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return "", err
	}
	minGas := uint8(gas.MinGasPrice)
	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID(coin)).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to sell coin %s: %v", coin, err)
	}

	encode, err := signedTransaction.Encode()
	if err != nil {
		return "", fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.SendTransaction(encode)
	if err != nil {
		return "", fmt.Errorf("unable to sell coin %s: %v", coin, err)
	}
	return res.Hash, nil
}

func (ma *MinterAdapter) SellAllSwapPool(c context.Context, w Wallet, coin string) (string, error) {
	data := transaction.NewSellAllSwapPoolData().
		AddCoin(ma.CoinID(coin)).
		AddCoin(ma.CoinID("BIP")).
		SetMinimumValueToBuy(new(big.Int))

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return "", err
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return "", err
	}
	minGas := uint8(gas.MinGasPrice)
	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID("BIP")).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to sell coin %s: %v", coin, err)
	}

	encode, err := signedTransaction.Encode()
	if err != nil {
		return "", fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.SendTransaction(encode)
	if err != nil {
		return "", fmt.Errorf("unable to sell coin %s: %v", coin, err)
	}
	return res.Hash, nil
}

func (ma *MinterAdapter) EstimateSendFee(c context.Context, w Wallet, coin string, amount float64, address string) (float64, float64, error) {
	signedTransaction, _, err := ma.prepareSendTx(c, w, coin, amount, address)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
	}
	encode, err := signedTransaction.Encode()
	if err != nil {
		return 0, 0, fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.EstimateTxCommission(encode)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to estimate send transatcion of %s to %s: %v", coin, address, err)
	}
	fee, _ := pipToBIP(res.Commission).Float64()
	return fee, 0, nil
}

func (ma *MinterAdapter) Send(c context.Context, w Wallet, coin string, amount float64, address string) (string, error) {
	var (
		signedTransaction transaction.Signed
		err               error
		nonce             = ^uint64(0)
		prevNonce         = ^uint64(0)
	)
	for {
		signedTransaction, nonce, err = ma.prepareSendTx(c, w, coin, amount, address)
		if err != nil {
			return "", fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
		}
		if prevNonce == nonce {
			continue
		}
		encode, err := signedTransaction.Encode()
		if err != nil {
			return "", fmt.Errorf("unable to encode transaction: %v", err)
		}
		res, err := ma.client.SendTransaction(encode)
		if err != nil {
			if isTransactionInMempool(err) {
				time.Sleep(1 * time.Second)
				continue
			}
			return "", fmt.Errorf("unable to send coin %s to %s: %v", coin, address, err)
		}
		return res.Hash, nil
	}
}

func (ma *MinterAdapter) prepareSendTx(c context.Context, w Wallet, coin string, amount float64, address string) (transaction.Signed, uint64, error) {
	data, _ := transaction.NewSendData().
		SetCoin(ma.CoinID(coin)).
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

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return nil, 0, fmt.Errorf("unable to generate nonce: %v", err)
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return nil, 0, err
	}
	minGas := uint8(gas.MinGasPrice)

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID("BIP")).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	return signedTransaction, nonce, nil
}

func (ma *MinterAdapter) EstimateMultiSendFee(c context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, float64, error) {
	signedTransaction, _, err := ma.prepareMultiSendTx(c, w, coin, addresses, amounts)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
	}
	encode, err := signedTransaction.Encode()
	if err != nil {
		return 0, 0, fmt.Errorf("unable to encode transaction: %v", err)
	}
	res, err := ma.client.EstimateTxCommission(encode)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to estimate multisend transaction of %s to %v: %v", coin, amounts, err)
	}
	fee, _ := pipToBIP(res.Commission).Float64()
	return fee, 0, nil
}

func (ma *MinterAdapter) MultiSend(c context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	var (
		signedTransaction transaction.Signed
		err               error
		nonce             = ^uint64(0)
		prevNonce         = ^uint64(0)
	)
	for {
		signedTransaction, nonce, err = ma.prepareMultiSendTx(c, w, coin, addresses, amounts)
		if err != nil {
			return "", fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
		}
		if prevNonce == nonce {
			continue
		}
		prevNonce = nonce

		encode, err := signedTransaction.Encode()
		if err != nil {
			return "", fmt.Errorf("unable to encode transaction: %v", err)
		}
		res, err := ma.client.SendTransaction(encode)
		if err != nil {
			if isTransactionInMempool(err) {
				time.Sleep(1 * time.Second)
				continue
			}
			return "", fmt.Errorf("unable to multisendsend coin %s to %v: %v", coin, addresses, err)
		}
		return res.Hash, nil
	}
}

func (ma *MinterAdapter) prepareMultiSendTx(c context.Context, w Wallet, coin string, addresses []string, amounts []float64) (transaction.Signed, uint64, error) {
	data := transaction.NewMultisendData()
	coinID := ma.CoinID(coin)
	for i := 0; i < len(addresses); i++ {
		item , err := transaction.NewSendData().
			SetCoin(coinID).
			SetValue(bipToCoin(amounts[i])).
			SetTo(addresses[i])
		if err != nil {
			return nil, 0, fmt.Errorf("unable to prepare transaction: %v", err)
		}
		data = data.AddItem(item)
	}

	newTransaction, err := transaction.NewBuilder(transaction.MainNetChainID).NewTransaction(data)
	if err != nil {
		return nil, 0, fmt.Errorf("unable to prepare transaction: %v", err)
	}
	mntWallet, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return nil, 0, err
	}

	nonce, err := ma.Nonce(c, mntWallet)
	if err != nil {
		return nil, 0, fmt.Errorf("unable to generate nonce: %v", err)
	}
	gas, err := ma.client.MinGasPrice()
	if err != nil {
		return nil, 0, err
	}
	minGas := uint8(gas.MinGasPrice)

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin(ma.CoinID("BIP")).
		SetGasPrice(minGas).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign(mntWallet.PrivateKey)
	return signedTransaction, nonce, nil
}

func (ma *MinterAdapter) IsTransactionComplete(c context.Context, hash string) bool {
	tx, err := ma.client.Transaction(api_service.NewTransactionParamsWithContext(c).WithHash(hash).Hash)
	return err == nil && tx != nil && tx.Nonce > 0
}

func (ma *MinterAdapter) Subscribe(_ context.Context, consumer listener.EventConsumer) error {
	ma.Unsubscribe()
	go func() {
		ma.ticker = time.NewTicker(ma.pollingDuration)
		for range ma.ticker.C {
			st, _ := ma.pollingClient.Status()
			if st == nil {
				continue
			}
			newLastBlockHeight := st.LatestBlockHeight
			startBlock := ma.lastBlockHeight + 1
			if ma.lastBlockHeight == 0 {
				startBlock = newLastBlockHeight
			}
			if newLastBlockHeight > ma.lastBlockHeight {
				for blockHeight := startBlock; blockHeight <= newLastBlockHeight; blockHeight++ {
					block, _ := ma.pollingClient.Block(blockHeight)
					if block != nil && block.Transactions != nil {
						for _, tx := range block.Transactions {
							bigFee := new(big.Int)
							bigFee.SetUint64(tx.Gas)
							fee, _ := pipToBIP(big.NewInt(0).Mul(bigFee, big.NewInt(1000000000000000)).String()).Float64()
							fee = fee * float64(tx.GasPrice)
							if tx.Type == uint64(transaction.TypeMultisend) {
								msend := new(models.MultiSendData)
								if err := tx.Data.UnmarshalTo(msend); err != nil {
									consumer.Consume(listener.Event{Error: err})
								}
								var items []listener.SendEvent
								for _, item := range msend.List {
									amount, _ := pipToBIP(item.Value).Float64()
									items = append(items, listener.SendEvent{
										To:     item.To,
										Coin:   item.Coin.Symbol,
										Amount: amount,
									})
								}
								consumer.Consume(listener.Event{
									Type:    listener.TypeMultisend,
									Hash:    tx.Hash,
									From:    tx.From,
									FeeCoin: tx.GasCoin.Symbol,
									Fee:     fee,
									Items:   items,
								})
							} else if tx.Type == uint64(transaction.TypeSend) {
								send := new(models.SendData)
								if err := tx.Data.UnmarshalTo(send); err != nil {
									consumer.Consume(listener.Event{Error: err})
								}
								amount, _ := pipToBIP(send.Value).Float64()
								consumer.Consume(listener.Event{Type: listener.TypeSend,
									Hash:    tx.Hash,
									From:    tx.From,
									FeeCoin: tx.GasCoin.Symbol,
									Fee:     fee,
									SendEvent: listener.SendEvent{
										To:     send.To,
										Coin:   send.Coin.Symbol,
										Amount: amount,
									}})
							} else if tx.Type == uint64(transaction.TypeBuyCoin) {
								buy := new(models.BuyCoinData)
								if err := tx.Data.UnmarshalTo(buy); err != nil {
									consumer.Consume(listener.Event{Error: err})
								}
								amount, _ := pipToBIP(buy.ValueToBuy).Float64()
								consumer.Consume(listener.Event{Type: listener.TypeBuy,
									Hash:    tx.Hash,
									From:    tx.From,
									FeeCoin: tx.GasCoin.Symbol,
									Fee:     fee,
									SendEvent: listener.SendEvent{
										ToCoin: buy.CoinToBuy.Symbol,
										Coin:   buy.CoinToSell.Symbol,
										Amount: amount,
									}})
							} else {
								continue
							}
						}
						ma.lastBlockHeight = block.Height
						for _, lnr := range ma.blockListeners {
							lnr(ma.lastBlockHeight, len(block.Transactions))
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

func (ma *MinterAdapter) GetLastBlockHeight() uint64 {
	return ma.lastBlockHeight
}

func (ma *MinterAdapter) SetLastBlockHeight(lastBlockHeight uint64) {
	ma.lastBlockHeight = lastBlockHeight
}

func (ma *MinterAdapter) AddBlockListener(listener func(lastBlockHeight uint64, transactionsCount int)) {
	ma.blockListeners = append(ma.blockListeners, listener)
}

func pipToBIP(pip string) *big.Float {
	weiFloat := new(big.Float)
	weiFloat.SetString(pip)
	return new(big.Float).Quo(weiFloat, big.NewFloat(math.Pow10(BIPDecimal)))
}

func bipToCoin(bip float64) *big.Int {
	value := big.NewFloat(bip)
	value.Mul(value, big.NewFloat(math.Pow10(ValuableDecimals)))
	result, _ := value.Int(nil)
	result.Mul(result, big.NewInt(int64(math.Pow10(BIPDecimal-ValuableDecimals))))
	return result
}

func isTransactionInMempool(err error) bool {
	_, body, e := http_client.ErrorBody(err)
	if e == nil && body != nil && body.Error.Code == "113" { //Tx already exists in mempool
		return true
	}
	return false
}
