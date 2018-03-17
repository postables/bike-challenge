// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bt_eth

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

// BtEthABI is the input ABI used to generate the binding from.
const BtEthABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"launch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"contributionCalculator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numContributions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"btWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numBikeTokensSold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfLaunched\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_purchaser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_contribution\",\"type\":\"uint256\"}],\"name\":\"TokensSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractLaunchedAndEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SalesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SalesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// BtEthBin is the compiled bytecode used for deploying new contracts.
const BtEthBin = `606060405260008054600160a060020a033316600160a060020a031991821681179092556001805490911690911790556109d08061003e6000396000f3006060604052600436106100d75763ffffffff60e060020a600035041663158ef93e8114610113578063214013ca1461013a578063238dafe01461015957806342e94c901461016c578063524592a11461019d57806361348708146101b3578063630b4989146101c6578063704b6c02146101d957806382f5decb146101f857806388a244ee1461020b5780638da5cb5b1461021e578063a6f9dae11461024d578063bf5839031461026c578063c63df5ef1461027f578063e3b172d714610292578063f851a440146102a5578063fddf0fc0146102b8575b600654610100900460ff1615156100ed57600080fd5b60065460ff1615156100fe57600080fd5b6101066102cb565b151561011157600080fd5b005b341561011e57600080fd5b6101266104be565b604051901515815260200160405180910390f35b341561014557600080fd5b610126600160a060020a03600435166104cc565b341561016457600080fd5b61012661065f565b341561017757600080fd5b61018b600160a060020a0360043516610668565b60405190815260200160405180910390f35b34156101a857600080fd5b61018b60043561067a565b34156101be57600080fd5b6101266106a9565b34156101d157600080fd5b61018b610729565b34156101e457600080fd5b610126600160a060020a036004351661072f565b341561020357600080fd5b61018b6107b8565b341561021657600080fd5b61018b6107c2565b341561022957600080fd5b6102316107c8565b604051600160a060020a03909116815260200160405180910390f35b341561025857600080fd5b610126600160a060020a03600435166107d7565b341561027757600080fd5b61018b610860565b341561028a57600080fd5b610126610866565b341561029d57600080fd5b6101266108e7565b34156102b057600080fd5b610231610923565b34156102c357600080fd5b61018b610932565b600080348190116102db57600080fd5b600254600090116102eb57600080fd5b61032a67016345785d8a000061031e651cf33b72a60061031234600a63ffffffff61093816565b9063ffffffff61096616565b9063ffffffff61093816565b90506000811161033657fe5b600160a060020a03331660009081526007602052604090205461035f908263ffffffff61097d16565b600160a060020a03331660009081526007602052604090205560025461038b908263ffffffff61098f16565b6002556003546103a1908263ffffffff61097d16565b6003556004546103b890600163ffffffff61097d16565b60045560025415156103fb576006805460ff191690557f3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e4760405160405180910390a15b33600160a060020a03167f57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e8260405190815260200160405180910390a2600654620100009004600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561049457600080fd5b5af115156104a157600080fd5b5050506040518051905015156104b657600080fd5b600191505090565b600654610100900460ff1681565b6000805433600160a060020a039081169116146104e857600080fd5b600654610100900460ff16156104fd57600080fd5b6006805475ffffffffffffffffffffffffffffffffffffffff0000191662010000600160a060020a0385811682029290921792839055651cf33b72a6009204166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561058157600080fd5b5af1151561058e57600080fd5b505050604051805190501115156105a457600080fd5b600654620100009004600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156105fa57600080fd5b5af1151561060757600080fd5b5050506040518051600255506006805460ff1961ff0019909116610100171660011790557f68f3e71400a107502ed4901f333a22781427bea3f4c32385e9c80cb46468ec2660405160405180910390a1506001919050565b60065460ff1681565b60076020526000908152604090205481565b60006106a367016345785d8a000061031e651cf33b72a60061031286600a63ffffffff61093816565b92915050565b6000805433600160a060020a039081169116146106c557600080fd5b60065460ff1615156106d657600080fd5b600654610100900460ff1615156106ec57600080fd5b6006805460ff191690557f3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e4760405160405180910390a15060015b90565b60045481565b6000805433600160a060020a0390811691161461074b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b651cf33b72a60081565b60035481565b600054600160a060020a031681565b6000805433600160a060020a039081169116146107f357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b60025481565b6000805433600160a060020a0390811691161461088257600080fd5b60065460ff161561089257600080fd5b600654610100900460ff1615156108a857600080fd5b6006805460ff191660011790557fec8f62fea11b2d0b268c0f08bc273cef8aeeb2347809eaf5ce4a1a04bb19c99160405160405180910390a150600190565b60065460009060ff610100909104161515600114801561090e575060065460ff1615156001145b1561091b57506001610726565b506000610726565b600154600160a060020a031681565b60055481565b6000828202831580610954575082848281151561095157fe5b04145b151561095f57600080fd5b9392505050565b600080828481151561097457fe5b04949350505050565b60008282018381101561095f57600080fd5b60008282111561099e57600080fd5b509003905600a165627a7a7230582080dcf037f7c8f12a4ffbcf8fe824f7d77caafd4876d7c4b64e4837764bc2fc170029`

