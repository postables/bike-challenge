// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
    "github.com/ethereum/go-ethereum/core/types" // this solve the undefined types issue
)

func main() {
	_, err := ethclient.Dial("/hoaame/solidity/Documents/Code/Postables/Sidechain-Bridge/poa/node1/geth.ipc")
	if err != nil {
		fmt.Printf("error detected")
	}
}

// RentalABI is the input ABI used to generate the binding from.
const RentalABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_bikeTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setErcInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"},{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"forceBikeAwol\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bikes\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"costPerDay\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rentals\",\"outputs\":[{\"name\":\"renter\",\"type\":\"address\"},{\"name\":\"bikeId\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"daysRented\",\"type\":\"uint256\"},{\"name\":\"returnDate\",\"type\":\"uint256\"},{\"name\":\"rented\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"},{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"bikeAwol\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_costPerDay\",\"type\":\"uint256\"}],\"name\":\"addBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"},{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"forceLateReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"returnBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_daysToRent\",\"type\":\"uint256\"},{\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"rentBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bikeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_costPerDay\",\"type\":\"uint256\"}],\"name\":\"BikeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_daysRented\",\"type\":\"uint256\"}],\"name\":\"BikeRented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_late\",\"type\":\"bool\"}],\"name\":\"BikeReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"BikeAwol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ErcInterfaceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"}]"

// Rental is an auto generated Go binding around an Ethereum contract.
type Rental struct {
	RentalCaller     // Read-only binding to the contract
	RentalTransactor // Write-only binding to the contract
	RentalFilterer   // Log filterer for contract events
}

// RentalCaller is an auto generated read-only Go binding around an Ethereum contract.
type RentalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RentalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RentalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RentalSession struct {
	Contract     *Rental           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RentalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RentalCallerSession struct {
	Contract *RentalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RentalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RentalTransactorSession struct {
	Contract     *RentalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RentalRaw is an auto generated low-level Go binding around an Ethereum contract.
type RentalRaw struct {
	Contract *Rental // Generic contract binding to access the raw methods on
}

// RentalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RentalCallerRaw struct {
	Contract *RentalCaller // Generic read-only contract binding to access the raw methods on
}

// RentalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RentalTransactorRaw struct {
	Contract *RentalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRental creates a new instance of Rental, bound to a specific deployed contract.
func NewRental(address common.Address, backend bind.ContractBackend) (*Rental, error) {
	contract, err := bindRental(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rental{RentalCaller: RentalCaller{contract: contract}, RentalTransactor: RentalTransactor{contract: contract}, RentalFilterer: RentalFilterer{contract: contract}}, nil
}

// NewRentalCaller creates a new read-only instance of Rental, bound to a specific deployed contract.
func NewRentalCaller(address common.Address, caller bind.ContractCaller) (*RentalCaller, error) {
	contract, err := bindRental(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RentalCaller{contract: contract}, nil
}

// NewRentalTransactor creates a new write-only instance of Rental, bound to a specific deployed contract.
func NewRentalTransactor(address common.Address, transactor bind.ContractTransactor) (*RentalTransactor, error) {
	contract, err := bindRental(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RentalTransactor{contract: contract}, nil
}

// NewRentalFilterer creates a new log filterer instance of Rental, bound to a specific deployed contract.
func NewRentalFilterer(address common.Address, filterer bind.ContractFilterer) (*RentalFilterer, error) {
	contract, err := bindRental(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RentalFilterer{contract: contract}, nil
}

// bindRental binds a generic wrapper to an already deployed contract.
func bindRental(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RentalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rental *RentalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rental.Contract.RentalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rental *RentalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rental.Contract.RentalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rental *RentalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rental.Contract.RentalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rental *RentalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rental.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rental *RentalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rental.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rental *RentalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rental.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Rental *RentalCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rental.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Rental *RentalSession) Admin() (common.Address, error) {
	return _Rental.Contract.Admin(&_Rental.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Rental *RentalCallerSession) Admin() (common.Address, error) {
	return _Rental.Contract.Admin(&_Rental.CallOpts)
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_Rental *RentalCaller) BikeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rental.contract.Call(opts, out, "bikeCount")
	return *ret0, err
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_Rental *RentalSession) BikeCount() (*big.Int, error) {
	return _Rental.Contract.BikeCount(&_Rental.CallOpts)
}

// BikeCount is a free data retrieval call binding the contract method 0xf4f3bc90.
//
// Solidity: function bikeCount() constant returns(uint256)
func (_Rental *RentalCallerSession) BikeCount() (*big.Int, error) {
	return _Rental.Contract.BikeCount(&_Rental.CallOpts)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_Rental *RentalCaller) Bikes(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Rental.contract.Call(opts, out, "bikes", arg0)
	return *ret, err
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_Rental *RentalSession) Bikes(arg0 *big.Int) (struct {
	Id         *big.Int
	CostPerDay *big.Int
	State      uint8
	Exists     bool
}, error) {
	return _Rental.Contract.Bikes(&_Rental.CallOpts, arg0)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, costPerDay uint256, state uint8, exists bool)
func (_Rental *RentalCallerSession) Bikes(arg0 *big.Int) (struct {
	Id         *big.Int
	CostPerDay *big.Int
	State      uint8
	Exists     bool
}, error) {
	return _Rental.Contract.Bikes(&_Rental.CallOpts, arg0)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_Rental *RentalCaller) Dev(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rental.contract.Call(opts, out, "dev")
	return *ret0, err
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_Rental *RentalSession) Dev() (bool, error) {
	return _Rental.Contract.Dev(&_Rental.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_Rental *RentalCallerSession) Dev() (bool, error) {
	return _Rental.Contract.Dev(&_Rental.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Rental *RentalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rental.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Rental *RentalSession) Owner() (common.Address, error) {
	return _Rental.Contract.Owner(&_Rental.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Rental *RentalCallerSession) Owner() (common.Address, error) {
	return _Rental.Contract.Owner(&_Rental.CallOpts)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_Rental *RentalCaller) Rentals(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Rental.contract.Call(opts, out, "rentals", arg0)
	return *ret, err
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_Rental *RentalSession) Rentals(arg0 common.Address) (struct {
	Renter     common.Address
	BikeId     *big.Int
	Deposit    *big.Int
	Cost       *big.Int
	DaysRented *big.Int
	ReturnDate *big.Int
	Rented     bool
}, error) {
	return _Rental.Contract.Rentals(&_Rental.CallOpts, arg0)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(renter address, bikeId uint256, deposit uint256, cost uint256, daysRented uint256, returnDate uint256, rented bool)
func (_Rental *RentalCallerSession) Rentals(arg0 common.Address) (struct {
	Renter     common.Address
	BikeId     *big.Int
	Deposit    *big.Int
	Cost       *big.Int
	DaysRented *big.Int
	ReturnDate *big.Int
	Rented     bool
}, error) {
	return _Rental.Contract.Rentals(&_Rental.CallOpts, arg0)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_Rental *RentalTransactor) AddBike(opts *bind.TransactOpts, _bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "addBike", _bikeId, _costPerDay)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_Rental *RentalSession) AddBike(_bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.AddBike(&_Rental.TransactOpts, _bikeId, _costPerDay)
}

// AddBike is a paid mutator transaction binding the contract method 0x9616734a.
//
// Solidity: function addBike(_bikeId uint256, _costPerDay uint256) returns(bool)
func (_Rental *RentalTransactorSession) AddBike(_bikeId *big.Int, _costPerDay *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.AddBike(&_Rental.TransactOpts, _bikeId, _costPerDay)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactor) BikeAwol(opts *bind.TransactOpts, _renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "bikeAwol", _renter, _bikeId)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalSession) BikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.BikeAwol(&_Rental.TransactOpts, _renter, _bikeId)
}

// BikeAwol is a paid mutator transaction binding the contract method 0x89e1476f.
//
// Solidity: function bikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactorSession) BikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.BikeAwol(&_Rental.TransactOpts, _renter, _bikeId)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_Rental *RentalTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_Rental *RentalSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Rental.Contract.ChangeOwner(&_Rental.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns(bool)
func (_Rental *RentalTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Rental.Contract.ChangeOwner(&_Rental.TransactOpts, _owner)
}

// ForceBikeAwol is a paid mutator transaction binding the contract method 0x52d11d4e.
//
// Solidity: function forceBikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactor) ForceBikeAwol(opts *bind.TransactOpts, _renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "forceBikeAwol", _renter, _bikeId)
}

// ForceBikeAwol is a paid mutator transaction binding the contract method 0x52d11d4e.
//
// Solidity: function forceBikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalSession) ForceBikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ForceBikeAwol(&_Rental.TransactOpts, _renter, _bikeId)
}

// ForceBikeAwol is a paid mutator transaction binding the contract method 0x52d11d4e.
//
// Solidity: function forceBikeAwol(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactorSession) ForceBikeAwol(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ForceBikeAwol(&_Rental.TransactOpts, _renter, _bikeId)
}

// ForceLateReturn is a paid mutator transaction binding the contract method 0xa373b597.
//
// Solidity: function forceLateReturn(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactor) ForceLateReturn(opts *bind.TransactOpts, _renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "forceLateReturn", _renter, _bikeId)
}

// ForceLateReturn is a paid mutator transaction binding the contract method 0xa373b597.
//
// Solidity: function forceLateReturn(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalSession) ForceLateReturn(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ForceLateReturn(&_Rental.TransactOpts, _renter, _bikeId)
}

// ForceLateReturn is a paid mutator transaction binding the contract method 0xa373b597.
//
// Solidity: function forceLateReturn(_renter address, _bikeId uint256) returns(bool)
func (_Rental *RentalTransactorSession) ForceLateReturn(_renter common.Address, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ForceLateReturn(&_Rental.TransactOpts, _renter, _bikeId)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_Rental *RentalTransactor) RentBike(opts *bind.TransactOpts, _bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "rentBike", _bikeId, _daysToRent, _deposit)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_Rental *RentalSession) RentBike(_bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.RentBike(&_Rental.TransactOpts, _bikeId, _daysToRent, _deposit)
}

// RentBike is a paid mutator transaction binding the contract method 0xcb24450d.
//
// Solidity: function rentBike(_bikeId uint256, _daysToRent uint256, _deposit uint256) returns(bool)
func (_Rental *RentalTransactorSession) RentBike(_bikeId *big.Int, _daysToRent *big.Int, _deposit *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.RentBike(&_Rental.TransactOpts, _bikeId, _daysToRent, _deposit)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_Rental *RentalTransactor) ReturnBike(opts *bind.TransactOpts, _bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "returnBike", _bikeId)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_Rental *RentalSession) ReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ReturnBike(&_Rental.TransactOpts, _bikeId)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa56a91c5.
//
// Solidity: function returnBike(_bikeId uint256) returns(bool)
func (_Rental *RentalTransactorSession) ReturnBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Rental.Contract.ReturnBike(&_Rental.TransactOpts, _bikeId)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_Rental *RentalTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_Rental *RentalSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Rental.Contract.SetAdmin(&_Rental.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_admin address) returns(bool)
func (_Rental *RentalTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Rental.Contract.SetAdmin(&_Rental.TransactOpts, _admin)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_Rental *RentalTransactor) SetErcInterface(opts *bind.TransactOpts, _bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rental.contract.Transact(opts, "setErcInterface", _bikeTokenContractAddress)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_Rental *RentalSession) SetErcInterface(_bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rental.Contract.SetErcInterface(&_Rental.TransactOpts, _bikeTokenContractAddress)
}

// SetErcInterface is a paid mutator transaction binding the contract method 0x04361569.
//
// Solidity: function setErcInterface(_bikeTokenContractAddress address) returns(bool)
func (_Rental *RentalTransactorSession) SetErcInterface(_bikeTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rental.Contract.SetErcInterface(&_Rental.TransactOpts, _bikeTokenContractAddress)
}

// RentalAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Rental contract.
type RentalAdminSetIterator struct {
	Event *RentalAdminSet // Event containing the contract specifics and raw log

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
func (it *RentalAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalAdminSet)
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
		it.Event = new(RentalAdminSet)
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
func (it *RentalAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalAdminSet represents a AdminSet event raised by the Rental contract.
type RentalAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_Rental *RentalFilterer) FilterAdminSet(opts *bind.FilterOpts) (*RentalAdminSetIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &RentalAdminSetIterator{contract: _Rental.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(_admin address)
func (_Rental *RentalFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *RentalAdminSet) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalAdminSet)
				if err := _Rental.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// RentalBikeAddedIterator is returned from FilterBikeAdded and is used to iterate over the raw logs and unpacked data for BikeAdded events raised by the Rental contract.
type RentalBikeAddedIterator struct {
	Event *RentalBikeAdded // Event containing the contract specifics and raw log

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
func (it *RentalBikeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalBikeAdded)
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
		it.Event = new(RentalBikeAdded)
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
func (it *RentalBikeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalBikeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalBikeAdded represents a BikeAdded event raised by the Rental contract.
type RentalBikeAdded struct {
	Id         *big.Int
	CostPerDay *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBikeAdded is a free log retrieval operation binding the contract event 0xea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416.
//
// Solidity: event BikeAdded(_id uint256, _costPerDay uint256)
func (_Rental *RentalFilterer) FilterBikeAdded(opts *bind.FilterOpts) (*RentalBikeAddedIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "BikeAdded")
	if err != nil {
		return nil, err
	}
	return &RentalBikeAddedIterator{contract: _Rental.contract, event: "BikeAdded", logs: logs, sub: sub}, nil
}

// WatchBikeAdded is a free log subscription operation binding the contract event 0xea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416.
//
// Solidity: event BikeAdded(_id uint256, _costPerDay uint256)
func (_Rental *RentalFilterer) WatchBikeAdded(opts *bind.WatchOpts, sink chan<- *RentalBikeAdded) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "BikeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalBikeAdded)
				if err := _Rental.contract.UnpackLog(event, "BikeAdded", log); err != nil {
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

// RentalBikeAwolIterator is returned from FilterBikeAwol and is used to iterate over the raw logs and unpacked data for BikeAwol events raised by the Rental contract.
type RentalBikeAwolIterator struct {
	Event *RentalBikeAwol // Event containing the contract specifics and raw log

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
func (it *RentalBikeAwolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalBikeAwol)
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
		it.Event = new(RentalBikeAwol)
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
func (it *RentalBikeAwolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalBikeAwolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalBikeAwol represents a BikeAwol event raised by the Rental contract.
type RentalBikeAwol struct {
	Renter common.Address
	BikeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeAwol is a free log retrieval operation binding the contract event 0x15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25.
//
// Solidity: event BikeAwol(_renter address, _bikeId uint256)
func (_Rental *RentalFilterer) FilterBikeAwol(opts *bind.FilterOpts) (*RentalBikeAwolIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "BikeAwol")
	if err != nil {
		return nil, err
	}
	return &RentalBikeAwolIterator{contract: _Rental.contract, event: "BikeAwol", logs: logs, sub: sub}, nil
}

// WatchBikeAwol is a free log subscription operation binding the contract event 0x15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25.
//
// Solidity: event BikeAwol(_renter address, _bikeId uint256)
func (_Rental *RentalFilterer) WatchBikeAwol(opts *bind.WatchOpts, sink chan<- *RentalBikeAwol) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "BikeAwol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalBikeAwol)
				if err := _Rental.contract.UnpackLog(event, "BikeAwol", log); err != nil {
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

// RentalBikeRentedIterator is returned from FilterBikeRented and is used to iterate over the raw logs and unpacked data for BikeRented events raised by the Rental contract.
type RentalBikeRentedIterator struct {
	Event *RentalBikeRented // Event containing the contract specifics and raw log

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
func (it *RentalBikeRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalBikeRented)
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
		it.Event = new(RentalBikeRented)
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
func (it *RentalBikeRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalBikeRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalBikeRented represents a BikeRented event raised by the Rental contract.
type RentalBikeRented struct {
	Renter     common.Address
	Id         *big.Int
	DaysRented *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBikeRented is a free log retrieval operation binding the contract event 0xabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd2408.
//
// Solidity: event BikeRented(_renter address, _id uint256, _daysRented uint256)
func (_Rental *RentalFilterer) FilterBikeRented(opts *bind.FilterOpts) (*RentalBikeRentedIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "BikeRented")
	if err != nil {
		return nil, err
	}
	return &RentalBikeRentedIterator{contract: _Rental.contract, event: "BikeRented", logs: logs, sub: sub}, nil
}

// WatchBikeRented is a free log subscription operation binding the contract event 0xabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd2408.
//
// Solidity: event BikeRented(_renter address, _id uint256, _daysRented uint256)
func (_Rental *RentalFilterer) WatchBikeRented(opts *bind.WatchOpts, sink chan<- *RentalBikeRented) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "BikeRented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalBikeRented)
				if err := _Rental.contract.UnpackLog(event, "BikeRented", log); err != nil {
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

// RentalBikeReturnedIterator is returned from FilterBikeReturned and is used to iterate over the raw logs and unpacked data for BikeReturned events raised by the Rental contract.
type RentalBikeReturnedIterator struct {
	Event *RentalBikeReturned // Event containing the contract specifics and raw log

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
func (it *RentalBikeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalBikeReturned)
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
		it.Event = new(RentalBikeReturned)
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
func (it *RentalBikeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalBikeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalBikeReturned represents a BikeReturned event raised by the Rental contract.
type RentalBikeReturned struct {
	Renter common.Address
	BikeId *big.Int
	Late   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeReturned is a free log retrieval operation binding the contract event 0xd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f.
//
// Solidity: event BikeReturned(_renter address, _bikeId uint256, _late bool)
func (_Rental *RentalFilterer) FilterBikeReturned(opts *bind.FilterOpts) (*RentalBikeReturnedIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "BikeReturned")
	if err != nil {
		return nil, err
	}
	return &RentalBikeReturnedIterator{contract: _Rental.contract, event: "BikeReturned", logs: logs, sub: sub}, nil
}

// WatchBikeReturned is a free log subscription operation binding the contract event 0xd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f.
//
// Solidity: event BikeReturned(_renter address, _bikeId uint256, _late bool)
func (_Rental *RentalFilterer) WatchBikeReturned(opts *bind.WatchOpts, sink chan<- *RentalBikeReturned) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "BikeReturned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalBikeReturned)
				if err := _Rental.contract.UnpackLog(event, "BikeReturned", log); err != nil {
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

// RentalErcInterfaceSetIterator is returned from FilterErcInterfaceSet and is used to iterate over the raw logs and unpacked data for ErcInterfaceSet events raised by the Rental contract.
type RentalErcInterfaceSetIterator struct {
	Event *RentalErcInterfaceSet // Event containing the contract specifics and raw log

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
func (it *RentalErcInterfaceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalErcInterfaceSet)
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
		it.Event = new(RentalErcInterfaceSet)
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
func (it *RentalErcInterfaceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalErcInterfaceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalErcInterfaceSet represents a ErcInterfaceSet event raised by the Rental contract.
type RentalErcInterfaceSet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterErcInterfaceSet is a free log retrieval operation binding the contract event 0xb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a25.
//
// Solidity: event ErcInterfaceSet()
func (_Rental *RentalFilterer) FilterErcInterfaceSet(opts *bind.FilterOpts) (*RentalErcInterfaceSetIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "ErcInterfaceSet")
	if err != nil {
		return nil, err
	}
	return &RentalErcInterfaceSetIterator{contract: _Rental.contract, event: "ErcInterfaceSet", logs: logs, sub: sub}, nil
}

// WatchErcInterfaceSet is a free log subscription operation binding the contract event 0xb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a25.
//
// Solidity: event ErcInterfaceSet()
func (_Rental *RentalFilterer) WatchErcInterfaceSet(opts *bind.WatchOpts, sink chan<- *RentalErcInterfaceSet) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "ErcInterfaceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalErcInterfaceSet)
				if err := _Rental.contract.UnpackLog(event, "ErcInterfaceSet", log); err != nil {
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

// RentalNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Rental contract.
type RentalNewOwnerIterator struct {
	Event *RentalNewOwner // Event containing the contract specifics and raw log

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
func (it *RentalNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalNewOwner)
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
		it.Event = new(RentalNewOwner)
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
func (it *RentalNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalNewOwner represents a NewOwner event raised by the Rental contract.
type RentalNewOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_Rental *RentalFilterer) FilterNewOwner(opts *bind.FilterOpts) (*RentalNewOwnerIterator, error) {

	logs, sub, err := _Rental.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &RentalNewOwnerIterator{contract: _Rental.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(_owner address)
func (_Rental *RentalFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *RentalNewOwner) (event.Subscription, error) {

	logs, sub, err := _Rental.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalNewOwner)
				if err := _Rental.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
