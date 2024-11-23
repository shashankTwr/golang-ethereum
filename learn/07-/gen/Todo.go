// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Task      string
	Completed bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Todo__NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_task\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"deleteTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"task\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"readTask\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"task\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"todo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"readTaskCompletedOnIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"readTaskContentOnIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"task\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"task\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"toggleTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506116268061005c5f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c8063893d20e811610064578063893d20e8146101425780638d97767214610160578063ab755faf14610191578063c213c527146101c1578063e1e29558146101f15761009c565b806350f2ba11146100a0578063560f3192146100d0578063613137a4146100ec57806367238562146101085780636a2b588414610124575b5f5ffd5b6100ba60048036038101906100b59190610cc3565b61020d565b6040516100c79190610d5e565b60405180910390f35b6100ea60048036038101906100e59190610cc3565b610344565b005b61010660048036038101906101019190610cc3565b6104c0565b005b610122600480360381019061011d9190610eaa565b6105b1565b005b61012c6106b2565b6040516101399190611048565b60405180910390f35b61014a61083c565b60405161015791906110a7565b60405180910390f35b61017a60048036038101906101759190610cc3565b610864565b6040516101889291906110cf565b60405180910390f35b6101ab60048036038101906101a69190610cc3565b610924565b6040516101b89190611137565b60405180910390f35b6101db60048036038101906101d69190610cc3565b610a8c565b6040516101e89190611157565b60405180910390f35b61020b60048036038101906102069190610cc3565b610b49565b005b606060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610295576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82815481106102a8576102a7611170565b5b905f5260205f2090600202015f0180546102c1906111ca565b80601f01602080910402602001604051908101604052809291908181526020018280546102ed906111ca565b80156103385780601f1061030f57610100808354040283529160200191610338565b820191905f5260205f20905b81548152906001019060200180831161031b57829003601f168201915b50505050509050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ca576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8190505b5f80549050811015610473575f81815481106103ee576103ed611170565b5b905f5260205f2090600202015f6001836104089190611227565b8154811061041957610418611170565b5b905f5260205f2090600202015f8201815f019081610437919061140f565b50600182015f9054906101000a900460ff16816001015f6101000a81548160ff02191690831515021790555090505080806001019150506103cf565b505f805480610485576104846114f4565b5b600190038181905f5260205f2090600202015f5f82015f6104a69190610c0c565b600182015f6101000a81549060ff02191690555050905550565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610546576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818154811061055957610558611170565b5b905f5260205f2090600202016001015f9054906101000a900460ff16155f828154811061058957610588611170565b5b905f5260205f2090600202016001015f6101000a81548160ff02191690831515021790555050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610637576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60405180604001604052808381526020015f151581525090505f81908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01908161068c9190611521565b506020820151816001015f6101000a81548160ff02191690831515021790555050505050565b606060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073a576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805480602002602001604051908101604052809291908181526020015f905b82821015610833578382905f5260205f2090600202016040518060400160405290815f8201805461078a906111ca565b80601f01602080910402602001604051908101604052809291908181526020018280546107b6906111ca565b80156108015780601f106107d857610100808354040283529160200191610801565b820191905f5260205f20905b8154815290600101906020018083116107e457829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250508152602001906001019061075a565b50505050905090565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f8181548110610872575f80fd5b905f5260205f2090600202015f91509050805f018054610891906111ca565b80601f01602080910402602001604051908101604052809291908181526020018280546108bd906111ca565b80156109085780601f106108df57610100808354040283529160200191610908565b820191905f5260205f20905b8154815290600101906020018083116108eb57829003601f168201915b505050505090806001015f9054906101000a900460ff16905082565b61092c610c49565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109b2576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82815481106109c5576109c4611170565b5b905f5260205f2090600202016040518060400160405290815f820180546109eb906111ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610a17906111ca565b8015610a625780601f10610a3957610100808354040283529160200191610a62565b820191905f5260205f20905b815481529060010190602001808311610a4557829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050919050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b13576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281548110610b2657610b25611170565b5b905f5260205f2090600202016001015f9054906101000a900460ff169050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bcf576040517f40eaa30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f8281548110610be457610be3611170565b5b905f5260205f2090600202016001015f6101000a81548160ff02191690831515021790555050565b508054610c18906111ca565b5f825580601f10610c295750610c46565b601f0160209004905f5260205f2090810190610c459190610c64565b5b50565b6040518060400160405280606081526020015f151581525090565b5b80821115610c7b575f815f905550600101610c65565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610ca281610c90565b8114610cac575f5ffd5b50565b5f81359050610cbd81610c99565b92915050565b5f60208284031215610cd857610cd7610c88565b5b5f610ce584828501610caf565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610d3082610cee565b610d3a8185610cf8565b9350610d4a818560208601610d08565b610d5381610d16565b840191505092915050565b5f6020820190508181035f830152610d768184610d26565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610dbc82610d16565b810181811067ffffffffffffffff82111715610ddb57610dda610d86565b5b80604052505050565b5f610ded610c7f565b9050610df98282610db3565b919050565b5f67ffffffffffffffff821115610e1857610e17610d86565b5b610e2182610d16565b9050602081019050919050565b828183375f83830152505050565b5f610e4e610e4984610dfe565b610de4565b905082815260208101848484011115610e6a57610e69610d82565b5b610e75848285610e2e565b509392505050565b5f82601f830112610e9157610e90610d7e565b5b8135610ea1848260208601610e3c565b91505092915050565b5f60208284031215610ebf57610ebe610c88565b5b5f82013567ffffffffffffffff811115610edc57610edb610c8c565b5b610ee884828501610e7d565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f610f3482610cee565b610f3e8185610f1a565b9350610f4e818560208601610d08565b610f5781610d16565b840191505092915050565b5f8115159050919050565b610f7681610f62565b82525050565b5f604083015f8301518482035f860152610f968282610f2a565b9150506020830151610fab6020860182610f6d565b508091505092915050565b5f610fc18383610f7c565b905092915050565b5f602082019050919050565b5f610fdf82610ef1565b610fe98185610efb565b935083602082028501610ffb85610f0b565b805f5b8581101561103657848403895281516110178582610fb6565b945061102283610fc9565b925060208a01995050600181019050610ffe565b50829750879550505050505092915050565b5f6020820190508181035f8301526110608184610fd5565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61109182611068565b9050919050565b6110a181611087565b82525050565b5f6020820190506110ba5f830184611098565b92915050565b6110c981610f62565b82525050565b5f6040820190508181035f8301526110e78185610d26565b90506110f660208301846110c0565b9392505050565b5f604083015f8301518482035f8601526111178282610f2a565b915050602083015161112c6020860182610f6d565b508091505092915050565b5f6020820190508181035f83015261114f81846110fd565b905092915050565b5f60208201905061116a5f8301846110c0565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111e157607f821691505b6020821081036111f4576111f361119d565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61123182610c90565b915061123c83610c90565b9250828203905081811115611254576112536111fa565b5b92915050565b5f81549050611268816111ca565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112cb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611290565b6112d58683611290565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61131061130b61130684610c90565b6112ed565b610c90565b9050919050565b5f819050919050565b611329836112f6565b61133d61133582611317565b84845461129c565b825550505050565b5f5f905090565b611354611345565b61135f818484611320565b505050565b5b81811015611382576113775f8261134c565b600181019050611365565b5050565b601f8211156113c7576113988161126f565b6113a184611281565b810160208510156113b0578190505b6113c46113bc85611281565b830182611364565b50505b505050565b5f82821c905092915050565b5f6113e75f19846008026113cc565b1980831691505092915050565b5f6113ff83836113d8565b9150826002028217905092915050565b81810361141d5750506114f2565b6114268261125a565b67ffffffffffffffff81111561143f5761143e610d86565b5b61144982546111ca565b611454828285611386565b5f601f831160018114611481575f841561146f578287015490505b61147985826113f4565b8655506114eb565b601f19841661148f8761126f565b965061149a8661126f565b5f5b828110156114c15784890154825560018201915060018501945060208101905061149c565b868310156114de57848901546114da601f8916826113d8565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b61152a82610cee565b67ffffffffffffffff81111561154357611542610d86565b5b61154d82546111ca565b611558828285611386565b5f60209050601f831160018114611589575f8415611577578287015190505b61158185826113f4565b8655506115e8565b601f1984166115978661126f565b5f5b828110156115be57848901518255600182019150602085019450602081019050611599565b868310156115db57848901516115d7601f8916826113d8565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220305aa4962cb3424634eef54e9b6fd80edeadcc98d02b49c417e35835a74f565364736f6c634300081c0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Todo *TodoCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Todo *TodoSession) GetOwner() (common.Address, error) {
	return _Todo.Contract.GetOwner(&_Todo.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Todo *TodoCallerSession) GetOwner() (common.Address, error) {
	return _Todo.Contract.GetOwner(&_Todo.CallOpts)
}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoCaller) ListTasks(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "listTasks")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoSession) ListTasks() ([]TodoTask, error) {
	return _Todo.Contract.ListTasks(&_Todo.CallOpts)
}

