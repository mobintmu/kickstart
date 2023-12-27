// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CampaignABI is the input ABI used to generate the binding from.
const CampaignABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"addRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approversCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"finalizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CampaignBin is the compiled bytecode used for deploying new contracts.
var CampaignBin = "0x60806040523480156200001157600080fd5b50604051620016b4380380620016b483398181016040528101906200003791906200009c565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060018190555050620000f7565b6000815190506200009681620000dd565b92915050565b600060208284031215620000b557620000b4620000d8565b5b6000620000c58482850162000085565b91505092915050565b6000819050919050565b600080fd5b620000e881620000ce565b8114620000f457600080fd5b50565b6115ad80620001076000396000f3fe6080604052600436106100a75760003560e01c806382fde0931161006457806382fde093146101d4578063937e09b1146101ff5780639fb42b1f1461022a578063ae86da0c14610255578063d7bb99ba1461027e578063d7d1bbdb14610288576100a7565b806303441006146100ac5780630a144391146100d557806312065fe014610112578063209ebc081461013d578063481c6a751461016857806381d12c5814610193575b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce9190610cd5565b6102b1565b005b3480156100e157600080fd5b506100fc60048036038101906100f79190610c39565b610514565b6040516101099190610eaa565b60405180910390f35b34801561011e57600080fd5b50610127610534565b604051610134919061101f565b60405180910390f35b34801561014957600080fd5b5061015261053c565b60405161015f9190610eaa565b60405180910390f35b34801561017457600080fd5b5061017d610590565b60405161018a9190610e8f565b60405180910390f35b34801561019f57600080fd5b506101ba60048036038101906101b59190610cd5565b6105b4565b6040516101cb959493929190610ec5565b60405180910390f35b3480156101e057600080fd5b506101e961069f565b6040516101f6919061101f565b60405180910390f35b34801561020b57600080fd5b506102146106a5565b604051610221919061101f565b60405180910390f35b34801561023657600080fd5b5061023f6106ab565b60405161024c919061101f565b60405180910390f35b34801561026157600080fd5b5061027c60048036038101906102779190610c66565b6106b1565b005b6102866107fe565b005b34801561029457600080fd5b506102af60048036038101906102aa9190610cd5565b6108b5565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033690610f3f565b60405180910390fd5b6004548110610383576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037a90610f1f565b60405180910390fd5b60008110156103c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103be90610f7f565b60405180910390fd5b60006005600083815260200190815260200160002090508060020160149054906101000a900460ff1615610430576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042790610fff565b60405180910390fd5b600260035461043f91906110ac565b816003015411610484576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047b90610fdf565b60405180910390fd5b8060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc82600101549081150290604051600060405180830381858888f193505050501580156104f2573d6000803e3d6000fd5b5060018160020160146101000a81548160ff0219169083151502179055505050565b60026020528060005260406000206000915054906101000a900460ff1681565b600047905090565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60056020528060005260406000206000915090508060000180546105d790611179565b80601f016020809104026020016040519081016040528092919081815260200182805461060390611179565b80156106505780601f1061062557610100808354040283529160200191610650565b820191906000526020600020905b81548152906001019060200180831161063357829003601f168201915b5050505050908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060030154905085565b60035481565b60015481565b60045481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073690610f3f565b60405180910390fd5b60006005600060046000815480929190610758906111dc565b919050558152602001908152602001600020905083816000019080519060200190610784929190610ae7565b50828160010181905550818160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008160020160146101000a81548160ff0219169083151502179055506000816003018190555050505050565b600154341015610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083a90610f5f565b60405180910390fd5b6001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600360008154809291906108ae906111dc565b9190505550565b6000600560008381526020019081526020016000209050600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094f90610fbf565b60405180910390fd5b8060040160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156109e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109de90610f9f565b60405180910390fd5b6004548210610a2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2290610f1f565b60405180910390fd5b6000821015610a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6690610f7f565b60405180910390fd5b60018160040160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550806003016000815480929190610ade906111dc565b91905055505050565b828054610af390611179565b90600052602060002090601f016020900481019282610b155760008555610b5c565b82601f10610b2e57805160ff1916838001178555610b5c565b82800160010185558215610b5c579182015b82811115610b5b578251825591602001919060010190610b40565b5b509050610b699190610b6d565b5090565b5b80821115610b86576000816000905550600101610b6e565b5090565b6000610b9d610b988461105f565b61103a565b905082815260208101848484011115610bb957610bb86112e6565b5b610bc4848285611137565b509392505050565b600081359050610bdb81611532565b92915050565b600081359050610bf081611549565b92915050565b600082601f830112610c0b57610c0a6112e1565b5b8135610c1b848260208601610b8a565b91505092915050565b600081359050610c3381611560565b92915050565b600060208284031215610c4f57610c4e6112f0565b5b6000610c5d84828501610bcc565b91505092915050565b600080600060608486031215610c7f57610c7e6112f0565b5b600084013567ffffffffffffffff811115610c9d57610c9c6112eb565b5b610ca986828701610bf6565b9350506020610cba86828701610c24565b9250506040610ccb86828701610be1565b9150509250925092565b600060208284031215610ceb57610cea6112f0565b5b6000610cf984828501610c24565b91505092915050565b610d0b816110ef565b82525050565b610d1a816110dd565b82525050565b610d2981611101565b82525050565b6000610d3a82611090565b610d44818561109b565b9350610d54818560208601611146565b610d5d816112f5565b840191505092915050565b6000610d7560318361109b565b9150610d8082611306565b604082019050919050565b6000610d9860318361109b565b9150610da382611355565b604082019050919050565b6000610dbb601f8361109b565b9150610dc6826113a4565b602082019050919050565b6000610dde602d8361109b565b9150610de9826113cd565b604082019050919050565b6000610e01602e8361109b565b9150610e0c8261141c565b604082019050919050565b6000610e2460388361109b565b9150610e2f8261146b565b604082019050919050565b6000610e47603c8361109b565b9150610e52826114ba565b604082019050919050565b6000610e6a60208361109b565b9150610e7582611509565b602082019050919050565b610e898161112d565b82525050565b6000602082019050610ea46000830184610d11565b92915050565b6000602082019050610ebf6000830184610d20565b92915050565b600060a0820190508181036000830152610edf8188610d2f565b9050610eee6020830187610e80565b610efb6040830186610d02565b610f086060830185610d20565b610f156080830184610e80565b9695505050505050565b60006020820190508181036000830152610f3881610d68565b9050919050565b60006020820190508181036000830152610f5881610d8b565b9050919050565b60006020820190508181036000830152610f7881610dae565b9050919050565b60006020820190508181036000830152610f9881610dd1565b9050919050565b60006020820190508181036000830152610fb881610df4565b9050919050565b60006020820190508181036000830152610fd881610e17565b9050919050565b60006020820190508181036000830152610ff881610e3a565b9050919050565b6000602082019050818103600083015261101881610e5d565b9050919050565b60006020820190506110346000830184610e80565b92915050565b6000611044611055565b905061105082826111ab565b919050565b6000604051905090565b600067ffffffffffffffff82111561107a576110796112b2565b5b611083826112f5565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006110b78261112d565b91506110c28361112d565b9250826110d2576110d1611254565b5b828204905092915050565b60006110e88261110d565b9050919050565b60006110fa8261110d565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611164578082015181840152602081019050611149565b83811115611173576000848401525b50505050565b6000600282049050600182168061119157607f821691505b602082108114156111a5576111a4611283565b5b50919050565b6111b4826112f5565b810181811067ffffffffffffffff821117156111d3576111d26112b2565b5b80604052505050565b60006111e78261112d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561121a57611219611225565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f596f752063616e20766f7465206f6e6c7920696e2076616c696420696e64657860008201527f2e206d6f7265207468616e2076616c6964000000000000000000000000000000602082015250565b7f4f6e6c79207468652063616d706169676e206d616e616765722063616e20636160008201527f6c6c20746869732066756e6374696f6e2e000000000000000000000000000000602082015250565b7f6d696e696d756d20636f6e747269627574652069732072657175697265642e00600082015250565b7f596f752063616e20766f7465206f6e6c7920696e2076616c696420696e64657860008201527f2e206d696e75732076616c756500000000000000000000000000000000000000602082015250565b7f596f75206861766520616c726561647920766f74656420746f20617070726f7660008201527f6520746869732072657175657374000000000000000000000000000000000000602082015250565b7f4f6e6c7920636f6e7472696275746f72732063616e20617070726f766520612060008201527f7370656369666963207061796d656e7420726571756573740000000000000000602082015250565b7f546869732072657175657374206e65656473206d6f726520617070726f76616c60008201527f73206265666f72652069742063616e2062652066696e616c697a656400000000602082015250565b7f54686973207265717565737420686173206265656e20636f6d706c657465642e600082015250565b61153b816110dd565b811461154657600080fd5b50565b611552816110ef565b811461155d57600080fd5b50565b6115698161112d565b811461157457600080fd5b5056fea264697066735822122043ae49bb9cd4f329a1cb8956a690960d17508feeba769b932618569c2f17d0de64736f6c63430008070033"

