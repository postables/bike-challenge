// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bike_token

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

// BikeTokenABI is the input ABI used to generate the binding from.
const BikeTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"transferredFrom\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"transferred\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// BikeTokenBin is the compiled bytecode used for deploying new contracts.
const BikeTokenBin = `6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319918216811790925560018054909116909117905560408051908101604052600981527f42696b65546f6b656e000000000000000000000000000000000000000000000060208201526002908051610081929160200190610106565b5060408051908101604052600281527f4254000000000000000000000000000000000000000000000000000000000000602082015260039080516100c9929160200190610106565b506005805460ff191660121790556a108b2a2c280290940000006004819055600160a060020a0333166000908152600660205260409020556101a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014757805160ff1916838001178555610174565b82800160010185558215610174579182015b82811115610174578251825591602001919060010190610159565b50610180929150610184565b5090565b61019e91905b80821115610180576000815560010161018a565b90565b610a65806101b06000396000f3006060604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100df578063095ea7b31461016957806318160ddd1461019f57806323b872dd146101c457806327e235e3146101ec578063313ce5671461020b5780635c65816514610234578063704b6c021461025957806370a08231146102785780638da5cb5b1461029757806395d89b41146102c6578063a6f9dae1146102d9578063a9059cbb146102f8578063dd62ed3e1461031a578063f851a4401461033f575b600080fd5b34156100ea57600080fd5b6100f2610352565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012e578082015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017457600080fd5b61018b600160a060020a03600435166024356103f0565b604051901515815260200160405180910390f35b34156101aa57600080fd5b6101b2610494565b60405190815260200160405180910390f35b34156101cf57600080fd5b61018b600160a060020a036004358116906024351660443561049a565b34156101f757600080fd5b6101b2600160a060020a036004351661063e565b341561021657600080fd5b61021e610650565b60405160ff909116815260200160405180910390f35b341561023f57600080fd5b6101b2600160a060020a0360043581169060243516610659565b341561026457600080fd5b61018b600160a060020a0360043516610676565b341561028357600080fd5b6101b2600160a060020a03600435166106ff565b34156102a257600080fd5b6102aa61071a565b604051600160a060020a03909116815260200160405180910390f35b34156102d157600080fd5b6100f2610729565b34156102e457600080fd5b61018b600160a060020a0360043516610794565b341561030357600080fd5b61018b600160a060020a036004351660243561081d565b341561032557600080fd5b6101b2600160a060020a03600435811690602435166108f4565b341561034a57600080fd5b6102aa61091f565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e85780601f106103bd576101008083540402835291602001916103e8565b820191906000526020600020905b8154815290600101906020018083116103cb57829003601f168201915b505050505081565b600160a060020a033381166000908152600760209081526040808320938616835292905290812054610428908363ffffffff61092e16565b600160a060020a03338116600081815260076020908152604080832094891680845294909152908190209390935590917f6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e9085905190815260200160405180910390a350600192915050565b60045490565b60006104a7848484610947565b15156104b257600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290522054829010156104e657600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290529081205461051f908463ffffffff610a2416565b101561052a57600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290522054610561908363ffffffff610a2416565b600160a060020a0380861660008181526007602090815260408083203390951683529381528382209490945590815260069092529020546105a8908363ffffffff610a2416565b600160a060020a0380861660009081526006602052604080822093909355908516815220546105dd908363ffffffff61092e16565b600160a060020a03808516600081815260066020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60066020526000908152604090205481565b60055460ff1681565b600760209081526000928352604080842090915290825290205481565b6000805433600160a060020a0390811691161461069257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b600160a060020a031660009081526006602052604090205490565b600054600160a060020a031681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e85780601f106103bd576101008083540402835291602001916103e8565b6000805433600160a060020a039081169116146107b057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b600061082a338484610947565b151561083557600080fd5b600160a060020a03331660009081526006602052604090205461085e908363ffffffff610a2416565b600160a060020a033381166000908152600660205260408082209390935590851681522054610893908363ffffffff61092e16565b600160a060020a0380851660008181526006602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260076020908152604080832093909416825291909152205490565b600154600160a060020a031681565b60008282018381101561094057600080fd5b9392505050565b6000600160a060020a038416158015906109695750600160a060020a03831615155b80156109755750600082115b151561098057600080fd5b600160a060020a0384166000908152600660205260408120546109a9908463ffffffff610a2416565b10156109b457600080fd5b600160a060020a0383166000908152600660205260408120546109dd908463ffffffff61092e16565b116109e757600080fd5b600160a060020a038316600090815260066020526040902054610a10818463ffffffff61092e16565b11610a1a57600080fd5b5060019392505050565b600082821115610a3357600080fd5b509003905600a165627a7a7230582082a389e55dfae8434b69adc9f31041d11d0bf7fbc751746e370d12edce5409b00029`

