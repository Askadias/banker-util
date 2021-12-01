package gateway

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway/listener"
	"github.com/cosmos/cosmos-sdk/x/auth"

	sdk "github.com/cosmos/cosmos-sdk/types"

	decapi "bitbucket.org/decimalteam/decimal-go-sdk/api"
	"bitbucket.org/decimalteam/decimal-go-sdk/wallet"
	types "bitbucket.org/decimalteam/go-node/x/coin"
	"math"
	"math/big"
	"time"
)

const (
	DELGasPrice = 0.001
	DELDecimal  = 18.0
)

type DecimalAdapter struct {
	baseCoin        string
	client          *decapi.API
	pollingClient   *decapi.API
	ticker          *time.Ticker
	pollingDuration time.Duration
	lastBlockHeight uint64
	blockListeners  []func(lastBlockHeight uint64, transactionsCount int)
}

func NewDecimalAdapter(decimalClient *decapi.API, decimalPollingClient *decapi.API, pollingDuration time.Duration) *DecimalAdapter {
	return &DecimalAdapter{baseCoin: "DEL", client: decimalClient, pollingClient: decimalPollingClient, pollingDuration: pollingDuration}
}

func (ma *DecimalAdapter) getWallet(mnemonic string) (*wallet.Account, error) {
	decWallet, err := wallet.NewAccountFromMnemonicWords(mnemonic, "")
	if err != nil {
		return nil, fmt.Errorf("unable create wallet: %v", err)
	}
	return decWallet, nil
}

func (ma *DecimalAdapter) IsValidAddress(_ context.Context, address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	return err == nil
}

func (ma *DecimalAdapter) FindWallet(_ context.Context, privateKey string) (Wallet, error) {
	emptyWallet := Wallet{"", "", ""}
	decWallet, err := ma.getWallet(privateKey)
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, decWallet.Address(), privateKey}, nil
}

func (ma *DecimalAdapter) NewWallet(_ context.Context) (Wallet, error) {
	mnemonic, err := wallet.NewMnemonic(256, "")
	emptyWallet := Wallet{"", "", ""}
	if err != nil {
		return emptyWallet, fmt.Errorf("unable to generate wallet mnemonic: %v", err)
	}
	decWallet, err := ma.getWallet(mnemonic.Words())
	if err != nil {
		return emptyWallet, err
	}
	return Wallet{ma.baseCoin, decWallet.Address(), mnemonic.Words()}, nil
}

func (ma *DecimalAdapter) GetBalance(_ context.Context, address string) (map[string]float64, error) {
	balance := map[string]float64{}
	addr, err := ma.client.Address(address)
	if err != nil {
		return balance, fmt.Errorf("unable to get balance: %v", err)
	}
	for coin, value := range addr.Balance {
		balanceVal, _, err := big.ParseFloat(value, 10, 0, big.ToNegativeInf)
		if err != nil {
			return balance, fmt.Errorf("unable to parse %s balance: %v", coin, err)
		}
		quotient := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(DELDecimal), nil))
		balanceAmount, _ := new(big.Float).Quo(balanceVal, quotient).Float64()
		balance[coin] = balanceAmount
	}
	return balance, nil
}

func (ma *DecimalAdapter) SellAll(c context.Context, w Wallet, coin string) (string, error) {
	account, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	tx, err := ma.prepareSellAllTx(c, account, coin)
	if err != nil {
		return "", fmt.Errorf("unable to prepare transaction: %v", err)
	}

	res, err := ma.client.BroadcastSignedTransactionJSON(tx, account)
	if err != nil {
		return "", fmt.Errorf("unable to sell coin %s: %v", coin, err)
	}

	return res.TxHash, nil
}

func (ma *DecimalAdapter) prepareSellAllTx(_ context.Context, account *wallet.Account, coin string) (auth.StdTx, error) {
	buyer, err := sdk.AccAddressFromBech32(account.Address())
	if err != nil {
		return auth.StdTx{}, err
	}

	msg := decapi.NewMsgSellAllCoin(buyer, sdk.NewCoin(coin, sdk.NewInt(0)), sdk.NewCoin("DEL", sdk.NewInt(0)))

	msgs := []sdk.Msg{msg}
	feeCoins := sdk.NewCoins(sdk.NewCoin("DEL", sdk.NewInt(0)))

	signedTransaction, err := ma.client.NewSignedTransaction(msgs, feeCoins, "", account)
	return signedTransaction, nil
}

func (ma *DecimalAdapter) EstimateSendFee(c context.Context, w Wallet, coin string, amount float64, address string) (float64, float64, error) {
	account, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return 0, 0, err
	}

	signedTransaction, err := ma.prepareSendTx(c, account, coin, amount, address)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
	}

	gasWanted, err := ma.client.EstimateTransactionGasWanted(signedTransaction)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to estimate send transatcion of %s to %s: %v", coin, address, err)
	}

	return float64(gasWanted) * DELGasPrice, 0, nil
}

