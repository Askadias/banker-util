// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TransactionABI is the input ABI used to generate the binding from.
const TransactionABI = "[{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"affiliates\",\"type\":\"address[]\"}],\"name\":\"bulkSendEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractToken\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"affiliates\",\"type\":\"address[]\"}],\"name\":\"bulkSendEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feePerc\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"internalFeePerc\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerc\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"_affiliatesPerc\",\"type\":\"uint8[]\"}],\"name\":\"setAffiliatesPerc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_feePerc\",\"type\":\"uint8\"}],\"name\":\"setFeePerc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_internalFeePerc\",\"type\":\"uint8\"}],\"name\":\"setInternalFeePerc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"_partners\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_partnersPerc\",\"type\":\"uint8[]\"}],\"name\":\"setPartners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rewardPerc\",\"type\":\"uint8\"}],\"name\":\"setRewardPerc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractToken\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransactionBin is the compiled bytecode used for deploying new contracts.
var TransactionBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612163806100536000396000f3fe6080604052600436106100f25760003560e01c80638534ec231161008a578063d0e30db011610059578063d0e30db0146106cf578063dc82b667146106f1578063e115e6ae14610862578063ef41ea7e146108b8576100f2565b80638534ec23146105c05780638da5cb5b146105f1578063b4223c7314610648578063b852010214610679576100f2565b806340b4d178116100c657806340b4d178146103b5578063522f68151461040b578063681161771461047e57806372a3bb54146104e3576100f2565b8062f55d9d146100f757806301e336671461014857806324501cd9146101db5780633e917cb2146102b3575b600080fd5b34801561010357600080fd5b506101466004803603602081101561011a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108e9565b005b34801561015457600080fd5b506101c16004803603606081101561016b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061095b565b604051808215151515815260200191505060405180910390f35b6102b1600480360360408110156101f157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561022e57600080fd5b82018360208201111561024057600080fd5b8035906020019184602083028401116401000000008311171561026257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610a84565b005b6103b3600480360360808110156102c957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561033057600080fd5b82018360208201111561034257600080fd5b8035906020019184602083028401116401000000008311171561036457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050611012565b005b3480156103c157600080fd5b506103f1600480360360208110156103d857600080fd5b81019080803560ff169060200190929190505050611805565b604051808215151515815260200191505060405180910390f35b34801561041757600080fd5b506104646004803603604081101561042e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118e0565b604051808215151515815260200191505060405180910390f35b34801561048a57600080fd5b506104cd600480360360208110156104a157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061198c565b6040518082815260200191505060405180910390f35b3480156104ef57600080fd5b506105a66004803603602081101561050657600080fd5b810190808035906020019064010000000081111561052357600080fd5b82018360208201111561053557600080fd5b8035906020019184602083028401116401000000008311171561055757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506119ad565b604051808215151515815260200191505060405180910390f35b3480156105cc57600080fd5b506105d5611adf565b604051808260ff1660ff16815260200191505060405180910390f35b3480156105fd57600080fd5b50610606611af2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561065457600080fd5b5061065d611b17565b604051808260ff1660ff16815260200191505060405180910390f35b34801561068557600080fd5b506106b56004803603602081101561069c57600080fd5b81019080803560ff169060200190929190505050611b2a565b604051808215151515815260200191505060405180910390f35b6106d7611c05565b604051808215151515815260200191505060405180910390f35b3480156106fd57600080fd5b506108486004803603604081101561071457600080fd5b810190808035906020019064010000000081111561073157600080fd5b82018360208201111561074357600080fd5b8035906020019184602083028401116401000000008311171561076557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156107c557600080fd5b8201836020820111156107d757600080fd5b803590602001918460208302840111640100000000831117156107f957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050611c0e565b604051808215151515815260200191505060405180910390f35b34801561086e57600080fd5b5061089e6004803603602081101561088557600080fd5b81019080803560ff169060200190929190505050611dcf565b604051808215151515815260200191505060405180910390f35b3480156108c457600080fd5b506108cd611eaa565b604051808260ff1660ff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461094257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109b657600080fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610a3d57600080fd5b505af1158015610a51573d6000803e3d6000fd5b505050506040513d6020811015610a6757600080fd5b810190808051906020019092919050505050600190509392505050565b60038054905081511115610b00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f746f6f206d616e7920616666696c69617465730000000000000000000000000081525060200191505060405180910390fd5b60003490506000610b3f6064610b31600060149054906101000a900460ff1660ff1685611ebd90919063ffffffff16565b611ef790919063ffffffff16565b90503a8110610b99576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806120e26027913960400191505060405180910390fd5b6000610bd36064610bc5600060159054906101000a900460ff1660ff1685611ebd90919063ffffffff16565b611ef790919063ffffffff16565b90506000610c0f6064610c01600060169054906101000a900460ff1660ff1686611ebd90919063ffffffff16565b611ef790919063ffffffff16565b90506000610c268486611f1d90919063ffffffff16565b90508673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610c6e573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de28782604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16000610cef8487611f1d90919063ffffffff16565b905060008090505b87518160ff161015610e60576000610d5d6064610d4f60038560ff1681548110610d1d57fe5b90600052602060002090602091828204019190069054906101000a900460ff1660ff1688611ebd90919063ffffffff16565b611ef790919063ffffffff16565b9050610d728184611f1d90919063ffffffff16565b9250888260ff1681518110610d8357fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610dd0573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2898360ff1681518110610e0157fe5b602002602001015182604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1508080600101915050610cf7565b5060008090505b6002805490508160ff161015611007576000610ed16064610ec360028560ff1681548110610e9157fe5b90600052602060002090602091828204019190069054906101000a900460ff1660ff1686611ebd90919063ffffffff16565b611ef790919063ffffffff16565b905060018260ff1681548110610ee357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610f53573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de260018360ff1681548110610f8557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1508080600101915050610e67565b505050505050505050565b6003805490508151111561108e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f746f6f206d616e7920616666696c69617465730000000000000000000000000081525060200191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b15801561113f57600080fd5b505afa158015611153573d6000803e3d6000fd5b505050506040513d602081101561116957600080fd5b81019080805190602001909291905050508311156111ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7420656e6f7567682062616c616e6365000000000000000000000000000081525060200191505060405180910390fd5b6000611229606461121b600060149054906101000a900460ff1660ff1687611ebd90919063ffffffff16565b611ef790919063ffffffff16565b905060006112656064611257600060159054906101000a900460ff1660ff1685611ebd90919063ffffffff16565b611ef790919063ffffffff16565b905060006112a16064611293600060169054906101000a900460ff1660ff1686611ebd90919063ffffffff16565b611ef790919063ffffffff16565b905060006112b88488611f1d90919063ffffffff16565b90508573ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611300573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de28682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a160006113818489611f1d90919063ffffffff16565b905060008090505b86518160ff1610156115a25760006113ef60646113e160038560ff16815481106113af57fe5b90600052602060002090602091828204019190069054906101000a900460ff1660ff1688611ebd90919063ffffffff16565b611ef790919063ffffffff16565b90506114048184611f1d90919063ffffffff16565b92508a73ffffffffffffffffffffffffffffffffffffffff166323b872dd338a8560ff168151811061143257fe5b6020026020010151846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156114d757600080fd5b505af11580156114eb573d6000803e3d6000fd5b505050506040513d602081101561150157600080fd5b8101908080519060200190929190505050507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2888360ff168151811061154357fe5b602002602001015182604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1508080600101915050611389565b5060008090505b6002805490508160ff1610156117f9576000611613606461160560028560ff16815481106115d357fe5b90600052602060002090602091828204019190069054906101000a900460ff1660ff1686611ebd90919063ffffffff16565b611ef790919063ffffffff16565b90508a73ffffffffffffffffffffffffffffffffffffffff166323b872dd3360018560ff168154811061164257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561170a57600080fd5b505af115801561171e573d6000803e3d6000fd5b505050506040513d602081101561173457600080fd5b8101908080519060200190929190505050507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de260018360ff168154811061177757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15080806001019150506115a9565b50505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461186057600080fd5b60648260ff16106118bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806121096026913960400191505060405180910390fd5b81600060156101000a81548160ff021916908360ff16021790555060019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461193b57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015611981573d6000803e3d6000fd5b506001905092915050565b60008173ffffffffffffffffffffffffffffffffffffffff16319050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a0857600080fd5b600080600090505b83518160ff161015611a4457838160ff1681518110611a2b57fe5b6020026020010151820191508080600101915050611a10565b5060648160ff1614611abe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f746f74616c2070657263656e746167652073686f756c6420626520313030000081525060200191505060405180910390fd5b8260039080519060200190611ad4929190611f3d565b506001915050919050565b600060169054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060159054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b8557600080fd5b60648260ff1610611be1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806121096026913960400191505060405180910390fd5b81600060166101000a81548160ff021916908360ff16021790555060019050919050565b60006001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c6957600080fd5b8151835114611ce0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f706172616d6574657273206e6f74206d6174636800000000000000000000000081525060200191505060405180910390fd5b600080600090505b83518160ff161015611d1c57838160ff1681518110611d0357fe5b6020026020010151820191508080600101915050611ce8565b5060648160ff1614611d96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f746f74616c2070657263656e746167652073686f756c6420626520313030000081525060200191505060405180910390fd5b8360019080519060200190611dac929190611fe4565b508260029080519060200190611dc3929190611f3d565b50600191505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e2a57600080fd5b60648260ff1610611e86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806121096026913960400191505060405180910390fd5b81600060146101000a81548160ff021916908360ff16021790555060019050919050565b600060149054906101000a900460ff1681565b600080831415611ed05760009050611ef1565b6000828402905082848281611ee157fe5b0414611eec57600080fd5b809150505b92915050565b6000808211611f0557600080fd5b6000828481611f1057fe5b0490508091505092915050565b600082821115611f2c57600080fd5b600082840390508091505092915050565b82805482825590600052602060002090601f01602090048101928215611fd35791602002820160005b83821115611fa457835183826101000a81548160ff021916908360ff1602179055509260200192600101602081600001049283019260010302611f66565b8015611fd15782816101000a81549060ff0219169055600101602081600001049283019260010302611fa4565b505b509050611fe0919061206e565b5090565b82805482825590600052602060002090810192821561205d579160200282015b8281111561205c5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612004565b5b50905061206a919061209e565b5090565b61209b91905b8082111561209757600081816101000a81549060ff021916905550600101612074565b5090565b90565b6120de91905b808211156120da57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016120a4565b5090565b9056fe6e6f7420656e6f7567682067617320746f20636f76657220636f6d6d697373696f6e20636f73746665652070657263656e746167652073686f756c64206265206c657373207468616e20313030a265627a7a72315820edc495c0388731a2a94c8ddef5c364cef1853aee08e817bb0ade55a9e1cde4fe64736f6c63430005110032"

