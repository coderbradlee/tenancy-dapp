// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PropertyABI is the input ABI used to generate the binding from.
const PropertyABI = "[{\"inputs\":[{\"name\":\"s\",\"type\":\"string\"},{\"name\":\"p\",\"type\":\"string\"}],\"type\":\"constructor\"}]"

// PropertyBin is the compiled bytecode used for deploying new contracts.
const PropertyBin = `0x606060405260405161015e38038061015e83398101604052805160805190820191018160006000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009957805160ff19168380011785555b506100c99291505b808211156101225760008155600101610074565b505050506008806101566000396000f35b8280016001018555821561006c579182015b8281111561006c5782518260005055916020019190600101906100ab565b50508060016000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012657805160ff19168380011785555b50610088929150610074565b5090565b82800160010185558215610116579182015b82811115610116578251826000505591602001919060010190610138566060604052600256`

// DeployProperty deploys a new Ethereum contract, binding an instance of Property to it.
func DeployProperty(auth *bind.TransactOpts, backend bind.ContractBackend, s string, p string) (common.Address, *types.Transaction, *Property, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PropertyBin), backend, s, p)
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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownership\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"property\",\"type\":\"address\"}],\"name\":\"isPropertyOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"property\",\"type\":\"address\"}],\"name\":\"assignOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"type\":\"constructor\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `0x6060604052600080546c0100000000000000000000000033810204600160a060020a031990911617905561014b806100376000396000f3606060405260e060020a600035046327d6d6e0811461003f5780632b20e3971461006557806386db57a71461007c578063bbd2b299146100b8575b610002565b34610002576100e1600435600160205260009081526040902054600160a060020a031681565b34610002576100e1600054600160a060020a031681565b3461000257600160a060020a03600435811660009081526001602052604090205481163390911614604080519115158252519081900360200190f35b34610002576100fd60043560243560005433600160a060020a039081169116146100ff57610002565b60408051600160a060020a039092168252519081900360200190f35b005b600160a060020a038116600090815260016020526040902080546c010000000000000000000000008085020473ffffffffffffffffffffffffffffffffffffffff19909116179055505056`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
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

// IsPropertyOwner is a paid mutator transaction binding the contract method 0x86db57a7.
//
// Solidity: function isPropertyOwner(property address) returns(bool)
func (_Registry *RegistryTransactor) IsPropertyOwner(opts *bind.TransactOpts, property common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "isPropertyOwner", property)
}

// IsPropertyOwner is a paid mutator transaction binding the contract method 0x86db57a7.
//
// Solidity: function isPropertyOwner(property address) returns(bool)
func (_Registry *RegistrySession) IsPropertyOwner(property common.Address) (*types.Transaction, error) {
	return _Registry.Contract.IsPropertyOwner(&_Registry.TransactOpts, property)
}

// IsPropertyOwner is a paid mutator transaction binding the contract method 0x86db57a7.
//
// Solidity: function isPropertyOwner(property address) returns(bool)
func (_Registry *RegistryTransactorSession) IsPropertyOwner(property common.Address) (*types.Transaction, error) {
	return _Registry.Contract.IsPropertyOwner(&_Registry.TransactOpts, property)
}

// TenancyABI is the input ABI used to generate the binding from.
const TenancyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"acceptNegotiationOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"negotiationId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"rejectNegotiation\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tenant\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"negotiate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"registry\",\"type\":\"address\"},{\"name\":\"property\",\"type\":\"address\"},{\"name\":\"rent\",\"type\":\"uint32\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"prospectiveTenant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Negotiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"prospectiveTenant\",\"type\":\"address\"}],\"name\":\"RejectNegotiation\",\"type\":\"event\"}]"