// DeployBikeToken deploys a new Ethereum contract, binding an instance of BikeToken to it.
func DeployBikeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BikeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BikeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BikeToken{BikeTokenCaller: BikeTokenCaller{contract: contract}, BikeTokenTransactor: BikeTokenTransactor{contract: contract}, BikeTokenFilterer: BikeTokenFilterer{contract: contract}}, nil
}

// BikeToken is an auto generated Go binding around an Ethereum contract.
type BikeToken struct {
	BikeTokenCaller     // Read-only binding to the contract
	BikeTokenTransactor // Write-only binding to the contract
	BikeTokenFilterer   // Log filterer for contract events
}

// BikeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeTokenSession struct {
	Contract     *BikeToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeTokenCallerSession struct {
	Contract *BikeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BikeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeTokenTransactorSession struct {
	Contract     *BikeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BikeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeTokenRaw struct {
	Contract *BikeToken // Generic contract binding to access the raw methods on
}

// BikeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeTokenCallerRaw struct {
	Contract *BikeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BikeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeTokenTransactorRaw struct {
	Contract *BikeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBikeToken creates a new instance of BikeToken, bound to a specific deployed contract.
func NewBikeToken(address common.Address, backend bind.ContractBackend) (*BikeToken, error) {
	contract, err := bindBikeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BikeToken{BikeTokenCaller: BikeTokenCaller{contract: contract}, BikeTokenTransactor: BikeTokenTransactor{contract: contract}, BikeTokenFilterer: BikeTokenFilterer{contract: contract}}, nil
}

// NewBikeTokenCaller creates a new read-only instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenCaller(address common.Address, caller bind.ContractCaller) (*BikeTokenCaller, error) {
	contract, err := bindBikeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTokenCaller{contract: contract}, nil
}

// NewBikeTokenTransactor creates a new write-only instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeTokenTransactor, error) {
	contract, err := bindBikeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTokenTransactor{contract: contract}, nil
}

// NewBikeTokenFilterer creates a new log filterer instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeTokenFilterer, error) {
	contract, err := bindBikeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeTokenFilterer{contract: contract}, nil
}