// ListTasks is a free data retrieval call binding the contract method 0x6a2b5884.
//
// Solidity: function listTasks() view returns((string,bool)[])
func (_Todo *TodoCallerSession) ListTasks() ([]TodoTask, error) {
	return _Todo.Contract.ListTasks(&_Todo.CallOpts)
}

// ReadTask is a free data retrieval call binding the contract method 0xab755faf.
//
// Solidity: function readTask(uint256 _index) view returns((string,bool) todo)
func (_Todo *TodoCaller) ReadTask(opts *bind.CallOpts, _index *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "readTask", _index)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// ReadTask is a free data retrieval call binding the contract method 0xab755faf.
//
// Solidity: function readTask(uint256 _index) view returns((string,bool) todo)
func (_Todo *TodoSession) ReadTask(_index *big.Int) (TodoTask, error) {
	return _Todo.Contract.ReadTask(&_Todo.CallOpts, _index)
}

// ReadTask is a free data retrieval call binding the contract method 0xab755faf.
//
// Solidity: function readTask(uint256 _index) view returns((string,bool) todo)
func (_Todo *TodoCallerSession) ReadTask(_index *big.Int) (TodoTask, error) {
	return _Todo.Contract.ReadTask(&_Todo.CallOpts, _index)
}

// ReadTaskCompletedOnIndex is a free data retrieval call binding the contract method 0xc213c527.
//
// Solidity: function readTaskCompletedOnIndex(uint256 _index) view returns(bool)
func (_Todo *TodoCaller) ReadTaskCompletedOnIndex(opts *bind.CallOpts, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "readTaskCompletedOnIndex", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReadTaskCompletedOnIndex is a free data retrieval call binding the contract method 0xc213c527.
//
// Solidity: function readTaskCompletedOnIndex(uint256 _index) view returns(bool)
func (_Todo *TodoSession) ReadTaskCompletedOnIndex(_index *big.Int) (bool, error) {
	return _Todo.Contract.ReadTaskCompletedOnIndex(&_Todo.CallOpts, _index)
}

// ReadTaskCompletedOnIndex is a free data retrieval call binding the contract method 0xc213c527.
//
// Solidity: function readTaskCompletedOnIndex(uint256 _index) view returns(bool)
func (_Todo *TodoCallerSession) ReadTaskCompletedOnIndex(_index *big.Int) (bool, error) {
	return _Todo.Contract.ReadTaskCompletedOnIndex(&_Todo.CallOpts, _index)
}

// ReadTaskContentOnIndex is a free data retrieval call binding the contract method 0x50f2ba11.
//
// Solidity: function readTaskContentOnIndex(uint256 _index) view returns(string task)
func (_Todo *TodoCaller) ReadTaskContentOnIndex(opts *bind.CallOpts, _index *big.Int) (string, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "readTaskContentOnIndex", _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadTaskContentOnIndex is a free data retrieval call binding the contract method 0x50f2ba11.
//
// Solidity: function readTaskContentOnIndex(uint256 _index) view returns(string task)
func (_Todo *TodoSession) ReadTaskContentOnIndex(_index *big.Int) (string, error) {
	return _Todo.Contract.ReadTaskContentOnIndex(&_Todo.CallOpts, _index)
}

// ReadTaskContentOnIndex is a free data retrieval call binding the contract method 0x50f2ba11.
//
// Solidity: function readTaskContentOnIndex(uint256 _index) view returns(string task)
func (_Todo *TodoCallerSession) ReadTaskContentOnIndex(_index *big.Int) (string, error) {
	return _Todo.Contract.ReadTaskContentOnIndex(&_Todo.CallOpts, _index)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string task, bool completed)
func (_Todo *TodoCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Task      string
	Completed bool
}, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Task      string
		Completed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Task = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Completed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string task, bool completed)
func (_Todo *TodoSession) Tasks(arg0 *big.Int) (struct {
	Task      string
	Completed bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string task, bool completed)
func (_Todo *TodoCallerSession) Tasks(arg0 *big.Int) (struct {
	Task      string
	Completed bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _task) returns()
func (_Todo *TodoTransactor) AddTask(opts *bind.TransactOpts, _task string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "addTask", _task)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _task) returns()
func (_Todo *TodoSession) AddTask(_task string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _task)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _task) returns()
func (_Todo *TodoTransactorSession) AddTask(_task string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _task)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _index) returns()
func (_Todo *TodoTransactor) CompleteTask(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "completeTask", _index)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _index) returns()
func (_Todo *TodoSession) CompleteTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.CompleteTask(&_Todo.TransactOpts, _index)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _index) returns()
func (_Todo *TodoTransactorSession) CompleteTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.CompleteTask(&_Todo.TransactOpts, _index)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _index) returns()
func (_Todo *TodoTransactor) DeleteTask(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "deleteTask", _index)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _index) returns()
func (_Todo *TodoSession) DeleteTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTask(&_Todo.TransactOpts, _index)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _index) returns()
func (_Todo *TodoTransactorSession) DeleteTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTask(&_Todo.TransactOpts, _index)
}

// ToggleTask is a paid mutator transaction binding the contract method 0x613137a4.
//
// Solidity: function toggleTask(uint256 _index) returns()
func (_Todo *TodoTransactor) ToggleTask(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggleTask", _index)
}

// ToggleTask is a paid mutator transaction binding the contract method 0x613137a4.
//
// Solidity: function toggleTask(uint256 _index) returns()
func (_Todo *TodoSession) ToggleTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.ToggleTask(&_Todo.TransactOpts, _index)
}

// ToggleTask is a paid mutator transaction binding the contract method 0x613137a4.
//
// Solidity: function toggleTask(uint256 _index) returns()
func (_Todo *TodoTransactorSession) ToggleTask(_index *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.ToggleTask(&_Todo.TransactOpts, _index)
}
