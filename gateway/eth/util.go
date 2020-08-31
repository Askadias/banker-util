package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

var (
	GasPriceKey = "GasPrice"
)

func GetGasPrice(ctx context.Context) float64 {
	value := ctx.Value(GasPriceKey)
	if gasPrice, ok := value.(float64); ok {
		return gasPrice
	} else {
		return 0.0
	}
}

func WithGasPrice(ctx context.Context, gasPrice float64) context.Context {
	return context.WithValue(ctx, GasPriceKey, gasPrice)
}

var TokenABIJSON, _ = abi.JSON(strings.NewReader(TokenABI))

func PackTransferData(_to common.Address, _value *big.Int) ([]byte, error) {
	return TokenABIJSON.Pack("transfer", _to, _value)
}

type TransferInput struct {
	Recipient common.Address
	Amount    *big.Int
}

func UnpackTransferData(data []byte) (TransferInput, error) {
	var transferInput TransferInput
	method, err := TokenABIJSON.MethodById(data)
	if err != nil {
		return TransferInput{}, err
	}
	if method.Name != "transfer" {
		return TransferInput{}, fmt.Errorf("wrong method")
	}
	err = method.Inputs.Unpack(&transferInput, data[4:])
	if err != nil {
		return TransferInput{}, err
	}
	return transferInput, nil
}

type BulkSendETHInput struct {
	Addresses []common.Address
	Amounts   []*big.Int
}

type BulkSendTokenInput struct {
	Token     common.Address
	Addresses []common.Address
	Amounts   []*big.Int
}

var MultisendABIJSON, _ = abi.JSON(strings.NewReader(MultisendABI))

func PackBulkSendEthData(addresses []common.Address, amounts []*big.Int) ([]byte, error) {
	return MultisendABIJSON.Pack("bulkSendEth", addresses, amounts)
}

func PackBulkSendTokenData(tokenAddr common.Address, addresses []common.Address, amounts []*big.Int) ([]byte, error) {
	return MultisendABIJSON.Pack("bulkSendToken", tokenAddr, addresses, amounts)
}

func GetMultisendMethod(data []byte) (string, error) {
	if method, err := MultisendABIJSON.MethodById(data); err != nil {
		return "", err
	} else {
		return method.Name, nil
	}
}

func UnpackBulkSendETHData(data []byte) (BulkSendETHInput, error) {
	method, err := MultisendABIJSON.MethodById(data)
	if err != nil {
		return BulkSendETHInput{}, err
	}
	if method.Name != "bulkSendEth" {
		return BulkSendETHInput{}, fmt.Errorf("wrong method")
	}
	var bulkSendETHInput BulkSendETHInput
	err = method.Inputs.Unpack(&bulkSendETHInput, data[4:])
	if err != nil {
		return BulkSendETHInput{}, err
	}
	return bulkSendETHInput, nil
}

func UnpackBulkSendTokenData(data []byte) (BulkSendTokenInput, error) {
	method, err := MultisendABIJSON.MethodById(data)
	if err != nil {
		return BulkSendTokenInput{}, err
	}
	if method.Name != "bulkSendToken" {
		return BulkSendTokenInput{}, fmt.Errorf("wrong method")
	}
	var bulkSendTokenInput BulkSendTokenInput
	err = method.Inputs.Unpack(&bulkSendTokenInput, data[4:])
	if err != nil {
		return BulkSendTokenInput{}, err
	}
	return bulkSendTokenInput, nil
}