func (ma *DecimalAdapter) Send(c context.Context, w Wallet, coin string, amount float64, address string) (string, error) {
	account, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	tx, err := ma.prepareSendTx(c, account, coin, amount, address)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to send coin %s to %s: %v", coin, address, err)
	}

	res, err := ma.client.BroadcastSignedTransactionJSON(tx, account)
	if err != nil {
		return "", fmt.Errorf("unable to send coin %s to %s: %v", coin, address, err)
	}

	return res.TxHash, nil
}

func (ma *DecimalAdapter) prepareSendTx(_ context.Context, account *wallet.Account, coin string, amount float64, address string) (auth.StdTx, error) {
	sender, err := sdk.AccAddressFromBech32(account.Address())
	if err != nil {
		return auth.StdTx{}, err
	}
	receiver, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return auth.StdTx{}, err
	}

	msg := decapi.NewMsgSendCoin(sender, sdk.NewCoin(coin, delToCoin(amount)), receiver)

	msgs := []sdk.Msg{msg}
	feeCoins := sdk.NewCoins(sdk.NewCoin("DEL", sdk.NewInt(0)))

	signedTransaction, err := ma.client.NewSignedTransaction(msgs, feeCoins, "", account)
	return signedTransaction, nil
}

func (ma *DecimalAdapter) EstimateMultiSendFee(c context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, float64, error) {
	account, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return 0, 0, err
	}

	tx, err := ma.prepareMultiSendTx(c, account, coin, addresses, amounts)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
	}

	gasWanted, err := ma.client.EstimateTransactionGasWanted(tx)
	if err != nil {
		return 0, 0, fmt.Errorf("unable to estimate multisend transaction of %s to %v: %v", coin, amounts, err)
	}

	return float64(gasWanted) * DELGasPrice, 0, nil
}

func (ma *DecimalAdapter) MultiSend(c context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	account, err := ma.getWallet(w.PrivateKey)
	if err != nil {
		return "", err
	}

	tx, err := ma.prepareMultiSendTx(c, account, coin, addresses, amounts)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction to multisend coin %s to %v: %v", coin, amounts, err)
	}

	res, err := ma.client.BroadcastSignedTransactionJSON(tx, account)
	if err != nil {
		return "", fmt.Errorf("unable to multisendsend coin %s to %v: %v", coin, addresses, err)
	}

	return res.TxHash, nil
}

func (ma *DecimalAdapter) prepareMultiSendTx(_ context.Context, account *wallet.Account, coin string, addresses []string, amounts []float64) (auth.StdTx, error) {
	sender, err := sdk.AccAddressFromBech32(account.Address())
	if err != nil {
		return auth.StdTx{}, err
	}
	var sends []types.Send
	for i := 0; i < len(addresses); i++ {
		receiver, err := sdk.AccAddressFromBech32(addresses[i])
		if err != nil {
			return auth.StdTx{}, err
		}
		sends = append(sends, types.Send{Coin: sdk.NewCoin(coin, delToCoin(amounts[i])), Receiver: receiver})
	}

	msg := decapi.NewMsgMultiSendCoin(sender, sends)

	msgs := []sdk.Msg{msg}
	feeCoins := sdk.NewCoins(sdk.NewCoin("DEL", sdk.NewInt(0)))

	signedTransaction, err := ma.client.NewSignedTransaction(msgs, feeCoins, "", account)
	return signedTransaction, nil
}

func (ma *DecimalAdapter) IsTransactionComplete(_ context.Context, hash string) bool {
	tx, err := ma.client.Transaction(hash)
	return err == nil && tx != nil && tx.TxResult.Code == 0
}