// DeployTransaction deploys a new Ethereum contract, binding an instance of Transaction to it.
func DeployTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// Transaction is an auto generated Go binding around an Ethereum contract.
type Transaction struct {
	TransactionCaller     // Read-only binding to the contract
	TransactionTransactor // Write-only binding to the contract
	TransactionFilterer   // Log filterer for contract events
}

// TransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionSession struct {
	Contract     *Transaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionCallerSession struct {
	Contract *TransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionTransactorSession struct {
	Contract     *TransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionRaw struct {
	Contract *Transaction // Generic contract binding to access the raw methods on
}

// TransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionCallerRaw struct {
	Contract *TransactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionTransactorRaw struct {
	Contract *TransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransaction creates a new instance of Transaction, bound to a specific deployed contract.
func NewTransaction(address common.Address, backend bind.ContractBackend) (*Transaction, error) {
	contract, err := bindTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// NewTransactionCaller creates a new read-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionCaller(address common.Address, caller bind.ContractCaller) (*TransactionCaller, error) {
	contract, err := bindTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionCaller{contract: contract}, nil
}

// NewTransactionTransactor creates a new write-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionTransactor, error) {
	contract, err := bindTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionTransactor{contract: contract}, nil
}

// NewTransactionFilterer creates a new log filterer instance of Transaction, bound to a specific deployed contract.
func NewTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionFilterer, error) {
	contract, err := bindTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionFilterer{contract: contract}, nil
}

// bindTransaction binds a generic wrapper to an already deployed contract.
func bindTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.TransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transact(opts, method, params...)
}

