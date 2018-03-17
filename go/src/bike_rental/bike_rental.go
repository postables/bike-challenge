// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bike_rental

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

// BikeRentalABI is the input ABI used to generate the binding from.
const BikeRentalABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_bikeTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setErcInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bikes\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"costPerDay\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rentals\",\"outputs\":[{\"name\":\"renter\",\"type\":\"address\"},{\"name\":\"bikeId\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"daysRented\",\"type\":\"uint256\"},{\"name\":\"returnDate\",\"type\":\"uint256\"},{\"name\":\"rented\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"checkIfBikeIsReturned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"},{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"bikeAwol\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_costPerDay\",\"type\":\"uint256\"}],\"name\":\"addBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"returnBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"},{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"checkIfRentingBikeId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_daysToRent\",\"type\":\"uint256\"},{\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"rentBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"bikeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bikeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_costPerDay\",\"type\":\"uint256\"}],\"name\":\"BikeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_daysRented\",\"type\":\"uint256\"}],\"name\":\"BikeRented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_late\",\"type\":\"bool\"}],\"name\":\"BikeReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"BikeAwol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ErcInterfaceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// BikeRental is an auto generated Go binding around an Ethereum contract.
type BikeRental struct {
	BikeRentalCaller     // Read-only binding to the contract
	BikeRentalTransactor // Write-only binding to the contract
	BikeRentalFilterer   // Log filterer for contract events
}

// BikeRentalCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeRentalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeRentalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeRentalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeRentalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeRentalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeRentalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeRentalSession struct {
	Contract     *BikeRental       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeRentalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeRentalCallerSession struct {
	Contract *BikeRentalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BikeRentalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeRentalTransactorSession struct {
	Contract     *BikeRentalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BikeRentalRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeRentalRaw struct {
	Contract *BikeRental // Generic contract binding to access the raw methods on
}

// BikeRentalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeRentalCallerRaw struct {
	Contract *BikeRentalCaller // Generic read-only contract binding to access the raw methods on
}

// BikeRentalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeRentalTransactorRaw struct {
	Contract *BikeRentalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBikeRental creates a new instance of BikeRental, bound to a specific deployed contract.
func NewBikeRental(address common.Address, backend bind.ContractBackend) (*BikeRental, error) {
	contract, err := bindBikeRental(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BikeRental{BikeRentalCaller: BikeRentalCaller{contract: contract}, BikeRentalTransactor: BikeRentalTransactor{contract: contract}, BikeRentalFilterer: BikeRentalFilterer{contract: contract}}, nil
}

// NewBikeRentalCaller creates a new read-only instance of BikeRental, bound to a specific deployed contract.
func NewBikeRentalCaller(address common.Address, caller bind.ContractCaller) (*BikeRentalCaller, error) {
	contract, err := bindBikeRental(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeRentalCaller{contract: contract}, nil
}

// NewBikeRentalTransactor creates a new write-only instance of BikeRental, bound to a specific deployed contract.
func NewBikeRentalTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeRentalTransactor, error) {
	contract, err := bindBikeRental(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeRentalTransactor{contract: contract}, nil
}

// NewBikeRentalFilterer creates a new log filterer instance of BikeRental, bound to a specific deployed contract.
func NewBikeRentalFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeRentalFilterer, error) {
	contract, err := bindBikeRental(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeRentalFilterer{contract: contract}, nil
}

// bindBikeRental binds a generic wrapper to an already deployed contract.
func bindBikeRental(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeRentalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeRental *BikeRentalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeRental.Contract.BikeRentalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeRental *BikeRentalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeRental.Contract.BikeRentalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeRental *BikeRentalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeRental.Contract.BikeRentalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeRental *BikeRentalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeRental.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeRental *BikeRentalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeRental.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeRental *BikeRentalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeRental.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeRental *BikeRentalCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeRental *BikeRentalSession) Admin() (common.Address, error) {
	return _BikeRental.Contract.Admin(&_BikeRental.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_BikeRental *BikeRentalCallerSession) Admin() (common.Address, error) {
	return _BikeRental.Contract.Admin(&_BikeRental.CallOpts)
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_BikeRental *BikeRentalCaller) BikeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "bikeCount")
	return *ret0, err
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_BikeRental *BikeRentalSession) BikeCount() (*big.Int, error) {
	return _BikeRental.Contract.BikeCount(&_BikeRental.CallOpts)
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_BikeRental *BikeRentalCallerSession) BikeCount() (*big.Int, error) {
	return _BikeRental.Contract.BikeCount(&_BikeRental.CallOpts)
}

// BikeExists is a free data retrieval call binding the contract method 0xd7d172fe.
//
// Solidity: function bikeExists(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalCaller) BikeExists(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "bikeExists", _id)
	return *ret0, err
}

// BikeExists is a free data retrieval call binding the contract method 0xd7d172fe.
//
// Solidity: function bikeExists(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalSession) BikeExists(_id *big.Int) (bool, error) {
	return _BikeRental.Contract.BikeExists(&_BikeRental.CallOpts, _id)
}

