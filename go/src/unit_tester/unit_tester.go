// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unit_tester

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

// UnitTesterABI is the input ABI used to generate the binding from.
const UnitTesterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"testTransferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testLaunchBtEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setErcI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_btAddress\",\"type\":\"address\"}],\"name\":\"setBtI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testSetRentalTokenInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testDisableSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_brAddress\",\"type\":\"address\"}],\"name\":\"setBrI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"testReturnBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"testTransferTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testEnableSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"testTokenApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rentalContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_costPerDay\",\"type\":\"uint256\"}],\"name\":\"testAddBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCreationNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_daysToRent\",\"type\":\"uint256\"},{\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"testRentBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_testName\",\"type\":\"string\"}],\"name\":\"TestPassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_contractName\",\"type\":\"string\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// UnitTesterBin is the compiled bytecode used for deploying new contracts.
const UnitTesterBin = `60606040526001805460008054600160a060020a033316600160a060020a0319918216811790925560a060020a60ff02199092167401000000000000000000000000000000000000000017919091161790556111ce806100606000396000f3006060604052600436106101035763ffffffff60e060020a60003504166305a37e9a8114610108578063175c71661461014457806328155d12146101575780632909b1f5146101765780632fe90e041461019557806355a373d6146101a85780635f9f07a8146101d75780636da15791146101ea578063704b6c02146102095780637ad4c643146102285780638da5cb5b1461023e578063a6f9dae114610251578063b2857e0714610270578063b3fc6cfd14610292578063d2c8eb8d146102a5578063daf6ca30146102c7578063e8f4e9d4146102da578063eb5ad85f146102ed578063eb611bab14610306578063f60058191461032f578063f851a4401461034b575b600080fd5b341561011357600080fd5b610130600160a060020a036004358116906024351660443561035e565b604051901515815260200160405180910390f35b341561014f57600080fd5b6101306104b2565b341561016257600080fd5b610130600160a060020a03600435166105e5565b341561018157600080fd5b610130600160a060020a0360043516610642565b34156101a057600080fd5b61013061069e565b34156101b357600080fd5b6101bb610798565b604051600160a060020a03909116815260200160405180910390f35b34156101e257600080fd5b6101306107a7565b34156101f557600080fd5b610130600160a060020a03600435166108bb565b341561021457600080fd5b610130600160a060020a0360043516610917565b341561023357600080fd5b6101306004356109a0565b341561024957600080fd5b6101bb610ac7565b341561025c57600080fd5b610130600160a060020a0360043516610ad6565b341561027b57600080fd5b610130600160a060020a0360043516602435610b5f565b341561029d57600080fd5b610130610ca5565b34156102b057600080fd5b610130600160a060020a0360043516602435610d9b565b34156102d257600080fd5b6101bb610ec3565b34156102e557600080fd5b6101bb610ed2565b34156102f857600080fd5b610130600435602435610ee1565b341561031157600080fd5b61031961100e565b60405160ff909116815260200160405180910390f35b341561033a57600080fd5b61013060043560243560443561102f565b341561035657600080fd5b6101bb611173565b600554600090600160a060020a03166323b872dd85858560405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156103c457600080fd5b5af115156103d157600080fd5b5050506040518051905015156103e657600080fd5b60008051602061118383398151915260405160208082526018908201527f746f6b656e207472616e736665722066726f6d207465737400000000000000006040808301919091526060909101905180910390a16005548290600160a060020a03166370a082318560405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561048b57600080fd5b5af1151561049857600080fd5b505050604051805190501415156104ab57fe5b9392505050565b600754600254600091600160a060020a039081169163214013ca911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561050c57600080fd5b5af1151561051957600080fd5b50505060405180519050151561052e57600080fd5b6000805160206111838339815191526040516020808252601a908201527f42542d45544820696e697469616c697a6174696f6e20746573740000000000006040808301919091526060909101905180910390a1600754600160a060020a031663e3b172d76040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156105c057600080fd5b5af115156105cd57600080fd5b5050506040518051151560011490506105e257fe5b90565b6000805433600160a060020a0390811691161461060157600080fd5b5060058054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19918216811790925560028054909116909117905560015b919050565b6000805433600160a060020a0390811691161461065e57600080fd5b5060078054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1991821681179092556003805490911690911790556001919050565b600654600254600091600160a060020a03908116916304361569911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156106f857600080fd5b5af1151561070557600080fd5b50505060405180519050151561071a57600080fd5b60008051602061118383398151915260405160208082526021908201527f65726320696e746572666163652073657420666f722062696b652072656e74616040808301919091527f6c0000000000000000000000000000000000000000000000000000000000000060608301526080909101905180910390a16105e2565b600254600160a060020a031681565b600754600090600160a060020a031663613487086040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156107e957600080fd5b5af115156107f657600080fd5b50505060405180519050151561080b57600080fd5b60008051602061118383398151915260405160208082526019908201527f42542d4554482073616c65732064697361626c652074657374000000000000006040808301919091526060909101905180910390a1600754600160a060020a031663238dafe06040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561089d57600080fd5b5af115156108aa57600080fd5b50505060405180511590506105e257fe5b6000805433600160a060020a039081169116146108d757600080fd5b5060068054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1991821681179092556004805490911690911790556001919050565b6000805433600160a060020a0390811691161461093357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b600654600090600160a060020a031663a56a91c58360405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156109eb57600080fd5b5af115156109f857600080fd5b505050604051805190501515610a0d57600080fd5b60008051602061118383398151915260405160208082526010908201527f62696b652072657475726e2074657374000000000000000000000000000000006040808301919091526060909101905180910390a1600654600160a060020a03166375f1473f8360405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610aa857600080fd5b5af11515610ab557600080fd5b50505060405180519050151561063d57fe5b600054600160a060020a031681565b6000805433600160a060020a03908116911614610af257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b600554600090600160a060020a031663a9059cbb848460405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610bb857600080fd5b5af11515610bc557600080fd5b505050604051805190501515610bda57600080fd5b60008051602061118383398151915260405160208082526013908201527f746f6b656e207472616e736665722074657374000000000000000000000000006040808301919091526060909101905180910390a16005548290600160a060020a03166370a082318560405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610c7f57600080fd5b5af11515610c8c57600080fd5b50505060405180519050141515610c9f57fe5b92915050565b600754600090600160a060020a031663c63df5ef6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ce757600080fd5b5af11515610cf457600080fd5b505050604051805190501515610d0957600080fd5b60008051602061118383398151915260405160208082526019908201527f4254482d4554482073616c657320656e61626c652074657374000000000000006040808301919091526060909101905180910390a1600754600160a060020a031663238dafe06040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156105c057600080fd5b600554600090600160a060020a031663095ea7b3848460405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610df457600080fd5b5af11515610e0157600080fd5b505050604051805190501515610e1657600080fd5b60008051602061118383398151915260405160208082526013908201527f746f6b656e20617070726f76616c2074657374000000000000000000000000006040808301919091526060909101905180910390a16005548290600160a060020a031663dd62ed3e308660405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b1515610c7f57600080fd5b600354600160a060020a031681565b600454600160a060020a031681565b600654600090600160a060020a0316639616734a848460405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610f3257600080fd5b5af11515610f3f57600080fd5b505050604051805190501515610f5457600080fd5b6000805160206111838339815191526040516020808252600d908201527f62696b65206164642074657374000000000000000000000000000000000000006040808301919091526060909101905180910390a1600654600160a060020a031663d7d172fe8460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610fef57600080fd5b5af11515610ffc57600080fd5b505050604051805190501515610c9f57fe5b60015474010000000000000000000000000000000000000000900460ff1681565b600654600090600160a060020a031663cb24450d85858560405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401602060405180830381600087803b151561108957600080fd5b5af1151561109657600080fd5b5050506040518051905015156110ab57600080fd5b6000805160206111838339815191526040516020808252600e908201527f62696b652072656e7420746573740000000000000000000000000000000000006040808301919091526060909101905180910390a1600654600160a060020a031663c4540dc1308660405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561115457600080fd5b5af1151561116157600080fd5b5050506040518051905015156104ab57fe5b600154600160a060020a03168156001e95ae6e49e291fc2ef8e08cfd9272a479a92b37348a3ad93e7207ca5c9b0b6ea165627a7a72305820854c47b9f0e29c9ef405bb530edbd1d2e7fccc307ba0674e9c0bd729fe54bdd80029`

