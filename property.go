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

// EjariRulesABI is the input ABI used to generate the binding from.
const EjariRulesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"latitude\",\"type\":\"string\"},{\"name\":\"longitude\",\"type\":\"string\"},{\"name\":\"incrementPercentage\",\"type\":\"uint256\"},{\"name\":\"maxRent\",\"type\":\"uint256\"}],\"name\":\"addEjariRule\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"latitude\",\"type\":\"string\"},{\"name\":\"longitude\",\"type\":\"string\"},{\"name\":\"oldRent\",\"type\":\"uint256\"},{\"name\":\"newRent\",\"type\":\"uint256\"}],\"name\":\"isValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"type\":\"constructor\"}]"

// EjariRulesBin is the compiled bytecode used for deploying new contracts.
const EjariRulesBin = `0x6060604052600080546c0100000000000000000000000033810204600160a060020a0319909116179055610320806100376000396000f3606060405260e060020a60003504635610486181146100295780636ed40f9d146100e4575b610002565b346100025761022f6004808035906020019082018035906020019191908080601f01602080910402602001604051908101604052809392919081815260200183838082843750506040805160208835808b0135601f81018390048302840183019094528383529799986044989297509190910194509092508291508401838280828437509496505093359350506064359150506000543373ffffffffffffffffffffffffffffffffffffffff90811691161461024557610002565b34610002576102316004808035906020019082018035906020019191908080601f01602080910402602001604051908101604052809392919081815260200183838082843750506040805160208835808b0135601f81018390048302840183019094528383529799986044989297509190910194509092508291508401838280828437509496505093359350506064359150506000600060006001600050600060028989600060405160200152604051808380519060200190808383829060006004602084601f0104600302600f01f1509050018280519060200190808383829060006004602084601f0104600302600f01f150905001925050506020604051808303816000866161da5a03f11561000257505060408051518252602082019290925201600020805460018201549193506064908101870204915084111561030f5760009250610305565b005b604080519115158252519081900360200190f35b604060405190810160405280838152602001828152602001506001600050600060028787600060405160200152604051808380519060200190808383829060006004602084601f0104600302600f01f1509050018280519060200190808383829060006004602084601f0104600302600f01f150905001925050506020604051808303816000866161da5a03f11561000257505060408051518252602082810193909352016000208251815591015160019091015550505050565b600192505b5050949350505050565b80841115610300576000925061030556`