// DeployCampaign deploys a new Core contract, binding an instance of Campaign to it.
func DeployCampaign(auth *bind.TransactOpts, backend bind.ContractBackend, min *big.Int) (common.Address, *types.Transaction, *Campaign, error) {
	parsed, err := abi.JSON(strings.NewReader(CampaignABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CampaignBin), backend, min)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Campaign{CampaignCaller: CampaignCaller{contract: contract}, CampaignTransactor: CampaignTransactor{contract: contract}, CampaignFilterer: CampaignFilterer{contract: contract}}, nil
}

// Campaign is an auto generated Go binding around an Core contract.
type Campaign struct {
	CampaignCaller     // Read-only binding to the contract
	CampaignTransactor // Write-only binding to the contract
	CampaignFilterer   // Log filterer for contract events
}

// CampaignCaller is an auto generated read-only Go binding around an Core contract.
type CampaignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignTransactor is an auto generated write-only Go binding around an Core contract.
type CampaignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFilterer is an auto generated log filtering Go binding around an Core contract events.
type CampaignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type CampaignSession struct {
	Contract     *Campaign         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CampaignCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type CampaignCallerSession struct {
	Contract *CampaignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CampaignTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type CampaignTransactorSession struct {
	Contract     *CampaignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CampaignRaw is an auto generated low-level Go binding around an Core contract.
type CampaignRaw struct {
	Contract *Campaign // Generic contract binding to access the raw methods on
}

// CampaignCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type CampaignCallerRaw struct {
	Contract *CampaignCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type CampaignTransactorRaw struct {
	Contract *CampaignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaign creates a new instance of Campaign, bound to a specific deployed contract.
func NewCampaign(address common.Address, backend bind.ContractBackend) (*Campaign, error) {
	contract, err := bindCampaign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Campaign{CampaignCaller: CampaignCaller{contract: contract}, CampaignTransactor: CampaignTransactor{contract: contract}, CampaignFilterer: CampaignFilterer{contract: contract}}, nil
}

// NewCampaignCaller creates a new read-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignCaller(address common.Address, caller bind.ContractCaller) (*CampaignCaller, error) {
	contract, err := bindCampaign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignCaller{contract: contract}, nil
}

// NewCampaignTransactor creates a new write-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignTransactor, error) {
	contract, err := bindCampaign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignTransactor{contract: contract}, nil
}

// NewCampaignFilterer creates a new log filterer instance of Campaign, bound to a specific deployed contract.
func NewCampaignFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignFilterer, error) {
	contract, err := bindCampaign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignFilterer{contract: contract}, nil
}

// bindCampaign binds a generic wrapper to an already deployed contract.
func bindCampaign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CampaignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.CampaignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transact(opts, method, params...)
}