// DeployUnitTester deploys a new Ethereum contract, binding an instance of UnitTester to it.
func DeployUnitTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnitTester, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnitTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnitTester{UnitTesterCaller: UnitTesterCaller{contract: contract}, UnitTesterTransactor: UnitTesterTransactor{contract: contract}, UnitTesterFilterer: UnitTesterFilterer{contract: contract}}, nil
}

// UnitTester is an auto generated Go binding around an Ethereum contract.
type UnitTester struct {
	UnitTesterCaller     // Read-only binding to the contract
	UnitTesterTransactor // Write-only binding to the contract
	UnitTesterFilterer   // Log filterer for contract events
}

// UnitTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnitTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitTesterSession struct {
	Contract     *UnitTester       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitTesterCallerSession struct {
	Contract *UnitTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnitTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitTesterTransactorSession struct {
	Contract     *UnitTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnitTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitTesterRaw struct {
	Contract *UnitTester // Generic contract binding to access the raw methods on
}

// UnitTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitTesterCallerRaw struct {
	Contract *UnitTesterCaller // Generic read-only contract binding to access the raw methods on
}

// UnitTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitTesterTransactorRaw struct {
	Contract *UnitTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnitTester creates a new instance of UnitTester, bound to a specific deployed contract.
func NewUnitTester(address common.Address, backend bind.ContractBackend) (*UnitTester, error) {
	contract, err := bindUnitTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnitTester{UnitTesterCaller: UnitTesterCaller{contract: contract}, UnitTesterTransactor: UnitTesterTransactor{contract: contract}, UnitTesterFilterer: UnitTesterFilterer{contract: contract}}, nil
}

// NewUnitTesterCaller creates a new read-only instance of UnitTester, bound to a specific deployed contract.
func NewUnitTesterCaller(address common.Address, caller bind.ContractCaller) (*UnitTesterCaller, error) {
	contract, err := bindUnitTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnitTesterCaller{contract: contract}, nil
}

// NewUnitTesterTransactor creates a new write-only instance of UnitTester, bound to a specific deployed contract.
func NewUnitTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitTesterTransactor, error) {
	contract, err := bindUnitTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnitTesterTransactor{contract: contract}, nil
}