// DeployEjariRules deploys a new Ethereum contract, binding an instance of EjariRules to it.
func DeployEjariRules(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EjariRules, error) {
	parsed, err := abi.JSON(strings.NewReader(EjariRulesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EjariRulesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EjariRules{EjariRulesCaller: EjariRulesCaller{contract: contract}, EjariRulesTransactor: EjariRulesTransactor{contract: contract}}, nil
}

// EjariRules is an auto generated Go binding around an Ethereum contract.
type EjariRules struct {
	EjariRulesCaller     // Read-only binding to the contract
	EjariRulesTransactor // Write-only binding to the contract
}

// EjariRulesCaller is an auto generated read-only Go binding around an Ethereum contract.
type EjariRulesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EjariRulesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EjariRulesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EjariRulesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EjariRulesSession struct {
	Contract     *EjariRules       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EjariRulesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EjariRulesCallerSession struct {
	Contract *EjariRulesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EjariRulesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EjariRulesTransactorSession struct {
	Contract     *EjariRulesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EjariRulesRaw is an auto generated low-level Go binding around an Ethereum contract.
type EjariRulesRaw struct {
	Contract *EjariRules // Generic contract binding to access the raw methods on
}

// EjariRulesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EjariRulesCallerRaw struct {
	Contract *EjariRulesCaller // Generic read-only contract binding to access the raw methods on
}

// EjariRulesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EjariRulesTransactorRaw struct {
	Contract *EjariRulesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEjariRules creates a new instance of EjariRules, bound to a specific deployed contract.
func NewEjariRules(address common.Address, backend bind.ContractBackend) (*EjariRules, error) {
	contract, err := bindEjariRules(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EjariRules{EjariRulesCaller: EjariRulesCaller{contract: contract}, EjariRulesTransactor: EjariRulesTransactor{contract: contract}}, nil
}

// NewEjariRulesCaller creates a new read-only instance of EjariRules, bound to a specific deployed contract.
func NewEjariRulesCaller(address common.Address, caller bind.ContractCaller) (*EjariRulesCaller, error) {
	contract, err := bindEjariRules(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EjariRulesCaller{contract: contract}, nil
}

// NewEjariRulesTransactor creates a new write-only instance of EjariRules, bound to a specific deployed contract.
func NewEjariRulesTransactor(address common.Address, transactor bind.ContractTransactor) (*EjariRulesTransactor, error) {
	contract, err := bindEjariRules(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EjariRulesTransactor{contract: contract}, nil
}

// bindEjariRules binds a generic wrapper to an already deployed contract.
func bindEjariRules(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EjariRulesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EjariRules *EjariRulesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EjariRules.Contract.EjariRulesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EjariRules *EjariRulesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EjariRules.Contract.EjariRulesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EjariRules *EjariRulesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EjariRules.Contract.EjariRulesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EjariRules *EjariRulesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EjariRules.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EjariRules *EjariRulesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EjariRules.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EjariRules *EjariRulesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EjariRules.Contract.contract.Transact(opts, method, params...)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0x56104861.
//
// Solidity: function addEjariRule(latitude string, longitude string, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRules *EjariRulesTransactor) AddEjariRule(opts *bind.TransactOpts, latitude string, longitude string, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.contract.Transact(opts, "addEjariRule", latitude, longitude, incrementPercentage, maxRent)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0x56104861.
//
// Solidity: function addEjariRule(latitude string, longitude string, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRules *EjariRulesSession) AddEjariRule(latitude string, longitude string, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.Contract.AddEjariRule(&_EjariRules.TransactOpts, latitude, longitude, incrementPercentage, maxRent)
}

// AddEjariRule is a paid mutator transaction binding the contract method 0x56104861.
//
// Solidity: function addEjariRule(latitude string, longitude string, incrementPercentage uint256, maxRent uint256) returns()
func (_EjariRules *EjariRulesTransactorSession) AddEjariRule(latitude string, longitude string, incrementPercentage *big.Int, maxRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.Contract.AddEjariRule(&_EjariRules.TransactOpts, latitude, longitude, incrementPercentage, maxRent)
}

// IsValid is a paid mutator transaction binding the contract method 0x6ed40f9d.
//
// Solidity: function isValid(latitude string, longitude string, oldRent uint256, newRent uint256) returns(bool)
func (_EjariRules *EjariRulesTransactor) IsValid(opts *bind.TransactOpts, latitude string, longitude string, oldRent *big.Int, newRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.contract.Transact(opts, "isValid", latitude, longitude, oldRent, newRent)
}

// IsValid is a paid mutator transaction binding the contract method 0x6ed40f9d.
//
// Solidity: function isValid(latitude string, longitude string, oldRent uint256, newRent uint256) returns(bool)
func (_EjariRules *EjariRulesSession) IsValid(latitude string, longitude string, oldRent *big.Int, newRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.Contract.IsValid(&_EjariRules.TransactOpts, latitude, longitude, oldRent, newRent)
}

// IsValid is a paid mutator transaction binding the contract method 0x6ed40f9d.
//
// Solidity: function isValid(latitude string, longitude string, oldRent uint256, newRent uint256) returns(bool)
func (_EjariRules *EjariRulesTransactorSession) IsValid(latitude string, longitude string, oldRent *big.Int, newRent *big.Int) (*types.Transaction, error) {
	return _EjariRules.Contract.IsValid(&_EjariRules.TransactOpts, latitude, longitude, oldRent, newRent)
}

// PropertyABI is the input ABI used to generate the binding from.
const PropertyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"propertyRating\",\"outputs\":[{\"name\":\"totalRatings\",\"type\":\"uint256\"},{\"name\":\"numberOfRatings\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_government\",\"type\":\"address\"}],\"name\":\"setGovernment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rateProperty\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerRating\",\"outputs\":[{\"name\":\"totalRatings\",\"type\":\"uint256\"},{\"name\":\"numberOfRatings\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"validate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rateTenant\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deduction\",\"type\":\"uint256\"}],\"name\":\"terminate\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"rateOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rent\",\"type\":\"uint256\"}],\"name\":\"updateRent\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_latitude\",\"type\":\"string\"},{\"name\":\"_longitude\",\"type\":\"string\"},{\"name\":\"_rent\",\"type\":\"uint256\"},{\"name\":\"_security\",\"type\":\"uint256\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"government\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"government\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Validated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tenant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Interested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tenant\",\"type\":\"address\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tenant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Payment\",\"type\":\"event\"}]"

// PropertyBin is the compiled bytecode used for deploying new contracts.
const PropertyBin = `0x606060408190526000805460a060020a60ff0219600160a060020a031990911673429d61dc95cac25a24feffcf7db98f76d6ab3796171690556106523881900390819083398101604052805160805160a05160c051928401939190910191600380546c0100000000000000000000000033810204600160a060020a0319909116179055835160018054600082905290916020601f60026000198587161561010002019094169390930483018190047fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690810193909189019083901061010757805160ff19168380011785555b506101379291505b8082111561019057600081556001016100f3565b828001600101855582156100eb579182015b828111156100eb578251826000505591602001919060010190610119565b50508260026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019457805160ff19168380011785555b506101c49291506100f3565b5090565b82800160010185558215610184579182015b828111156101845782518260005055916020019190600101906101a6565b50506004829055600581905560035460005460408051600160a060020a03938416815292909116602083015280517f0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e559281900390910190a1505050506104248061022e6000396000f3606060405236156100825760e060020a600035046306096b6981146100875780630f4b1cfc1461009857806339c0ea6e146100c55780636758a1a0146100eb5780636901f668146100fc5780636e9c4d0a1461011f5780637a828b281461014857806388968b9414610169578063eb9220ab1461018f578063ef48eee6146101b5575b610002565b34610002576101d1600c54600d5482565b346100025760008054600160a060020a0319166c010000000000000000000000006004358102041790555b005b34610002576100c360043560065433600160a060020a039081169116146101ea57610002565b34610002576101d1600a54600b5482565b34610002576100c360005433600160a060020a039081169116146101fe57610002565b34610002576100c360043560035460009033600160a060020a0390811691161461029c57610002565b6100c360043560035433600160a060020a039081169116146102cd57610002565b34610002576100c360043560065433600160a060020a0390811691161461035457610002565b34610002576100c360043560035433600160a060020a0390811691161461036857610002565b6100c36004356024356005546004540134101561037c57610002565b6040805192835260208301919091528051918290030190f35b600c805482019055600d8054600101905550565b60005433600160a060020a0390811691161461021957610002565b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179081905560035460408051600160a060020a03938416815291909216602082015281517fd6c06c253ccd3a2c9b0c848536c88ce68a82ec904b839430912363c1f5be321c929181900390910190a1565b50600654600160a060020a031660009081526009602052604090208054820181556001818101805490910190555050565b600654600554604051600160a060020a03909216919083900380156108fc02916000818181858888f19350505050801561032c5750600354604051600160a060020a039091169082156108fc029083906000818181858888f193505050505b151561033757610002565b60068054600160a060020a03191690556000600781905560085550565b600a805482019055600b8054600101905550565b60085442101561037757610002565b600455565b600354600454604051600160a060020a039092169181156108fc0291906000818181858888f1935050505015156103b257610002565b60068054600160a060020a0319166c0100000000000000000000000033810204179081905560035460408051600160a060020a03938416815292909116602083015280517f36c4359614c37ef31afa649625f064e2ca57f0faa6dbc8ca2548356e23d944bc9281900390910190a1505056`

// DeployProperty deploys a new Ethereum contract, binding an instance of Property to it.
func DeployProperty(auth *bind.TransactOpts, backend bind.ContractBackend, _latitude string, _longitude string, _rent *big.Int, _security *big.Int) (common.Address, *types.Transaction, *Property, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PropertyBin), backend, _latitude, _longitude, _rent, _security)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Property{PropertyCaller: PropertyCaller{contract: contract}, PropertyTransactor: PropertyTransactor{contract: contract}}, nil
}

// Property is an auto generated Go binding around an Ethereum contract.
type Property struct {
	PropertyCaller     // Read-only binding to the contract
	PropertyTransactor // Write-only binding to the contract
}

// PropertyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertySession struct {
	Contract     *Property         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyCallerSession struct {
	Contract *PropertyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PropertyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyTransactorSession struct {
	Contract     *PropertyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PropertyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyRaw struct {
	Contract *Property // Generic contract binding to access the raw methods on
}

// PropertyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyCallerRaw struct {
	Contract *PropertyCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyTransactorRaw struct {
	Contract *PropertyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProperty creates a new instance of Property, bound to a specific deployed contract.
func NewProperty(address common.Address, backend bind.ContractBackend) (*Property, error) {
	contract, err := bindProperty(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Property{PropertyCaller: PropertyCaller{contract: contract}, PropertyTransactor: PropertyTransactor{contract: contract}}, nil
}

// NewPropertyCaller creates a new read-only instance of Property, bound to a specific deployed contract.
func NewPropertyCaller(address common.Address, caller bind.ContractCaller) (*PropertyCaller, error) {
	contract, err := bindProperty(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyCaller{contract: contract}, nil
}

// NewPropertyTransactor creates a new write-only instance of Property, bound to a specific deployed contract.
func NewPropertyTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyTransactor, error) {
	contract, err := bindProperty(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PropertyTransactor{contract: contract}, nil
}

// bindProperty binds a generic wrapper to an already deployed contract.
func bindProperty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Property *PropertyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Property.Contract.PropertyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Property *PropertyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.Contract.PropertyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Property *PropertyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Property.Contract.PropertyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Property *PropertyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Property.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Property *PropertyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Property *PropertyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Property.Contract.contract.Transact(opts, method, params...)
}

// OwnerRating is a free data retrieval call binding the contract method 0x6758a1a0.
//
// Solidity: function ownerRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertyCaller) OwnerRating(opts *bind.CallOpts) (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	ret := new(struct {
		TotalRatings    *big.Int
		NumberOfRatings *big.Int
	})
	out := ret
	err := _Property.contract.Call(opts, out, "ownerRating")
	return *ret, err
}

// OwnerRating is a free data retrieval call binding the contract method 0x6758a1a0.
//
// Solidity: function ownerRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertySession) OwnerRating() (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	return _Property.Contract.OwnerRating(&_Property.CallOpts)
}

// OwnerRating is a free data retrieval call binding the contract method 0x6758a1a0.
//
// Solidity: function ownerRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertyCallerSession) OwnerRating() (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	return _Property.Contract.OwnerRating(&_Property.CallOpts)
}

// PropertyRating is a free data retrieval call binding the contract method 0x06096b69.
//
// Solidity: function propertyRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertyCaller) PropertyRating(opts *bind.CallOpts) (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	ret := new(struct {
		TotalRatings    *big.Int
		NumberOfRatings *big.Int
	})
	out := ret
	err := _Property.contract.Call(opts, out, "propertyRating")
	return *ret, err
}

// PropertyRating is a free data retrieval call binding the contract method 0x06096b69.
//
// Solidity: function propertyRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertySession) PropertyRating() (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	return _Property.Contract.PropertyRating(&_Property.CallOpts)
}