// BikeExists is a free data retrieval call binding the contract method 0xd7d172fe.
//
// Solidity: function bikeExists(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalCallerSession) BikeExists(_id *big.Int) (bool, error) {
	return _BikeRental.Contract.BikeExists(&_BikeRental.CallOpts, _id)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_BikeRental *BikeRentalCaller) Bikes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id         *big.Int
	CostPerDay *big.Int
	State      uint8
	Exists     bool
}, error) {
	ret := new(struct {
		Id         *big.Int
		CostPerDay *big.Int
		State      uint8
		Exists     bool
	})
	out := ret
	err := _BikeRental.contract.Call(opts, out, "bikes", arg0)
	return *ret, err
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_BikeRental *BikeRentalSession) Bikes(arg0 *big.Int) (struct {
	Id         *big.Int
	CostPerDay *big.Int
	State      uint8
	Exists     bool
}, error) {
	return _BikeRental.Contract.Bikes(&_BikeRental.CallOpts, arg0)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_BikeRental *BikeRentalCallerSession) Bikes(arg0 *big.Int) (struct {
	Id         *big.Int
	CostPerDay *big.Int
	State      uint8
	Exists     bool
}, error) {
	return _BikeRental.Contract.Bikes(&_BikeRental.CallOpts, arg0)
}

// CheckIfBikeIsReturned is a free data retrieval call binding the contract method 0x75f1473f.
//
// Solidity: function checkIfBikeIsReturned(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalCaller) CheckIfBikeIsReturned(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "checkIfBikeIsReturned", _id)
	return *ret0, err
}

// CheckIfBikeIsReturned is a free data retrieval call binding the contract method 0x75f1473f.
//
// Solidity: function checkIfBikeIsReturned(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalSession) CheckIfBikeIsReturned(_id *big.Int) (bool, error) {
	return _BikeRental.Contract.CheckIfBikeIsReturned(&_BikeRental.CallOpts, _id)
}

// CheckIfBikeIsReturned is a free data retrieval call binding the contract method 0x75f1473f.
//
// Solidity: function checkIfBikeIsReturned(_id uint256) constant returns(bool)
func (_BikeRental *BikeRentalCallerSession) CheckIfBikeIsReturned(_id *big.Int) (bool, error) {
	return _BikeRental.Contract.CheckIfBikeIsReturned(&_BikeRental.CallOpts, _id)
}

// CheckIfRentingBikeId is a free data retrieval call binding the contract method 0xc4540dc1.
//
// Solidity: function checkIfRentingBikeId(_renter address, _bikeId uint256) constant returns(bool)
func (_BikeRental *BikeRentalCaller) CheckIfRentingBikeId(opts *bind.CallOpts, _renter common.Address, _bikeId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "checkIfRentingBikeId", _renter, _bikeId)
	return *ret0, err
}