// NewUnitTesterFilterer creates a new log filterer instance of UnitTester, bound to a specific deployed contract.
func NewUnitTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*UnitTesterFilterer, error) {
	contract, err := bindUnitTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnitTesterFilterer{contract: contract}, nil
}

// bindUnitTester binds a generic wrapper to an already deployed contract.
func bindUnitTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitTester *UnitTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitTester.Contract.UnitTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitTester *UnitTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.Contract.UnitTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitTester *UnitTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitTester.Contract.UnitTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitTester *UnitTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitTester *UnitTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitTester *UnitTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitTester.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_UnitTester *UnitTesterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_UnitTester *UnitTesterSession) Admin() (common.Address, error) {
	return _UnitTester.Contract.Admin(&_UnitTester.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_UnitTester *UnitTesterCallerSession) Admin() (common.Address, error) {
	return _UnitTester.Contract.Admin(&_UnitTester.CallOpts)
}

// ContractCreationNonce is a free data retrieval call binding the contract method 0xeb611bab.
//
// Solidity: function contractCreationNonce() constant returns(uint8)
func (_UnitTester *UnitTesterCaller) ContractCreationNonce(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "contractCreationNonce")
	return *ret0, err
}

// ContractCreationNonce is a free data retrieval call binding the contract method 0xeb611bab.
//
// Solidity: function contractCreationNonce() constant returns(uint8)
func (_UnitTester *UnitTesterSession) ContractCreationNonce() (uint8, error) {
	return _UnitTester.Contract.ContractCreationNonce(&_UnitTester.CallOpts)
}

