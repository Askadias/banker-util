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
	USDTContractAddress      = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	MultiSendContractAddress = "0xF62e93Eb9e92A0Cd9B9311ea01f1c74cbA5650F2"
	USDTDecimal              = 6
	GWEIDecimal              = 9
	ETHDecimal               = 18
	ETHDerivationPath        = "m/44'/60'/0'/0/0"
)

type EthereumAdapter struct {
	Client *ethclient.Client
}

func NewEthereumAdapter(ethereumClient *ethclient.Client) *EthereumAdapter {
	return &EthereumAdapter{Client: ethereumClient}
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

func (ea *EthereumAdapter) NewWallet(ctx context.Context) (Wallet, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return Wallet{}, fmt.Errorf("unable to generate mnemonic: %v", err)
	}
	_, account, err := ea.resolveMnemonic(mnemonic)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{account.Address.Hex(), mnemonic}, nil
}

func (ea *EthereumAdapter) GetBalance(ctx context.Context, address string) (map[string]float64, error) {
	balance := map[string]float64{}
	ethBalance, err := ea.Client.PendingBalanceAt(ctx, common.HexToAddress(address))
	if err != nil {
		return balance, err
	}

	float64Value, _ := weiToEther(ethBalance, ETHDecimal).Float64()
	balance["ETH"] = float64Value
	usdtBalance, err := ea.getTokenBalance(USDTContractAddress, address)
	if err != nil {
		return balance, fmt.Errorf("unable to get USDT balance: %v", err)
	}

	float64Value, _ = weiToEther(usdtBalance, USDTDecimal).Float64()
	balance["USDT"] = float64Value
	return balance, nil
}

func (ea *EthereumAdapter) getTokenBalance(tokenContractAddress, address string) (*big.Int, error) {
	caller, err := eth.NewToken(common.HexToAddress(tokenContractAddress), ea.Client)
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
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return 0, fmt.Errorf("unable to get suggested gasPrice: %v", err)
	}
	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))

	if coin == "USDT" {
		data, err := eth.PackTransferData(common.HexToAddress(address), etherToWei(amount, USDTDecimal))
		if err != nil {
			return 0, fmt.Errorf("unable to pack contract data: %v", err)
		}

		to := common.HexToAddress(USDTContractAddress)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: acceleratedGasPrice, Value: big.NewInt(0), Data: data}
		estimatedFee, err := ea.Client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(int64(estimatedFee)), GWEIDecimal).Float64()
		return fee * 2, nil
	} else if coin == "ETH" {
		to := common.HexToAddress(address)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: acceleratedGasPrice, Value: etherToWei(amount, ETHDecimal), Data: nil}
		estimatedFee, err := ea.Client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(int64(estimatedFee)), GWEIDecimal).Float64()
		return fee * 2, nil
	} else {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	}
}

func (ea *EthereumAdapter) Send(ctx context.Context, w Wallet, coin string, amount float64, address string) (string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to parse wallet private key: %v", err)
	}
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to get suggested gasPrice: %v", err)
	}

	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))

	if coin == "USDT" {
		caller, err := eth.NewToken(common.HexToAddress(USDTContractAddress), ea.Client)
		if err != nil {
			return "", fmt.Errorf("unable to bind token contract: %v", err)
		}

		opts := bind.NewKeyedTransactor(key)
		opts.Context = ctx
		opts.GasPrice = acceleratedGasPrice
		tx, err := caller.Transfer(opts, common.HexToAddress(address), etherToWei(amount, USDTDecimal))
		if err != nil {
			return "", fmt.Errorf("unable to sent transaction: %v", err)
		}
		return tx.Hash().Hex(), nil
	} else if coin == "ETH" {
		publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
		if !ok {
			return "", fmt.Errorf("unable to get public key for wallet: %s", w.Address)
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := ea.Client.PendingNonceAt(ctx, fromAddress)
		if err != nil {
			return "", fmt.Errorf("unable to get nonce: %v", err)
		}

		gasLimit := uint64(21000) // in units

		toAddress := common.HexToAddress(address)
		tx := types.NewTransaction(nonce, toAddress, etherToWei(amount, ETHDecimal), gasLimit, acceleratedGasPrice, nil)
		chainID, err := ea.Client.NetworkID(ctx)
		if err != nil {
			return "", fmt.Errorf("unable to get ETH networkID: %v", err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key)
		if err != nil {
			return "", fmt.Errorf("unable to sign transaction: %v", err)
		}

		err = ea.Client.SendTransaction(ctx, signedTx)
		if err != nil {
			return "", fmt.Errorf("unable to send transaction: %v", err)
		}
		return signedTx.Hash().Hex(), nil
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
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return 0, fmt.Errorf("unable to get suggested gasPrice: %v", err)
	}
	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))

	if coin != "USDT" && coin != "ETH" {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	} else {
		var destAddresses []common.Address
		var weiAmounts []*big.Int
		total := big.NewInt(0)
		for i := 0; i< len(addresses); i++ {
			destAddresses = append(destAddresses, common.HexToAddress(addresses[i]))
			if coin == "ETH" {
				weiETH := etherToWei(amounts[i], ETHDecimal)
				total.Add(total, weiETH)
				weiAmounts = append(weiAmounts, weiETH)
			} else if coin == "USDT" {
				weiAmounts = append(weiAmounts, etherToWei(amounts[i], USDTDecimal))
			}
		}
		var data []byte
		var err error
		if coin == "ETH" {
			data, err = eth.PackBulkSendEthData(destAddresses, weiAmounts)
		} else if coin == "USDT" {
			data, err = eth.PackBulkSendTokenData(common.HexToAddress(USDTContractAddress), destAddresses, weiAmounts)
		}
		if err != nil {
			return 0, fmt.Errorf("unable to pack contract data: %v", err)
		}

		to := common.HexToAddress(MultiSendContractAddress)
		msg := ethereum.CallMsg{From: from, To: &to, GasPrice: acceleratedGasPrice, Value: total, Data: data}
		estimatedFee, err := ea.Client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("unable to estimate gas price: %v", err)
		}

		fee, _ := weiToEther(big.NewInt(int64(estimatedFee)), GWEIDecimal).Float64()
		return fee * 2, nil
	}
}