// PropertyRating is a free data retrieval call binding the contract method 0x06096b69.
//
// Solidity: function propertyRating() constant returns(totalRatings uint256, numberOfRatings uint256)
func (_Property *PropertyCallerSession) PropertyRating() (struct {
	TotalRatings    *big.Int
	NumberOfRatings *big.Int
}, error) {
	return _Property.Contract.PropertyRating(&_Property.CallOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xef48eee6.
//
// Solidity: function pay(startTime uint256, endTime uint256) returns()
func (_Property *PropertyTransactor) Pay(opts *bind.TransactOpts, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "pay", startTime, endTime)
}

// Pay is a paid mutator transaction binding the contract method 0xef48eee6.
//
// Solidity: function pay(startTime uint256, endTime uint256) returns()
func (_Property *PropertySession) Pay(startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Pay(&_Property.TransactOpts, startTime, endTime)
}

// Pay is a paid mutator transaction binding the contract method 0xef48eee6.
//
// Solidity: function pay(startTime uint256, endTime uint256) returns()
func (_Property *PropertyTransactorSession) Pay(startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Pay(&_Property.TransactOpts, startTime, endTime)
}

// RateOwner is a paid mutator transaction binding the contract method 0x88968b94.
//
// Solidity: function rateOwner(rating uint256) returns()
func (_Property *PropertyTransactor) RateOwner(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "rateOwner", rating)
}

