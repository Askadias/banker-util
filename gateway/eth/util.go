package eth

import (
	"context"
	"fmt"
	"github.com/Askadias/banker-util/gateway/listener"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

const (
	// TokenABI is the input ABI used to generate the binding from.
	MultiSendContractAddress = "0x96AF6B6c38636512075754066327d96F5cEDc81c"
	GWEIDecimal              = 9
	Decimal                  = 18
	DerivationPath           = "m/44'/60'/0'/0/0"
	GasPriceKey              = "GasPrice"
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

func PackApproveData(_spender common.Address, _value *big.Int) ([]byte, error) {
	return TokenABIJSON.Pack("approve", _spender, _value)
}

type EventInput struct {
	Address common.Address
	Amount  *big.Int
	Type    listener.EventType
}

func UnpackEventData(data []byte) (EventInput, error) {
	var input EventInput
	method, err := TokenABIJSON.MethodById(data)
	if err != nil {
		return EventInput{}, err
	}
	if method.Name == "transfer" {
		res, err := method.Inputs.Unpack(data[4:])
		if err != nil {
			return EventInput{}, err
		}
		input.Address = res[0].(common.Address)
		input.Amount = res[1].(*big.Int)
		input.Type = listener.TypeSend
	} else if method.Name == "approve" && input.Address.Hex() == MultiSendContractAddress {
		res, err := method.Inputs.Unpack(data[4:])
		if err != nil {
			return EventInput{}, err
		}
		input.Address = res[0].(common.Address)
		input.Amount = res[1].(*big.Int)
		input.Type = listener.TypeApprove
	} else {
		return EventInput{}, fmt.Errorf("unsupported event")
	}
	return input, nil
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
	res, err := method.Inputs.Unpack(data[4:])
	var bulkSendETHInput BulkSendETHInput
	bulkSendETHInput.Addresses = res[0].([]common.Address)
	bulkSendETHInput.Amounts = res[1].([]*big.Int)
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
	res, err := method.Inputs.Unpack(data[4:])
	var bulkSendTokenInput BulkSendTokenInput
	bulkSendTokenInput.Token = res[0].(common.Address)
	bulkSendTokenInput.Addresses = res[1].([]common.Address)
	bulkSendTokenInput.Amounts = res[2].([]*big.Int)
	if err != nil {
		return BulkSendTokenInput{}, err
	}
	return bulkSendTokenInput, nil
}
