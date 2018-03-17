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

// BikeRentalBin is the compiled bytecode used for deploying new contracts.
const BikeRentalBin = `60606040526003805460ff1916600190811761010060b060020a0319169091556000805433600160a060020a0316600160a060020a03199182168117835583549091161790915561124a90819061005690396000f3006060604052600436106100cc5763ffffffff60e060020a6000350416630436156981146100d15780635c526e19146101045780636ba85eb414610155578063704b6c02146101bc57806375f1473f146101db57806389e1476f146101f15780638da5cb5b1461021357806391cca3db146102425780639616734a14610255578063a56a91c51461026e578063a6f9dae114610284578063c4540dc1146102a3578063cb24450d146102c5578063d7d172fe146102e1578063f4f3bc90146102f7578063f851a4401461031c575b600080fd5b34156100dc57600080fd5b6100f0600160a060020a036004351661032f565b604051901515815260200160405180910390f35b341561010f57600080fd5b61011a6004356103af565b6040518085815260200184815260200183600381111561013657fe5b60ff168152911515602083015250604090810193509150505180910390f35b341561016057600080fd5b610174600160a060020a03600435166103dc565b604051600160a060020a03909716875260208701959095526040808701949094526060860192909252608085015260a084015290151560c083015260e0909101905180910390f35b34156101c757600080fd5b6100f0600160a060020a0360043516610427565b34156101e657600080fd5b6100f06004356104b0565b34156101fc57600080fd5b6100f0600160a060020a0360043516602435610513565b341561021e57600080fd5b610226610778565b604051600160a060020a03909116815260200160405180910390f35b341561024d57600080fd5b6100f0610787565b341561026057600080fd5b6100f0600435602435610790565b341561027957600080fd5b6100f06004356108f2565b341561028f57600080fd5b6100f0600160a060020a03600435166109c9565b34156102ae57600080fd5b6100f0600160a060020a0360043516602435610a52565b34156102d057600080fd5b6100f0600435602435604435610ade565b34156102ec57600080fd5b6100f0600435610e55565b341561030257600080fd5b61030a610e72565b60405190815260200160405180910390f35b341561032757600080fd5b610226610e78565b6000805433600160a060020a0390811691161461034b57600080fd5b6003805475ffffffffffffffffffffffffffffffffffffffff0000191662010000600160a060020a038516021790557fb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a2560405160405180910390a15060015b919050565b60046020526000908152604090208054600182015460029092015490919060ff8082169161010090041684565b60056020819052600091825260409091208054600182015460028301546003840154600485015495850154600690950154600160a060020a0390941695929491939092919060ff1687565b6000805433600160a060020a0390811691161461044357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b60008060008381526004602052604090206002015460ff1660038111156104d357fe5b14806104fe5750600260008381526004602052604090206002015460ff1660038111156104fc57fe5b145b1561050b575060016103aa565b5060006103aa565b60008054819033600160a060020a0390811691161461053157600080fd5b600160a060020a0384166000908152600560205260409020600601548490849060ff161515600114801561057f5750600160a060020a03821660009081526005602052604090206001015481145b80156105aa5750600160008281526004602052604090206002015460ff1660038111156105a857fe5b145b15156105b557600080fd5b600160a060020a038616600090815260056020819052604090912001546105e59062278d0063ffffffff610e8716565b4210156105f157600080fd5b600160a060020a038616600090815260056020526040902060038101546002909101546106239163ffffffff610ea016565b600160a060020a0387166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff19168155600181018590556002808201869055600380830187905560048084018890559583018790556006909201805460ff199081169091558c87529490935293819020909101805490921690921790559093507f15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25908790879051600160a060020a03909216825260208201526040908101905180910390a1600354620100009004600160a060020a031663a9059cbb338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561074a57600080fd5b5af1151561075757600080fd5b50505060405180519050151561076c57600080fd5b50600195945050505050565b600054600160a060020a031681565b60035460ff1681565b6000805433600160a060020a039081169116146107ac57600080fd5b826000808281526004602052604090206002015460ff1660038111156107ce57fe5b1480156107f25750600081815260046020526040902060020154610100900460ff16155b15156107fd57600080fd5b608060405190810160405280858152602001848152602001600360019054906101000a900460ff16600381111561083057fe5b8152600160209182015260008681526004909152604090208151815560208201518160010155604082015160028201805460ff1916600183600381111561087357fe5b02179055506060820151600291820180549115156101000261ff0019909216919091179055546108ab9150600163ffffffff610e8716565b6002557fea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416848460405191825260208201526040908101905180910390a15060019392505050565b600160a060020a0333908116600090815260056020526040812060060154909190839060ff16151560011480156109435750600160a060020a03821660009081526005602052604090206001015481145b801561096e5750600160008281526004602052604090206002015460ff16600381111561096c57fe5b145b151561097957600080fd5b600160a060020a0333166000908152600560208190526040909120015442116109b9576109a584610eb5565b15156109b057600080fd5b600192506109c2565b6109a584611068565b5050919050565b6000805433600160a060020a039081169116146109e557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b600160a060020a03821660009081526005602052604081206006015460ff1615156001148015610a9c5750600160a060020a03831660009081526005602052604090206001015482145b8015610ac75750600160008381526004602052604090206002015460ff166003811115610ac557fe5b145b15610ad457506001610ad8565b5060005b92915050565b60008080858160008281526004602052604090206002015460ff166003811115610b0457fe5b148015610b2d575060008181526004602052604090206002015460ff6101009091041615156001145b1515610b3857600080fd5b33600160a060020a03811660009081526005602052604090206006015460ff1615610b6257600080fd5b600088815260046020526040902060010154610b9790600390610b8b908a63ffffffff6111f716565b9063ffffffff6111f716565b8614610ba257600080fd5b600088815260046020526040902060010154610bc4908863ffffffff6111f716565b9350610bd9876201518063ffffffff6111f716565b9250610beb428463ffffffff610e8716565b60008981526004602052604090819020600201805460ff1916600117905590935060e090519081016040908152600160a060020a03331680835260208084018c90528284018a905260608401889052608084018b905260a08401879052600160c0850152600091825260059052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151600691909101805460ff1916911515919091179055507fabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd24083389896040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1600354620100009004600160a060020a03166323b872dd33308960405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610d9257600080fd5b5af11515610d9f57600080fd5b505050604051805190501515610db457600080fd5b600354600054600160a060020a03620100009092048216916323b872dd913391168760405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610e2557600080fd5b5af11515610e3257600080fd5b505050604051805190501515610e4757600080fd5b506001979650505050505050565b600090815260046020526040902060020154610100900460ff1690565b60025481565b600154600160a060020a031681565b600082820183811015610e9957600080fd5b9392505050565b600082821115610eaf57600080fd5b50900390565b600160a060020a033316600090815260056020526040812060038101546002909101548291610eea919063ffffffff610ea016565b600160a060020a0333166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916815560018082018690556002808301879055600380840188905560048085018990559684018890556006909301805460ff1990811690915583548c89529690955292909520909101805495965061010090930460ff169492939290911691908490811115610f8857fe5b02179055507fd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f33846000604051600160a060020a039093168352602083019190915215156040808301919091526060909101905180910390a1600354620100009004600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561103d57600080fd5b5af1151561104a57600080fd5b50505060405180519050151561105f57600080fd5b50600192915050565b600160a060020a03331660009081526005602052604081206003810154600290910154829161109d919063ffffffff610ea016565b600160a060020a0333166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916815560018082018690556002808301879055600380840188905560048085018990559684018890556006909301805460ff1990811690915583548c89529690955292909520909101805495965061010090930460ff16949293929091169190849081111561113b57fe5b02179055507fd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f33846001604051600160a060020a039093168352602083019190915215156040808301919091526060909101905180910390a1600354600054600160a060020a036201000090920482169163a9059cbb91168360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561103d57600080fd5b6000828202831580611213575082848281151561121057fe5b04145b1515610e9957600080fd00a165627a7a723058209e9746140e4005f9f5b51bcf8cf3c8c83d577b96c36101788585c729236656310029`

// DeployBikeRental deploys a new Ethereum contract, binding an instance of BikeRental to it.
func DeployBikeRental(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BikeRental, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeRentalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BikeRentalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BikeRental{BikeRentalCaller: BikeRentalCaller{contract: contract}, BikeRentalTransactor: BikeRentalTransactor{contract: contract}, BikeRentalFilterer: BikeRentalFilterer{contract: contract}}, nil
}

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