// DeployBtEth deploys a new Ethereum contract, binding an instance of BtEth to it.
func DeployBtEth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BtEth, error) {
	parsed, err := abi.JSON(strings.NewReader(BtEthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BtEthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BtEth{BtEthCaller: BtEthCaller{contract: contract}, BtEthTransactor: BtEthTransactor{contract: contract}, BtEthFilterer: BtEthFilterer{contract: contract}}, nil
}

// BtEth is an auto generated Go binding around an Ethereum contract.
type BtEth struct {
	BtEthCaller     // Read-only binding to the contract
	BtEthTransactor // Write-only binding to the contract
	BtEthFilterer   // Log filterer for contract events
}

// BtEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtEthSession struct {
	Contract     *BtEth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtEthCallerSession struct {
	Contract *BtEthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BtEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtEthTransactorSession struct {
	Contract     *BtEthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtEthRaw struct {
	Contract *BtEth // Generic contract binding to access the raw methods on
}

// BtEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtEthCallerRaw struct {
	Contract *BtEthCaller // Generic read-only contract binding to access the raw methods on
}

// BtEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtEthTransactorRaw struct {
	Contract *BtEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtEth creates a new instance of BtEth, bound to a specific deployed contract.
func NewBtEth(address common.Address, backend bind.ContractBackend) (*BtEth, error) {
	contract, err := bindBtEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BtEth{BtEthCaller: BtEthCaller{contract: contract}, BtEthTransactor: BtEthTransactor{contract: contract}, BtEthFilterer: BtEthFilterer{contract: contract}}, nil
}

// NewBtEthCaller creates a new read-only instance of BtEth, bound to a specific deployed contract.
func NewBtEthCaller(address common.Address, caller bind.ContractCaller) (*BtEthCaller, error) {
	contract, err := bindBtEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtEthCaller{contract: contract}, nil
}

// NewBtEthTransactor creates a new write-only instance of BtEth, bound to a specific deployed contract.
func NewBtEthTransactor(address common.Address, transactor bind.ContractTransactor) (*BtEthTransactor, error) {
	contract, err := bindBtEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtEthTransactor{contract: contract}, nil
}

// NewBtEthFilterer creates a new log filterer instance of BtEth, bound to a specific deployed contract.
func NewBtEthFilterer(address common.Address, filterer bind.ContractFilterer) (*BtEthFilterer, error) {
	contract, err := bindBtEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtEthFilterer{contract: contract}, nil
}

// bindBtEth binds a generic wrapper to an already deployed contract.
func bindBtEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BtEthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtEth *BtEthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BtEth.Contract.BtEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtEth *BtEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtEth.Contract.BtEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtEth *BtEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtEth.Contract.BtEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtEth *BtEthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BtEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtEth *BtEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtEth *BtEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtEth.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BtEth *BtEthCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BtEth *BtEthSession) Admin() (common.Address, error) {
	return _BtEth.Contract.Admin(&_BtEth.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BtEth *BtEthCallerSession) Admin() (common.Address, error) {
	return _BtEth.Contract.Admin(&_BtEth.CallOpts)
}

// BtWei is a free data retrieval call binding the contract method 0x82f5decb.
//
// Solidity: function btWei() constant returns(uint256)
func (_BtEth *BtEthCaller) BtWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "btWei")
	return *ret0, err
}

// BtWei is a free data retrieval call binding the contract method 0x82f5decb.
//
// Solidity: function btWei() constant returns(uint256)
func (_BtEth *BtEthSession) BtWei() (*big.Int, error) {
	return _BtEth.Contract.BtWei(&_BtEth.CallOpts)
}

// BtWei is a free data retrieval call binding the contract method 0x82f5decb.
//
// Solidity: function btWei() constant returns(uint256)
func (_BtEth *BtEthCallerSession) BtWei() (*big.Int, error) {
	return _BtEth.Contract.BtWei(&_BtEth.CallOpts)
}

// CheckIfLaunched is a free data retrieval call binding the contract method 0xe3b172d7.
//
// Solidity: function checkIfLaunched() constant returns(bool)
func (_BtEth *BtEthCaller) CheckIfLaunched(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "checkIfLaunched")
	return *ret0, err
}

