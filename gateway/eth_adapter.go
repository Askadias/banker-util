package gateway

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Askadias/banker-util/gateway/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"math"
	"math/big"
)

const (
	// TokenABI is the input ABI used to generate the binding from.
	MultiSendContractAddress = "0x96AF6B6c38636512075754066327d96F5cEDc81c"
	GWEIDecimal              = 9
	ETHDecimal               = 18
	ETHDerivationPath        = "m/44'/60'/0'/0/0"
)

type tokenConfig struct {
	address  string
	decimals int
}

var (
	tokens = map[string]tokenConfig{
		"USDT": {"0xdac17f958d2ee523a2206206994597c13d831ec7", 6},
	}
	MaxUint256 = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
)

type EthereumAdapter struct {
	baseCoin string
	client   *ethclient.Client
	gasPrice *big.Int
}

func NewEthereumAdapter(ethereumClient *ethclient.Client, gasPriceGwei float64) *EthereumAdapter {
	return &EthereumAdapter{
		baseCoin: "ETH",
		client:   ethereumClient,
		gasPrice: etherToWei(gasPriceGwei, GWEIDecimal),
	}
}

func (ea *EthereumAdapter) resolveMnemonic(mnemonic string) (*hdwallet.Wallet, accounts.Account, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, accounts.Account{}, fmt.Errorf("unable to get wallet from mnemonic: %v", err)
	}

	path := hdwallet.MustParseDerivationPath(ETHDerivationPath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, accounts.Account{}, fmt.Errorf("unable to derive a new account at path: %v", err)
	}
	return wallet, account, nil
}

func (ea *EthereumAdapter) getWalletKey(privateKey string) (*ecdsa.PrivateKey, error) {
	if bip39.IsMnemonicValid(privateKey) {
		wallet, account, err := ea.resolveMnemonic(privateKey)
		if err != nil {
			return nil, err
		}
		return wallet.PrivateKey(account)
	} else {
		return crypto.HexToECDSA(privateKey)
	}
}

func (ea *EthereumAdapter) FindWallet(_ context.Context, privateKey string) (Wallet, error) {
	emptyWallet := Wallet{"", "", ""}
	key, err := ea.getWalletKey(privateKey)
	if err != nil {
		return emptyWallet, fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return emptyWallet, fmt.Errorf("unable to get public key for wallet")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return Wallet{ea.baseCoin, address.Hex(), privateKey}, nil
}

func (ea *EthereumAdapter) NewWallet(_ context.Context) (Wallet, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return Wallet{}, fmt.Errorf("unable to generate mnemonic: %v", err)
	}
	_, account, err := ea.resolveMnemonic(mnemonic)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{ea.baseCoin, account.Address.Hex(), mnemonic}, nil
}

func (ea *EthereumAdapter) GetBalance(ctx context.Context, address string) (map[string]float64, error) {
	balance := map[string]float64{}
	ethBalance, err := ea.client.PendingBalanceAt(ctx, common.HexToAddress(address))
	if err != nil {
		return balance, err
	}

	float64Value, _ := weiToEther(ethBalance, ETHDecimal).Float64()
	balance["ETH"] = float64Value
	for token, conf := range tokens {
		tokenBalance, err := ea.getTokenBalance(conf.address, address)
		if err != nil {
			return balance, fmt.Errorf("unable to get USDT balance: %v", err)
		}

		float64Value, _ = weiToEther(tokenBalance, conf.decimals).Float64()
		if float64Value > 0 {
			balance[token] = float64Value
		}
	}
	return balance, nil
}

func (ea *EthereumAdapter) getTokenBalance(tokenContractAddress, address string) (*big.Int, error) {
	caller, err := eth.NewToken(common.HexToAddress(tokenContractAddress), ea.client)
	if err != nil {
		return nil, fmt.Errorf("unable to bind token contract: %v", err)
	}
	return caller.BalanceOf(nil, common.HexToAddress(address))
}

func (ea *EthereumAdapter) EstimateSendFee(ctx context.Context, w Wallet, coin string, amount float64, address string) (float64, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return 0, fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return 0, fmt.Errorf("unable to get public key for wallet: %s", w.Address)
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	if coin == "ETH" {
		to := common.HexToAddress(address)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: ea.gasPrice, Value: etherToWei(amount, ETHDecimal), Data: nil}
		estimatedFee, err := ea.client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(0).Mul(big.NewInt(int64(estimatedFee)), ea.gasPrice), ETHDecimal).Float64()
		return fee, nil
	} else if tokenConf, ok := tokens[coin]; ok {
		data, err := eth.PackTransferData(common.HexToAddress(address), etherToWei(amount, tokenConf.decimals))
		if err != nil {
			return 0, fmt.Errorf("unable to pack contract data: %v", err)
		}

		to := common.HexToAddress(tokenConf.address)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: ea.gasPrice, Value: big.NewInt(0), Data: data}
		estimatedFee, err := ea.client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(0).Mul(big.NewInt(int64(estimatedFee)), ea.gasPrice), ETHDecimal).Float64()
		return fee, nil
	} else {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	}
}