// FeePerc is a free data retrieval call binding the contract method 0xef41ea7e.
//
// Solidity: function feePerc() constant returns(uint8)
func (_Transaction *TransactionCaller) FeePerc(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Transaction.contract.Call(opts, out, "feePerc")
	return *ret0, err
}

// FeePerc is a free data retrieval call binding the contract method 0xef41ea7e.
//
// Solidity: function feePerc() constant returns(uint8)
func (_Transaction *TransactionSession) FeePerc() (uint8, error) {
	return _Transaction.Contract.FeePerc(&_Transaction.CallOpts)
}

// FeePerc is a free data retrieval call binding the contract method 0xef41ea7e.
//
// Solidity: function feePerc() constant returns(uint8)
func (_Transaction *TransactionCallerSession) FeePerc() (uint8, error) {
	return _Transaction.Contract.FeePerc(&_Transaction.CallOpts)
}

// Getbalance is a free data retrieval call binding the contract method 0x68116177.
//
// Solidity: function getbalance(address addr) constant returns(uint256 value)
func (_Transaction *TransactionCaller) Getbalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transaction.contract.Call(opts, out, "getbalance", addr)
	return *ret0, err
}

// Getbalance is a free data retrieval call binding the contract method 0x68116177.
//
// Solidity: function getbalance(address addr) constant returns(uint256 value)
func (_Transaction *TransactionSession) Getbalance(addr common.Address) (*big.Int, error) {
	return _Transaction.Contract.Getbalance(&_Transaction.CallOpts, addr)
}