// CheckIfLaunched is a free data retrieval call binding the contract method 0xe3b172d7.
//
// Solidity: function checkIfLaunched() constant returns(bool)
func (_BtEth *BtEthSession) CheckIfLaunched() (bool, error) {
	return _BtEth.Contract.CheckIfLaunched(&_BtEth.CallOpts)
}

// CheckIfLaunched is a free data retrieval call binding the contract method 0xe3b172d7.
//
// Solidity: function checkIfLaunched() constant returns(bool)
func (_BtEth *BtEthCallerSession) CheckIfLaunched() (bool, error) {
	return _BtEth.Contract.CheckIfLaunched(&_BtEth.CallOpts)
}

// ContributionCalculator is a free data retrieval call binding the contract method 0x524592a1.
//
// Solidity: function contributionCalculator(_eth uint256) constant returns(uint256)
func (_BtEth *BtEthCaller) ContributionCalculator(opts *bind.CallOpts, _eth *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "contributionCalculator", _eth)
	return *ret0, err
}

// ContributionCalculator is a free data retrieval call binding the contract method 0x524592a1.
//
// Solidity: function contributionCalculator(_eth uint256) constant returns(uint256)
func (_BtEth *BtEthSession) ContributionCalculator(_eth *big.Int) (*big.Int, error) {
	return _BtEth.Contract.ContributionCalculator(&_BtEth.CallOpts, _eth)
}

// ContributionCalculator is a free data retrieval call binding the contract method 0x524592a1.
//
// Solidity: function contributionCalculator(_eth uint256) constant returns(uint256)
func (_BtEth *BtEthCallerSession) ContributionCalculator(_eth *big.Int) (*big.Int, error) {
	return _BtEth.Contract.ContributionCalculator(&_BtEth.CallOpts, _eth)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_BtEth *BtEthCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "contributions", arg0)
	return *ret0, err
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_BtEth *BtEthSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _BtEth.Contract.Contributions(&_BtEth.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_BtEth *BtEthCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _BtEth.Contract.Contributions(&_BtEth.CallOpts, arg0)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_BtEth *BtEthCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_BtEth *BtEthSession) Enabled() (bool, error) {
	return _BtEth.Contract.Enabled(&_BtEth.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_BtEth *BtEthCallerSession) Enabled() (bool, error) {
	return _BtEth.Contract.Enabled(&_BtEth.CallOpts)
}

// EthRaised is a free data retrieval call binding the contract method 0xfddf0fc0.
//
// Solidity: function ethRaised() constant returns(uint256)
func (_BtEth *BtEthCaller) EthRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "ethRaised")
	return *ret0, err
}

