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
const CampaignABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"addRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approversCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contibute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"finalizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaied\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CampaignBin is the compiled bytecode used for deploying new contracts.
var CampaignBin = "0x608060405234801562000010575f80fd5b50604051620017f2380380620017f28339818101604052810190620000369190620000bf565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060018190555050620000ef565b5f80fd5b5f819050919050565b6200009b8162000087565b8114620000a6575f80fd5b50565b5f81519050620000b98162000090565b92915050565b5f60208284031215620000d757620000d662000083565b5b5f620000e684828501620000a9565b91505092915050565b6116f580620000fd5f395ff3fe6080604052600436106100a6575f3560e01c8063937e09b111610063578063937e09b1146101cc5780639fb42b1f146101f6578063ae86da0c14610220578063d7d1bbdb14610248578063ea25213f14610270578063ea7b625e1461027a576100a6565b806303441006146100aa5780630a144391146100d257806312065fe01461010e578063481c6a751461013857806381d12c581461016257806382fde093146101a2575b5f80fd5b3480156100b5575f80fd5b506100d060048036038101906100cb9190610ae3565b6102a4565b005b3480156100dd575f80fd5b506100f860048036038101906100f39190610b68565b6104fe565b6040516101059190610bad565b60405180910390f35b348015610119575f80fd5b5061012261051b565b60405161012f9190610bd5565b60405180910390f35b348015610143575f80fd5b5061014c610522565b6040516101599190610bfd565b60405180910390f35b34801561016d575f80fd5b5061018860048036038101906101839190610ae3565b610545565b604051610199959493929190610cc0565b60405180910390f35b3480156101ad575f80fd5b506101b6610629565b6040516101c39190610bd5565b60405180910390f35b3480156101d7575f80fd5b506101e061062f565b6040516101ed9190610bd5565b60405180910390f35b348015610201575f80fd5b5061020a610635565b6040516102179190610bd5565b60405180910390f35b34801561022b575f80fd5b5061024660048036038101906102419190610e6e565b61063b565b005b348015610253575f80fd5b5061026e60048036038101906102699190610ae3565b610778565b005b61027861099c565b005b348015610285575f80fd5b5061028e610a4f565b60405161029b9190610bad565b60405180910390f35b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610331576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032890610f4a565b60405180910390fd5b6004548110610375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036c90610fd8565b60405180910390fd5b5f8110156103b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103af90611066565b60405180910390fd5b5f60055f8381526020019081526020015f2090508060020160149054906101000a900460ff161561041e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610415906110ce565b60405180910390fd5b600260035461042d9190611146565b816003015411610472576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610469906111e6565b60405180910390fd5b806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc826001015490811502906040515f60405180830381858888f193505050501580156104dc573d5f803e3d5ffd5b5060018160020160146101000a81548160ff0219169083151502179055505050565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f47905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005602052805f5260405f205f91509050805f01805461056490611231565b80601f016020809104026020016040519081016040528092919081815260200182805461059090611231565b80156105db5780601f106105b2576101008083540402835291602001916105db565b820191905f5260205f20905b8154815290600101906020018083116105be57829003601f168201915b505050505090806001015490806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060030154905085565b60035481565b60015481565b60045481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bf90610f4a565b60405180910390fd5b5f60055f60045f8154809291906106de90611261565b9190505581526020019081526020015f20905083815f0190816107019190611445565b5082816001018190555081816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f8160020160146101000a81548160ff0219169083151502179055505f816003018190555050505050565b5f60055f8381526020019081526020015f20905060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080c90611584565b60405180910390fd5b806004015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156108a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089890611612565b60405180910390fd5b60045482106108e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dc90610fd8565b60405180910390fd5b5f821015610928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091f90611066565b60405180910390fd5b6001816004015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550806003015f81548092919061099390611261565b91905055505050565b6001543410156109e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d89061167a565b60405180910390fd5b600160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060035f815480929190610a4890611261565b9190505550565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610ac281610ab0565b8114610acc575f80fd5b50565b5f81359050610add81610ab9565b92915050565b5f60208284031215610af857610af7610aa8565b5b5f610b0584828501610acf565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b3782610b0e565b9050919050565b610b4781610b2d565b8114610b51575f80fd5b50565b5f81359050610b6281610b3e565b92915050565b5f60208284031215610b7d57610b7c610aa8565b5b5f610b8a84828501610b54565b91505092915050565b5f8115159050919050565b610ba781610b93565b82525050565b5f602082019050610bc05f830184610b9e565b92915050565b610bcf81610ab0565b82525050565b5f602082019050610be85f830184610bc6565b92915050565b610bf781610b2d565b82525050565b5f602082019050610c105f830184610bee565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610c4d578082015181840152602081019050610c32565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610c7282610c16565b610c7c8185610c20565b9350610c8c818560208601610c30565b610c9581610c58565b840191505092915050565b5f610caa82610b0e565b9050919050565b610cba81610ca0565b82525050565b5f60a0820190508181035f830152610cd88188610c68565b9050610ce76020830187610bc6565b610cf46040830186610cb1565b610d016060830185610b9e565b610d0e6080830184610bc6565b9695505050505050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d5682610c58565b810181811067ffffffffffffffff82111715610d7557610d74610d20565b5b80604052505050565b5f610d87610a9f565b9050610d938282610d4d565b919050565b5f67ffffffffffffffff821115610db257610db1610d20565b5b610dbb82610c58565b9050602081019050919050565b828183375f83830152505050565b5f610de8610de384610d98565b610d7e565b905082815260208101848484011115610e0457610e03610d1c565b5b610e0f848285610dc8565b509392505050565b5f82601f830112610e2b57610e2a610d18565b5b8135610e3b848260208601610dd6565b91505092915050565b610e4d81610ca0565b8114610e57575f80fd5b50565b5f81359050610e6881610e44565b92915050565b5f805f60608486031215610e8557610e84610aa8565b5b5f84013567ffffffffffffffff811115610ea257610ea1610aac565b5b610eae86828701610e17565b9350506020610ebf86828701610acf565b9250506040610ed086828701610e5a565b9150509250925092565b7f4f6e6c79207468652063616d706169676e206d616e616765722063616e2063615f8201527f6c6c20746869732066756e6374696f6e2e000000000000000000000000000000602082015250565b5f610f34603183610c20565b9150610f3f82610eda565b604082019050919050565b5f6020820190508181035f830152610f6181610f28565b9050919050565b7f596f752063616e20766f7465206f6e6c7920696e2076616c696420696e6465785f8201527f2e206d6f7265207468616e2076616c6964000000000000000000000000000000602082015250565b5f610fc2603183610c20565b9150610fcd82610f68565b604082019050919050565b5f6020820190508181035f830152610fef81610fb6565b9050919050565b7f596f752063616e20766f7465206f6e6c7920696e2076616c696420696e6465785f8201527f2e206d696e75732076616c756500000000000000000000000000000000000000602082015250565b5f611050602d83610c20565b915061105b82610ff6565b604082019050919050565b5f6020820190508181035f83015261107d81611044565b9050919050565b7f54686973207265717565737420686173206265656e20636f6d706c657465642e5f82015250565b5f6110b8602083610c20565b91506110c382611084565b602082019050919050565b5f6020820190508181035f8301526110e5816110ac565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61115082610ab0565b915061115b83610ab0565b92508261116b5761116a6110ec565b5b828204905092915050565b7f546869732072657175657374206e65656473206d6f726520617070726f76616c5f8201527f73206265666f72652069742063616e2062652066696e616c697a656400000000602082015250565b5f6111d0603c83610c20565b91506111db82611176565b604082019050919050565b5f6020820190508181035f8301526111fd816111c4565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061124857607f821691505b60208210810361125b5761125a611204565b5b50919050565b5f61126b82610ab0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361129d5761129c611119565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826112c9565b61130e86836112c9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61134961134461133f84610ab0565b611326565b610ab0565b9050919050565b5f819050919050565b6113628361132f565b61137661136e82611350565b8484546112d5565b825550505050565b5f90565b61138a61137e565b611395818484611359565b505050565b5b818110156113b8576113ad5f82611382565b60018101905061139b565b5050565b601f8211156113fd576113ce816112a8565b6113d7846112ba565b810160208510156113e6578190505b6113fa6113f2856112ba565b83018261139a565b50505b505050565b5f82821c905092915050565b5f61141d5f1984600802611402565b1980831691505092915050565b5f611435838361140e565b9150826002028217905092915050565b61144e82610c16565b67ffffffffffffffff81111561146757611466610d20565b5b6114718254611231565b61147c8282856113bc565b5f60209050601f8311600181146114ad575f841561149b578287015190505b6114a5858261142a565b86555061150c565b601f1984166114bb866112a8565b5f5b828110156114e2578489015182556001820191506020850194506020810190506114bd565b868310156114ff57848901516114fb601f89168261140e565b8355505b6001600288020188555050505b505050505050565b7f4f6e6c7920636f6e7472696275746f72732063616e20617070726f76652061205f8201527f7370656369666963207061796d656e7420726571756573740000000000000000602082015250565b5f61156e603883610c20565b915061157982611514565b604082019050919050565b5f6020820190508181035f83015261159b81611562565b9050919050565b7f596f75206861766520616c726561647920766f74656420746f20617070726f765f8201527f6520746869732072657175657374000000000000000000000000000000000000602082015250565b5f6115fc602e83610c20565b9150611607826115a2565b604082019050919050565b5f6020820190508181035f830152611629816115f0565b9050919050565b7f6d696e696d756d20636f6e747269627574652069732072657175697265642e005f82015250565b5f611664601f83610c20565b915061166f82611630565b602082019050919050565b5f6020820190508181035f83015261169181611658565b905091905056fea2646970667358221220c03a9713e929980965af701f1cdcc46662bfcfda56a5b3133084db66066c605d64736f6c637829302e382e32342d646576656c6f702e323032332e31322e31392b636f6d6d69742e3932663338336438005a"

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

