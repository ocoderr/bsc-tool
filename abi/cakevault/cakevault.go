// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cakevault

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CakevaultABI is the input ABI used to generate the binding from.
const CakevaultABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_receiptToken\",\"type\":\"address\"},{\"internalType\":\"contractIMasterChef\",\"name\":\"_masterchef\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callFee\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CALL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PERFORMANCE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateHarvestCakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateTotalPendingCakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"inCaseTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHarvestedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterchef\",\"outputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_callFee\",\"type\":\"uint256\"}],\"name\":\"setCallFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFee\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFeePeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cakeAtLastUserAction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUserActionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Cakevault is an auto generated Go binding around an Ethereum contract.
type Cakevault struct {
	CakevaultCaller     // Read-only binding to the contract
	CakevaultTransactor // Write-only binding to the contract
	CakevaultFilterer   // Log filterer for contract events
}

// CakevaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CakevaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakevaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CakevaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakevaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CakevaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakevaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CakevaultSession struct {
	Contract     *Cakevault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakevaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CakevaultCallerSession struct {
	Contract *CakevaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CakevaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CakevaultTransactorSession struct {
	Contract     *CakevaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CakevaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CakevaultRaw struct {
	Contract *Cakevault // Generic contract binding to access the raw methods on
}

// CakevaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CakevaultCallerRaw struct {
	Contract *CakevaultCaller // Generic read-only contract binding to access the raw methods on
}

// CakevaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CakevaultTransactorRaw struct {
	Contract *CakevaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCakevault creates a new instance of Cakevault, bound to a specific deployed contract.
func NewCakevault(address common.Address, backend bind.ContractBackend) (*Cakevault, error) {
	contract, err := bindCakevault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cakevault{CakevaultCaller: CakevaultCaller{contract: contract}, CakevaultTransactor: CakevaultTransactor{contract: contract}, CakevaultFilterer: CakevaultFilterer{contract: contract}}, nil
}

// NewCakevaultCaller creates a new read-only instance of Cakevault, bound to a specific deployed contract.
func NewCakevaultCaller(address common.Address, caller bind.ContractCaller) (*CakevaultCaller, error) {
	contract, err := bindCakevault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CakevaultCaller{contract: contract}, nil
}

// NewCakevaultTransactor creates a new write-only instance of Cakevault, bound to a specific deployed contract.
func NewCakevaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CakevaultTransactor, error) {
	contract, err := bindCakevault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CakevaultTransactor{contract: contract}, nil
}

// NewCakevaultFilterer creates a new log filterer instance of Cakevault, bound to a specific deployed contract.
func NewCakevaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CakevaultFilterer, error) {
	contract, err := bindCakevault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CakevaultFilterer{contract: contract}, nil
}

// bindCakevault binds a generic wrapper to an already deployed contract.
func bindCakevault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CakevaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cakevault *CakevaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cakevault.Contract.CakevaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cakevault *CakevaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.Contract.CakevaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cakevault *CakevaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cakevault.Contract.CakevaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cakevault *CakevaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cakevault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cakevault *CakevaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cakevault *CakevaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cakevault.Contract.contract.Transact(opts, method, params...)
}

// MAXCALLFEE is a free data retrieval call binding the contract method 0x2ad5a53f.
//
// Solidity: function MAX_CALL_FEE() view returns(uint256)
func (_Cakevault *CakevaultCaller) MAXCALLFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "MAX_CALL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCALLFEE is a free data retrieval call binding the contract method 0x2ad5a53f.
//
// Solidity: function MAX_CALL_FEE() view returns(uint256)
func (_Cakevault *CakevaultSession) MAXCALLFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXCALLFEE(&_Cakevault.CallOpts)
}

