// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EjariRulesCABI is the input ABI used to generate the binding from.
const EjariRulesCABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"locality\",\"type\":\"address\"},{\"name\":\"incrementPercentage\",\"type\":\"uint256\"},{\"name\":\"maxRent\",\"type\":\"uint256\"}],\"name\":\"addEjariRule\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locality\",\"type\":\"address\"},{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"incrementPercentage\",\"type\":\"uint256\"}],\"name\":\"isValidCondition\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"type\":\"constructor\"}]"

// EjariRulesCBin is the compiled bytecode used for deploying new contracts.
const EjariRulesCBin = `0x6060604052600080546c0100000000000000000000000033810204600160a060020a031990911617905560eb806100366000396000f3606060405260e060020a6000350463d4cae86481146026578063f02ce71614604e575b6002565b34600257609260043560243560443560005433600160a060020a0390811691161460a8576002565b346002576094600435602435604435600160a060020a0383166000908152600160205260408120805483118060865750600181015484115b1560de576000915060e3565b005b604080519115158252519081900360200190f35b6040805180820182529283526020808401928352600160a060020a03949094166000908152600194859052209151825551910155565b600191505b50939250505056`

// DeployEjariRulesC deploys a new Ethereum contract, binding an instance of EjariRulesC to it.
func DeployEjariRulesC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EjariRulesC, error) {
	parsed, err := abi.JSON(strings.NewReader(EjariRulesCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EjariRulesCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EjariRulesC{EjariRulesCCaller: EjariRulesCCaller{contract: contract}, EjariRulesCTransactor: EjariRulesCTransactor{contract: contract}}, nil
}

// EjariRulesC is an auto generated Go binding around an Ethereum contract.
type EjariRulesC struct {
	EjariRulesCCaller     // Read-only binding to the contract
	EjariRulesCTransactor // Write-only binding to the contract
}

// EjariRulesCCaller is an auto generated read-only Go binding around an Ethereum contract.
type EjariRulesCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EjariRulesCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EjariRulesCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EjariRulesCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EjariRulesCSession struct {
	Contract     *EjariRulesC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EjariRulesCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EjariRulesCCallerSession struct {
	Contract *EjariRulesCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EjariRulesCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EjariRulesCTransactorSession struct {
	Contract     *EjariRulesCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EjariRulesCRaw is an auto generated low-level Go binding around an Ethereum contract.
type EjariRulesCRaw struct {
	Contract *EjariRulesC // Generic contract binding to access the raw methods on
}

// EjariRulesCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EjariRulesCCallerRaw struct {
	Contract *EjariRulesCCaller // Generic read-only contract binding to access the raw methods on
}

// EjariRulesCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EjariRulesCTransactorRaw struct {
	Contract *EjariRulesCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEjariRulesC creates a new instance of EjariRulesC, bound to a specific deployed contract.
func NewEjariRulesC(address common.Address, backend bind.ContractBackend) (*EjariRulesC, error) {
	contract, err := bindEjariRulesC(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EjariRulesC{EjariRulesCCaller: EjariRulesCCaller{contract: contract}, EjariRulesCTransactor: EjariRulesCTransactor{contract: contract}}, nil
}

// NewEjariRulesCCaller creates a new read-only instance of EjariRulesC, bound to a specific deployed contract.
func NewEjariRulesCCaller(address common.Address, caller bind.ContractCaller) (*EjariRulesCCaller, error) {
	contract, err := bindEjariRulesC(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EjariRulesCCaller{contract: contract}, nil
}

// NewEjariRulesCTransactor creates a new write-only instance of EjariRulesC, bound to a specific deployed contract.
func NewEjariRulesCTransactor(address common.Address, transactor bind.ContractTransactor) (*EjariRulesCTransactor, error) {
	contract, err := bindEjariRulesC(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EjariRulesCTransactor{contract: contract}, nil
}

// bindEjariRulesC binds a generic wrapper to an already deployed contract.
func bindEjariRulesC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EjariRulesCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EjariRulesC *EjariRulesCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EjariRulesC.Contract.EjariRulesCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EjariRulesC *EjariRulesCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EjariRulesC.Contract.EjariRulesCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EjariRulesC *EjariRulesCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EjariRulesC.Contract.EjariRulesCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EjariRulesC *EjariRulesCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EjariRulesC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EjariRulesC *EjariRulesCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EjariRulesC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EjariRulesC *EjariRulesCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EjariRulesC.Contract.contract.Transact(opts, method, params...)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0xd4cae864.
//
// Solidity: function addEjariRule(locality address, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRulesC *EjariRulesCTransactor) AddEjariRule(opts *bind.TransactOpts, locality common.Address, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.contract.Transact(opts, "addEjariRule", locality, incrementPercentage, maxRent)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0xd4cae864.
//
// Solidity: function addEjariRule(locality address, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRulesC *EjariRulesCSession) AddEjariRule(locality common.Address, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.Contract.AddEjariRule(&_EjariRulesC.TransactOpts, locality, incrementPercentage, maxRent)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0xd4cae864.
//
// Solidity: function addEjariRule(locality address, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRulesC *EjariRulesCTransactorSession) AddEjariRule(locality common.Address, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.Contract.AddEjariRule(&_EjariRulesC.TransactOpts, locality, incrementPercentage, maxRent)
}

// IsValidCondition is a paid mutator transaction binding the contract method 0xf02ce716.
//
// Solidity: function isValidCondition(locality address, rent uint256, incrementPercentage uint256) returns(bool)
func (_EjariRulesC *EjariRulesCTransactor) IsValidCondition(opts *bind.TransactOpts, locality common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.contract.Transact(opts, "isValidCondition", locality, rent, incrementPercentage)
}

// IsValidCondition is a paid mutator transaction binding the contract method 0xf02ce716.
//
// Solidity: function isValidCondition(locality address, rent uint256, incrementPercentage uint256) returns(bool)
func (_EjariRulesC *EjariRulesCSession) IsValidCondition(locality common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.Contract.IsValidCondition(&_EjariRulesC.TransactOpts, locality, rent, incrementPercentage)
}

// IsValidCondition is a paid mutator transaction binding the contract method 0xf02ce716.
//
// Solidity: function isValidCondition(locality address, rent uint256, incrementPercentage uint256) returns(bool)
func (_EjariRulesC *EjariRulesCTransactorSession) IsValidCondition(locality common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _EjariRulesC.Contract.IsValidCondition(&_EjariRulesC.TransactOpts, locality, rent, incrementPercentage)
}

// LocalityABI is the input ABI used to generate the binding from.
const LocalityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"n\",\"type\":\"string\"}],\"type\":\"constructor\"}]"

// LocalityBin is the compiled bytecode used for deploying new contracts.
const LocalityBin = `0x6060604052604051610130380380610130833981016040528051018060026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10608d57805160ff19168380011785555b50607d9291505b8082111560ba5760008155600101606b565b5050506072806100be6000396000f35b828001600101855582156064579182015b828111156064578251826000505591602001919060010190609e565b509056606060405260e060020a6000350463172d9773811460305780632826f6a514603c578063e7ee6ad6146048575b6002565b34600257606060015481565b34600257606060005481565b34600257600080546004350190556001805481019055005b60408051918252519081900360200190f3`

// DeployLocality deploys a new Ethereum contract, binding an instance of Locality to it.
func DeployLocality(auth *bind.TransactOpts, backend bind.ContractBackend, n string) (common.Address, *types.Transaction, *Locality, error) {
	parsed, err := abi.JSON(strings.NewReader(LocalityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LocalityBin), backend, n)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Locality{LocalityCaller: LocalityCaller{contract: contract}, LocalityTransactor: LocalityTransactor{contract: contract}}, nil
}

// Locality is an auto generated Go binding around an Ethereum contract.
type Locality struct {
	LocalityCaller     // Read-only binding to the contract
	LocalityTransactor // Write-only binding to the contract
}

// LocalityCaller is an auto generated read-only Go binding around an Ethereum contract.
type LocalityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LocalityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LocalitySession struct {
	Contract     *Locality         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LocalityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LocalityCallerSession struct {
	Contract *LocalityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LocalityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LocalityTransactorSession struct {
	Contract     *LocalityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LocalityRaw is an auto generated low-level Go binding around an Ethereum contract.
type LocalityRaw struct {
	Contract *Locality // Generic contract binding to access the raw methods on
}

// LocalityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LocalityCallerRaw struct {
	Contract *LocalityCaller // Generic read-only contract binding to access the raw methods on
}

// LocalityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LocalityTransactorRaw struct {
	Contract *LocalityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocality creates a new instance of Locality, bound to a specific deployed contract.
func NewLocality(address common.Address, backend bind.ContractBackend) (*Locality, error) {
	contract, err := bindLocality(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locality{LocalityCaller: LocalityCaller{contract: contract}, LocalityTransactor: LocalityTransactor{contract: contract}}, nil
}

// NewLocalityCaller creates a new read-only instance of Locality, bound to a specific deployed contract.
func NewLocalityCaller(address common.Address, caller bind.ContractCaller) (*LocalityCaller, error) {
	contract, err := bindLocality(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &LocalityCaller{contract: contract}, nil
}

// NewLocalityTransactor creates a new write-only instance of Locality, bound to a specific deployed contract.
func NewLocalityTransactor(address common.Address, transactor bind.ContractTransactor) (*LocalityTransactor, error) {
	contract, err := bindLocality(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &LocalityTransactor{contract: contract}, nil
}

// bindLocality binds a generic wrapper to an already deployed contract.
func bindLocality(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LocalityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locality *LocalityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Locality.Contract.LocalityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locality *LocalityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locality.Contract.LocalityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locality *LocalityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locality.Contract.LocalityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locality *LocalityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Locality.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locality *LocalityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locality.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locality *LocalityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locality.Contract.contract.Transact(opts, method, params...)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Locality *LocalityCaller) NumberOfRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Locality.contract.Call(opts, out, "numberOfRatings")
	return *ret0, err
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Locality *LocalitySession) NumberOfRatings() (*big.Int, error) {
	return _Locality.Contract.NumberOfRatings(&_Locality.CallOpts)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Locality *LocalityCallerSession) NumberOfRatings() (*big.Int, error) {
	return _Locality.Contract.NumberOfRatings(&_Locality.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Locality *LocalityCaller) TotalRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Locality.contract.Call(opts, out, "totalRatings")
	return *ret0, err
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Locality *LocalitySession) TotalRatings() (*big.Int, error) {
	return _Locality.Contract.TotalRatings(&_Locality.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Locality *LocalityCallerSession) TotalRatings() (*big.Int, error) {
	return _Locality.Contract.TotalRatings(&_Locality.CallOpts)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Locality *LocalityTransactor) Rate(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Locality.contract.Transact(opts, "rate", rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Locality *LocalitySession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Locality.Contract.Rate(&_Locality.TransactOpts, rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Locality *LocalityTransactorSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Locality.Contract.Rate(&_Locality.TransactOpts, rating)
}

// PersonABI is the input ABI used to generate the binding from.
const PersonABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAccountAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"n\",\"type\":\"string\"},{\"name\":\"a\",\"type\":\"address\"}],\"type\":\"constructor\"}]"

// PersonBin is the compiled bytecode used for deploying new contracts.
const PersonBin = `0x60606040526040516101b13803806101b18339810160405280516080519101908160026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1060b857805160ff19168380011785555b5060829291505b8082111560e557600081556001016070565b5050600380546c0100000000000000000000000080840204600160a060020a0319909116179055505060c8806100e96000396000f35b828001600101855582156069579182015b82811115606957825182600050559160200191906001019060c9565b509056606060405260e060020a60003504630e2562d981146044578063172d97731460715780632826f6a514607d5780635dab2420146089578063e7ee6ad614609e575b6002565b34600257600354600160a060020a03165b60408051600160a060020a039092168252519081900360200190f35b3460025760b660015481565b3460025760b660005481565b346002576055600354600160a060020a031681565b34600257600080546004350190556001805481019055005b60408051918252519081900360200190f3`

// DeployPerson deploys a new Ethereum contract, binding an instance of Person to it.
func DeployPerson(auth *bind.TransactOpts, backend bind.ContractBackend, n string, a common.Address) (common.Address, *types.Transaction, *Person, error) {
	parsed, err := abi.JSON(strings.NewReader(PersonABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PersonBin), backend, n, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Person{PersonCaller: PersonCaller{contract: contract}, PersonTransactor: PersonTransactor{contract: contract}}, nil
}

// Person is an auto generated Go binding around an Ethereum contract.
type Person struct {
	PersonCaller     // Read-only binding to the contract
	PersonTransactor // Write-only binding to the contract
}

// PersonCaller is an auto generated read-only Go binding around an Ethereum contract.
type PersonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PersonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PersonSession struct {
	Contract     *Person           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PersonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PersonCallerSession struct {
	Contract *PersonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PersonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PersonTransactorSession struct {
	Contract     *PersonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PersonRaw is an auto generated low-level Go binding around an Ethereum contract.
type PersonRaw struct {
	Contract *Person // Generic contract binding to access the raw methods on
}

// PersonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PersonCallerRaw struct {
	Contract *PersonCaller // Generic read-only contract binding to access the raw methods on
}

// PersonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PersonTransactorRaw struct {
	Contract *PersonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerson creates a new instance of Person, bound to a specific deployed contract.
func NewPerson(address common.Address, backend bind.ContractBackend) (*Person, error) {
	contract, err := bindPerson(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Person{PersonCaller: PersonCaller{contract: contract}, PersonTransactor: PersonTransactor{contract: contract}}, nil
}

// NewPersonCaller creates a new read-only instance of Person, bound to a specific deployed contract.
func NewPersonCaller(address common.Address, caller bind.ContractCaller) (*PersonCaller, error) {
	contract, err := bindPerson(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PersonCaller{contract: contract}, nil
}

// NewPersonTransactor creates a new write-only instance of Person, bound to a specific deployed contract.
func NewPersonTransactor(address common.Address, transactor bind.ContractTransactor) (*PersonTransactor, error) {
	contract, err := bindPerson(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PersonTransactor{contract: contract}, nil
}

// bindPerson binds a generic wrapper to an already deployed contract.
func bindPerson(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PersonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Person *PersonRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Person.Contract.PersonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Person *PersonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Person.Contract.PersonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Person *PersonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Person.Contract.PersonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Person *PersonCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Person.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Person *PersonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Person.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Person *PersonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Person.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Person *PersonCaller) Account(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Person.contract.Call(opts, out, "account")
	return *ret0, err
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Person *PersonSession) Account() (common.Address, error) {
	return _Person.Contract.Account(&_Person.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Person *PersonCallerSession) Account() (common.Address, error) {
	return _Person.Contract.Account(&_Person.CallOpts)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Person *PersonCaller) NumberOfRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Person.contract.Call(opts, out, "numberOfRatings")
	return *ret0, err
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Person *PersonSession) NumberOfRatings() (*big.Int, error) {
	return _Person.Contract.NumberOfRatings(&_Person.CallOpts)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Person *PersonCallerSession) NumberOfRatings() (*big.Int, error) {
	return _Person.Contract.NumberOfRatings(&_Person.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Person *PersonCaller) TotalRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Person.contract.Call(opts, out, "totalRatings")
	return *ret0, err
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Person *PersonSession) TotalRatings() (*big.Int, error) {
	return _Person.Contract.TotalRatings(&_Person.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Person *PersonCallerSession) TotalRatings() (*big.Int, error) {
	return _Person.Contract.TotalRatings(&_Person.CallOpts)
}

// GetAccountAddress is a paid mutator transaction binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() returns(address)
func (_Person *PersonTransactor) GetAccountAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Person.contract.Transact(opts, "getAccountAddress")
}

// GetAccountAddress is a paid mutator transaction binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() returns(address)
func (_Person *PersonSession) GetAccountAddress() (*types.Transaction, error) {
	return _Person.Contract.GetAccountAddress(&_Person.TransactOpts)
}

// GetAccountAddress is a paid mutator transaction binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() returns(address)
func (_Person *PersonTransactorSession) GetAccountAddress() (*types.Transaction, error) {
	return _Person.Contract.GetAccountAddress(&_Person.TransactOpts)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Person *PersonTransactor) Rate(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Person.contract.Transact(opts, "rate", rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Person *PersonSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Person.Contract.Rate(&_Person.TransactOpts, rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Person *PersonTransactorSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Person.Contract.Rate(&_Person.TransactOpts, rating)
}

// PropertyCABI is the input ABI used to generate the binding from.
const PropertyCABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getLocality\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"lt\",\"type\":\"string\"},{\"name\":\"ln\",\"type\":\"string\"},{\"name\":\"loc\",\"type\":\"address\"}],\"type\":\"constructor\"}]"

// PropertyCBin is the compiled bytecode used for deploying new contracts.
const PropertyCBin = `0x606060405260405161024238038061024283398101604052805160805160a0519183019201908260026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c357805160ff19168380011785555b506100f39291505b8082111561014c5760008155600101610078565b5050600480546c0100000000000000000000000080840204600160a060020a031990911617905550505060c2806101806000396000f35b82800160010185558215610070579182015b828111156100705782518260005055916020019190600101906100d5565b50508160036000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015057805160ff19168380011785555b5061008c929150610078565b5090565b82800160010185558215610140579182015b8281111561014057825182600050559160200191906001019061016256606060405260e060020a6000350463172d97738114603a5780632826f6a5146046578063b2b5f86f146052578063e7ee6ad6146098575b6002565b3460025760b060015481565b3460025760b060005481565b3460025760045473ffffffffffffffffffffffffffffffffffffffff166040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34600257600080546004350190556001805481019055005b60408051918252519081900360200190f3`

// DeployPropertyC deploys a new Ethereum contract, binding an instance of PropertyC to it.
func DeployPropertyC(auth *bind.TransactOpts, backend bind.ContractBackend, lt string, ln string, loc common.Address) (common.Address, *types.Transaction, *PropertyC, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PropertyCBin), backend, lt, ln, loc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PropertyC{PropertyCCaller: PropertyCCaller{contract: contract}, PropertyCTransactor: PropertyCTransactor{contract: contract}}, nil
}

// PropertyC is an auto generated Go binding around an Ethereum contract.
type PropertyC struct {
	PropertyCCaller     // Read-only binding to the contract
	PropertyCTransactor // Write-only binding to the contract
}

// PropertyCCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertyCSession struct {
	Contract     *PropertyC        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyCCallerSession struct {
	Contract *PropertyCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PropertyCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyCTransactorSession struct {
	Contract     *PropertyCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PropertyCRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyCRaw struct {
	Contract *PropertyC // Generic contract binding to access the raw methods on
}

// PropertyCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyCCallerRaw struct {
	Contract *PropertyCCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyCTransactorRaw struct {
	Contract *PropertyCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropertyC creates a new instance of PropertyC, bound to a specific deployed contract.
func NewPropertyC(address common.Address, backend bind.ContractBackend) (*PropertyC, error) {
	contract, err := bindPropertyC(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropertyC{PropertyCCaller: PropertyCCaller{contract: contract}, PropertyCTransactor: PropertyCTransactor{contract: contract}}, nil
}

// NewPropertyCCaller creates a new read-only instance of PropertyC, bound to a specific deployed contract.
func NewPropertyCCaller(address common.Address, caller bind.ContractCaller) (*PropertyCCaller, error) {
	contract, err := bindPropertyC(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyCCaller{contract: contract}, nil
}

// NewPropertyCTransactor creates a new write-only instance of PropertyC, bound to a specific deployed contract.
func NewPropertyCTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyCTransactor, error) {
	contract, err := bindPropertyC(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PropertyCTransactor{contract: contract}, nil
}

// bindPropertyC binds a generic wrapper to an already deployed contract.
func bindPropertyC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyC *PropertyCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropertyC.Contract.PropertyCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyC *PropertyCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyC.Contract.PropertyCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyC *PropertyCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyC.Contract.PropertyCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyC *PropertyCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropertyC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyC *PropertyCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyC *PropertyCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyC.Contract.contract.Transact(opts, method, params...)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_PropertyC *PropertyCCaller) NumberOfRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropertyC.contract.Call(opts, out, "numberOfRatings")
	return *ret0, err
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_PropertyC *PropertyCSession) NumberOfRatings() (*big.Int, error) {
	return _PropertyC.Contract.NumberOfRatings(&_PropertyC.CallOpts)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_PropertyC *PropertyCCallerSession) NumberOfRatings() (*big.Int, error) {
	return _PropertyC.Contract.NumberOfRatings(&_PropertyC.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_PropertyC *PropertyCCaller) TotalRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropertyC.contract.Call(opts, out, "totalRatings")
	return *ret0, err
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_PropertyC *PropertyCSession) TotalRatings() (*big.Int, error) {
	return _PropertyC.Contract.TotalRatings(&_PropertyC.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_PropertyC *PropertyCCallerSession) TotalRatings() (*big.Int, error) {
	return _PropertyC.Contract.TotalRatings(&_PropertyC.CallOpts)
}

// GetLocality is a paid mutator transaction binding the contract method 0xb2b5f86f.
//
// Solidity: function getLocality() returns(address)
func (_PropertyC *PropertyCTransactor) GetLocality(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyC.contract.Transact(opts, "getLocality")
}

// GetLocality is a paid mutator transaction binding the contract method 0xb2b5f86f.
//
// Solidity: function getLocality() returns(address)
func (_PropertyC *PropertyCSession) GetLocality() (*types.Transaction, error) {
	return _PropertyC.Contract.GetLocality(&_PropertyC.TransactOpts)
}

// GetLocality is a paid mutator transaction binding the contract method 0xb2b5f86f.
//
// Solidity: function getLocality() returns(address)
func (_PropertyC *PropertyCTransactorSession) GetLocality() (*types.Transaction, error) {
	return _PropertyC.Contract.GetLocality(&_PropertyC.TransactOpts)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_PropertyC *PropertyCTransactor) Rate(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _PropertyC.contract.Transact(opts, "rate", rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_PropertyC *PropertyCSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _PropertyC.Contract.Rate(&_PropertyC.TransactOpts, rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_PropertyC *PropertyCTransactorSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _PropertyC.Contract.Rate(&_PropertyC.TransactOpts, rating)
}

// RatedABI is the input ABI used to generate the binding from.
const RatedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRatings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// RatedBin is the compiled bytecode used for deploying new contracts.
const RatedBin = `0x606060405260728060106000396000f3606060405260e060020a6000350463172d9773811460305780632826f6a514603c578063e7ee6ad6146048575b6002565b34600257606060015481565b34600257606060005481565b34600257600080546004350190556001805481019055005b60408051918252519081900360200190f3`

// DeployRated deploys a new Ethereum contract, binding an instance of Rated to it.
func DeployRated(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rated, error) {
	parsed, err := abi.JSON(strings.NewReader(RatedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RatedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rated{RatedCaller: RatedCaller{contract: contract}, RatedTransactor: RatedTransactor{contract: contract}}, nil
}

// Rated is an auto generated Go binding around an Ethereum contract.
type Rated struct {
	RatedCaller     // Read-only binding to the contract
	RatedTransactor // Write-only binding to the contract
}

// RatedCaller is an auto generated read-only Go binding around an Ethereum contract.
type RatedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RatedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RatedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RatedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RatedSession struct {
	Contract     *Rated            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RatedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RatedCallerSession struct {
	Contract *RatedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RatedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RatedTransactorSession struct {
	Contract     *RatedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RatedRaw is an auto generated low-level Go binding around an Ethereum contract.
type RatedRaw struct {
	Contract *Rated // Generic contract binding to access the raw methods on
}

// RatedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RatedCallerRaw struct {
	Contract *RatedCaller // Generic read-only contract binding to access the raw methods on
}

// RatedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RatedTransactorRaw struct {
	Contract *RatedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRated creates a new instance of Rated, bound to a specific deployed contract.
func NewRated(address common.Address, backend bind.ContractBackend) (*Rated, error) {
	contract, err := bindRated(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rated{RatedCaller: RatedCaller{contract: contract}, RatedTransactor: RatedTransactor{contract: contract}}, nil
}

// NewRatedCaller creates a new read-only instance of Rated, bound to a specific deployed contract.
func NewRatedCaller(address common.Address, caller bind.ContractCaller) (*RatedCaller, error) {
	contract, err := bindRated(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RatedCaller{contract: contract}, nil
}

// NewRatedTransactor creates a new write-only instance of Rated, bound to a specific deployed contract.
func NewRatedTransactor(address common.Address, transactor bind.ContractTransactor) (*RatedTransactor, error) {
	contract, err := bindRated(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RatedTransactor{contract: contract}, nil
}

// bindRated binds a generic wrapper to an already deployed contract.
func bindRated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RatedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rated *RatedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rated.Contract.RatedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rated *RatedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rated.Contract.RatedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rated *RatedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rated.Contract.RatedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rated *RatedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rated *RatedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rated *RatedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rated.Contract.contract.Transact(opts, method, params...)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Rated *RatedCaller) NumberOfRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rated.contract.Call(opts, out, "numberOfRatings")
	return *ret0, err
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Rated *RatedSession) NumberOfRatings() (*big.Int, error) {
	return _Rated.Contract.NumberOfRatings(&_Rated.CallOpts)
}

// NumberOfRatings is a free data retrieval call binding the contract method 0x172d9773.
//
// Solidity: function numberOfRatings() constant returns(uint256)
func (_Rated *RatedCallerSession) NumberOfRatings() (*big.Int, error) {
	return _Rated.Contract.NumberOfRatings(&_Rated.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Rated *RatedCaller) TotalRatings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rated.contract.Call(opts, out, "totalRatings")
	return *ret0, err
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Rated *RatedSession) TotalRatings() (*big.Int, error) {
	return _Rated.Contract.TotalRatings(&_Rated.CallOpts)
}

// TotalRatings is a free data retrieval call binding the contract method 0x2826f6a5.
//
// Solidity: function totalRatings() constant returns(uint256)
func (_Rated *RatedCallerSession) TotalRatings() (*big.Int, error) {
	return _Rated.Contract.TotalRatings(&_Rated.CallOpts)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Rated *RatedTransactor) Rate(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Rated.contract.Transact(opts, "rate", rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Rated *RatedSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Rated.Contract.Rate(&_Rated.TransactOpts, rating)
}

// Rate is a paid mutator transaction binding the contract method 0xe7ee6ad6.
//
// Solidity: function rate(rating uint256) returns()
func (_Rated *RatedTransactorSession) Rate(rating *big.Int) (*types.Transaction, error) {
	return _Rated.Contract.Rate(&_Rated.TransactOpts, rating)
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownership\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"property\",\"type\":\"address\"},{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"incrementPercentage\",\"type\":\"uint256\"}],\"name\":\"isValidTenancy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ejariRules\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"property\",\"type\":\"address\"}],\"name\":\"assignOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ejariRule\",\"type\":\"address\"}],\"type\":\"constructor\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `0x606060405260405160208061031b83395050608060405260008054600160a060020a0319166c01000000000000000000000000338102041781556102d390819061004890396000f3606060405260e060020a600035046327d6d6e0811461004a5780632b20e397146100705780633faa78b61461008757806346d71ab414610136578063bbd2b2991461014d575b610002565b3461000257610176600435600260205260009081526040902054600160a060020a031681565b3461000257610176600054600160a060020a031681565b3461000257610192600435602435604435600160a060020a038084166000908152600260209081526040808320548151830184905281517f0e2562d90000000000000000000000000000000000000000000000000000000081529151939416928392630e2562d9926004808201939182900301818887803b156100025760325a03f1156100025750506040515133600160a060020a0390811691161490506101a857600091505b509392505050565b3461000257610176600154600160a060020a031681565b34610002576101a660043560243560005433600160a060020a0390811691161461028757610002565b60408051600160a060020a039092168252519081900360200190f35b604080519115158252519081900360200190f35b005b600160009054906101000a9004600160a060020a0316600160a060020a031663f02ce71686600160a060020a031663b2b5f86f6000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f115610002575050506040518051906020015086866000604051602001526040518460e060020a0281526004018084600160a060020a031681526020018381526020018281526020019350505050602060405180830381600087803b156100025760325a03f11561000257505060405151925061012e9050565b600160a060020a038116600090815260026020526040902080546c010000000000000000000000008085020473ffffffffffffffffffffffffffffffffffffffff19909116179055505056`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, ejariRule common.Address) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, ejariRule)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// EjariRules is a free data retrieval call binding the contract method 0x46d71ab4.
//
// Solidity: function ejariRules() constant returns(address)
func (_Registry *RegistryCaller) EjariRules(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "ejariRules")
	return *ret0, err
}

// EjariRules is a free data retrieval call binding the contract method 0x46d71ab4.
//
// Solidity: function ejariRules() constant returns(address)
func (_Registry *RegistrySession) EjariRules() (common.Address, error) {
	return _Registry.Contract.EjariRules(&_Registry.CallOpts)
}

// EjariRules is a free data retrieval call binding the contract method 0x46d71ab4.
//
// Solidity: function ejariRules() constant returns(address)
func (_Registry *RegistryCallerSession) EjariRules() (common.Address, error) {
	return _Registry.Contract.EjariRules(&_Registry.CallOpts)
}

// Ownership is a free data retrieval call binding the contract method 0x27d6d6e0.
//
// Solidity: function ownership( address) constant returns(address)
func (_Registry *RegistryCaller) Ownership(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "ownership", arg0)
	return *ret0, err
}

// Ownership is a free data retrieval call binding the contract method 0x27d6d6e0.
//
// Solidity: function ownership( address) constant returns(address)
func (_Registry *RegistrySession) Ownership(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.Ownership(&_Registry.CallOpts, arg0)
}

// Ownership is a free data retrieval call binding the contract method 0x27d6d6e0.
//
// Solidity: function ownership( address) constant returns(address)
func (_Registry *RegistryCallerSession) Ownership(arg0 common.Address) (common.Address, error) {
	return _Registry.Contract.Ownership(&_Registry.CallOpts, arg0)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_Registry *RegistryCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "registrar")
	return *ret0, err
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_Registry *RegistrySession) Registrar() (common.Address, error) {
	return _Registry.Contract.Registrar(&_Registry.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_Registry *RegistryCallerSession) Registrar() (common.Address, error) {
	return _Registry.Contract.Registrar(&_Registry.CallOpts)
}

// AssignOwnership is a paid mutator transaction binding the contract method 0xbbd2b299.
//
// Solidity: function assignOwnership(owner address, property address) returns()
func (_Registry *RegistryTransactor) AssignOwnership(opts *bind.TransactOpts, owner common.Address, property common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "assignOwnership", owner, property)
}

// AssignOwnership is a paid mutator transaction binding the contract method 0xbbd2b299.
//
// Solidity: function assignOwnership(owner address, property address) returns()
func (_Registry *RegistrySession) AssignOwnership(owner common.Address, property common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AssignOwnership(&_Registry.TransactOpts, owner, property)
}

// AssignOwnership is a paid mutator transaction binding the contract method 0xbbd2b299.
//
// Solidity: function assignOwnership(owner address, property address) returns()
func (_Registry *RegistryTransactorSession) AssignOwnership(owner common.Address, property common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AssignOwnership(&_Registry.TransactOpts, owner, property)
}

// IsValidTenancy is a paid mutator transaction binding the contract method 0x3faa78b6.
//
// Solidity: function isValidTenancy(property address, rent uint256, incrementPercentage uint256) returns(bool)
func (_Registry *RegistryTransactor) IsValidTenancy(opts *bind.TransactOpts, property common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "isValidTenancy", property, rent, incrementPercentage)
}

// IsValidTenancy is a paid mutator transaction binding the contract method 0x3faa78b6.
//
// Solidity: function isValidTenancy(property address, rent uint256, incrementPercentage uint256) returns(bool)
func (_Registry *RegistrySession) IsValidTenancy(property common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.IsValidTenancy(&_Registry.TransactOpts, property, rent, incrementPercentage)
}

// IsValidTenancy is a paid mutator transaction binding the contract method 0x3faa78b6.
//
// Solidity: function isValidTenancy(property address, rent uint256, incrementPercentage uint256) returns(bool)
func (_Registry *RegistryTransactorSession) IsValidTenancy(property common.Address, rent *big.Int, incrementPercentage *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.IsValidTenancy(&_Registry.TransactOpts, property, rent, incrementPercentage)
}

// TenancyABI is the input ABI used to generate the binding from.
const TenancyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"terminate\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"property\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"acceptNegotiationOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptNegotiationTenant\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"rejectNegotiation\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tenant\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"security\",\"type\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"updateCondition\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"condition\",\"outputs\":[{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"security\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"security\",\"type\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"negotiate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"registry\",\"type\":\"address\"},{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"property\",\"type\":\"address\"},{\"name\":\"rent\",\"type\":\"uint256\"},{\"name\":\"security\",\"type\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"prospectiveTenant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Negotiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"prospectiveTenant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"RejectNegotiation\",\"type\":\"event\"}]"

// TenancyBin is the compiled bytecode used for deploying new contracts.
const TenancyBin = `0x606060405260405160e080610bc383396101406040819052915160805160a05160c0519351610100516101205160006101608190527f3faa78b6000000000000000000000000000000000000000000000000000000008952600160a060020a03808616610144526101648990526101848290529698959794969495939492939192891691633faa78b6916101a491602091606490829087803b156100025760325a03f115610002575050604051511590506101275760018054600160a060020a0319166c010000000000000000000000009788029790970496909617909555604080516080810182528481526020810184905290810182905260600185905260039290925560045560055550600655506007805460ff19169055610a978061012c6000396000f35b61000256606060405236156100985760e060020a60003504630c08bf88811461009d578063176fd3f01461019c578063199a620a146101b35780633ccfd60b146102325780637ed02af9146102625780638da5cb5b1461027d5780639bcc912314610294578063adf0779114610313578063b2e2c1c91461032a578063c19d93fb146103ba578063c5031331146103cb578063d5c69585146103e2575b610002565b61040f60006000600160009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f1156100025750506040515133600160a060020a039081169116149050806101915750600260009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f1156100025750506040515133600160a060020a0390811691161490505b151561046557610002565b3461000257610411600054600160a060020a031681565b346100025761040f6004356001546040805160006020918201819052825160e060020a630e2562d902815292519093600160a060020a031692630e2562d992600480830193919282900301818787803b156100025760325a03f1156100025750506040515133600160a060020a03908116911614905061051157610002565b346100025761040f600160a060020a03331660009081526008602052604081206005810154151561064257610002565b61040f600754600090819060ff16600214156106bc57610002565b3461000257610411600154600160a060020a031681565b346100025761040f6004356001546040805160006020918201819052825160e060020a630e2562d902815292519093600160a060020a031692630e2562d992600480830193919282900301818787803b156100025760325a03f1156100025750506040515133600160a060020a03908116911614905061083357610002565b3461000257610411600254600160a060020a031681565b346100025761040f600435602435604435606435600160009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f1156100025750506040515133600160a060020a03908116911614905061086e57610002565b346100025761042d60075460ff1681565b346100025761043f60035460045460055460065484565b346100025761040f60043560243560443560643560843560075460009060ff16600214156108b557610002565b005b60408051600160a060020a039092168252519081900360200190f35b60408051918252519081900360200190f35b604080519485526020850193909352838301919091526060830152519081900360800190f35b60065442111561009857600260009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f1156100025750506040518051600160a060020a0316915083830380156108fc02916000818181858888f19350505050151561050057610002565b50506007805460ff19166002179055565b60075460ff166002141561052457610002565b600160009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f1156100025750506040515133600160a060020a0390811691161490506105a057610002565b50600160a060020a0381166000908152600860205260409020600581015415156105c957610002565b805460098054600160a060020a031916606060020a600160a060020a039093168302929092049190911790556001810154600a556002810154600b556003810154600c556004810154600d556005810154600e5560060154600f805460ff191660f860020a60ff90931683029290920491909117905550565b60075460ff166002141561065557610002565b600160a060020a033381166000818152600860209081526040808320600581019390935560015481519485529094169083015282519094507f34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f929181900390910190a15050565b6009546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100025760325a03f1156100025750506040515133600160a060020a03908116911614905061073157610002565b600b54600a540134101561074457610002565b5050600a546001546040805160006020918201819052825160e060020a630e2562d90281529251606460028702049594600160a060020a031693630e2562d9936004808301949193928390030190829087803b156100025760325a03f1156100025750506040518051600160a060020a0316915083830380156108fc02916000818181858888f1935050505015156107db57610002565b6007805460ff1916600117905560095460028054606060020a600160a060020a03909316830292909204600160a060020a0319909216919091179055600a54600355600b54600455600c54600555600d546006555050565b60075460ff166002141561084657610002565b50600160a060020a03166000908152600860205260409020600601805460ff19166001179055565b60075460ff166002141561088157610002565b6040805160808101825285815260208101859052908101839052606001819052600393909355600491909155600555600655565b33600160a060020a031686600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f11561000257505060405151600160a060020a031691909114905061092457610002565b60075460ff16600114156109b65733600160a060020a0316600260009054906101000a9004600160a060020a0316600160a060020a0316630e2562d96000604051602001526040518160e060020a028152600401809050602060405180830381600087803b156100025760325a03f11561000257505060405151600160a060020a03169190911490506109b657610002565b50600160a060020a033381166000818152600860208181526040808420805482516080810184528c81528085018c90528084018b905260600189905260018083018d9055600283018c9055600383018b9055600483018a905560068301805460ff191690556005830181905595879052938352606060020a808d02819004600160a060020a0319958616178089168202919091049416939093178355925483519485529094169383019390935280517f5bfa20921b9a30cfbdde3c55ae53d357aa9fa3ec5ddb8e0700767a8e604ad8949281900390910190a150505050505056`

// DeployTenancy deploys a new Ethereum contract, binding an instance of Tenancy to it.
func DeployTenancy(auth *bind.TransactOpts, backend bind.ContractBackend, registry common.Address, person common.Address, property common.Address, rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (common.Address, *types.Transaction, *Tenancy, error) {
	parsed, err := abi.JSON(strings.NewReader(TenancyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TenancyBin), backend, registry, person, property, rent, security, start, end)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tenancy{TenancyCaller: TenancyCaller{contract: contract}, TenancyTransactor: TenancyTransactor{contract: contract}}, nil
}

// Tenancy is an auto generated Go binding around an Ethereum contract.
type Tenancy struct {
	TenancyCaller     // Read-only binding to the contract
	TenancyTransactor // Write-only binding to the contract
}

// TenancyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TenancyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TenancyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TenancyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TenancySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TenancySession struct {
	Contract     *Tenancy          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TenancyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TenancyCallerSession struct {
	Contract *TenancyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TenancyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TenancyTransactorSession struct {
	Contract     *TenancyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TenancyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TenancyRaw struct {
	Contract *Tenancy // Generic contract binding to access the raw methods on
}

// TenancyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TenancyCallerRaw struct {
	Contract *TenancyCaller // Generic read-only contract binding to access the raw methods on
}

// TenancyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TenancyTransactorRaw struct {
	Contract *TenancyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTenancy creates a new instance of Tenancy, bound to a specific deployed contract.
func NewTenancy(address common.Address, backend bind.ContractBackend) (*Tenancy, error) {
	contract, err := bindTenancy(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tenancy{TenancyCaller: TenancyCaller{contract: contract}, TenancyTransactor: TenancyTransactor{contract: contract}}, nil
}

// NewTenancyCaller creates a new read-only instance of Tenancy, bound to a specific deployed contract.
func NewTenancyCaller(address common.Address, caller bind.ContractCaller) (*TenancyCaller, error) {
	contract, err := bindTenancy(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TenancyCaller{contract: contract}, nil
}

// NewTenancyTransactor creates a new write-only instance of Tenancy, bound to a specific deployed contract.
func NewTenancyTransactor(address common.Address, transactor bind.ContractTransactor) (*TenancyTransactor, error) {
	contract, err := bindTenancy(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TenancyTransactor{contract: contract}, nil
}

// bindTenancy binds a generic wrapper to an already deployed contract.
func bindTenancy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TenancyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tenancy *TenancyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tenancy.Contract.TenancyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tenancy *TenancyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenancy.Contract.TenancyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tenancy *TenancyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tenancy.Contract.TenancyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tenancy *TenancyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tenancy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tenancy *TenancyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenancy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tenancy *TenancyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tenancy.Contract.contract.Transact(opts, method, params...)
}

// Condition is a free data retrieval call binding the contract method 0xc5031331.
//
// Solidity: function condition() constant returns(rent uint256, security uint256, startTime uint256, endTime uint256)
func (_Tenancy *TenancyCaller) Condition(opts *bind.CallOpts) (struct {
	Rent      *big.Int
	Security  *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	ret := new(struct {
		Rent      *big.Int
		Security  *big.Int
		StartTime *big.Int
		EndTime   *big.Int
	})
	out := ret
	err := _Tenancy.contract.Call(opts, out, "condition")
	return *ret, err
}

// Condition is a free data retrieval call binding the contract method 0xc5031331.
//
// Solidity: function condition() constant returns(rent uint256, security uint256, startTime uint256, endTime uint256)
func (_Tenancy *TenancySession) Condition() (struct {
	Rent      *big.Int
	Security  *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Tenancy.Contract.Condition(&_Tenancy.CallOpts)
}

// Condition is a free data retrieval call binding the contract method 0xc5031331.
//
// Solidity: function condition() constant returns(rent uint256, security uint256, startTime uint256, endTime uint256)
func (_Tenancy *TenancyCallerSession) Condition() (struct {
	Rent      *big.Int
	Security  *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Tenancy.Contract.Condition(&_Tenancy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tenancy *TenancyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tenancy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tenancy *TenancySession) Owner() (common.Address, error) {
	return _Tenancy.Contract.Owner(&_Tenancy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tenancy *TenancyCallerSession) Owner() (common.Address, error) {
	return _Tenancy.Contract.Owner(&_Tenancy.CallOpts)
}

// Property is a free data retrieval call binding the contract method 0x176fd3f0.
//
// Solidity: function property() constant returns(address)
func (_Tenancy *TenancyCaller) Property(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tenancy.contract.Call(opts, out, "property")
	return *ret0, err
}

// Property is a free data retrieval call binding the contract method 0x176fd3f0.
//
// Solidity: function property() constant returns(address)
func (_Tenancy *TenancySession) Property() (common.Address, error) {
	return _Tenancy.Contract.Property(&_Tenancy.CallOpts)
}

// Property is a free data retrieval call binding the contract method 0x176fd3f0.
//
// Solidity: function property() constant returns(address)
func (_Tenancy *TenancyCallerSession) Property() (common.Address, error) {
	return _Tenancy.Contract.Property(&_Tenancy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Tenancy *TenancyCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Tenancy.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Tenancy *TenancySession) State() (uint8, error) {
	return _Tenancy.Contract.State(&_Tenancy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Tenancy *TenancyCallerSession) State() (uint8, error) {
	return _Tenancy.Contract.State(&_Tenancy.CallOpts)
}

// Tenant is a free data retrieval call binding the contract method 0xadf07791.
//
// Solidity: function tenant() constant returns(address)
func (_Tenancy *TenancyCaller) Tenant(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tenancy.contract.Call(opts, out, "tenant")
	return *ret0, err
}

// Tenant is a free data retrieval call binding the contract method 0xadf07791.
//
// Solidity: function tenant() constant returns(address)
func (_Tenancy *TenancySession) Tenant() (common.Address, error) {
	return _Tenancy.Contract.Tenant(&_Tenancy.CallOpts)
}

// Tenant is a free data retrieval call binding the contract method 0xadf07791.
//
// Solidity: function tenant() constant returns(address)
func (_Tenancy *TenancyCallerSession) Tenant() (common.Address, error) {
	return _Tenancy.Contract.Tenant(&_Tenancy.CallOpts)
}

// AcceptNegotiationOwner is a paid mutator transaction binding the contract method 0x199a620a.
//
// Solidity: function acceptNegotiationOwner(prospectiveTenant address) returns()
func (_Tenancy *TenancyTransactor) AcceptNegotiationOwner(opts *bind.TransactOpts, prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "acceptNegotiationOwner", prospectiveTenant)
}

// AcceptNegotiationOwner is a paid mutator transaction binding the contract method 0x199a620a.
//
// Solidity: function acceptNegotiationOwner(prospectiveTenant address) returns()
func (_Tenancy *TenancySession) AcceptNegotiationOwner(prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.Contract.AcceptNegotiationOwner(&_Tenancy.TransactOpts, prospectiveTenant)
}

// AcceptNegotiationOwner is a paid mutator transaction binding the contract method 0x199a620a.
//
// Solidity: function acceptNegotiationOwner(prospectiveTenant address) returns()
func (_Tenancy *TenancyTransactorSession) AcceptNegotiationOwner(prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.Contract.AcceptNegotiationOwner(&_Tenancy.TransactOpts, prospectiveTenant)
}

// AcceptNegotiationTenant is a paid mutator transaction binding the contract method 0x7ed02af9.
//
// Solidity: function acceptNegotiationTenant() returns()
func (_Tenancy *TenancyTransactor) AcceptNegotiationTenant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "acceptNegotiationTenant")
}

// AcceptNegotiationTenant is a paid mutator transaction binding the contract method 0x7ed02af9.
//
// Solidity: function acceptNegotiationTenant() returns()
func (_Tenancy *TenancySession) AcceptNegotiationTenant() (*types.Transaction, error) {
	return _Tenancy.Contract.AcceptNegotiationTenant(&_Tenancy.TransactOpts)
}

// AcceptNegotiationTenant is a paid mutator transaction binding the contract method 0x7ed02af9.
//
// Solidity: function acceptNegotiationTenant() returns()
func (_Tenancy *TenancyTransactorSession) AcceptNegotiationTenant() (*types.Transaction, error) {
	return _Tenancy.Contract.AcceptNegotiationTenant(&_Tenancy.TransactOpts)
}

// Negotiate is a paid mutator transaction binding the contract method 0xd5c69585.
//
// Solidity: function negotiate(person address, rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancyTransactor) Negotiate(opts *bind.TransactOpts, person common.Address, rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "negotiate", person, rent, security, start, end)
}

// Negotiate is a paid mutator transaction binding the contract method 0xd5c69585.
//
// Solidity: function negotiate(person address, rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancySession) Negotiate(person common.Address, rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.Contract.Negotiate(&_Tenancy.TransactOpts, person, rent, security, start, end)
}

// Negotiate is a paid mutator transaction binding the contract method 0xd5c69585.
//
// Solidity: function negotiate(person address, rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancyTransactorSession) Negotiate(person common.Address, rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.Contract.Negotiate(&_Tenancy.TransactOpts, person, rent, security, start, end)
}

// RejectNegotiation is a paid mutator transaction binding the contract method 0x9bcc9123.
//
// Solidity: function rejectNegotiation(prospectiveTenant address) returns()
func (_Tenancy *TenancyTransactor) RejectNegotiation(opts *bind.TransactOpts, prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "rejectNegotiation", prospectiveTenant)
}

// RejectNegotiation is a paid mutator transaction binding the contract method 0x9bcc9123.
//
// Solidity: function rejectNegotiation(prospectiveTenant address) returns()
func (_Tenancy *TenancySession) RejectNegotiation(prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.Contract.RejectNegotiation(&_Tenancy.TransactOpts, prospectiveTenant)
}

// RejectNegotiation is a paid mutator transaction binding the contract method 0x9bcc9123.
//
// Solidity: function rejectNegotiation(prospectiveTenant address) returns()
func (_Tenancy *TenancyTransactorSession) RejectNegotiation(prospectiveTenant common.Address) (*types.Transaction, error) {
	return _Tenancy.Contract.RejectNegotiation(&_Tenancy.TransactOpts, prospectiveTenant)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Tenancy *TenancyTransactor) Terminate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Tenancy *TenancySession) Terminate() (*types.Transaction, error) {
	return _Tenancy.Contract.Terminate(&_Tenancy.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Tenancy *TenancyTransactorSession) Terminate() (*types.Transaction, error) {
	return _Tenancy.Contract.Terminate(&_Tenancy.TransactOpts)
}

// UpdateCondition is a paid mutator transaction binding the contract method 0xb2e2c1c9.
//
// Solidity: function updateCondition(rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancyTransactor) UpdateCondition(opts *bind.TransactOpts, rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "updateCondition", rent, security, start, end)
}

// UpdateCondition is a paid mutator transaction binding the contract method 0xb2e2c1c9.
//
// Solidity: function updateCondition(rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancySession) UpdateCondition(rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.Contract.UpdateCondition(&_Tenancy.TransactOpts, rent, security, start, end)
}

// UpdateCondition is a paid mutator transaction binding the contract method 0xb2e2c1c9.
//
// Solidity: function updateCondition(rent uint256, security uint256, start uint256, end uint256) returns()
func (_Tenancy *TenancyTransactorSession) UpdateCondition(rent *big.Int, security *big.Int, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Tenancy.Contract.UpdateCondition(&_Tenancy.TransactOpts, rent, security, start, end)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tenancy *TenancyTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tenancy *TenancySession) Withdraw() (*types.Transaction, error) {
	return _Tenancy.Contract.Withdraw(&_Tenancy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tenancy *TenancyTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Tenancy.Contract.Withdraw(&_Tenancy.TransactOpts)
}