// Approvers is a free data retrieval call binding the contract method 0xc9e85ba3.
//
// Ylem: function approvers(address ) view returns(bool)
func (_Campaign *CampaignCaller) Approvers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "approvers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvers is a free data retrieval call binding the contract method 0xc9e85ba3.
//
// Ylem: function approvers(address ) view returns(bool)
func (_Campaign *CampaignSession) Approvers(arg0 common.Address) (bool, error) {
	return _Campaign.Contract.Approvers(&_Campaign.CallOpts, arg0)
}

// Approvers is a free data retrieval call binding the contract method 0xc9e85ba3.
//
// Ylem: function approvers(address ) view returns(bool)
func (_Campaign *CampaignCallerSession) Approvers(arg0 common.Address) (bool, error) {
	return _Campaign.Contract.Approvers(&_Campaign.CallOpts, arg0)
}

// ApproversCount is a free data retrieval call binding the contract method 0xc6422197.
//
// Ylem: function approversCount() view returns(uint256)
func (_Campaign *CampaignCaller) ApproversCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "approversCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproversCount is a free data retrieval call binding the contract method 0xc6422197.
//
// Ylem: function approversCount() view returns(uint256)
func (_Campaign *CampaignSession) ApproversCount() (*big.Int, error) {
	return _Campaign.Contract.ApproversCount(&_Campaign.CallOpts)
}