// MAXCALLFEE is a free data retrieval call binding the contract method 0x2ad5a53f.
//
// Solidity: function MAX_CALL_FEE() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) MAXCALLFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXCALLFEE(&_Cakevault.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_Cakevault *CakevaultCaller) MAXPERFORMANCEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "MAX_PERFORMANCE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_Cakevault *CakevaultSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXPERFORMANCEFEE(&_Cakevault.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXPERFORMANCEFEE(&_Cakevault.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_Cakevault *CakevaultCaller) MAXWITHDRAWFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "MAX_WITHDRAW_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_Cakevault *CakevaultSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXWITHDRAWFEE(&_Cakevault.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _Cakevault.Contract.MAXWITHDRAWFEE(&_Cakevault.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_Cakevault *CakevaultCaller) MAXWITHDRAWFEEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "MAX_WITHDRAW_FEE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_Cakevault *CakevaultSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _Cakevault.Contract.MAXWITHDRAWFEEPERIOD(&_Cakevault.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _Cakevault.Contract.MAXWITHDRAWFEEPERIOD(&_Cakevault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cakevault *CakevaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cakevault *CakevaultSession) Admin() (common.Address, error) {
	return _Cakevault.Contract.Admin(&_Cakevault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cakevault *CakevaultCallerSession) Admin() (common.Address, error) {
	return _Cakevault.Contract.Admin(&_Cakevault.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Cakevault *CakevaultCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Cakevault *CakevaultSession) Available() (*big.Int, error) {
	return _Cakevault.Contract.Available(&_Cakevault.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) Available() (*big.Int, error) {
	return _Cakevault.Contract.Available(&_Cakevault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Cakevault *CakevaultCaller) BalanceOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Cakevault *CakevaultSession) BalanceOf() (*big.Int, error) {
	return _Cakevault.Contract.BalanceOf(&_Cakevault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) BalanceOf() (*big.Int, error) {
	return _Cakevault.Contract.BalanceOf(&_Cakevault.CallOpts)
}

// CalculateHarvestCakeRewards is a free data retrieval call binding the contract method 0x9d72596b.
//
// Solidity: function calculateHarvestCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultCaller) CalculateHarvestCakeRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "calculateHarvestCakeRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateHarvestCakeRewards is a free data retrieval call binding the contract method 0x9d72596b.
//
// Solidity: function calculateHarvestCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultSession) CalculateHarvestCakeRewards() (*big.Int, error) {
	return _Cakevault.Contract.CalculateHarvestCakeRewards(&_Cakevault.CallOpts)
}

// CalculateHarvestCakeRewards is a free data retrieval call binding the contract method 0x9d72596b.
//
// Solidity: function calculateHarvestCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) CalculateHarvestCakeRewards() (*big.Int, error) {
	return _Cakevault.Contract.CalculateHarvestCakeRewards(&_Cakevault.CallOpts)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultCaller) CalculateTotalPendingCakeRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "calculateTotalPendingCakeRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _Cakevault.Contract.CalculateTotalPendingCakeRewards(&_Cakevault.CallOpts)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _Cakevault.Contract.CalculateTotalPendingCakeRewards(&_Cakevault.CallOpts)
}

// CallFee is a free data retrieval call binding the contract method 0x90321e1a.
//
// Solidity: function callFee() view returns(uint256)
func (_Cakevault *CakevaultCaller) CallFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "callFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallFee is a free data retrieval call binding the contract method 0x90321e1a.
//
// Solidity: function callFee() view returns(uint256)
func (_Cakevault *CakevaultSession) CallFee() (*big.Int, error) {
	return _Cakevault.Contract.CallFee(&_Cakevault.CallOpts)
}

// CallFee is a free data retrieval call binding the contract method 0x90321e1a.
//
// Solidity: function callFee() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) CallFee() (*big.Int, error) {
	return _Cakevault.Contract.CallFee(&_Cakevault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Cakevault *CakevaultCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Cakevault *CakevaultSession) GetPricePerFullShare() (*big.Int, error) {
	return _Cakevault.Contract.GetPricePerFullShare(&_Cakevault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Cakevault.Contract.GetPricePerFullShare(&_Cakevault.CallOpts)
}

// LastHarvestedTime is a free data retrieval call binding the contract method 0xb60f0531.
//
// Solidity: function lastHarvestedTime() view returns(uint256)
func (_Cakevault *CakevaultCaller) LastHarvestedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "lastHarvestedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHarvestedTime is a free data retrieval call binding the contract method 0xb60f0531.
//
// Solidity: function lastHarvestedTime() view returns(uint256)
func (_Cakevault *CakevaultSession) LastHarvestedTime() (*big.Int, error) {
	return _Cakevault.Contract.LastHarvestedTime(&_Cakevault.CallOpts)
}

// LastHarvestedTime is a free data retrieval call binding the contract method 0xb60f0531.
//
// Solidity: function lastHarvestedTime() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) LastHarvestedTime() (*big.Int, error) {
	return _Cakevault.Contract.LastHarvestedTime(&_Cakevault.CallOpts)
}

