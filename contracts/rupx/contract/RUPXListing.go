// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/rupayaproject/rupaya/accounts/abi"
	"github.com/rupayaproject/rupaya/accounts/abi/bind"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/core/types"
)

// RUPXListingABI is the input ABI used to generate the binding from.
const RUPXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// RUPXListingBin is the compiled bytecode used for deploying new contracts.
const RUPXListingBin = `0x608060405234801561001057600080fd5b506102bf806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea000003410156101e357600080fd5b6040516068903480156108fc02916000818181858888f19350505050158015610210573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a72305820a258c7f15c7c6507a28499e1a95c7e7ca19f22f78bcf25bf0b842006720fd85d0029`

// DeployRUPXListing deploys a new Ethereum contract, binding an instance of RUPXListing to it.
func DeployRUPXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RUPXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(RUPXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RUPXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RUPXListing{RUPXListingCaller: RUPXListingCaller{contract: contract}, RUPXListingTransactor: RUPXListingTransactor{contract: contract}, RUPXListingFilterer: RUPXListingFilterer{contract: contract}}, nil
}

// RUPXListing is an auto generated Go binding around an Ethereum contract.
type RUPXListing struct {
	RUPXListingCaller     // Read-only binding to the contract
	RUPXListingTransactor // Write-only binding to the contract
	RUPXListingFilterer   // Log filterer for contract events
}

// RUPXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RUPXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RUPXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RUPXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RUPXListingSession struct {
	Contract     *RUPXListing      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RUPXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RUPXListingCallerSession struct {
	Contract *RUPXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RUPXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RUPXListingTransactorSession struct {
	Contract     *RUPXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RUPXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RUPXListingRaw struct {
	Contract *RUPXListing // Generic contract binding to access the raw methods on
}

// RUPXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RUPXListingCallerRaw struct {
	Contract *RUPXListingCaller // Generic read-only contract binding to access the raw methods on
}

// RUPXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RUPXListingTransactorRaw struct {
	Contract *RUPXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRUPXListing creates a new instance of RUPXListing, bound to a specific deployed contract.
func NewRUPXListing(address common.Address, backend bind.ContractBackend) (*RUPXListing, error) {
	contract, err := bindRUPXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RUPXListing{RUPXListingCaller: RUPXListingCaller{contract: contract}, RUPXListingTransactor: RUPXListingTransactor{contract: contract}, RUPXListingFilterer: RUPXListingFilterer{contract: contract}}, nil
}

// NewRUPXListingCaller creates a new read-only instance of RUPXListing, bound to a specific deployed contract.
func NewRUPXListingCaller(address common.Address, caller bind.ContractCaller) (*RUPXListingCaller, error) {
	contract, err := bindRUPXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RUPXListingCaller{contract: contract}, nil
}

// NewRUPXListingTransactor creates a new write-only instance of RUPXListing, bound to a specific deployed contract.
func NewRUPXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*RUPXListingTransactor, error) {
	contract, err := bindRUPXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RUPXListingTransactor{contract: contract}, nil
}

// NewRUPXListingFilterer creates a new log filterer instance of RUPXListing, bound to a specific deployed contract.
func NewRUPXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*RUPXListingFilterer, error) {
	contract, err := bindRUPXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RUPXListingFilterer{contract: contract}, nil
}

// bindRUPXListing binds a generic wrapper to an already deployed contract.
func bindRUPXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RUPXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RUPXListing *RUPXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RUPXListing.Contract.RUPXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RUPXListing *RUPXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RUPXListing.Contract.RUPXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RUPXListing *RUPXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RUPXListing.Contract.RUPXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RUPXListing *RUPXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RUPXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RUPXListing *RUPXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RUPXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RUPXListing *RUPXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RUPXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPXListing *RUPXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RUPXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPXListing *RUPXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _RUPXListing.Contract.GetTokenStatus(&_RUPXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPXListing *RUPXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _RUPXListing.Contract.GetTokenStatus(&_RUPXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPXListing *RUPXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RUPXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPXListing *RUPXListingSession) Tokens() ([]common.Address, error) {
	return _RUPXListing.Contract.Tokens(&_RUPXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPXListing *RUPXListingCallerSession) Tokens() ([]common.Address, error) {
	return _RUPXListing.Contract.Tokens(&_RUPXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPXListing *RUPXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RUPXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPXListing *RUPXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _RUPXListing.Contract.Apply(&_RUPXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPXListing *RUPXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _RUPXListing.Contract.Apply(&_RUPXListing.TransactOpts, token)
}