// ApproversCount is a free data retrieval call binding the contract method 0xc6422197.
//
// Ylem: function approversCount() view returns(uint256)
func (_Campaign *CampaignCallerSession) ApproversCount() (*big.Int, error) {
	return _Campaign.Contract.ApproversCount(&_Campaign.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xec4a3135.
//
// Ylem: function getBalance() view returns(uint256)
func (_Campaign *CampaignCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xec4a3135.
//
// Ylem: function getBalance() view returns(uint256)
func (_Campaign *CampaignSession) GetBalance() (*big.Int, error) {
	return _Campaign.Contract.GetBalance(&_Campaign.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xec4a3135.
//
// Ylem: function getBalance() view returns(uint256)
func (_Campaign *CampaignCallerSession) GetBalance() (*big.Int, error) {
	return _Campaign.Contract.GetBalance(&_Campaign.CallOpts)
}

// IsPaid is a free data retrieval call binding the contract method 0x6d48d1ba.
//
// Ylem: function isPaid() view returns(bool)
func (_Campaign *CampaignCaller) IsPaid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "isPaid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaid is a free data retrieval call binding the contract method 0x6d48d1ba.
//
// Ylem: function isPaid() view returns(bool)
func (_Campaign *CampaignSession) IsPaid() (bool, error) {
	return _Campaign.Contract.IsPaid(&_Campaign.CallOpts)
}

// IsPaid is a free data retrieval call binding the contract method 0x6d48d1ba.
//
// Ylem: function isPaid() view returns(bool)
func (_Campaign *CampaignCallerSession) IsPaid() (bool, error) {
	return _Campaign.Contract.IsPaid(&_Campaign.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0xcad39ebc.
//
// Ylem: function manager() view returns(address)
func (_Campaign *CampaignCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0xcad39ebc.
//
// Ylem: function manager() view returns(address)
func (_Campaign *CampaignSession) Manager() (common.Address, error) {
	return _Campaign.Contract.Manager(&_Campaign.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0xcad39ebc.
//
// Ylem: function manager() view returns(address)
func (_Campaign *CampaignCallerSession) Manager() (common.Address, error) {
	return _Campaign.Contract.Manager(&_Campaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x3fc05c70.
//
// Ylem: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignCaller) MinimumContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "minimumContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumContribution is a free data retrieval call binding the contract method 0x3fc05c70.
//
// Ylem: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignSession) MinimumContribution() (*big.Int, error) {
	return _Campaign.Contract.MinimumContribution(&_Campaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x3fc05c70.
//
// Ylem: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignCallerSession) MinimumContribution() (*big.Int, error) {
	return _Campaign.Contract.MinimumContribution(&_Campaign.CallOpts)
}

// NumRequests is a free data retrieval call binding the contract method 0x21aaecaa.
//
// Ylem: function numRequests() view returns(uint256)
func (_Campaign *CampaignCaller) NumRequests(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "numRequests")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRequests is a free data retrieval call binding the contract method 0x21aaecaa.
//
// Ylem: function numRequests() view returns(uint256)
func (_Campaign *CampaignSession) NumRequests() (*big.Int, error) {
	return _Campaign.Contract.NumRequests(&_Campaign.CallOpts)
}

// NumRequests is a free data retrieval call binding the contract method 0x21aaecaa.
//
// Ylem: function numRequests() view returns(uint256)
func (_Campaign *CampaignCallerSession) NumRequests() (*big.Int, error) {
	return _Campaign.Contract.NumRequests(&_Campaign.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0xbdb8bb84.
//
// Ylem: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Description   string
		Value         *big.Int
		Recipient     common.Address
		Complete      bool
		ApprovalCount *big.Int
	})

	outstruct.Description = out[0].(string)
	outstruct.Value = out[1].(*big.Int)
	outstruct.Recipient = out[2].(common.Address)
	outstruct.Complete = out[3].(bool)
	outstruct.ApprovalCount = out[4].(*big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0xbdb8bb84.
//
// Ylem: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _Campaign.Contract.Requests(&_Campaign.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0xbdb8bb84.
//
// Ylem: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignCallerSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _Campaign.Contract.Requests(&_Campaign.CallOpts, arg0)
}

// AddRequest is a paid mutator transaction binding the contract method 0x85679124.
//
// Ylem: function addRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignTransactor) AddRequest(opts *bind.TransactOpts, description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "addRequest", description, value, recipient)
}

// AddRequest is a paid mutator transaction binding the contract method 0x85679124.
//
// Ylem: function addRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignSession) AddRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.AddRequest(&_Campaign.TransactOpts, description, value, recipient)
}

// AddRequest is a paid mutator transaction binding the contract method 0x85679124.
//
// Ylem: function addRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignTransactorSession) AddRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.AddRequest(&_Campaign.TransactOpts, description, value, recipient)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x75c6dc18.
//
// Ylem: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignTransactor) ApproveRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "approveRequest", index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x75c6dc18.
//
// Ylem: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveRequest(&_Campaign.TransactOpts, index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x75c6dc18.
//
// Ylem: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignTransactorSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveRequest(&_Campaign.TransactOpts, index)
}

// Contribute is a paid mutator transaction binding the contract method 0xe94c5af6.
//
// Ylem: function contribute() payable returns()
func (_Campaign *CampaignTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xe94c5af6.
//
// Ylem: function contribute() payable returns()
func (_Campaign *CampaignSession) Contribute() (*types.Transaction, error) {
	return _Campaign.Contract.Contribute(&_Campaign.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xe94c5af6.
//
// Ylem: function contribute() payable returns()
func (_Campaign *CampaignTransactorSession) Contribute() (*types.Transaction, error) {
	return _Campaign.Contract.Contribute(&_Campaign.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0xcbbb8949.
//
// Ylem: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignTransactor) FinalizeRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "finalizeRequest", index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0xcbbb8949.
//
// Ylem: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.FinalizeRequest(&_Campaign.TransactOpts, index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0xcbbb8949.
//
// Ylem: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignTransactorSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.FinalizeRequest(&_Campaign.TransactOpts, index)
}