// IsPaied is a free data retrieval call binding the contract method 0x30378bf0.
//
// Ylem: function isPaied() view returns(bool)
func (_Campaign *CampaignCaller) IsPaied(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "isPaied")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaied is a free data retrieval call binding the contract method 0x30378bf0.
//
// Ylem: function isPaied() view returns(bool)
func (_Campaign *CampaignSession) IsPaied() (bool, error) {
	return _Campaign.Contract.IsPaied(&_Campaign.CallOpts)
}

// IsPaied is a free data retrieval call binding the contract method 0x30378bf0.
//
// Ylem: function isPaied() view returns(bool)
func (_Campaign *CampaignCallerSession) IsPaied() (bool, error) {
	return _Campaign.Contract.IsPaied(&_Campaign.CallOpts)
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

// Contibute is a paid mutator transaction binding the contract method 0x8435fdb1.
//
// Ylem: function contibute() payable returns()
func (_Campaign *CampaignTransactor) Contibute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "contibute")
}

// Contibute is a paid mutator transaction binding the contract method 0x8435fdb1.
//
// Ylem: function contibute() payable returns()
func (_Campaign *CampaignSession) Contibute() (*types.Transaction, error) {
	return _Campaign.Contract.Contibute(&_Campaign.TransactOpts)
}

// Contibute is a paid mutator transaction binding the contract method 0x8435fdb1.
//
// Ylem: function contibute() payable returns()
func (_Campaign *CampaignTransactorSession) Contibute() (*types.Transaction, error) {
	return _Campaign.Contract.Contibute(&_Campaign.TransactOpts)
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