func (ea *EthereumAdapter) MultiSend(ctx context.Context, w Wallet, coin string, addresses []string, amounts []float64) (string, error) {
	key, err := ea.getWalletKey(w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to parse wallet private key: %v", err)
	}
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to get suggested gasPrice: %v", err)
	}

	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))

	if coin != "USDT" && coin != "ETH" {
		return "", fmt.Errorf("coin %s is not supported", coin)
	} else {
		var err error
		caller, err := eth.NewMultisend(common.HexToAddress(MultiSendContractAddress), ea.Client)
		if err != nil {
			return "", fmt.Errorf("unable to bind token contract: %v", err)
		}

		var destAddresses []common.Address
		var weiAmounts []*big.Int
		total := big.NewInt(0)
		for i := 0; i< len(addresses); i++ {
			destAddresses = append(destAddresses, common.HexToAddress(addresses[i]))
			if coin == "ETH" {
				weiETH := etherToWei(amounts[i], ETHDecimal)
				total.Add(total, weiETH)
				weiAmounts = append(weiAmounts, weiETH)
			} else if coin == "USDT" {
				weiAmounts = append(weiAmounts, etherToWei(amounts[i], USDTDecimal))
			}
		}

		opts := bind.NewKeyedTransactor(key)
		opts.Context = ctx
		opts.GasPrice = acceleratedGasPrice
		opts.Value = total

		//estimate, err := ea.EstimateMultiSendFee(ctx, w, coin, amounts)
		//		//if err != nil {
		//		//	return "", err
		//		//}

		//opts.GasLimit = etherToWei(estimate, ETHDecimal).Uint64() * 1000

		var tx *types.Transaction
		if coin == "ETH" {
			tx, err = caller.BulkSendEth(opts, destAddresses, weiAmounts)
		} else if coin == "USDT" {
			tx, err = caller.BulkSendToken(opts, common.HexToAddress(USDTContractAddress), destAddresses, weiAmounts)
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
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return "", "", fmt.Errorf("unable to suggest gas price: %v", err)
	}

	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))
	nonce, err := ea.Client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", "", fmt.Errorf("unable to get nonce: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = acceleratedGasPrice
	address, tx, _, err := eth.DeployMultisend(auth, ea.Client)
	if err != nil {
		return "", "", fmt.Errorf("unable to deploy contract: %v", err)
	}

	return address.Hex(), tx.Hash().Hex(), nil
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
	gasPrice, err := ea.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get suggested gasPrice: %v", err)
	}

	acceleratedGasPrice := gasPrice.Mul(gasPrice, big.NewInt(2))

	caller, err := eth.NewTransaction(common.HexToAddress(MultiSendContractAddress), ea.Client)
	if err != nil {
		return nil, fmt.Errorf("unable to bind token contract: %v", err)
	}

	opts := bind.NewKeyedTransactor(key)
	opts.Context = ctx
	opts.GasPrice = acceleratedGasPrice

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