// Masterchef is a free data retrieval call binding the contract method 0xfb1db278.
//
// Solidity: function masterchef() view returns(address)
func (_Cakevault *CakevaultCaller) Masterchef(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "masterchef")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Masterchef is a free data retrieval call binding the contract method 0xfb1db278.
//
// Solidity: function masterchef() view returns(address)
func (_Cakevault *CakevaultSession) Masterchef() (common.Address, error) {
	return _Cakevault.Contract.Masterchef(&_Cakevault.CallOpts)
}

// Masterchef is a free data retrieval call binding the contract method 0xfb1db278.
//
// Solidity: function masterchef() view returns(address)
func (_Cakevault *CakevaultCallerSession) Masterchef() (common.Address, error) {
	return _Cakevault.Contract.Masterchef(&_Cakevault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cakevault *CakevaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cakevault *CakevaultSession) Owner() (common.Address, error) {
	return _Cakevault.Contract.Owner(&_Cakevault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cakevault *CakevaultCallerSession) Owner() (common.Address, error) {
	return _Cakevault.Contract.Owner(&_Cakevault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cakevault *CakevaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cakevault *CakevaultSession) Paused() (bool, error) {
	return _Cakevault.Contract.Paused(&_Cakevault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cakevault *CakevaultCallerSession) Paused() (bool, error) {
	return _Cakevault.Contract.Paused(&_Cakevault.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Cakevault *CakevaultCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Cakevault *CakevaultSession) PerformanceFee() (*big.Int, error) {
	return _Cakevault.Contract.PerformanceFee(&_Cakevault.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) PerformanceFee() (*big.Int, error) {
	return _Cakevault.Contract.PerformanceFee(&_Cakevault.CallOpts)
}

// ReceiptToken is a free data retrieval call binding the contract method 0xec78e832.
//
// Solidity: function receiptToken() view returns(address)
func (_Cakevault *CakevaultCaller) ReceiptToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "receiptToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiptToken is a free data retrieval call binding the contract method 0xec78e832.
//
// Solidity: function receiptToken() view returns(address)
func (_Cakevault *CakevaultSession) ReceiptToken() (common.Address, error) {
	return _Cakevault.Contract.ReceiptToken(&_Cakevault.CallOpts)
}

// ReceiptToken is a free data retrieval call binding the contract method 0xec78e832.
//
// Solidity: function receiptToken() view returns(address)
func (_Cakevault *CakevaultCallerSession) ReceiptToken() (common.Address, error) {
	return _Cakevault.Contract.ReceiptToken(&_Cakevault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cakevault *CakevaultCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cakevault *CakevaultSession) Token() (common.Address, error) {
	return _Cakevault.Contract.Token(&_Cakevault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cakevault *CakevaultCallerSession) Token() (common.Address, error) {
	return _Cakevault.Contract.Token(&_Cakevault.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cakevault *CakevaultCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cakevault *CakevaultSession) TotalShares() (*big.Int, error) {
	return _Cakevault.Contract.TotalShares(&_Cakevault.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) TotalShares() (*big.Int, error) {
	return _Cakevault.Contract.TotalShares(&_Cakevault.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Cakevault *CakevaultCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Cakevault *CakevaultSession) Treasury() (common.Address, error) {
	return _Cakevault.Contract.Treasury(&_Cakevault.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Cakevault *CakevaultCallerSession) Treasury() (common.Address, error) {
	return _Cakevault.Contract.Treasury(&_Cakevault.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime)
func (_Cakevault *CakevaultCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
}, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Shares               *big.Int
		LastDepositedTime    *big.Int
		CakeAtLastUserAction *big.Int
		LastUserActionTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastDepositedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CakeAtLastUserAction = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastUserActionTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime)
func (_Cakevault *CakevaultSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
}, error) {
	return _Cakevault.Contract.UserInfo(&_Cakevault.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime)
func (_Cakevault *CakevaultCallerSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
}, error) {
	return _Cakevault.Contract.UserInfo(&_Cakevault.CallOpts, arg0)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_Cakevault *CakevaultCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_Cakevault *CakevaultSession) WithdrawFee() (*big.Int, error) {
	return _Cakevault.Contract.WithdrawFee(&_Cakevault.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) WithdrawFee() (*big.Int, error) {
	return _Cakevault.Contract.WithdrawFee(&_Cakevault.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_Cakevault *CakevaultCaller) WithdrawFeePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cakevault.contract.Call(opts, &out, "withdrawFeePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_Cakevault *CakevaultSession) WithdrawFeePeriod() (*big.Int, error) {
	return _Cakevault.Contract.WithdrawFeePeriod(&_Cakevault.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_Cakevault *CakevaultCallerSession) WithdrawFeePeriod() (*big.Int, error) {
	return _Cakevault.Contract.WithdrawFeePeriod(&_Cakevault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Cakevault *CakevaultTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Cakevault *CakevaultSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.Deposit(&_Cakevault.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Cakevault *CakevaultTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.Deposit(&_Cakevault.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Cakevault *CakevaultTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Cakevault *CakevaultSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Cakevault.Contract.EmergencyWithdraw(&_Cakevault.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Cakevault *CakevaultTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Cakevault.Contract.EmergencyWithdraw(&_Cakevault.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Cakevault *CakevaultTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Cakevault *CakevaultSession) Harvest() (*types.Transaction, error) {
	return _Cakevault.Contract.Harvest(&_Cakevault.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Cakevault *CakevaultTransactorSession) Harvest() (*types.Transaction, error) {
	return _Cakevault.Contract.Harvest(&_Cakevault.TransactOpts)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Cakevault *CakevaultTransactor) InCaseTokensGetStuck(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "inCaseTokensGetStuck", _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Cakevault *CakevaultSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.InCaseTokensGetStuck(&_Cakevault.TransactOpts, _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Cakevault *CakevaultTransactorSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.InCaseTokensGetStuck(&_Cakevault.TransactOpts, _token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cakevault *CakevaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cakevault *CakevaultSession) Pause() (*types.Transaction, error) {
	return _Cakevault.Contract.Pause(&_Cakevault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cakevault *CakevaultTransactorSession) Pause() (*types.Transaction, error) {
	return _Cakevault.Contract.Pause(&_Cakevault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cakevault *CakevaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cakevault *CakevaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cakevault.Contract.RenounceOwnership(&_Cakevault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cakevault *CakevaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cakevault.Contract.RenounceOwnership(&_Cakevault.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Cakevault *CakevaultTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Cakevault *CakevaultSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.SetAdmin(&_Cakevault.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Cakevault *CakevaultTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.SetAdmin(&_Cakevault.TransactOpts, _admin)
}

// SetCallFee is a paid mutator transaction binding the contract method 0x26465826.
//
// Solidity: function setCallFee(uint256 _callFee) returns()
func (_Cakevault *CakevaultTransactor) SetCallFee(opts *bind.TransactOpts, _callFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setCallFee", _callFee)
}

// SetCallFee is a paid mutator transaction binding the contract method 0x26465826.
//
// Solidity: function setCallFee(uint256 _callFee) returns()
func (_Cakevault *CakevaultSession) SetCallFee(_callFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetCallFee(&_Cakevault.TransactOpts, _callFee)
}

// SetCallFee is a paid mutator transaction binding the contract method 0x26465826.
//
// Solidity: function setCallFee(uint256 _callFee) returns()
func (_Cakevault *CakevaultTransactorSession) SetCallFee(_callFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetCallFee(&_Cakevault.TransactOpts, _callFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_Cakevault *CakevaultTransactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_Cakevault *CakevaultSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetPerformanceFee(&_Cakevault.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_Cakevault *CakevaultTransactorSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetPerformanceFee(&_Cakevault.TransactOpts, _performanceFee)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Cakevault *CakevaultTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Cakevault *CakevaultSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.SetTreasury(&_Cakevault.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Cakevault *CakevaultTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.SetTreasury(&_Cakevault.TransactOpts, _treasury)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_Cakevault *CakevaultTransactor) SetWithdrawFee(opts *bind.TransactOpts, _withdrawFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setWithdrawFee", _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_Cakevault *CakevaultSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetWithdrawFee(&_Cakevault.TransactOpts, _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_Cakevault *CakevaultTransactorSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetWithdrawFee(&_Cakevault.TransactOpts, _withdrawFee)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_Cakevault *CakevaultTransactor) SetWithdrawFeePeriod(opts *bind.TransactOpts, _withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "setWithdrawFeePeriod", _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_Cakevault *CakevaultSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetWithdrawFeePeriod(&_Cakevault.TransactOpts, _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_Cakevault *CakevaultTransactorSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.SetWithdrawFeePeriod(&_Cakevault.TransactOpts, _withdrawFeePeriod)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cakevault *CakevaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cakevault *CakevaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.TransferOwnership(&_Cakevault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cakevault *CakevaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cakevault.Contract.TransferOwnership(&_Cakevault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cakevault *CakevaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cakevault *CakevaultSession) Unpause() (*types.Transaction, error) {
	return _Cakevault.Contract.Unpause(&_Cakevault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cakevault *CakevaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _Cakevault.Contract.Unpause(&_Cakevault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Cakevault *CakevaultTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Cakevault *CakevaultSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.Withdraw(&_Cakevault.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Cakevault *CakevaultTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Cakevault.Contract.Withdraw(&_Cakevault.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Cakevault *CakevaultTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cakevault.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Cakevault *CakevaultSession) WithdrawAll() (*types.Transaction, error) {
	return _Cakevault.Contract.WithdrawAll(&_Cakevault.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Cakevault *CakevaultTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Cakevault.Contract.WithdrawAll(&_Cakevault.TransactOpts)
}

// CakevaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Cakevault contract.
type CakevaultDepositIterator struct {
	Event *CakevaultDeposit // Event containing the contract specifics and raw log

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
func (it *CakevaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultDeposit)
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
		it.Event = new(CakevaultDeposit)
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
func (it *CakevaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultDeposit represents a Deposit event raised by the Cakevault contract.
type CakevaultDeposit struct {
	Sender            common.Address
	Amount            *big.Int
	Shares            *big.Int
	LastDepositedTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 lastDepositedTime)
func (_Cakevault *CakevaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*CakevaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakevaultDepositIterator{contract: _Cakevault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 lastDepositedTime)
func (_Cakevault *CakevaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CakevaultDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultDeposit)
				if err := _Cakevault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 lastDepositedTime)
func (_Cakevault *CakevaultFilterer) ParseDeposit(log types.Log) (*CakevaultDeposit, error) {
	event := new(CakevaultDeposit)
	if err := _Cakevault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Cakevault contract.
type CakevaultHarvestIterator struct {
	Event *CakevaultHarvest // Event containing the contract specifics and raw log

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
func (it *CakevaultHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultHarvest)
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
		it.Event = new(CakevaultHarvest)
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
func (it *CakevaultHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultHarvest represents a Harvest event raised by the Cakevault contract.
type CakevaultHarvest struct {
	Sender         common.Address
	PerformanceFee *big.Int
	CallFee        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed sender, uint256 performanceFee, uint256 callFee)
func (_Cakevault *CakevaultFilterer) FilterHarvest(opts *bind.FilterOpts, sender []common.Address) (*CakevaultHarvestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakevaultHarvestIterator{contract: _Cakevault.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed sender, uint256 performanceFee, uint256 callFee)
func (_Cakevault *CakevaultFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *CakevaultHarvest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultHarvest)
				if err := _Cakevault.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed sender, uint256 performanceFee, uint256 callFee)
func (_Cakevault *CakevaultFilterer) ParseHarvest(log types.Log) (*CakevaultHarvest, error) {
	event := new(CakevaultHarvest)
	if err := _Cakevault.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cakevault contract.
type CakevaultOwnershipTransferredIterator struct {
	Event *CakevaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CakevaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultOwnershipTransferred)
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
		it.Event = new(CakevaultOwnershipTransferred)
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
func (it *CakevaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultOwnershipTransferred represents a OwnershipTransferred event raised by the Cakevault contract.
type CakevaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cakevault *CakevaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CakevaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CakevaultOwnershipTransferredIterator{contract: _Cakevault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cakevault *CakevaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CakevaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultOwnershipTransferred)
				if err := _Cakevault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cakevault *CakevaultFilterer) ParseOwnershipTransferred(log types.Log) (*CakevaultOwnershipTransferred, error) {
	event := new(CakevaultOwnershipTransferred)
	if err := _Cakevault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Cakevault contract.
type CakevaultPauseIterator struct {
	Event *CakevaultPause // Event containing the contract specifics and raw log

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
func (it *CakevaultPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultPause)
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
		it.Event = new(CakevaultPause)
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
func (it *CakevaultPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultPause represents a Pause event raised by the Cakevault contract.
type CakevaultPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Cakevault *CakevaultFilterer) FilterPause(opts *bind.FilterOpts) (*CakevaultPauseIterator, error) {

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CakevaultPauseIterator{contract: _Cakevault.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Cakevault *CakevaultFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CakevaultPause) (event.Subscription, error) {

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultPause)
				if err := _Cakevault.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Cakevault *CakevaultFilterer) ParsePause(log types.Log) (*CakevaultPause, error) {
	event := new(CakevaultPause)
	if err := _Cakevault.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Cakevault contract.
type CakevaultPausedIterator struct {
	Event *CakevaultPaused // Event containing the contract specifics and raw log

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
func (it *CakevaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultPaused)
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
		it.Event = new(CakevaultPaused)
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
func (it *CakevaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultPaused represents a Paused event raised by the Cakevault contract.
type CakevaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cakevault *CakevaultFilterer) FilterPaused(opts *bind.FilterOpts) (*CakevaultPausedIterator, error) {

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CakevaultPausedIterator{contract: _Cakevault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cakevault *CakevaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CakevaultPaused) (event.Subscription, error) {

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultPaused)
				if err := _Cakevault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cakevault *CakevaultFilterer) ParsePaused(log types.Log) (*CakevaultPaused, error) {
	event := new(CakevaultPaused)
	if err := _Cakevault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Cakevault contract.
type CakevaultUnpauseIterator struct {
	Event *CakevaultUnpause // Event containing the contract specifics and raw log

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
func (it *CakevaultUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultUnpause)
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
		it.Event = new(CakevaultUnpause)
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
func (it *CakevaultUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultUnpause represents a Unpause event raised by the Cakevault contract.
type CakevaultUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Cakevault *CakevaultFilterer) FilterUnpause(opts *bind.FilterOpts) (*CakevaultUnpauseIterator, error) {

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CakevaultUnpauseIterator{contract: _Cakevault.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Cakevault *CakevaultFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CakevaultUnpause) (event.Subscription, error) {

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultUnpause)
				if err := _Cakevault.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Cakevault *CakevaultFilterer) ParseUnpause(log types.Log) (*CakevaultUnpause, error) {
	event := new(CakevaultUnpause)
	if err := _Cakevault.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Cakevault contract.
type CakevaultUnpausedIterator struct {
	Event *CakevaultUnpaused // Event containing the contract specifics and raw log

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
func (it *CakevaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultUnpaused)
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
		it.Event = new(CakevaultUnpaused)
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
func (it *CakevaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultUnpaused represents a Unpaused event raised by the Cakevault contract.
type CakevaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cakevault *CakevaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CakevaultUnpausedIterator, error) {

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CakevaultUnpausedIterator{contract: _Cakevault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cakevault *CakevaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CakevaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultUnpaused)
				if err := _Cakevault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cakevault *CakevaultFilterer) ParseUnpaused(log types.Log) (*CakevaultUnpaused, error) {
	event := new(CakevaultUnpaused)
	if err := _Cakevault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakevaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Cakevault contract.
type CakevaultWithdrawIterator struct {
	Event *CakevaultWithdraw // Event containing the contract specifics and raw log

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
func (it *CakevaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakevaultWithdraw)
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
		it.Event = new(CakevaultWithdraw)
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
func (it *CakevaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakevaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakevaultWithdraw represents a Withdraw event raised by the Cakevault contract.
type CakevaultWithdraw struct {
	Sender common.Address
	Amount *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_Cakevault *CakevaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*CakevaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakevaultWithdrawIterator{contract: _Cakevault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_Cakevault *CakevaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CakevaultWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cakevault.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakevaultWithdraw)
				if err := _Cakevault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_Cakevault *CakevaultFilterer) ParseWithdraw(log types.Log) (*CakevaultWithdraw, error) {
	event := new(CakevaultWithdraw)
	if err := _Cakevault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