// CheckIfRentingBikeId is a free data retrieval call binding the contract method 0xc4540dc1.
//
// Solidity: function checkIfRentingBikeId(_renter address, _bikeId uint256) constant returns(bool)
func (_BikeRental *BikeRentalSession) CheckIfRentingBikeId(_renter common.Address, _bikeId *big.Int) (bool, error) {
	return _BikeRental.Contract.CheckIfRentingBikeId(&_BikeRental.CallOpts, _renter, _bikeId)
}

// CheckIfRentingBikeId is a free data retrieval call binding the contract method 0xc4540dc1.
//
// Solidity: function checkIfRentingBikeId(_renter address, _bikeId uint256) constant returns(bool)
func (_BikeRental *BikeRentalCallerSession) CheckIfRentingBikeId(_renter common.Address, _bikeId *big.Int) (bool, error) {
	return _BikeRental.Contract.CheckIfRentingBikeId(&_BikeRental.CallOpts, _renter, _bikeId)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_BikeRental *BikeRentalCaller) Dev(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "dev")
	return *ret0, err
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_BikeRental *BikeRentalSession) Dev() (bool, error) {
	return _BikeRental.Contract.Dev(&_BikeRental.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_BikeRental *BikeRentalCallerSession) Dev() (bool, error) {
	return _BikeRental.Contract.Dev(&_BikeRental.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeRental *BikeRentalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeRental.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeRental *BikeRentalSession) Owner() (common.Address, error) {
	return _BikeRental.Contract.Owner(&_BikeRental.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BikeRental *BikeRentalCallerSession) Owner() (common.Address, error) {
	return _BikeRental.Contract.Owner(&_BikeRental.CallOpts)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_BikeRental *BikeRentalCaller) Rentals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Renter     common.Address
	BikeId     *big.Int
	Deposit    *big.Int
	Cost       *big.Int
	DaysRented *big.Int
	ReturnDate *big.Int
	Rented     bool
}, error) {
	ret := new(struct {
		Renter     common.Address
		BikeId     *big.Int
		Deposit    *big.Int
		Cost       *big.Int
		DaysRented *big.Int
		ReturnDate *big.Int
		Rented     bool
	})
	out := ret
	err := _BikeRental.contract.Call(opts, out, "rentals", arg0)
	return *ret, err
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_BikeRental *BikeRentalSession) Rentals(arg0 common.Address) (struct {
	Renter     common.Address
	BikeId     *big.Int
	Deposit    *big.Int
	Cost       *big.Int
	DaysRented *big.Int
	ReturnDate *big.Int
	Rented     bool
}, error) {
	return _BikeRental.Contract.Rentals(&_BikeRental.CallOpts, arg0)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_BikeRental *BikeRentalCallerSession) Rentals(arg0 common.Address) (struct {
	Renter     common.Address
	BikeId     *big.Int
	Deposit    *big.Int
	Cost       *big.Int
	DaysRented *big.Int
	ReturnDate *big.Int
	Rented     bool
}, error) {
	return _BikeRental.Contract.Rentals(&_BikeRental.CallOpts, arg0)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_BikeRental *BikeRentalTransactor) AddBike(opts *bind.TransactOpts, _bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "addBike", _bikeId, _costPerDay)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_BikeRental *BikeRentalSession) AddBike(_bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.AddBike(&_BikeRental.TransactOpts, _bikeId, _costPerDay)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) AddBike(_bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.AddBike(&_BikeRental.TransactOpts, _bikeId, _costPerDay)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalTransactor) BikeAwol(opts *bind.TransactOpts, _renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "bikeAwol", _renter, _bikeId)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalSession) BikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.BikeAwol(&_BikeRental.TransactOpts, _renter, _bikeId)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) BikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.BikeAwol(&_BikeRental.TransactOpts, _renter, _bikeId)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeRental *BikeRentalTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeRental *BikeRentalSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.ChangeOwner(&_BikeRental.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.ChangeOwner(&_BikeRental.TransactOpts, _owner)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_BikeRental *BikeRentalTransactor) RentBike(opts *bind.TransactOpts, _bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "rentBike", _bikeId, _daysToRent, _deposit)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_BikeRental *BikeRentalSession) RentBike(_bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.RentBike(&_BikeRental.TransactOpts, _bikeId, _daysToRent, _deposit)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) RentBike(_bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.RentBike(&_BikeRental.TransactOpts, _bikeId, _daysToRent, _deposit)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalTransactor) ReturnBike(opts *bind.TransactOpts, _bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "returnBike", _bikeId)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalSession) ReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.ReturnBike(&_BikeRental.TransactOpts, _bikeId)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) ReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _BikeRental.Contract.ReturnBike(&_BikeRental.TransactOpts, _bikeId)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeRental *BikeRentalTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeRental *BikeRentalSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.SetAdmin(&_BikeRental.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.SetAdmin(&_BikeRental.TransactOpts, _admin)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_BikeRental *BikeRentalTransactor) SetErcInterface(opts *bind.TransactOpts, _bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _BikeRental.contract.Transact(opts, "setErcInterface", _bikeTokenContractAddress)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_BikeRental *BikeRentalSession) SetErcInterface(_bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.SetErcInterface(&_BikeRental.TransactOpts, _bikeTokenContractAddress)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_BikeRental *BikeRentalTransactorSession) SetErcInterface(_bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _BikeRental.Contract.SetErcInterface(&_BikeRental.TransactOpts, _bikeTokenContractAddress)
}

// BikeRentalAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the BikeRental contract.
type BikeRentalAdminSetIterator struct {
	Event *BikeRentalAdminSet // Event containing the contract specifics and raw log

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
func (it *BikeRentalAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalAdminSet)
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
		it.Event = new(BikeRentalAdminSet)
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
func (it *BikeRentalAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalAdminSet represents a AdminSet event raised by the BikeRental contract.
type BikeRentalAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BikeRental *BikeRentalFilterer) FilterAdminSet(opts *bind.FilterOpts) (*BikeRentalAdminSetIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &BikeRentalAdminSetIterator{contract: _BikeRental.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_BikeRental *BikeRentalFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *BikeRentalAdminSet) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalAdminSet)
				if err := _BikeRental.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// BikeRentalBikeAddedIterator is returned from FilterBikeAdded and is used to iterate over the raw logs and unpacked data for BikeAdded events raised by the BikeRental contract.
type BikeRentalBikeAddedIterator struct {
	Event *BikeRentalBikeAdded // Event containing the contract specifics and raw log

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
func (it *BikeRentalBikeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalBikeAdded)
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
		it.Event = new(BikeRentalBikeAdded)
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
func (it *BikeRentalBikeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalBikeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalBikeAdded represents a BikeAdded event raised by the BikeRental contract.
type BikeRentalBikeAdded struct {
	Id         *big.Int
	CostPerDay *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBikeAdded is a free log retrieval operation binding the contract event 0xea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416.
//
// Solidity: event BikeAdded(_id uint256, _costPerDay uint256)
func (_BikeRental *BikeRentalFilterer) FilterBikeAdded(opts *bind.FilterOpts) (*BikeRentalBikeAddedIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "BikeAdded")
	if err != nil {
		return nil, err
	}
	return &BikeRentalBikeAddedIterator{contract: _BikeRental.contract, event: "BikeAdded", logs: logs, sub: sub}, nil
}

// WatchBikeAdded is a free log subscription operation binding the contract event 0xea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416.
//
// Solidity: event BikeAdded(_id uint256, _costPerDay uint256)
func (_BikeRental *BikeRentalFilterer) WatchBikeAdded(opts *bind.WatchOpts, sink chan<- *BikeRentalBikeAdded) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "BikeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalBikeAdded)
				if err := _BikeRental.contract.UnpackLog(event, "BikeAdded", log); err != nil {
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

// BikeRentalBikeAwolIterator is returned from FilterBikeAwol and is used to iterate over the raw logs and unpacked data for BikeAwol events raised by the BikeRental contract.
type BikeRentalBikeAwolIterator struct {
	Event *BikeRentalBikeAwol // Event containing the contract specifics and raw log

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
func (it *BikeRentalBikeAwolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalBikeAwol)
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
		it.Event = new(BikeRentalBikeAwol)
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
func (it *BikeRentalBikeAwolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalBikeAwolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalBikeAwol represents a BikeAwol event raised by the BikeRental contract.
type BikeRentalBikeAwol struct {
	Renter common.Address
	BikeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeAwol is a free log retrieval operation binding the contract event 0x15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25.
//
// Solidity: event BikeAwol(_renter address, _bikeId uint256)
func (_BikeRental *BikeRentalFilterer) FilterBikeAwol(opts *bind.FilterOpts) (*BikeRentalBikeAwolIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "BikeAwol")
	if err != nil {
		return nil, err
	}
	return &BikeRentalBikeAwolIterator{contract: _BikeRental.contract, event: "BikeAwol", logs: logs, sub: sub}, nil
}

// WatchBikeAwol is a free log subscription operation binding the contract event 0x15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25.
//
// Solidity: event BikeAwol(_renter address, _bikeId uint256)
func (_BikeRental *BikeRentalFilterer) WatchBikeAwol(opts *bind.WatchOpts, sink chan<- *BikeRentalBikeAwol) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "BikeAwol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalBikeAwol)
				if err := _BikeRental.contract.UnpackLog(event, "BikeAwol", log); err != nil {
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

// BikeRentalBikeRentedIterator is returned from FilterBikeRented and is used to iterate over the raw logs and unpacked data for BikeRented events raised by the BikeRental contract.
type BikeRentalBikeRentedIterator struct {
	Event *BikeRentalBikeRented // Event containing the contract specifics and raw log

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
func (it *BikeRentalBikeRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalBikeRented)
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
		it.Event = new(BikeRentalBikeRented)
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
func (it *BikeRentalBikeRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalBikeRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalBikeRented represents a BikeRented event raised by the BikeRental contract.
type BikeRentalBikeRented struct {
	Renter     common.Address
	Id         *big.Int
	DaysRented *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBikeRented is a free log retrieval operation binding the contract event 0xabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd2408.
//
// Solidity: event BikeRented(_renter address, _id uint256, _daysRented uint256)
func (_BikeRental *BikeRentalFilterer) FilterBikeRented(opts *bind.FilterOpts) (*BikeRentalBikeRentedIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "BikeRented")
	if err != nil {
		return nil, err
	}
	return &BikeRentalBikeRentedIterator{contract: _BikeRental.contract, event: "BikeRented", logs: logs, sub: sub}, nil
}

// WatchBikeRented is a free log subscription operation binding the contract event 0xabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd2408.
//
// Solidity: event BikeRented(_renter address, _id uint256, _daysRented uint256)
func (_BikeRental *BikeRentalFilterer) WatchBikeRented(opts *bind.WatchOpts, sink chan<- *BikeRentalBikeRented) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "BikeRented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalBikeRented)
				if err := _BikeRental.contract.UnpackLog(event, "BikeRented", log); err != nil {
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

// BikeRentalBikeReturnedIterator is returned from FilterBikeReturned and is used to iterate over the raw logs and unpacked data for BikeReturned events raised by the BikeRental contract.
type BikeRentalBikeReturnedIterator struct {
	Event *BikeRentalBikeReturned // Event containing the contract specifics and raw log

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
func (it *BikeRentalBikeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalBikeReturned)
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
		it.Event = new(BikeRentalBikeReturned)
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
func (it *BikeRentalBikeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalBikeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalBikeReturned represents a BikeReturned event raised by the BikeRental contract.
type BikeRentalBikeReturned struct {
	Renter common.Address
	BikeId *big.Int
	Late   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeReturned is a free log retrieval operation binding the contract event 0xd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f.
//
// Solidity: event BikeReturned(_renter address, _bikeId uint256, _late bool)
func (_BikeRental *BikeRentalFilterer) FilterBikeReturned(opts *bind.FilterOpts) (*BikeRentalBikeReturnedIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "BikeReturned")
	if err != nil {
		return nil, err
	}
	return &BikeRentalBikeReturnedIterator{contract: _BikeRental.contract, event: "BikeReturned", logs: logs, sub: sub}, nil
}

// WatchBikeReturned is a free log subscription operation binding the contract event 0xd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f.
//
// Solidity: event BikeReturned(_renter address, _bikeId uint256, _late bool)
func (_BikeRental *BikeRentalFilterer) WatchBikeReturned(opts *bind.WatchOpts, sink chan<- *BikeRentalBikeReturned) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "BikeReturned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalBikeReturned)
				if err := _BikeRental.contract.UnpackLog(event, "BikeReturned", log); err != nil {
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

// BikeRentalErcInterfaceSetIterator is returned from FilterErcInterfaceSet and is used to iterate over the raw logs and unpacked data for ErcInterfaceSet events raised by the BikeRental contract.
type BikeRentalErcInterfaceSetIterator struct {
	Event *BikeRentalErcInterfaceSet // Event containing the contract specifics and raw log

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
func (it *BikeRentalErcInterfaceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalErcInterfaceSet)
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
		it.Event = new(BikeRentalErcInterfaceSet)
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
func (it *BikeRentalErcInterfaceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalErcInterfaceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalErcInterfaceSet represents a ErcInterfaceSet event raised by the BikeRental contract.
type BikeRentalErcInterfaceSet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterErcInterfaceSet is a free log retrieval operation binding the contract event 0xb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a25.
//
// Solidity: event ErcInterfaceSet()
func (_BikeRental *BikeRentalFilterer) FilterErcInterfaceSet(opts *bind.FilterOpts) (*BikeRentalErcInterfaceSetIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "ErcInterfaceSet")
	if err != nil {
		return nil, err
	}
	return &BikeRentalErcInterfaceSetIterator{contract: _BikeRental.contract, event: "ErcInterfaceSet", logs: logs, sub: sub}, nil
}

// WatchErcInterfaceSet is a free log subscription operation binding the contract event 0xb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a25.
//
// Solidity: event ErcInterfaceSet()
func (_BikeRental *BikeRentalFilterer) WatchErcInterfaceSet(opts *bind.WatchOpts, sink chan<- *BikeRentalErcInterfaceSet) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "ErcInterfaceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalErcInterfaceSet)
				if err := _BikeRental.contract.UnpackLog(event, "ErcInterfaceSet", log); err != nil {
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

// BikeRentalNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the BikeRental contract.
type BikeRentalNewOwnerIterator struct {
	Event *BikeRentalNewOwner // Event containing the contract specifics and raw log

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
func (it *BikeRentalNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRentalNewOwner)
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
		it.Event = new(BikeRentalNewOwner)
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
func (it *BikeRentalNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentalNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRentalNewOwner represents a NewOwner event raised by the BikeRental contract.
type BikeRentalNewOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BikeRental *BikeRentalFilterer) FilterNewOwner(opts *bind.FilterOpts) (*BikeRentalNewOwnerIterator, error) {

	logs, sub, err := _BikeRental.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &BikeRentalNewOwnerIterator{contract: _BikeRental.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_BikeRental *BikeRentalFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *BikeRentalNewOwner) (event.Subscription, error) {

	logs, sub, err := _BikeRental.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRentalNewOwner)
				if err := _BikeRental.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