// Getbalance is a free data retrieval call binding the contract method 0x68116177.
//
// Solidity: function getbalance(address addr) constant returns(uint256 value)
func (_Transaction *TransactionCallerSession) Getbalance(addr common.Address) (*big.Int, error) {
	return _Transaction.Contract.Getbalance(&_Transaction.CallOpts, addr)
}

// InternalFeePerc is a free data retrieval call binding the contract method 0xb4223c73.
//
// Solidity: function internalFeePerc() constant returns(uint8)
func (_Transaction *TransactionCaller) InternalFeePerc(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Transaction.contract.Call(opts, out, "internalFeePerc")
	return *ret0, err
}

// InternalFeePerc is a free data retrieval call binding the contract method 0xb4223c73.
//
// Solidity: function internalFeePerc() constant returns(uint8)
func (_Transaction *TransactionSession) InternalFeePerc() (uint8, error) {
	return _Transaction.Contract.InternalFeePerc(&_Transaction.CallOpts)
}

// InternalFeePerc is a free data retrieval call binding the contract method 0xb4223c73.
//
// Solidity: function internalFeePerc() constant returns(uint8)
func (_Transaction *TransactionCallerSession) InternalFeePerc() (uint8, error) {
	return _Transaction.Contract.InternalFeePerc(&_Transaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transaction *TransactionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Transaction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transaction *TransactionSession) Owner() (common.Address, error) {
	return _Transaction.Contract.Owner(&_Transaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transaction *TransactionCallerSession) Owner() (common.Address, error) {
	return _Transaction.Contract.Owner(&_Transaction.CallOpts)
}

// RewardPerc is a free data retrieval call binding the contract method 0x8534ec23.
//
// Solidity: function rewardPerc() constant returns(uint8)
func (_Transaction *TransactionCaller) RewardPerc(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Transaction.contract.Call(opts, out, "rewardPerc")
	return *ret0, err
}

// RewardPerc is a free data retrieval call binding the contract method 0x8534ec23.
//
// Solidity: function rewardPerc() constant returns(uint8)
func (_Transaction *TransactionSession) RewardPerc() (uint8, error) {
	return _Transaction.Contract.RewardPerc(&_Transaction.CallOpts)
}

// RewardPerc is a free data retrieval call binding the contract method 0x8534ec23.
//
// Solidity: function rewardPerc() constant returns(uint8)
func (_Transaction *TransactionCallerSession) RewardPerc() (uint8, error) {
	return _Transaction.Contract.RewardPerc(&_Transaction.CallOpts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x24501cd9.
//
// Solidity: function bulkSendEth(address target, address[] affiliates) returns()
func (_Transaction *TransactionTransactor) BulkSendEth(opts *bind.TransactOpts, target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "bulkSendEth", target, affiliates)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x24501cd9.
//
// Solidity: function bulkSendEth(address target, address[] affiliates) returns()
func (_Transaction *TransactionSession) BulkSendEth(target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.BulkSendEth(&_Transaction.TransactOpts, target, affiliates)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x24501cd9.
//
// Solidity: function bulkSendEth(address target, address[] affiliates) returns()
func (_Transaction *TransactionTransactorSession) BulkSendEth(target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.BulkSendEth(&_Transaction.TransactOpts, target, affiliates)
}

// BulkSendEth0 is a paid mutator transaction binding the contract method 0x3e917cb2.
//
// Solidity: function bulkSendEth(address tokenAddr, uint256 amount, address target, address[] affiliates) returns()
func (_Transaction *TransactionTransactor) BulkSendEth0(opts *bind.TransactOpts, tokenAddr common.Address, amount *big.Int, target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "bulkSendEth0", tokenAddr, amount, target, affiliates)
}

// BulkSendEth0 is a paid mutator transaction binding the contract method 0x3e917cb2.
//
// Solidity: function bulkSendEth(address tokenAddr, uint256 amount, address target, address[] affiliates) returns()
func (_Transaction *TransactionSession) BulkSendEth0(tokenAddr common.Address, amount *big.Int, target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.BulkSendEth0(&_Transaction.TransactOpts, tokenAddr, amount, target, affiliates)
}

// BulkSendEth0 is a paid mutator transaction binding the contract method 0x3e917cb2.
//
// Solidity: function bulkSendEth(address tokenAddr, uint256 amount, address target, address[] affiliates) returns()
func (_Transaction *TransactionTransactorSession) BulkSendEth0(tokenAddr common.Address, amount *big.Int, target common.Address, affiliates []common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.BulkSendEth0(&_Transaction.TransactOpts, tokenAddr, amount, target, affiliates)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_Transaction *TransactionTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_Transaction *TransactionSession) Deposit() (*types.Transaction, error) {
	return _Transaction.Contract.Deposit(&_Transaction.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_Transaction *TransactionTransactorSession) Deposit() (*types.Transaction, error) {
	return _Transaction.Contract.Deposit(&_Transaction.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Transaction *TransactionTransactor) Destroy(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "destroy", _to)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Transaction *TransactionSession) Destroy(_to common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.Destroy(&_Transaction.TransactOpts, _to)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Transaction *TransactionTransactorSession) Destroy(_to common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.Destroy(&_Transaction.TransactOpts, _to)
}

// SetAffiliatesPerc is a paid mutator transaction binding the contract method 0x72a3bb54.
//
// Solidity: function setAffiliatesPerc(uint8[] _affiliatesPerc) returns(bool success)
func (_Transaction *TransactionTransactor) SetAffiliatesPerc(opts *bind.TransactOpts, _affiliatesPerc []uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "setAffiliatesPerc", _affiliatesPerc)
}

// SetAffiliatesPerc is a paid mutator transaction binding the contract method 0x72a3bb54.
//
// Solidity: function setAffiliatesPerc(uint8[] _affiliatesPerc) returns(bool success)
func (_Transaction *TransactionSession) SetAffiliatesPerc(_affiliatesPerc []uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetAffiliatesPerc(&_Transaction.TransactOpts, _affiliatesPerc)
}

// SetAffiliatesPerc is a paid mutator transaction binding the contract method 0x72a3bb54.
//
// Solidity: function setAffiliatesPerc(uint8[] _affiliatesPerc) returns(bool success)
func (_Transaction *TransactionTransactorSession) SetAffiliatesPerc(_affiliatesPerc []uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetAffiliatesPerc(&_Transaction.TransactOpts, _affiliatesPerc)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xe115e6ae.
//
// Solidity: function setFeePerc(uint8 _feePerc) returns(bool success)
func (_Transaction *TransactionTransactor) SetFeePerc(opts *bind.TransactOpts, _feePerc uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "setFeePerc", _feePerc)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xe115e6ae.
//
// Solidity: function setFeePerc(uint8 _feePerc) returns(bool success)
func (_Transaction *TransactionSession) SetFeePerc(_feePerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetFeePerc(&_Transaction.TransactOpts, _feePerc)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xe115e6ae.
//
// Solidity: function setFeePerc(uint8 _feePerc) returns(bool success)
func (_Transaction *TransactionTransactorSession) SetFeePerc(_feePerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetFeePerc(&_Transaction.TransactOpts, _feePerc)
}

// SetInternalFeePerc is a paid mutator transaction binding the contract method 0x40b4d178.
//
// Solidity: function setInternalFeePerc(uint8 _internalFeePerc) returns(bool success)
func (_Transaction *TransactionTransactor) SetInternalFeePerc(opts *bind.TransactOpts, _internalFeePerc uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "setInternalFeePerc", _internalFeePerc)
}

// SetInternalFeePerc is a paid mutator transaction binding the contract method 0x40b4d178.
//
// Solidity: function setInternalFeePerc(uint8 _internalFeePerc) returns(bool success)
func (_Transaction *TransactionSession) SetInternalFeePerc(_internalFeePerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetInternalFeePerc(&_Transaction.TransactOpts, _internalFeePerc)
}

// SetInternalFeePerc is a paid mutator transaction binding the contract method 0x40b4d178.
//
// Solidity: function setInternalFeePerc(uint8 _internalFeePerc) returns(bool success)
func (_Transaction *TransactionTransactorSession) SetInternalFeePerc(_internalFeePerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetInternalFeePerc(&_Transaction.TransactOpts, _internalFeePerc)
}

// SetPartners is a paid mutator transaction binding the contract method 0xdc82b667.
//
// Solidity: function setPartners(address[] _partners, uint8[] _partnersPerc) returns(bool success)
func (_Transaction *TransactionTransactor) SetPartners(opts *bind.TransactOpts, _partners []common.Address, _partnersPerc []uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "setPartners", _partners, _partnersPerc)
}

// SetPartners is a paid mutator transaction binding the contract method 0xdc82b667.
//
// Solidity: function setPartners(address[] _partners, uint8[] _partnersPerc) returns(bool success)
func (_Transaction *TransactionSession) SetPartners(_partners []common.Address, _partnersPerc []uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetPartners(&_Transaction.TransactOpts, _partners, _partnersPerc)
}

// SetPartners is a paid mutator transaction binding the contract method 0xdc82b667.
//
// Solidity: function setPartners(address[] _partners, uint8[] _partnersPerc) returns(bool success)
func (_Transaction *TransactionTransactorSession) SetPartners(_partners []common.Address, _partnersPerc []uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetPartners(&_Transaction.TransactOpts, _partners, _partnersPerc)
}

// SetRewardPerc is a paid mutator transaction binding the contract method 0xb8520102.
//
// Solidity: function setRewardPerc(uint8 _rewardPerc) returns(bool success)
func (_Transaction *TransactionTransactor) SetRewardPerc(opts *bind.TransactOpts, _rewardPerc uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "setRewardPerc", _rewardPerc)
}

// SetRewardPerc is a paid mutator transaction binding the contract method 0xb8520102.
//
// Solidity: function setRewardPerc(uint8 _rewardPerc) returns(bool success)
func (_Transaction *TransactionSession) SetRewardPerc(_rewardPerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetRewardPerc(&_Transaction.TransactOpts, _rewardPerc)
}

// SetRewardPerc is a paid mutator transaction binding the contract method 0xb8520102.
//
// Solidity: function setRewardPerc(uint8 _rewardPerc) returns(bool success)
func (_Transaction *TransactionTransactorSession) SetRewardPerc(_rewardPerc uint8) (*types.Transaction, error) {
	return _Transaction.Contract.SetRewardPerc(&_Transaction.TransactOpts, _rewardPerc)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address addr, uint256 amount) returns(bool success)
func (_Transaction *TransactionTransactor) WithdrawEther(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "withdrawEther", addr, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address addr, uint256 amount) returns(bool success)
func (_Transaction *TransactionSession) WithdrawEther(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transaction.Contract.WithdrawEther(&_Transaction.TransactOpts, addr, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address addr, uint256 amount) returns(bool success)
func (_Transaction *TransactionTransactorSession) WithdrawEther(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transaction.Contract.WithdrawEther(&_Transaction.TransactOpts, addr, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address tokenAddr, address _to, uint256 _amount) returns(bool success)
func (_Transaction *TransactionTransactor) WithdrawToken(opts *bind.TransactOpts, tokenAddr common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "withdrawToken", tokenAddr, _to, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address tokenAddr, address _to, uint256 _amount) returns(bool success)
func (_Transaction *TransactionSession) WithdrawToken(tokenAddr common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Transaction.Contract.WithdrawToken(&_Transaction.TransactOpts, tokenAddr, _to, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address tokenAddr, address _to, uint256 _amount) returns(bool success)
func (_Transaction *TransactionTransactorSession) WithdrawToken(tokenAddr common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Transaction.Contract.WithdrawToken(&_Transaction.TransactOpts, tokenAddr, _to, _amount)
}

// TransactionRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Transaction contract.
type TransactionRefundIterator struct {
	Event *TransactionRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TransactionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TransactionRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TransactionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionRefund represents a Refund event raised by the Transaction contract.
type TransactionRefund struct {
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Transaction *TransactionFilterer) FilterRefund(opts *bind.FilterOpts) (*TransactionRefundIterator, error) {

	logs, sub, err := _Transaction.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &TransactionRefundIterator{contract: _Transaction.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Transaction *TransactionFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *TransactionRefund) (event.Subscription, error) {

	logs, sub, err := _Transaction.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionRefund)
				if err := _Transaction.contract.UnpackLog(event, "Refund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefund is a log parse operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Transaction *TransactionFilterer) ParseRefund(log types.Log) (*TransactionRefund, error) {
	event := new(TransactionRefund)
	if err := _Transaction.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransactionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Transaction contract.
type TransactionTransferIterator struct {
	Event *TransactionTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TransactionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TransactionTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TransactionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionTransfer represents a Transfer event raised by the Transaction contract.
type TransactionTransfer struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Transaction *TransactionFilterer) FilterTransfer(opts *bind.FilterOpts) (*TransactionTransferIterator, error) {

	logs, sub, err := _Transaction.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &TransactionTransferIterator{contract: _Transaction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Transaction *TransactionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TransactionTransfer) (event.Subscription, error) {

	logs, sub, err := _Transaction.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionTransfer)
				if err := _Transaction.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Transaction *TransactionFilterer) ParseTransfer(log types.Log) (*TransactionTransfer, error) {
	event := new(TransactionTransfer)
	if err := _Transaction.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