// ContractCreationNonce is a free data retrieval call binding the contract method 0xeb611bab.
//
// Solidity: function contractCreationNonce() constant returns(uint8)
func (_UnitTester *UnitTesterCallerSession) ContractCreationNonce() (uint8, error) {
	return _UnitTester.Contract.ContractCreationNonce(&_UnitTester.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitTester *UnitTesterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitTester *UnitTesterSession) Owner() (common.Address, error) {
	return _UnitTester.Contract.Owner(&_UnitTester.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitTester *UnitTesterCallerSession) Owner() (common.Address, error) {
	return _UnitTester.Contract.Owner(&_UnitTester.CallOpts)
}

// RentalContract is a free data retrieval call binding the contract method 0xe8f4e9d4.
//
// Solidity: function rentalContract() constant returns(address)
func (_UnitTester *UnitTesterCaller) RentalContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "rentalContract")
	return *ret0, err
}

// RentalContract is a free data retrieval call binding the contract method 0xe8f4e9d4.
//
// Solidity: function rentalContract() constant returns(address)
func (_UnitTester *UnitTesterSession) RentalContract() (common.Address, error) {
	return _UnitTester.Contract.RentalContract(&_UnitTester.CallOpts)
}

// RentalContract is a free data retrieval call binding the contract method 0xe8f4e9d4.
//
// Solidity: function rentalContract() constant returns(address)
func (_UnitTester *UnitTesterCallerSession) RentalContract() (common.Address, error) {
	return _UnitTester.Contract.RentalContract(&_UnitTester.CallOpts)
}

// SaleContract is a free data retrieval call binding the contract method 0xdaf6ca30.
//
// Solidity: function saleContract() constant returns(address)
func (_UnitTester *UnitTesterCaller) SaleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "saleContract")
	return *ret0, err
}

// SaleContract is a free data retrieval call binding the contract method 0xdaf6ca30.
//
// Solidity: function saleContract() constant returns(address)
func (_UnitTester *UnitTesterSession) SaleContract() (common.Address, error) {
	return _UnitTester.Contract.SaleContract(&_UnitTester.CallOpts)
}