func (ea *EthereumAdapter) Send(ctx context.Context, w Wallet, coin string, amount float64, address string) (string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	if coin == "ETH" {
		publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
		if !ok {
			return "", fmt.Errorf("unable to get public key for wallet: %s", w.Address)
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := ea.client.PendingNonceAt(ctx, fromAddress)
		if err != nil {
			return "", fmt.Errorf("unable to get nonce: %v", err)
		}

		gasLimit := uint64(21000) // in units
		toAddress := common.HexToAddress(address)
		tx := types.NewTransaction(nonce, toAddress, etherToWei(amount, ETHDecimal), gasLimit, ea.gasPrice, nil)
		chainID, err := ea.client.NetworkID(ctx)
		if err != nil {
			return "", fmt.Errorf("unable to get ETH networkID: %v", err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key)
		if err != nil {
			return "", fmt.Errorf("unable to sign transaction: %v", err)
		}

		err = ea.client.SendTransaction(ctx, signedTx)
		if err != nil {
			return "", fmt.Errorf("unable to send transaction: %v", err)
		}

		return signedTx.Hash().Hex(), nil
	} else if tokenConf, ok := tokens[coin]; ok {
		caller, err := eth.NewToken(common.HexToAddress(tokenConf.address), ea.client)
		if err != nil {
			return "", fmt.Errorf("unable to bind token contract: %v", err)
		}

		opts := bind.NewKeyedTransactor(key)
		opts.Context = ctx
		opts.GasPrice = ea.gasPrice
		tx, err := caller.Transfer(opts, common.HexToAddress(address), etherToWei(amount, tokenConf.decimals))
		if err != nil {
			return "", fmt.Errorf("unable to sent transaction: %v", err)
		}

		return tx.Hash().Hex(), nil
	} else {
		return "", fmt.Errorf("coin %s is not supported", coin)
	}
}

func (ea *EthereumAdapter) EstimateMultiSendFee(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (float64, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return 0, fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return 0, fmt.Errorf("unable to get public key for wallet: %s", w.Address)
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	if tokenConf, ok := tokens[coin]; !ok && coin != "ETH" {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	} else {
		var destAddresses []common.Address
		var weiAmounts []*big.Int
		total := big.NewInt(0)
		for i := 0; i < len(addresses); i++ {
			destAddresses = append(destAddresses, common.HexToAddress(addresses[i]))
			if coin == "ETH" {
				weiETH := etherToWei(amounts[i], ETHDecimal)
				total.Add(total, weiETH)
				weiAmounts = append(weiAmounts, weiETH)
			} else {
				weiAmounts = append(weiAmounts, etherToWei(amounts[i], tokenConf.decimals))
			}
		}
		var data []byte
		var err error
		if coin == "ETH" {
			data, err = eth.PackBulkSendEthData(destAddresses, weiAmounts)
		} else {
			data, err = eth.PackBulkSendTokenData(common.HexToAddress(tokenConf.address), destAddresses, weiAmounts)
		}
		if err != nil {
			return 0, fmt.Errorf("unable to pack contract data: %v", err)
		}

		to := common.HexToAddress(MultiSendContractAddress)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: ea.gasPrice, Value: total, Data: data}
		estimatedFee, err := ea.client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(0).Mul(big.NewInt(int64(estimatedFee)), ea.gasPrice), ETHDecimal).Float64()
		return fee, nil
	}
}

func (ea *EthereumAdapter) MultiSend(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	if tokenConf, ok := tokens[coin]; !ok && coin != "ETH" {
		return "", fmt.Errorf("coin %s is not supported", coin)
	} else {
		var err error
		caller, err := eth.NewMultisend(common.HexToAddress(MultiSendContractAddress), ea.client)
		if err != nil {
			return "", fmt.Errorf("unable to bind token contract: %v", err)
		}

		var destAddresses []common.Address
		var weiAmounts []*big.Int
		total := big.NewInt(0)
		for i := 0; i < len(addresses); i++ {
			destAddresses = append(destAddresses, common.HexToAddress(addresses[i]))
			if coin == "ETH" {
				weiETH := etherToWei(amounts[i], ETHDecimal)
				total.Add(total, weiETH)
				weiAmounts = append(weiAmounts, weiETH)
			} else {
				weiAmounts = append(weiAmounts, etherToWei(amounts[i], tokenConf.decimals))
			}
		}

		opts := bind.NewKeyedTransactor(key)
		opts.Context = ctx
		opts.GasPrice = ea.gasPrice
		opts.Value = total

		var tx *types.Transaction
		if coin == "ETH" {
			tx, err = caller.BulkSendEth(opts, destAddresses, weiAmounts)
		} else {
			tx, err = caller.BulkSendToken(opts, common.HexToAddress(tokenConf.address), destAddresses, weiAmounts)
		}
		if err != nil {
			return "", fmt.Errorf("unable to sent transaction: %v", err)
		}
		return tx.Hash().Hex(), nil
	}
}

func (ea *EthereumAdapter) DeployMultiSendContract(ctx context.Context, w Wallet) (string, string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", "", fmt.Errorf("unable to bind token contract: %v", err)
	}

	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("unable to get public key for wallet: %s", w.Address)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ea.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", "", fmt.Errorf("unable to get nonce: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = ea.gasPrice
	address, tx, _, err := eth.DeployMultisend(auth, ea.client)
	if err != nil {
		return "", "", fmt.Errorf("unable to deploy contract: %v", err)
	}

	return address.Hex(), tx.Hash().Hex(), nil
}

