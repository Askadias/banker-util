package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func PackTransferData(_to common.Address, _value *big.Int) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return parsed.Pack("transfer", _to, _value)
}

func PackBulkSendEthData(addresses []common.Address, amounts []*big.Int) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return parsed.Pack("bulkSendEth", addresses, amounts)
}

func PackBulkSendTokenData(tokenAddr common.Address, addresses []common.Address, amounts []*big.Int) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return parsed.Pack("bulkSendToken", tokenAddr, addresses, amounts)
}