// SaleContract is a free data retrieval call binding the contract method 0xdaf6ca30.
//
// Solidity: function saleContract() constant returns(address)
func (_UnitTester *UnitTesterCallerSession) SaleContract() (common.Address, error) {
	return _UnitTester.Contract.SaleContract(&_UnitTester.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_UnitTester *UnitTesterCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitTester.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_UnitTester *UnitTesterSession) TokenContract() (common.Address, error) {
	return _UnitTester.Contract.TokenContract(&_UnitTester.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_UnitTester *UnitTesterCallerSession) TokenContract() (common.Address, error) {
	return _UnitTester.Contract.TokenContract(&_UnitTester.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_UnitTester *UnitTesterTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_UnitTester *UnitTesterSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.ChangeOwner(&_UnitTester.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.ChangeOwner(&_UnitTester.TransactOpts, _owner)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_UnitTester *UnitTesterTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_UnitTester *UnitTesterSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetAdmin(&_UnitTester.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetAdmin(&_UnitTester.TransactOpts, _admin)
}

// SetBrI is a paid mutator transaction binding the contract method 0x6da15791.
//
// Solidity: function setBrI(_brAddress address) returns(bool)
func (_UnitTester *UnitTesterTransactor) SetBrI(opts *bind.TransactOpts, _brAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "setBrI", _brAddress)
}

// SetBrI is a paid mutator transaction binding the contract method 0x6da15791.
//
// Solidity: function setBrI(_brAddress address) returns(bool)
func (_UnitTester *UnitTesterSession) SetBrI(_brAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetBrI(&_UnitTester.TransactOpts, _brAddress)
}

// SetBrI is a paid mutator transaction binding the contract method 0x6da15791.
//
// Solidity: function setBrI(_brAddress address) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) SetBrI(_brAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetBrI(&_UnitTester.TransactOpts, _brAddress)
}

// SetBtI is a paid mutator transaction binding the contract method 0x2909b1f5.
//
// Solidity: function setBtI(_btAddress address) returns(bool)
func (_UnitTester *UnitTesterTransactor) SetBtI(opts *bind.TransactOpts, _btAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "setBtI", _btAddress)
}

// SetBtI is a paid mutator transaction binding the contract method 0x2909b1f5.
//
// Solidity: function setBtI(_btAddress address) returns(bool)
func (_UnitTester *UnitTesterSession) SetBtI(_btAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetBtI(&_UnitTester.TransactOpts, _btAddress)
}

// SetBtI is a paid mutator transaction binding the contract method 0x2909b1f5.
//
// Solidity: function setBtI(_btAddress address) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) SetBtI(_btAddress common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetBtI(&_UnitTester.TransactOpts, _btAddress)
}

// SetErcI is a paid mutator transaction binding the contract method 0x28155d12.
//
// Solidity: function setErcI(_token address) returns(bool)
func (_UnitTester *UnitTesterTransactor) SetErcI(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "setErcI", _token)
}

// SetErcI is a paid mutator transaction binding the contract method 0x28155d12.
//
// Solidity: function setErcI(_token address) returns(bool)
func (_UnitTester *UnitTesterSession) SetErcI(_token common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetErcI(&_UnitTester.TransactOpts, _token)
}

// SetErcI is a paid mutator transaction binding the contract method 0x28155d12.
//
// Solidity: function setErcI(_token address) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) SetErcI(_token common.Address) (*types.Transaction, error) {
	return _UnitTester.Contract.SetErcI(&_UnitTester.TransactOpts, _token)
}

// TestAddBike is a paid mutator transaction binding the contract method 0xeb5ad85f.
//
// Solidity: function testAddBike(_id uint256, _costPerDay uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestAddBike(opts *bind.TransactOpts, _id *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testAddBike", _id, _costPerDay)
}

// TestAddBike is a paid mutator transaction binding the contract method 0xeb5ad85f.
//
// Solidity: function testAddBike(_id uint256, _costPerDay uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestAddBike(_id *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestAddBike(&_UnitTester.TransactOpts, _id, _costPerDay)
}

// TestAddBike is a paid mutator transaction binding the contract method 0xeb5ad85f.
//
// Solidity: function testAddBike(_id uint256, _costPerDay uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestAddBike(_id *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestAddBike(&_UnitTester.TransactOpts, _id, _costPerDay)
}

// TestDisableSales is a paid mutator transaction binding the contract method 0x5f9f07a8.
//
// Solidity: function testDisableSales() returns(bool)
func (_UnitTester *UnitTesterTransactor) TestDisableSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testDisableSales")
}

// TestDisableSales is a paid mutator transaction binding the contract method 0x5f9f07a8.
//
// Solidity: function testDisableSales() returns(bool)
func (_UnitTester *UnitTesterSession) TestDisableSales() (*types.Transaction, error) {
	return _UnitTester.Contract.TestDisableSales(&_UnitTester.TransactOpts)
}

// TestDisableSales is a paid mutator transaction binding the contract method 0x5f9f07a8.
//
// Solidity: function testDisableSales() returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestDisableSales() (*types.Transaction, error) {
	return _UnitTester.Contract.TestDisableSales(&_UnitTester.TransactOpts)
}