func (ea *EthereumAdapter) ApproveTokenMultisend(ctx context.Context, w Wallet, coin string) (string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	if tokenConf, ok := tokens[coin]; ok {
		caller, err := eth.NewToken(common.HexToAddress(tokenConf.address), ea.client)
		if err != nil {
			return "", fmt.Errorf("unable to bind token contract: %v", err)
		}

		opts := bind.NewKeyedTransactor(key)
		opts.Context = ctx
		opts.GasPrice = ea.gasPrice

		fmt.Println("MaxUint256: " + MaxUint256.String())
		tx, err := caller.Approve(opts, common.HexToAddress(MultiSendContractAddress), MaxUint256)
		if err != nil {
			return "", fmt.Errorf("unable to approve transaction: %v", err)
		}

		return tx.Hash().Hex(), nil
	} else {
		return "", fmt.Errorf("coin %s is not supported", coin)
	}
}

func (ea *EthereumAdapter) ConfigureTransactionContract(
	ctx context.Context,
	w Wallet,
	feePerc uint8,
	internalFeePerc uint8,
	rewardPerc uint8,
	affiliatesPerc []uint8,
	partnersPerc []uint8,
	partnersWallets []string,
) ([]string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse wallet private key: %v", err)
	}

	caller, err := eth.NewTransaction(common.HexToAddress(MultiSendContractAddress), ea.client)
	if err != nil {
		return nil, fmt.Errorf("unable to bind token contract: %v", err)
	}

	opts := bind.NewKeyedTransactor(key)
	opts.Context = ctx
	opts.GasPrice = ea.gasPrice

	if len(partnersPerc) != len(partnersWallets) {
		return nil, fmt.Errorf("number of partner wallets should be equal to number of their percentages")
	}

	tx1, err := caller.SetFeePerc(opts, feePerc)
	if err != nil {
		return nil, fmt.Errorf("unable to sent transaction: %v", err)
	}
	tx2, err := caller.SetInternalFeePerc(opts, internalFeePerc)
	if err != nil {
		return nil, fmt.Errorf("unable to sent transaction: %v", err)
	}
	tx3, err := caller.SetRewardPerc(opts, rewardPerc)
	if err != nil {
		return nil, fmt.Errorf("unable to sent transaction: %v", err)
	}
	tx4, err := caller.SetAffiliatesPerc(opts, affiliatesPerc)
	if err != nil {
		return nil, fmt.Errorf("unable to sent transaction: %v", err)
	}
	var partners []common.Address
	for i := 0; i < len(partnersPerc); i++ {
		partners = append(partners, common.HexToAddress(partnersWallets[i]))
	}
	tx5, err := caller.SetPartners(opts, partners, partnersPerc)
	if err != nil {
		return nil, fmt.Errorf("unable to sent transaction: %v", err)
	}
	return []string{
		tx1.Hash().Hex(),
		tx2.Hash().Hex(),
		tx3.Hash().Hex(),
		tx4.Hash().Hex(),
		tx5.Hash().Hex(),
	}, nil
}

func weiToEther(wei *big.Int, decimal int) *big.Float {
	weiFloat := new(big.Float)
	weiFloat.SetString(wei.String())
	return new(big.Float).Quo(weiFloat, big.NewFloat(math.Pow10(decimal)))
}

func etherToWei(ether float64, decimal int) *big.Int {
	weiInt64 := int64(ether * math.Pow(10, float64(decimal)))
	return big.NewInt(weiInt64)
}
