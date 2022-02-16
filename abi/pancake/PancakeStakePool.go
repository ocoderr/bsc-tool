// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancake

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

// PancakeABI is the input ABI used to generate the binding from.
const PancakeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenRecovered\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AdminTokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"NewPoolLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewStartAndEndBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RewardsStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SMART_CHEF_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasUserLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLimitPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverWrongTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_hasUserLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"updatePoolLimitPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pancake is an auto generated Go binding around an Ethereum contract.
type Pancake struct {
	PancakeCaller     // Read-only binding to the contract
	PancakeTransactor // Write-only binding to the contract
	PancakeFilterer   // Log filterer for contract events
}

// PancakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSession struct {
	Contract     *Pancake          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeCallerSession struct {
	Contract *PancakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PancakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeTransactorSession struct {
	Contract     *PancakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PancakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeRaw struct {
	Contract *Pancake // Generic contract binding to access the raw methods on
}

// PancakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeCallerRaw struct {
	Contract *PancakeCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeTransactorRaw struct {
	Contract *PancakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancake creates a new instance of Pancake, bound to a specific deployed contract.
func NewPancake(address common.Address, backend bind.ContractBackend) (*Pancake, error) {
	contract, err := bindPancake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancake{PancakeCaller: PancakeCaller{contract: contract}, PancakeTransactor: PancakeTransactor{contract: contract}, PancakeFilterer: PancakeFilterer{contract: contract}}, nil
}

// NewPancakeCaller creates a new read-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeCaller(address common.Address, caller bind.ContractCaller) (*PancakeCaller, error) {
	contract, err := bindPancake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeCaller{contract: contract}, nil
}

// NewPancakeTransactor creates a new write-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeTransactor, error) {
	contract, err := bindPancake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeTransactor{contract: contract}, nil
}

// NewPancakeFilterer creates a new log filterer instance of Pancake, bound to a specific deployed contract.
func NewPancakeFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeFilterer, error) {
	contract, err := bindPancake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeFilterer{contract: contract}, nil
}