// TestEnableSales is a paid mutator transaction binding the contract method 0xb3fc6cfd.
//
// Solidity: function testEnableSales() returns(bool)
func (_UnitTester *UnitTesterTransactor) TestEnableSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testEnableSales")
}

// TestEnableSales is a paid mutator transaction binding the contract method 0xb3fc6cfd.
//
// Solidity: function testEnableSales() returns(bool)
func (_UnitTester *UnitTesterSession) TestEnableSales() (*types.Transaction, error) {
	return _UnitTester.Contract.TestEnableSales(&_UnitTester.TransactOpts)
}

// TestEnableSales is a paid mutator transaction binding the contract method 0xb3fc6cfd.
//
// Solidity: function testEnableSales() returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestEnableSales() (*types.Transaction, error) {
	return _UnitTester.Contract.TestEnableSales(&_UnitTester.TransactOpts)
}

// TestLaunchBtEth is a paid mutator transaction binding the contract method 0x175c7166.
//
// Solidity: function testLaunchBtEth() returns(bool)
func (_UnitTester *UnitTesterTransactor) TestLaunchBtEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testLaunchBtEth")
}

// TestLaunchBtEth is a paid mutator transaction binding the contract method 0x175c7166.
//
// Solidity: function testLaunchBtEth() returns(bool)
func (_UnitTester *UnitTesterSession) TestLaunchBtEth() (*types.Transaction, error) {
	return _UnitTester.Contract.TestLaunchBtEth(&_UnitTester.TransactOpts)
}

// TestLaunchBtEth is a paid mutator transaction binding the contract method 0x175c7166.
//
// Solidity: function testLaunchBtEth() returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestLaunchBtEth() (*types.Transaction, error) {
	return _UnitTester.Contract.TestLaunchBtEth(&_UnitTester.TransactOpts)
}

// TestRentBike is a paid mutator transaction binding the contract method 0xf6005819.
//
// Solidity: function testRentBike(_id uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestRentBike(opts *bind.TransactOpts, _id *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testRentBike", _id, _daysToRent, _deposit)
}

// TestRentBike is a paid mutator transaction binding the contract method 0xf6005819.
//
// Solidity: function testRentBike(_id uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestRentBike(_id *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestRentBike(&_UnitTester.TransactOpts, _id, _daysToRent, _deposit)
}

// TestRentBike is a paid mutator transaction binding the contract method 0xf6005819.
//
// Solidity: function testRentBike(_id uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestRentBike(_id *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestRentBike(&_UnitTester.TransactOpts, _id, _daysToRent, _deposit)
}

// TestReturnBike is a paid mutator transaction binding the contract method 0x7ad4c643.
//
// Solidity: function testReturnBike(_bikeId uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestReturnBike(opts *bind.TransactOpts, _bikeId *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testReturnBike", _bikeId)
}

// TestReturnBike is a paid mutator transaction binding the contract method 0x7ad4c643.
//
// Solidity: function testReturnBike(_bikeId uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestReturnBike(&_UnitTester.TransactOpts, _bikeId)
}

// TestReturnBike is a paid mutator transaction binding the contract method 0x7ad4c643.
//
// Solidity: function testReturnBike(_bikeId uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestReturnBike(&_UnitTester.TransactOpts, _bikeId)
}

// TestSetRentalTokenInterface is a paid mutator transaction binding the contract method 0x2fe90e04.
//
// Solidity: function testSetRentalTokenInterface() returns(bool)
func (_UnitTester *UnitTesterTransactor) TestSetRentalTokenInterface(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testSetRentalTokenInterface")
}

// TestSetRentalTokenInterface is a paid mutator transaction binding the contract method 0x2fe90e04.
//
// Solidity: function testSetRentalTokenInterface() returns(bool)
func (_UnitTester *UnitTesterSession) TestSetRentalTokenInterface() (*types.Transaction, error) {
	return _UnitTester.Contract.TestSetRentalTokenInterface(&_UnitTester.TransactOpts)
}