func (ma *DecimalAdapter) Subscribe(_ context.Context, consumer listener.EventConsumer) error {
	ma.Unsubscribe()
	go func() {
		ma.ticker = time.NewTicker(ma.pollingDuration)
		//for range ma.ticker.C {
		//	st, _ := ma.pollingClient.()
		//	if st == nil {
		//		continue
		//	}
		//	newLastBlockHeight := st.LatestBlockHeight
		//	startBlock := ma.lastBlockHeight + 1
		//	if ma.lastBlockHeight == 0 {
		//		startBlock = newLastBlockHeight
		//	}
		//	if newLastBlockHeight > ma.lastBlockHeight {
		//		for blockHeight := startBlock; blockHeight <= newLastBlockHeight; blockHeight++ {
		//			block, _ := ma.pollingClient.Block(blockHeight)
		//			if block != nil && block.Transactions != nil {
		//				for _, tx := range block.Transactions {
		//					bigFee := new(big.Int)
		//					bigFee.SetUint64(tx.Gas)
		//					fee, _ := pipToDEL(big.NewInt(0).Mul(bigFee, big.NewInt(1000000000000000)).String()).Float64()
		//					fee = fee * float64(tx.GasPrice)
		//					if tx.Type == uint64(transaction.TypeMultisend) {
		//						msend := new(models.MultiSendData)
		//						if err := tx.Data.UnmarshalTo(msend); err != nil {
		//							consumer.Consume(listener.Event{Error: err})
		//						}
		//						var items []listener.SendEvent
		//						for _, item := range msend.List {
		//							amount, _ := pipToDEL(item.Value).Float64()
		//							items = append(items, listener.SendEvent{
		//								To:     item.To,
		//								Coin:   item.Coin.Symbol,
		//								Amount: amount,
		//							})
		//						}
		//						consumer.Consume(listener.Event{
		//							Type:    listener.TypeMultisend,
		//							Hash:    tx.Hash,
		//							From:    tx.From,
		//							FeeCoin: tx.GasCoin.Symbol,
		//							Fee:     fee,
		//							Items:   items,
		//						})
		//					} else if tx.Type == uint64(transaction.TypeSend) {
		//						send := new(models.SendData)
		//						if err := tx.Data.UnmarshalTo(send); err != nil {
		//							consumer.Consume(listener.Event{Error: err})
		//						}
		//						amount, _ := pipToDEL(send.Value).Float64()
		//						consumer.Consume(listener.Event{Type: listener.TypeSend,
		//							Hash:    tx.Hash,
		//							From:    tx.From,
		//							FeeCoin: tx.GasCoin.Symbol,
		//							Fee:     fee,
		//							SendEvent: listener.SendEvent{
		//								To:     send.To,
		//								Coin:   send.Coin.Symbol,
		//								Amount: amount,
		//							}})
		//					} else if tx.Type == uint64(transaction.TypeBuyCoin) {
		//						buy := new(models.BuyCoinData)
		//						if err := tx.Data.UnmarshalTo(buy); err != nil {
		//							consumer.Consume(listener.Event{Error: err})
		//						}
		//						amount, _ := pipToDEL(buy.ValueToBuy).Float64()
		//						consumer.Consume(listener.Event{Type: listener.TypeBuy,
		//							Hash:    tx.Hash,
		//							From:    tx.From,
		//							FeeCoin: tx.GasCoin.Symbol,
		//							Fee:     fee,
		//							SendEvent: listener.SendEvent{
		//								ToCoin: buy.CoinToBuy.Symbol,
		//								Coin:   buy.CoinToSell.Symbol,
		//								Amount: amount,
		//							}})
		//					} else if tx.Type == uint64(transaction.TypeBuySwapPool) {
		//						buy := new(models.BuySwapPoolData)
		//						if err := tx.Data.UnmarshalTo(buy); err != nil {
		//							consumer.Consume(listener.Event{Error: err})
		//						}
		//						amount, _ := pipToDEL(buy.ValueToBuy).Float64()
		//						consumer.Consume(listener.Event{Type: listener.TypeBuy,
		//							Hash:    tx.Hash,
		//							From:    tx.From,
		//							FeeCoin: tx.GasCoin.Symbol,
		//							Fee:     fee,
		//							SendEvent: listener.SendEvent{
		//								Coin:   buy.Coins[0].Symbol,
		//								ToCoin: buy.Coins[1].Symbol,
		//								Amount: amount,
		//							}})
		//					} else {
		//						continue
		//					}
		//				}
		//				ma.lastBlockHeight = block.Height
		//				for _, lnr := range ma.blockListeners {
		//					lnr(ma.lastBlockHeight, len(block.Transactions))
		//				}
		//			}
		//		}
		//	}
		//}
	}()
	return nil
}

func (ma *DecimalAdapter) Unsubscribe() {
	if ma.ticker != nil {
		ma.ticker.Stop()
		ma.ticker = nil
	}
}

func (ma *DecimalAdapter) GetLastBlockHeight() uint64 {
	return ma.lastBlockHeight
}

func (ma *DecimalAdapter) SetLastBlockHeight(lastBlockHeight uint64) {
	ma.lastBlockHeight = lastBlockHeight
}

func (ma *DecimalAdapter) AddBlockListener(listener func(lastBlockHeight uint64, transactionsCount int)) {
	ma.blockListeners = append(ma.blockListeners, listener)
}

func pipToDEL(pip string) *big.Float {
	weiFloat := new(big.Float)
	weiFloat.SetString(pip)
	return new(big.Float).Quo(weiFloat, big.NewFloat(math.Pow10(DELDecimal)))
}

func delToCoin(del float64) sdk.Int {
	value := big.NewFloat(del)
	value.Mul(value, big.NewFloat(math.Pow10(ValuableDecimals)))
	result, _ := value.Int(nil)
	result.Mul(result, big.NewInt(int64(math.Pow10(DELDecimal-ValuableDecimals))))
	return sdk.NewIntFromBigInt(result)
}