// bindPancake binds a generic wrapper to an already deployed contract.
func bindPancake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.PancakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pancake *PancakeCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pancake *PancakeSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Pancake.Contract.PRECISIONFACTOR(&_Pancake.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pancake *PancakeCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Pancake.Contract.PRECISIONFACTOR(&_Pancake.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pancake *PancakeCaller) SMARTCHEFFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "SMART_CHEF_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pancake *PancakeSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _Pancake.Contract.SMARTCHEFFACTORY(&_Pancake.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pancake *PancakeCallerSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _Pancake.Contract.SMARTCHEFFACTORY(&_Pancake.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pancake *PancakeCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pancake *PancakeSession) AccTokenPerShare() (*big.Int, error) {
	return _Pancake.Contract.AccTokenPerShare(&_Pancake.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pancake *PancakeCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _Pancake.Contract.AccTokenPerShare(&_Pancake.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pancake *PancakeCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pancake *PancakeSession) BonusEndBlock() (*big.Int, error) {
	return _Pancake.Contract.BonusEndBlock(&_Pancake.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pancake *PancakeCallerSession) BonusEndBlock() (*big.Int, error) {
	return _Pancake.Contract.BonusEndBlock(&_Pancake.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pancake *PancakeCaller) HasUserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "hasUserLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pancake *PancakeSession) HasUserLimit() (bool, error) {
	return _Pancake.Contract.HasUserLimit(&_Pancake.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pancake *PancakeCallerSession) HasUserLimit() (bool, error) {
	return _Pancake.Contract.HasUserLimit(&_Pancake.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pancake *PancakeCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pancake *PancakeSession) IsInitialized() (bool, error) {
	return _Pancake.Contract.IsInitialized(&_Pancake.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pancake *PancakeCallerSession) IsInitialized() (bool, error) {
	return _Pancake.Contract.IsInitialized(&_Pancake.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pancake *PancakeCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pancake *PancakeSession) LastRewardBlock() (*big.Int, error) {
	return _Pancake.Contract.LastRewardBlock(&_Pancake.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pancake *PancakeCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Pancake.Contract.LastRewardBlock(&_Pancake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancake *PancakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancake *PancakeSession) Owner() (common.Address, error) {
	return _Pancake.Contract.Owner(&_Pancake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancake *PancakeCallerSession) Owner() (common.Address, error) {
	return _Pancake.Contract.Owner(&_Pancake.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pancake *PancakeCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "pendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pancake *PancakeSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Pancake.Contract.PendingReward(&_Pancake.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pancake *PancakeCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Pancake.Contract.PendingReward(&_Pancake.CallOpts, _user)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pancake *PancakeCaller) PoolLimitPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "poolLimitPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pancake *PancakeSession) PoolLimitPerUser() (*big.Int, error) {
	return _Pancake.Contract.PoolLimitPerUser(&_Pancake.CallOpts)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pancake *PancakeCallerSession) PoolLimitPerUser() (*big.Int, error) {
	return _Pancake.Contract.PoolLimitPerUser(&_Pancake.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pancake *PancakeCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pancake *PancakeSession) RewardPerBlock() (*big.Int, error) {
	return _Pancake.Contract.RewardPerBlock(&_Pancake.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pancake *PancakeCallerSession) RewardPerBlock() (*big.Int, error) {
	return _Pancake.Contract.RewardPerBlock(&_Pancake.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pancake *PancakeCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pancake *PancakeSession) RewardToken() (common.Address, error) {
	return _Pancake.Contract.RewardToken(&_Pancake.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pancake *PancakeCallerSession) RewardToken() (common.Address, error) {
	return _Pancake.Contract.RewardToken(&_Pancake.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pancake *PancakeCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pancake *PancakeSession) StakedToken() (common.Address, error) {
	return _Pancake.Contract.StakedToken(&_Pancake.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pancake *PancakeCallerSession) StakedToken() (common.Address, error) {
	return _Pancake.Contract.StakedToken(&_Pancake.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pancake *PancakeCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pancake *PancakeSession) StartBlock() (*big.Int, error) {
	return _Pancake.Contract.StartBlock(&_Pancake.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pancake *PancakeCallerSession) StartBlock() (*big.Int, error) {
	return _Pancake.Contract.StartBlock(&_Pancake.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Pancake *PancakeCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Pancake *PancakeSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Pancake.Contract.UserInfo(&_Pancake.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Pancake *PancakeCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Pancake.Contract.UserInfo(&_Pancake.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pancake *PancakeTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pancake *PancakeSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Deposit(&_Pancake.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pancake *PancakeTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Deposit(&_Pancake.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pancake *PancakeTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pancake *PancakeSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.EmergencyRewardWithdraw(&_Pancake.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pancake *PancakeTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.EmergencyRewardWithdraw(&_Pancake.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pancake *PancakeTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pancake *PancakeSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Pancake.Contract.EmergencyWithdraw(&_Pancake.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pancake *PancakeTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Pancake.Contract.EmergencyWithdraw(&_Pancake.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pancake *PancakeTransactor) Initialize(opts *bind.TransactOpts, _stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "initialize", _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pancake *PancakeSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Initialize(&_Pancake.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pancake *PancakeTransactorSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Initialize(&_Pancake.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pancake *PancakeTransactor) RecoverWrongTokens(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "recoverWrongTokens", _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pancake *PancakeSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.RecoverWrongTokens(&_Pancake.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pancake *PancakeTransactorSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.RecoverWrongTokens(&_Pancake.TransactOpts, _tokenAddress, _tokenAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pancake *PancakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pancake *PancakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pancake.Contract.RenounceOwnership(&_Pancake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pancake *PancakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pancake.Contract.RenounceOwnership(&_Pancake.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pancake *PancakeTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pancake *PancakeSession) StopReward() (*types.Transaction, error) {
	return _Pancake.Contract.StopReward(&_Pancake.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pancake *PancakeTransactorSession) StopReward() (*types.Transaction, error) {
	return _Pancake.Contract.StopReward(&_Pancake.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancake *PancakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancake *PancakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.TransferOwnership(&_Pancake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancake *PancakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.TransferOwnership(&_Pancake.TransactOpts, newOwner)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pancake *PancakeTransactor) UpdatePoolLimitPerUser(opts *bind.TransactOpts, _hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "updatePoolLimitPerUser", _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pancake *PancakeSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdatePoolLimitPerUser(&_Pancake.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pancake *PancakeTransactorSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdatePoolLimitPerUser(&_Pancake.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pancake *PancakeTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "updateRewardPerBlock", _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pancake *PancakeSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdateRewardPerBlock(&_Pancake.TransactOpts, _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pancake *PancakeTransactorSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdateRewardPerBlock(&_Pancake.TransactOpts, _rewardPerBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pancake *PancakeTransactor) UpdateStartAndEndBlocks(opts *bind.TransactOpts, _startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "updateStartAndEndBlocks", _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pancake *PancakeSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdateStartAndEndBlocks(&_Pancake.TransactOpts, _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pancake *PancakeTransactorSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.UpdateStartAndEndBlocks(&_Pancake.TransactOpts, _startBlock, _bonusEndBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pancake *PancakeTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pancake *PancakeSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Withdraw(&_Pancake.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pancake *PancakeTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Withdraw(&_Pancake.TransactOpts, _amount)
}

// PancakeAdminTokenRecoveryIterator is returned from FilterAdminTokenRecovery and is used to iterate over the raw logs and unpacked data for AdminTokenRecovery events raised by the Pancake contract.
type PancakeAdminTokenRecoveryIterator struct {
	Event *PancakeAdminTokenRecovery // Event containing the contract specifics and raw log

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
func (it *PancakeAdminTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeAdminTokenRecovery)
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
		it.Event = new(PancakeAdminTokenRecovery)
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
func (it *PancakeAdminTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeAdminTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeAdminTokenRecovery represents a AdminTokenRecovery event raised by the Pancake contract.
type PancakeAdminTokenRecovery struct {
	TokenRecovered common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdminTokenRecovery is a free log retrieval operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pancake *PancakeFilterer) FilterAdminTokenRecovery(opts *bind.FilterOpts) (*PancakeAdminTokenRecoveryIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return &PancakeAdminTokenRecoveryIterator{contract: _Pancake.contract, event: "AdminTokenRecovery", logs: logs, sub: sub}, nil
}

// WatchAdminTokenRecovery is a free log subscription operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pancake *PancakeFilterer) WatchAdminTokenRecovery(opts *bind.WatchOpts, sink chan<- *PancakeAdminTokenRecovery) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeAdminTokenRecovery)
				if err := _Pancake.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
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

// ParseAdminTokenRecovery is a log parse operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pancake *PancakeFilterer) ParseAdminTokenRecovery(log types.Log) (*PancakeAdminTokenRecovery, error) {
	event := new(PancakeAdminTokenRecovery)
	if err := _Pancake.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pancake contract.
type PancakeDepositIterator struct {
	Event *PancakeDeposit // Event containing the contract specifics and raw log

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
func (it *PancakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeDeposit)
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
		it.Event = new(PancakeDeposit)
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
func (it *PancakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeDeposit represents a Deposit event raised by the Pancake contract.
type PancakeDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*PancakeDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeDepositIterator{contract: _Pancake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PancakeDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeDeposit)
				if err := _Pancake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) ParseDeposit(log types.Log) (*PancakeDeposit, error) {
	event := new(PancakeDeposit)
	if err := _Pancake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Pancake contract.
type PancakeEmergencyWithdrawIterator struct {
	Event *PancakeEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeEmergencyWithdraw)
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
		it.Event = new(PancakeEmergencyWithdraw)
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
func (it *PancakeEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeEmergencyWithdraw represents a EmergencyWithdraw event raised by the Pancake contract.
type PancakeEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*PancakeEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeEmergencyWithdrawIterator{contract: _Pancake.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeEmergencyWithdraw)
				if err := _Pancake.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) ParseEmergencyWithdraw(log types.Log) (*PancakeEmergencyWithdraw, error) {
	event := new(PancakeEmergencyWithdraw)
	if err := _Pancake.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeNewPoolLimitIterator is returned from FilterNewPoolLimit and is used to iterate over the raw logs and unpacked data for NewPoolLimit events raised by the Pancake contract.
type PancakeNewPoolLimitIterator struct {
	Event *PancakeNewPoolLimit // Event containing the contract specifics and raw log

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
func (it *PancakeNewPoolLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeNewPoolLimit)
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
		it.Event = new(PancakeNewPoolLimit)
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
func (it *PancakeNewPoolLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeNewPoolLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeNewPoolLimit represents a NewPoolLimit event raised by the Pancake contract.
type PancakeNewPoolLimit struct {
	PoolLimitPerUser *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPoolLimit is a free log retrieval operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pancake *PancakeFilterer) FilterNewPoolLimit(opts *bind.FilterOpts) (*PancakeNewPoolLimitIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return &PancakeNewPoolLimitIterator{contract: _Pancake.contract, event: "NewPoolLimit", logs: logs, sub: sub}, nil
}

// WatchNewPoolLimit is a free log subscription operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pancake *PancakeFilterer) WatchNewPoolLimit(opts *bind.WatchOpts, sink chan<- *PancakeNewPoolLimit) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeNewPoolLimit)
				if err := _Pancake.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
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

// ParseNewPoolLimit is a log parse operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pancake *PancakeFilterer) ParseNewPoolLimit(log types.Log) (*PancakeNewPoolLimit, error) {
	event := new(PancakeNewPoolLimit)
	if err := _Pancake.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeNewRewardPerBlockIterator is returned from FilterNewRewardPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardPerBlock events raised by the Pancake contract.
type PancakeNewRewardPerBlockIterator struct {
	Event *PancakeNewRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *PancakeNewRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeNewRewardPerBlock)
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
		it.Event = new(PancakeNewRewardPerBlock)
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
func (it *PancakeNewRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeNewRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeNewRewardPerBlock represents a NewRewardPerBlock event raised by the Pancake contract.
type PancakeNewRewardPerBlock struct {
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPerBlock is a free log retrieval operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pancake *PancakeFilterer) FilterNewRewardPerBlock(opts *bind.FilterOpts) (*PancakeNewRewardPerBlockIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &PancakeNewRewardPerBlockIterator{contract: _Pancake.contract, event: "NewRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardPerBlock is a free log subscription operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pancake *PancakeFilterer) WatchNewRewardPerBlock(opts *bind.WatchOpts, sink chan<- *PancakeNewRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeNewRewardPerBlock)
				if err := _Pancake.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
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

// ParseNewRewardPerBlock is a log parse operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pancake *PancakeFilterer) ParseNewRewardPerBlock(log types.Log) (*PancakeNewRewardPerBlock, error) {
	event := new(PancakeNewRewardPerBlock)
	if err := _Pancake.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeNewStartAndEndBlocksIterator is returned from FilterNewStartAndEndBlocks and is used to iterate over the raw logs and unpacked data for NewStartAndEndBlocks events raised by the Pancake contract.
type PancakeNewStartAndEndBlocksIterator struct {
	Event *PancakeNewStartAndEndBlocks // Event containing the contract specifics and raw log

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
func (it *PancakeNewStartAndEndBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeNewStartAndEndBlocks)
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
		it.Event = new(PancakeNewStartAndEndBlocks)
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
func (it *PancakeNewStartAndEndBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeNewStartAndEndBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeNewStartAndEndBlocks represents a NewStartAndEndBlocks event raised by the Pancake contract.
type PancakeNewStartAndEndBlocks struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewStartAndEndBlocks is a free log retrieval operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pancake *PancakeFilterer) FilterNewStartAndEndBlocks(opts *bind.FilterOpts) (*PancakeNewStartAndEndBlocksIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return &PancakeNewStartAndEndBlocksIterator{contract: _Pancake.contract, event: "NewStartAndEndBlocks", logs: logs, sub: sub}, nil
}

// WatchNewStartAndEndBlocks is a free log subscription operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pancake *PancakeFilterer) WatchNewStartAndEndBlocks(opts *bind.WatchOpts, sink chan<- *PancakeNewStartAndEndBlocks) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeNewStartAndEndBlocks)
				if err := _Pancake.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
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

// ParseNewStartAndEndBlocks is a log parse operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pancake *PancakeFilterer) ParseNewStartAndEndBlocks(log types.Log) (*PancakeNewStartAndEndBlocks, error) {
	event := new(PancakeNewStartAndEndBlocks)
	if err := _Pancake.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pancake contract.
type PancakeOwnershipTransferredIterator struct {
	Event *PancakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeOwnershipTransferred)
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
		it.Event = new(PancakeOwnershipTransferred)
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
func (it *PancakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeOwnershipTransferred represents a OwnershipTransferred event raised by the Pancake contract.
type PancakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pancake *PancakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeOwnershipTransferredIterator{contract: _Pancake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pancake *PancakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeOwnershipTransferred)
				if err := _Pancake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pancake *PancakeFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeOwnershipTransferred, error) {
	event := new(PancakeOwnershipTransferred)
	if err := _Pancake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeRewardsStopIterator is returned from FilterRewardsStop and is used to iterate over the raw logs and unpacked data for RewardsStop events raised by the Pancake contract.
type PancakeRewardsStopIterator struct {
	Event *PancakeRewardsStop // Event containing the contract specifics and raw log

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
func (it *PancakeRewardsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeRewardsStop)
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
		it.Event = new(PancakeRewardsStop)
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
func (it *PancakeRewardsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeRewardsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeRewardsStop represents a RewardsStop event raised by the Pancake contract.
type PancakeRewardsStop struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsStop is a free log retrieval operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pancake *PancakeFilterer) FilterRewardsStop(opts *bind.FilterOpts) (*PancakeRewardsStopIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return &PancakeRewardsStopIterator{contract: _Pancake.contract, event: "RewardsStop", logs: logs, sub: sub}, nil
}

// WatchRewardsStop is a free log subscription operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pancake *PancakeFilterer) WatchRewardsStop(opts *bind.WatchOpts, sink chan<- *PancakeRewardsStop) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeRewardsStop)
				if err := _Pancake.contract.UnpackLog(event, "RewardsStop", log); err != nil {
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

// ParseRewardsStop is a log parse operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pancake *PancakeFilterer) ParseRewardsStop(log types.Log) (*PancakeRewardsStop, error) {
	event := new(PancakeRewardsStop)
	if err := _Pancake.contract.UnpackLog(event, "RewardsStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pancake contract.
type PancakeWithdrawIterator struct {
	Event *PancakeWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeWithdraw)
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
		it.Event = new(PancakeWithdraw)
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
func (it *PancakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeWithdraw represents a Withdraw event raised by the Pancake contract.
type PancakeWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*PancakeWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeWithdrawIterator{contract: _Pancake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeWithdraw)
				if err := _Pancake.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Pancake *PancakeFilterer) ParseWithdraw(log types.Log) (*PancakeWithdraw, error) {
	event := new(PancakeWithdraw)
	if err := _Pancake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