// TestSetRentalTokenInterface is a paid mutator transaction binding the contract method 0x2fe90e04.
//
// Solidity: function testSetRentalTokenInterface() returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestSetRentalTokenInterface() (*types.Transaction, error) {
	return _UnitTester.Contract.TestSetRentalTokenInterface(&_UnitTester.TransactOpts)
}

// TestTokenApproval is a paid mutator transaction binding the contract method 0xd2c8eb8d.
//
// Solidity: function testTokenApproval(_spender address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestTokenApproval(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testTokenApproval", _spender, _amount)
}

// TestTokenApproval is a paid mutator transaction binding the contract method 0xd2c8eb8d.
//
// Solidity: function testTokenApproval(_spender address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestTokenApproval(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTokenApproval(&_UnitTester.TransactOpts, _spender, _amount)
}

// TestTokenApproval is a paid mutator transaction binding the contract method 0xd2c8eb8d.
//
// Solidity: function testTokenApproval(_spender address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestTokenApproval(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTokenApproval(&_UnitTester.TransactOpts, _spender, _amount)
}

// TestTransferFrom is a paid mutator transaction binding the contract method 0x05a37e9a.
//
// Solidity: function testTransferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestTransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testTransferFrom", _owner, _recipient, _amount)
}

// TestTransferFrom is a paid mutator transaction binding the contract method 0x05a37e9a.
//
// Solidity: function testTransferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestTransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTransferFrom(&_UnitTester.TransactOpts, _owner, _recipient, _amount)
}

// TestTransferFrom is a paid mutator transaction binding the contract method 0x05a37e9a.
//
// Solidity: function testTransferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestTransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTransferFrom(&_UnitTester.TransactOpts, _owner, _recipient, _amount)
}

// TestTransferTokens is a paid mutator transaction binding the contract method 0xb2857e07.
//
// Solidity: function testTransferTokens(_recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactor) TestTransferTokens(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.contract.Transact(opts, "testTransferTokens", _recipient, _amount)
}

// TestTransferTokens is a paid mutator transaction binding the contract method 0xb2857e07.
//
// Solidity: function testTransferTokens(_recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterSession) TestTransferTokens(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTransferTokens(&_UnitTester.TransactOpts, _recipient, _amount)
}

// TestTransferTokens is a paid mutator transaction binding the contract method 0xb2857e07.
//
// Solidity: function testTransferTokens(_recipient address, _amount uint256) returns(bool)
func (_UnitTester *UnitTesterTransactorSession) TestTransferTokens(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UnitTester.Contract.TestTransferTokens(&_UnitTester.TransactOpts, _recipient, _amount)
}

// UnitTesterAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the UnitTester contract.
type UnitTesterAdminSetIterator struct {
	Event *UnitTesterAdminSet // Event containing the contract specifics and raw log

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
func (it *UnitTesterAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitTesterAdminSet)
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
		it.Event = new(UnitTesterAdminSet)
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
func (it *UnitTesterAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitTesterAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitTesterAdminSet represents a AdminSet event raised by the UnitTester contract.
type UnitTesterAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_UnitTester *UnitTesterFilterer) FilterAdminSet(opts *bind.FilterOpts) (*UnitTesterAdminSetIterator, error) {

	logs, sub, err := _UnitTester.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &UnitTesterAdminSetIterator{contract: _UnitTester.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_UnitTester *UnitTesterFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *UnitTesterAdminSet) (event.Subscription, error) {

	logs, sub, err := _UnitTester.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitTesterAdminSet)
				if err := _UnitTester.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// UnitTesterContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the UnitTester contract.
type UnitTesterContractCreatedIterator struct {
	Event *UnitTesterContractCreated // Event containing the contract specifics and raw log

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
func (it *UnitTesterContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitTesterContractCreated)
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
		it.Event = new(UnitTesterContractCreated)
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
func (it *UnitTesterContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitTesterContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitTesterContractCreated represents a ContractCreated event raised by the UnitTester contract.
type UnitTesterContractCreated struct {
	ContractAddress common.Address
	ContractName    string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x5748f17320a7bfc4dad968e541c61b9fa9923765f9efdfa3dffc76763e02d196.
//
// Solidity: event ContractCreated(_contractAddress address, _contractName string)
func (_UnitTester *UnitTesterFilterer) FilterContractCreated(opts *bind.FilterOpts) (*UnitTesterContractCreatedIterator, error) {

	logs, sub, err := _UnitTester.contract.FilterLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return &UnitTesterContractCreatedIterator{contract: _UnitTester.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x5748f17320a7bfc4dad968e541c61b9fa9923765f9efdfa3dffc76763e02d196.
//
// Solidity: event ContractCreated(_contractAddress address, _contractName string)
func (_UnitTester *UnitTesterFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *UnitTesterContractCreated) (event.Subscription, error) {

	logs, sub, err := _UnitTester.contract.WatchLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitTesterContractCreated)
				if err := _UnitTester.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// UnitTesterNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the UnitTester contract.
type UnitTesterNewOwnerIterator struct {
	Event *UnitTesterNewOwner // Event containing the contract specifics and raw log

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
func (it *UnitTesterNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitTesterNewOwner)
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
		it.Event = new(UnitTesterNewOwner)
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
func (it *UnitTesterNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitTesterNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitTesterNewOwner represents a NewOwner event raised by the UnitTester contract.
type UnitTesterNewOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_UnitTester *UnitTesterFilterer) FilterNewOwner(opts *bind.FilterOpts) (*UnitTesterNewOwnerIterator, error) {

	logs, sub, err := _UnitTester.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &UnitTesterNewOwnerIterator{contract: _UnitTester.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_UnitTester *UnitTesterFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *UnitTesterNewOwner) (event.Subscription, error) {

	logs, sub, err := _UnitTester.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitTesterNewOwner)
				if err := _UnitTester.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// UnitTesterTestPassedIterator is returned from FilterTestPassed and is used to iterate over the raw logs and unpacked data for TestPassed events raised by the UnitTester contract.
type UnitTesterTestPassedIterator struct {
	Event *UnitTesterTestPassed // Event containing the contract specifics and raw log

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
func (it *UnitTesterTestPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitTesterTestPassed)
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
		it.Event = new(UnitTesterTestPassed)
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
func (it *UnitTesterTestPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitTesterTestPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitTesterTestPassed represents a TestPassed event raised by the UnitTester contract.
type UnitTesterTestPassed struct {
	TestName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTestPassed is a free log retrieval operation binding the contract event 0x1e95ae6e49e291fc2ef8e08cfd9272a479a92b37348a3ad93e7207ca5c9b0b6e.
//
// Solidity: event TestPassed(_testName string)
func (_UnitTester *UnitTesterFilterer) FilterTestPassed(opts *bind.FilterOpts) (*UnitTesterTestPassedIterator, error) {

	logs, sub, err := _UnitTester.contract.FilterLogs(opts, "TestPassed")
	if err != nil {
		return nil, err
	}
	return &UnitTesterTestPassedIterator{contract: _UnitTester.contract, event: "TestPassed", logs: logs, sub: sub}, nil
}

// WatchTestPassed is a free log subscription operation binding the contract event 0x1e95ae6e49e291fc2ef8e08cfd9272a479a92b37348a3ad93e7207ca5c9b0b6e.
//
// Solidity: event TestPassed(_testName string)
func (_UnitTester *UnitTesterFilterer) WatchTestPassed(opts *bind.WatchOpts, sink chan<- *UnitTesterTestPassed) (event.Subscription, error) {

	logs, sub, err := _UnitTester.contract.WatchLogs(opts, "TestPassed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitTesterTestPassed)
				if err := _UnitTester.contract.UnpackLog(event, "TestPassed", log); err != nil {
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