// RateOwner is a paid mutator transaction binding the contract method 0x88968b94.
//
// Solidity: function rateOwner(rating uint256) returns()
func (_Property *PropertySession) RateOwner(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateOwner(&_Property.TransactOpts, rating)
}

// RateOwner is a paid mutator transaction binding the contract method 0x88968b94.
//
// Solidity: function rateOwner(rating uint256) returns()
func (_Property *PropertyTransactorSession) RateOwner(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateOwner(&_Property.TransactOpts, rating)
}

// RateProperty is a paid mutator transaction binding the contract method 0x39c0ea6e.
//
// Solidity: function rateProperty(rating uint256) returns()
func (_Property *PropertyTransactor) RateProperty(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "rateProperty", rating)
}

// RateProperty is a paid mutator transaction binding the contract method 0x39c0ea6e.
//
// Solidity: function rateProperty(rating uint256) returns()
func (_Property *PropertySession) RateProperty(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateProperty(&_Property.TransactOpts, rating)
}

// RateProperty is a paid mutator transaction binding the contract method 0x39c0ea6e.
//
// Solidity: function rateProperty(rating uint256) returns()
func (_Property *PropertyTransactorSession) RateProperty(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateProperty(&_Property.TransactOpts, rating)
}

// RateTenant is a paid mutator transaction binding the contract method 0x6e9c4d0a.
//
// Solidity: function rateTenant(rating uint256) returns()
func (_Property *PropertyTransactor) RateTenant(opts *bind.TransactOpts, rating *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "rateTenant", rating)
}