// EthRaised is a free data retrieval call binding the contract method 0xfddf0fc0.
//
// Solidity: function ethRaised() constant returns(uint256)
func (_BtEth *BtEthSession) EthRaised() (*big.Int, error) {
	return _BtEth.Contract.EthRaised(&_BtEth.CallOpts)
}

// EthRaised is a free data retrieval call binding the contract method 0xfddf0fc0.
//
// Solidity: function ethRaised() constant returns(uint256)
func (_BtEth *BtEthCallerSession) EthRaised() (*big.Int, error) {
	return _BtEth.Contract.EthRaised(&_BtEth.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BtEth *BtEthCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BtEth *BtEthSession) Initialized() (bool, error) {
	return _BtEth.Contract.Initialized(&_BtEth.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BtEth *BtEthCallerSession) Initialized() (bool, error) {
	return _BtEth.Contract.Initialized(&_BtEth.CallOpts)
}

// NumBikeTokensSold is a free data retrieval call binding the contract method 0x88a244ee.
//
// Solidity: function numBikeTokensSold() constant returns(uint256)
func (_BtEth *BtEthCaller) NumBikeTokensSold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "numBikeTokensSold")
	return *ret0, err
}

// NumBikeTokensSold is a free data retrieval call binding the contract method 0x88a244ee.
//
// Solidity: function numBikeTokensSold() constant returns(uint256)
func (_BtEth *BtEthSession) NumBikeTokensSold() (*big.Int, error) {
	return _BtEth.Contract.NumBikeTokensSold(&_BtEth.CallOpts)
}

// NumBikeTokensSold is a free data retrieval call binding the contract method 0x88a244ee.
//
// Solidity: function numBikeTokensSold() constant returns(uint256)
func (_BtEth *BtEthCallerSession) NumBikeTokensSold() (*big.Int, error) {
	return _BtEth.Contract.NumBikeTokensSold(&_BtEth.CallOpts)
}

// NumContributions is a free data retrieval call binding the contract method 0x630b4989.
//
// Solidity: function numContributions() constant returns(uint256)
func (_BtEth *BtEthCaller) NumContributions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "numContributions")
	return *ret0, err
}

// NumContributions is a free data retrieval call binding the contract method 0x630b4989.
//
// Solidity: function numContributions() constant returns(uint256)
func (_BtEth *BtEthSession) NumContributions() (*big.Int, error) {
	return _BtEth.Contract.NumContributions(&_BtEth.CallOpts)
}