// bindBikeToken binds a generic wrapper to an already deployed contract.
func bindBikeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeToken *BikeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeToken.Contract.BikeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeToken *BikeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.Contract.BikeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeToken *BikeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeToken.Contract.BikeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeToken *BikeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeToken *BikeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeToken *BikeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeToken.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeToken *BikeTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeToken *BikeTokenSession) Admin() (common.Address, error) {
	return _BikeToken.Contract.Admin(&_BikeToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeToken *BikeTokenCallerSession) Admin() (common.Address, error) {
	return _BikeToken.Contract.Admin(&_BikeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowance(&_BikeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowance(&_BikeToken.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowed(&_BikeToken.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowed(&_BikeToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) BalanceOf(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "balanceOf", _holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _BikeToken.Contract.BalanceOf(&_BikeToken.CallOpts, _holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _BikeToken.Contract.BalanceOf(&_BikeToken.CallOpts, _holder)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Balances(&_BikeToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Balances(&_BikeToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BikeToken *BikeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BikeToken *BikeTokenSession) Decimals() (uint8, error) {
	return _BikeToken.Contract.Decimals(&_BikeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BikeToken *BikeTokenCallerSession) Decimals() (uint8, error) {
	return _BikeToken.Contract.Decimals(&_BikeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenSession) Name() (string, error) {
	return _BikeToken.Contract.Name(&_BikeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenCallerSession) Name() (string, error) {
	return _BikeToken.Contract.Name(&_BikeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeToken *BikeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeToken *BikeTokenSession) Owner() (common.Address, error) {
	return _BikeToken.Contract.Owner(&_BikeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeToken *BikeTokenCallerSession) Owner() (common.Address, error) {
	return _BikeToken.Contract.Owner(&_BikeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenSession) Symbol() (string, error) {
	return _BikeToken.Contract.Symbol(&_BikeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenCallerSession) Symbol() (string, error) {
	return _BikeToken.Contract.Symbol(&_BikeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenSession) TotalSupply() (*big.Int, error) {
	return _BikeToken.Contract.TotalSupply(&_BikeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BikeToken.Contract.TotalSupply(&_BikeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_BikeToken *BikeTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_BikeToken *BikeTokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Approve(&_BikeToken.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_BikeToken *BikeTokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Approve(&_BikeToken.TransactOpts, _spender, _amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeToken *BikeTokenTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeToken *BikeTokenSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BikeToken.Contract.ChangeOwner(&_BikeToken.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeToken *BikeTokenTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BikeToken.Contract.ChangeOwner(&_BikeToken.TransactOpts, _owner)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeToken *BikeTokenTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeToken *BikeTokenSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BikeToken.Contract.SetAdmin(&_BikeToken.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeToken *BikeTokenTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BikeToken.Contract.SetAdmin(&_BikeToken.TransactOpts, _admin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_BikeToken *BikeTokenTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_BikeToken *BikeTokenSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Transfer(&_BikeToken.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_BikeToken *BikeTokenTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Transfer(&_BikeToken.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_BikeToken *BikeTokenTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "transferFrom", _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_BikeToken *BikeTokenSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.TransferFrom(&_BikeToken.TransactOpts, _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_BikeToken *BikeTokenTransactorSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.TransferFrom(&_BikeToken.TransactOpts, _owner, _recipient, _amount)
}

// BikeTokenAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the BikeToken contract.
type BikeTokenAdminSetIterator struct {
	Event *BikeTokenAdminSet // Event containing the contract specifics and raw log

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
func (it *BikeTokenAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenAdminSet)
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
		it.Event = new(BikeTokenAdminSet)
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
func (it *BikeTokenAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenAdminSet represents a AdminSet event raised by the BikeToken contract.
type BikeTokenAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BikeToken *BikeTokenFilterer) FilterAdminSet(opts *bind.FilterOpts) (*BikeTokenAdminSetIterator, error) {

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &BikeTokenAdminSetIterator{contract: _BikeToken.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BikeToken *BikeTokenFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *BikeTokenAdminSet) (event.Subscription, error) {

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenAdminSet)
				if err := _BikeToken.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// BikeTokenApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the BikeToken contract.
type BikeTokenApproveIterator struct {
	Event *BikeTokenApprove // Event containing the contract specifics and raw log

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
func (it *BikeTokenApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenApprove)
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
		it.Event = new(BikeTokenApprove)
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
func (it *BikeTokenApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenApprove represents a Approve event raised by the BikeToken contract.
type BikeTokenApprove struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(_owner indexed address, _spender indexed address, _amount uint256)
func (_BikeToken *BikeTokenFilterer) FilterApprove(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*BikeTokenApproveIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "Approve", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &BikeTokenApproveIterator{contract: _BikeToken.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(_owner indexed address, _spender indexed address, _amount uint256)
func (_BikeToken *BikeTokenFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *BikeTokenApprove, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "Approve", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenApprove)
				if err := _BikeToken.contract.UnpackLog(event, "Approve", log); err != nil {
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

// BikeTokenNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the BikeToken contract.
type BikeTokenNewOwnerIterator struct {
	Event *BikeTokenNewOwner // Event containing the contract specifics and raw log

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
func (it *BikeTokenNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenNewOwner)
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
		it.Event = new(BikeTokenNewOwner)
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
func (it *BikeTokenNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenNewOwner represents a NewOwner event raised by the BikeToken contract.
type BikeTokenNewOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BikeToken *BikeTokenFilterer) FilterNewOwner(opts *bind.FilterOpts) (*BikeTokenNewOwnerIterator, error) {

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &BikeTokenNewOwnerIterator{contract: _BikeToken.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BikeToken *BikeTokenFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *BikeTokenNewOwner) (event.Subscription, error) {

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenNewOwner)
				if err := _BikeToken.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// BikeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BikeToken contract.
type BikeTokenTransferIterator struct {
	Event *BikeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BikeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenTransfer)
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
		it.Event = new(BikeTokenTransfer)
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
func (it *BikeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenTransfer represents a Transfer event raised by the BikeToken contract.
type BikeTokenTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_BikeToken *BikeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*BikeTokenTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &BikeTokenTransferIterator{contract: _BikeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_BikeToken *BikeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BikeTokenTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenTransfer)
				if err := _BikeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