// RateTenant is a paid mutator transaction binding the contract method 0x6e9c4d0a.
//
// Solidity: function rateTenant(rating uint256) returns()
func (_Property *PropertySession) RateTenant(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateTenant(&_Property.TransactOpts, rating)
}

// RateTenant is a paid mutator transaction binding the contract method 0x6e9c4d0a.
//
// Solidity: function rateTenant(rating uint256) returns()
func (_Property *PropertyTransactorSession) RateTenant(rating *big.Int) (*types.Transaction, error) {
	return _Property.Contract.RateTenant(&_Property.TransactOpts, rating)
}

// SetGovernment is a paid mutator transaction binding the contract method 0x0f4b1cfc.
//
// Solidity: function setGovernment(_government address) returns()
func (_Property *PropertyTransactor) SetGovernment(opts *bind.TransactOpts, _government common.Address) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "setGovernment", _government)
}

// SetGovernment is a paid mutator transaction binding the contract method 0x0f4b1cfc.
//
// Solidity: function setGovernment(_government address) returns()
func (_Property *PropertySession) SetGovernment(_government common.Address) (*types.Transaction, error) {
	return _Property.Contract.SetGovernment(&_Property.TransactOpts, _government)
}

// SetGovernment is a paid mutator transaction binding the contract method 0x0f4b1cfc.
//
// Solidity: function setGovernment(_government address) returns()
func (_Property *PropertyTransactorSession) SetGovernment(_government common.Address) (*types.Transaction, error) {
	return _Property.Contract.SetGovernment(&_Property.TransactOpts, _government)
}

// Terminate is a paid mutator transaction binding the contract method 0x7a828b28.
//
// Solidity: function terminate(deduction uint256) returns()
func (_Property *PropertyTransactor) Terminate(opts *bind.TransactOpts, deduction *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "terminate", deduction)
}

// Terminate is a paid mutator transaction binding the contract method 0x7a828b28.
//
// Solidity: function terminate(deduction uint256) returns()
func (_Property *PropertySession) Terminate(deduction *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Terminate(&_Property.TransactOpts, deduction)
}

// Terminate is a paid mutator transaction binding the contract method 0x7a828b28.
//
// Solidity: function terminate(deduction uint256) returns()
func (_Property *PropertyTransactorSession) Terminate(deduction *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Terminate(&_Property.TransactOpts, deduction)
}

// UpdateRent is a paid mutator transaction binding the contract method 0xeb9220ab.
//
// Solidity: function updateRent(_rent uint256) returns()
func (_Property *PropertyTransactor) UpdateRent(opts *bind.TransactOpts, _rent *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "updateRent", _rent)
}

// UpdateRent is a paid mutator transaction binding the contract method 0xeb9220ab.
//
// Solidity: function updateRent(_rent uint256) returns()
func (_Property *PropertySession) UpdateRent(_rent *big.Int) (*types.Transaction, error) {
	return _Property.Contract.UpdateRent(&_Property.TransactOpts, _rent)
}

// UpdateRent is a paid mutator transaction binding the contract method 0xeb9220ab.
//
// Solidity: function updateRent(_rent uint256) returns()
func (_Property *PropertyTransactorSession) UpdateRent(_rent *big.Int) (*types.Transaction, error) {
	return _Property.Contract.UpdateRent(&_Property.TransactOpts, _rent)
}

// Validate is a paid mutator transaction binding the contract method 0x6901f668.
//
// Solidity: function validate() returns()
func (_Property *PropertyTransactor) Validate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "validate")
}

// Validate is a paid mutator transaction binding the contract method 0x6901f668.
//
// Solidity: function validate() returns()
func (_Property *PropertySession) Validate() (*types.Transaction, error) {
	return _Property.Contract.Validate(&_Property.TransactOpts)
}

// Validate is a paid mutator transaction binding the contract method 0x6901f668.
//
// Solidity: function validate() returns()
func (_Property *PropertyTransactorSession) Validate() (*types.Transaction, error) {
	return _Property.Contract.Validate(&_Property.TransactOpts)
}