// NumContributions is a free data retrieval call binding the contract method 0x630b4989.
//
// Solidity: function numContributions() constant returns(uint256)
func (_BtEth *BtEthCallerSession) NumContributions() (*big.Int, error) {
	return _BtEth.Contract.NumContributions(&_BtEth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BtEth *BtEthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BtEth *BtEthSession) Owner() (common.Address, error) {
	return _BtEth.Contract.Owner(&_BtEth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BtEth *BtEthCallerSession) Owner() (common.Address, error) {
	return _BtEth.Contract.Owner(&_BtEth.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_BtEth *BtEthCaller) RemainingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BtEth.contract.Call(opts, out, "remainingTokens")
	return *ret0, err
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_BtEth *BtEthSession) RemainingTokens() (*big.Int, error) {
	return _BtEth.Contract.RemainingTokens(&_BtEth.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_BtEth *BtEthCallerSession) RemainingTokens() (*big.Int, error) {
	return _BtEth.Contract.RemainingTokens(&_BtEth.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BtEth *BtEthTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BtEth.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BtEth *BtEthSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.ChangeOwner(&_BtEth.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BtEth *BtEthTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.ChangeOwner(&_BtEth.TransactOpts, _owner)
}

// DisableSales is a paid mutator transaction binding the contract method 0x61348708.
//
// Solidity: function disableSales() returns(bool)
func (_BtEth *BtEthTransactor) DisableSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtEth.contract.Transact(opts, "disableSales")
}

// DisableSales is a paid mutator transaction binding the contract method 0x61348708.
//
// Solidity: function disableSales() returns(bool)
func (_BtEth *BtEthSession) DisableSales() (*types.Transaction, error) {
	return _BtEth.Contract.DisableSales(&_BtEth.TransactOpts)
}

// DisableSales is a paid mutator transaction binding the contract method 0x61348708.
//
// Solidity: function disableSales() returns(bool)
func (_BtEth *BtEthTransactorSession) DisableSales() (*types.Transaction, error) {
	return _BtEth.Contract.DisableSales(&_BtEth.TransactOpts)
}

// EnableSales is a paid mutator transaction binding the contract method 0xc63df5ef.
//
// Solidity: function enableSales() returns(bool)
func (_BtEth *BtEthTransactor) EnableSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtEth.contract.Transact(opts, "enableSales")
}

// EnableSales is a paid mutator transaction binding the contract method 0xc63df5ef.
//
// Solidity: function enableSales() returns(bool)
func (_BtEth *BtEthSession) EnableSales() (*types.Transaction, error) {
	return _BtEth.Contract.EnableSales(&_BtEth.TransactOpts)
}

// EnableSales is a paid mutator transaction binding the contract method 0xc63df5ef.
//
// Solidity: function enableSales() returns(bool)
func (_BtEth *BtEthTransactorSession) EnableSales() (*types.Transaction, error) {
	return _BtEth.Contract.EnableSales(&_BtEth.TransactOpts)
}

// Launch is a paid mutator transaction binding the contract method 0x214013ca.
//
// Solidity: function launch(_tokenAddress address) returns(bool)
func (_BtEth *BtEthTransactor) Launch(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _BtEth.contract.Transact(opts, "launch", _tokenAddress)
}

// Launch is a paid mutator transaction binding the contract method 0x214013ca.
//
// Solidity: function launch(_tokenAddress address) returns(bool)
func (_BtEth *BtEthSession) Launch(_tokenAddress common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.Launch(&_BtEth.TransactOpts, _tokenAddress)
}

// Launch is a paid mutator transaction binding the contract method 0x214013ca.
//
// Solidity: function launch(_tokenAddress address) returns(bool)
func (_BtEth *BtEthTransactorSession) Launch(_tokenAddress common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.Launch(&_BtEth.TransactOpts, _tokenAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BtEth *BtEthTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BtEth.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BtEth *BtEthSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.SetAdmin(&_BtEth.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BtEth *BtEthTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BtEth.Contract.SetAdmin(&_BtEth.TransactOpts, _admin)
}

// BtEthAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the BtEth contract.
type BtEthAdminSetIterator struct {
	Event *BtEthAdminSet // Event containing the contract specifics and raw log

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
func (it *BtEthAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthAdminSet)
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
		it.Event = new(BtEthAdminSet)
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
func (it *BtEthAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthAdminSet represents a AdminSet event raised by the BtEth contract.
type BtEthAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BtEth *BtEthFilterer) FilterAdminSet(opts *bind.FilterOpts) (*BtEthAdminSetIterator, error) {

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &BtEthAdminSetIterator{contract: _BtEth.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BtEth *BtEthFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *BtEthAdminSet) (event.Subscription, error) {

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthAdminSet)
				if err := _BtEth.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// BtEthContractLaunchedAndEnabledIterator is returned from FilterContractLaunchedAndEnabled and is used to iterate over the raw logs and unpacked data for ContractLaunchedAndEnabled events raised by the BtEth contract.
type BtEthContractLaunchedAndEnabledIterator struct {
	Event *BtEthContractLaunchedAndEnabled // Event containing the contract specifics and raw log

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
func (it *BtEthContractLaunchedAndEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthContractLaunchedAndEnabled)
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
		it.Event = new(BtEthContractLaunchedAndEnabled)
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
func (it *BtEthContractLaunchedAndEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthContractLaunchedAndEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthContractLaunchedAndEnabled represents a ContractLaunchedAndEnabled event raised by the BtEth contract.
type BtEthContractLaunchedAndEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractLaunchedAndEnabled is a free log retrieval operation binding the contract event 0x68f3e71400a107502ed4901f333a22781427bea3f4c32385e9c80cb46468ec26.
//
// Solidity: event ContractLaunchedAndEnabled()
func (_BtEth *BtEthFilterer) FilterContractLaunchedAndEnabled(opts *bind.FilterOpts) (*BtEthContractLaunchedAndEnabledIterator, error) {

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "ContractLaunchedAndEnabled")
	if err != nil {
		return nil, err
	}
	return &BtEthContractLaunchedAndEnabledIterator{contract: _BtEth.contract, event: "ContractLaunchedAndEnabled", logs: logs, sub: sub}, nil
}

// WatchContractLaunchedAndEnabled is a free log subscription operation binding the contract event 0x68f3e71400a107502ed4901f333a22781427bea3f4c32385e9c80cb46468ec26.
//
// Solidity: event ContractLaunchedAndEnabled()
func (_BtEth *BtEthFilterer) WatchContractLaunchedAndEnabled(opts *bind.WatchOpts, sink chan<- *BtEthContractLaunchedAndEnabled) (event.Subscription, error) {

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "ContractLaunchedAndEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthContractLaunchedAndEnabled)
				if err := _BtEth.contract.UnpackLog(event, "ContractLaunchedAndEnabled", log); err != nil {
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

// BtEthNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the BtEth contract.
type BtEthNewOwnerIterator struct {
	Event *BtEthNewOwner // Event containing the contract specifics and raw log

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
func (it *BtEthNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthNewOwner)
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
		it.Event = new(BtEthNewOwner)
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
func (it *BtEthNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthNewOwner represents a NewOwner event raised by the BtEth contract.
type BtEthNewOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BtEth *BtEthFilterer) FilterNewOwner(opts *bind.FilterOpts) (*BtEthNewOwnerIterator, error) {

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &BtEthNewOwnerIterator{contract: _BtEth.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BtEth *BtEthFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *BtEthNewOwner) (event.Subscription, error) {

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthNewOwner)
				if err := _BtEth.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// BtEthSalesDisabledIterator is returned from FilterSalesDisabled and is used to iterate over the raw logs and unpacked data for SalesDisabled events raised by the BtEth contract.
type BtEthSalesDisabledIterator struct {
	Event *BtEthSalesDisabled // Event containing the contract specifics and raw log

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
func (it *BtEthSalesDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthSalesDisabled)
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
		it.Event = new(BtEthSalesDisabled)
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
func (it *BtEthSalesDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthSalesDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthSalesDisabled represents a SalesDisabled event raised by the BtEth contract.
type BtEthSalesDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSalesDisabled is a free log retrieval operation binding the contract event 0x3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e47.
//
// Solidity: event SalesDisabled()
func (_BtEth *BtEthFilterer) FilterSalesDisabled(opts *bind.FilterOpts) (*BtEthSalesDisabledIterator, error) {

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "SalesDisabled")
	if err != nil {
		return nil, err
	}
	return &BtEthSalesDisabledIterator{contract: _BtEth.contract, event: "SalesDisabled", logs: logs, sub: sub}, nil
}

// WatchSalesDisabled is a free log subscription operation binding the contract event 0x3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e47.
//
// Solidity: event SalesDisabled()
func (_BtEth *BtEthFilterer) WatchSalesDisabled(opts *bind.WatchOpts, sink chan<- *BtEthSalesDisabled) (event.Subscription, error) {

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "SalesDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthSalesDisabled)
				if err := _BtEth.contract.UnpackLog(event, "SalesDisabled", log); err != nil {
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

// BtEthSalesEnabledIterator is returned from FilterSalesEnabled and is used to iterate over the raw logs and unpacked data for SalesEnabled events raised by the BtEth contract.
type BtEthSalesEnabledIterator struct {
	Event *BtEthSalesEnabled // Event containing the contract specifics and raw log

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
func (it *BtEthSalesEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthSalesEnabled)
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
		it.Event = new(BtEthSalesEnabled)
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
func (it *BtEthSalesEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthSalesEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthSalesEnabled represents a SalesEnabled event raised by the BtEth contract.
type BtEthSalesEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSalesEnabled is a free log retrieval operation binding the contract event 0xec8f62fea11b2d0b268c0f08bc273cef8aeeb2347809eaf5ce4a1a04bb19c991.
//
// Solidity: event SalesEnabled()
func (_BtEth *BtEthFilterer) FilterSalesEnabled(opts *bind.FilterOpts) (*BtEthSalesEnabledIterator, error) {

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "SalesEnabled")
	if err != nil {
		return nil, err
	}
	return &BtEthSalesEnabledIterator{contract: _BtEth.contract, event: "SalesEnabled", logs: logs, sub: sub}, nil
}

// WatchSalesEnabled is a free log subscription operation binding the contract event 0xec8f62fea11b2d0b268c0f08bc273cef8aeeb2347809eaf5ce4a1a04bb19c991.
//
// Solidity: event SalesEnabled()
func (_BtEth *BtEthFilterer) WatchSalesEnabled(opts *bind.WatchOpts, sink chan<- *BtEthSalesEnabled) (event.Subscription, error) {

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "SalesEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthSalesEnabled)
				if err := _BtEth.contract.UnpackLog(event, "SalesEnabled", log); err != nil {
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

// BtEthTokensSoldIterator is returned from FilterTokensSold and is used to iterate over the raw logs and unpacked data for TokensSold events raised by the BtEth contract.
type BtEthTokensSoldIterator struct {
	Event *BtEthTokensSold // Event containing the contract specifics and raw log

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
func (it *BtEthTokensSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtEthTokensSold)
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
		it.Event = new(BtEthTokensSold)
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
func (it *BtEthTokensSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtEthTokensSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtEthTokensSold represents a TokensSold event raised by the BtEth contract.
type BtEthTokensSold struct {
	Purchaser    common.Address
	Contribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensSold is a free log retrieval operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(_purchaser indexed address, _contribution uint256)
func (_BtEth *BtEthFilterer) FilterTokensSold(opts *bind.FilterOpts, _purchaser []common.Address) (*BtEthTokensSoldIterator, error) {

	var _purchaserRule []interface{}
	for _, _purchaserItem := range _purchaser {
		_purchaserRule = append(_purchaserRule, _purchaserItem)
	}

	logs, sub, err := _BtEth.contract.FilterLogs(opts, "TokensSold", _purchaserRule)
	if err != nil {
		return nil, err
	}
	return &BtEthTokensSoldIterator{contract: _BtEth.contract, event: "TokensSold", logs: logs, sub: sub}, nil
}

// WatchTokensSold is a free log subscription operation binding the contract event 0x57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e.
//
// Solidity: event TokensSold(_purchaser indexed address, _contribution uint256)
func (_BtEth *BtEthFilterer) WatchTokensSold(opts *bind.WatchOpts, sink chan<- *BtEthTokensSold, _purchaser []common.Address) (event.Subscription, error) {

	var _purchaserRule []interface{}
	for _, _purchaserItem := range _purchaser {
		_purchaserRule = append(_purchaserRule, _purchaserItem)
	}

	logs, sub, err := _BtEth.contract.WatchLogs(opts, "TokensSold", _purchaserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtEthTokensSold)
				if err := _BtEth.contract.UnpackLog(event, "TokensSold", log); err != nil {
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