// TenancyBin is the compiled bytecode used for deploying new contracts.
const TenancyBin = `0x606060408190526005805463ffffffff19166001179055808061065881395060c06040819052905160805160a051600060e08190527f86db57a7000000000000000000000000000000000000000000000000000000008552600160a060020a0380841660c452939492939192918516916386db57a79160e491602091602490829087803b1560025760325a03f11560025750506040515115905060d057505060018054600160a060020a0319166c010000000000000000000000003381020417905550610584806100d46000396000f35b600256606060405236156100565760e060020a6000350463199a620a811461005b5780632a302669146100845780638da5cb5b146100a75780639bcc9123146100be578063adf0779114610188578063f6b94cb41461019f575b610002565b346100025761026760043560015460009033600160a060020a0390811691161461029e57610002565b346100025761026960043560036020526000908152604090205463ffffffff1681565b3461000257610282600154600160a060020a031681565b3461000257610267600435600160a060020a03811660009081526003602052604081205463ffffffff169081111561005657600160046000508263ffffffff168154811015610002579060005260206000209001600050805460f860020a9283029290920460c060020a0260c060020a60ff021990921691909117905560015460408051600160a060020a039283168152918416602083015280517f20760a4e390a305962dd840b66c77f19509f4aaec0ae5ee1e2e0b9ab892fd04e9281900390910190a16103c8565b3461000257610282600254600160a060020a031681565b3461000257610267600435600160a060020a03331660009081526003602052604081205463ffffffff16908111156103cc578160046000508263ffffffff1681548110156100025760009182526020822001805460e060020a9384029390930460a060020a0260a060020a63ffffffff0219909316929092179091556004805463ffffffff841690811015610002579060005260206000209001600050805460f860020a9283029290920460c060020a0260c060020a60ff021990921691909117905561052e565b005b6040805163ffffffff9092168252519081900360200190f35b60408051600160a060020a039092168252519081900360200190f35b50600160a060020a0381166000908152600360205260409020546004805463ffffffff9092169182908110156100025790600052602060002090016000508054600680546c01000000000000000000000000600160a060020a0390931683029290920473ffffffffffffffffffffffffffffffffffffffff1990921691909117808255825460e060020a63ffffffff60a060020a92839004168102040260a060020a63ffffffff021990911617808255825460f860020a60ff60c060020a928390048116820282900490920260c060020a60ff021990931692909217808455845460c860020a90819004831684028490040260c860020a60ff021990911617808455935460d060020a908190049091168202919091040260d060020a60ff02199092169190911790555b5050565b6005805433600160a060020a0381166000908152600360209081526040808320805463ffffffff191660e060020a63ffffffff978816810204179055805160a081018252938452908301879052820181905260608201819052608082015291546004805490929190911690811015610002576000918252602091829020835191018054928401516040850151606086015160809096015173ffffffffffffffffffffffffffffffffffffffff199095166c01000000000000000000000000948502949094049390931760a060020a63ffffffff02191660a060020a60e060020a928302839004021760c060020a60ff02191660c060020a60f860020a948502859004021760c860020a60ff02191660c860020a958402849004959095029490941760d060020a60ff02191660d060020a93830292909204929092021790556005805463ffffffff19811663ffffffff9091166001018302929092049190911790555b60015460408051600160a060020a033381168252909216602083015263ffffffff841682820152517fd0d0546c1578257629800deb0b45ef2b35d76057132b69ac2be057062fc8e3b89181900360600190a1505056`

// DeployTenancy deploys a new Ethereum contract, binding an instance of Tenancy to it.
func DeployTenancy(auth *bind.TransactOpts, backend bind.ContractBackend, registry common.Address, property common.Address, rent uint32) (common.Address, *types.Transaction, *Tenancy, error) {
	parsed, err := abi.JSON(strings.NewReader(TenancyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TenancyBin), backend, registry, property, rent)
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

// NegotiationId is a free data retrieval call binding the contract method 0x2a302669.
//
// Solidity: function negotiationId( address) constant returns(uint32)
func (_Tenancy *TenancyCaller) NegotiationId(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Tenancy.contract.Call(opts, out, "negotiationId", arg0)
	return *ret0, err
}

// NegotiationId is a free data retrieval call binding the contract method 0x2a302669.
//
// Solidity: function negotiationId( address) constant returns(uint32)
func (_Tenancy *TenancySession) NegotiationId(arg0 common.Address) (uint32, error) {
	return _Tenancy.Contract.NegotiationId(&_Tenancy.CallOpts, arg0)
}

// NegotiationId is a free data retrieval call binding the contract method 0x2a302669.
//
// Solidity: function negotiationId( address) constant returns(uint32)
func (_Tenancy *TenancyCallerSession) NegotiationId(arg0 common.Address) (uint32, error) {
	return _Tenancy.Contract.NegotiationId(&_Tenancy.CallOpts, arg0)
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

// Negotiate is a paid mutator transaction binding the contract method 0xf6b94cb4.
//
// Solidity: function negotiate(amount uint32) returns()
func (_Tenancy *TenancyTransactor) Negotiate(opts *bind.TransactOpts, amount uint32) (*types.Transaction, error) {
	return _Tenancy.contract.Transact(opts, "negotiate", amount)
}

// Negotiate is a paid mutator transaction binding the contract method 0xf6b94cb4.
//
// Solidity: function negotiate(amount uint32) returns()
func (_Tenancy *TenancySession) Negotiate(amount uint32) (*types.Transaction, error) {
	return _Tenancy.Contract.Negotiate(&_Tenancy.TransactOpts, amount)
}

// Negotiate is a paid mutator transaction binding the contract method 0xf6b94cb4.
//
// Solidity: function negotiate(amount uint32) returns()
func (_Tenancy *TenancyTransactorSession) Negotiate(amount uint32) (*types.Transaction, error) {
	return _Tenancy.Contract.Negotiate(&_Tenancy.TransactOpts, amount)
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
