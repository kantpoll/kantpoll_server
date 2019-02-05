// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Mgz27ABI is the input ABI used to generate the binding from.
const Mgz27ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"checkHashcode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"howManyCandidatesInBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"incrementCancellations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"addGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getTor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"}],\"name\":\"getPreVote\",\"outputs\":[{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getCampaignIpfsInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCancellations\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"website\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stoppingAccessionToGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"iterateCandidatesCounter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disableCandidateLink\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ballots\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"closed\",\"type\":\"bool\"},{\"name\":\"stopped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"unstopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyBallots\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"}],\"name\":\"removeRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"addVoterToGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineTor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"addPreVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"first_number\",\"type\":\"bytes32\"},{\"name\":\"the_candidate\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[27]\"},{\"name\":\"\",\"type\":\"uint256[27]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[27]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groups\",\"outputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupPubkeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[27]\"},{\"name\":\"\",\"type\":\"bytes32[27]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingRounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"website\",\"type\":\"bytes32\"}],\"name\":\"addCandidateIntoBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineCanCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"stopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"getVotesPerBallotCandidateCategory\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"defineCurrentMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mgz\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineDisableCandidateLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"closeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canCancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"hasGroup\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"defineCurrentBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committedStatistics\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"str\",\"type\":\"bytes32\"}],\"name\":\"defineStoppingAccessionToGroups\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"category\",\"type\":\"bytes32\"}],\"name\":\"addGroupCategory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineCampaignIpfsInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationStatisticsPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Mgz27 is an auto generated Go binding around an Ethereum contract.
type Mgz27 struct {
	Mgz27Caller     // Read-only binding to the contract
	Mgz27Transactor // Write-only binding to the contract
	Mgz27Filterer   // Log filterer for contract events
}

// Mgz27Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mgz27Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz27Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mgz27Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz27Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mgz27Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz27Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mgz27Session struct {
	Contract     *Mgz27            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz27CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mgz27CallerSession struct {
	Contract *Mgz27Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Mgz27TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mgz27TransactorSession struct {
	Contract     *Mgz27Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz27Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mgz27Raw struct {
	Contract *Mgz27 // Generic contract binding to access the raw methods on
}

// Mgz27CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mgz27CallerRaw struct {
	Contract *Mgz27Caller // Generic read-only contract binding to access the raw methods on
}

// Mgz27TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mgz27TransactorRaw struct {
	Contract *Mgz27Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMgz27 creates a new instance of Mgz27, bound to a specific deployed contract.
func NewMgz27(address common.Address, backend bind.ContractBackend) (*Mgz27, error) {
	contract, err := bindMgz27(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mgz27{Mgz27Caller: Mgz27Caller{contract: contract}, Mgz27Transactor: Mgz27Transactor{contract: contract}, Mgz27Filterer: Mgz27Filterer{contract: contract}}, nil
}

// NewMgz27Caller creates a new read-only instance of Mgz27, bound to a specific deployed contract.
func NewMgz27Caller(address common.Address, caller bind.ContractCaller) (*Mgz27Caller, error) {
	contract, err := bindMgz27(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz27Caller{contract: contract}, nil
}

// NewMgz27Transactor creates a new write-only instance of Mgz27, bound to a specific deployed contract.
func NewMgz27Transactor(address common.Address, transactor bind.ContractTransactor) (*Mgz27Transactor, error) {
	contract, err := bindMgz27(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz27Transactor{contract: contract}, nil
}

// NewMgz27Filterer creates a new log filterer instance of Mgz27, bound to a specific deployed contract.
func NewMgz27Filterer(address common.Address, filterer bind.ContractFilterer) (*Mgz27Filterer, error) {
	contract, err := bindMgz27(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mgz27Filterer{contract: contract}, nil
}

// bindMgz27 binds a generic wrapper to an already deployed contract.
func bindMgz27(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Mgz27ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz27 *Mgz27Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz27.Contract.Mgz27Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz27 *Mgz27Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz27.Contract.Mgz27Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz27 *Mgz27Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz27.Contract.Mgz27Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz27 *Mgz27CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz27.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz27 *Mgz27TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz27.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz27 *Mgz27TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz27.Contract.contract.Transact(opts, method, params...)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz27 *Mgz27Caller) Ballots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	ret := new(struct {
		Id      [32]byte
		Closed  bool
		Stopped bool
	})
	out := ret
	err := _Mgz27.contract.Call(opts, out, "ballots", arg0)
	return *ret, err
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz27 *Mgz27Session) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz27.Contract.Ballots(&_Mgz27.CallOpts, arg0)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz27 *Mgz27CallerSession) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz27.Contract.Ballots(&_Mgz27.CallOpts, arg0)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz27 *Mgz27Caller) CanCancel(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "canCancel")
	return *ret0, err
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz27 *Mgz27Session) CanCancel() (bool, error) {
	return _Mgz27.Contract.CanCancel(&_Mgz27.CallOpts)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz27 *Mgz27CallerSession) CanCancel() (bool, error) {
	return _Mgz27.Contract.CanCancel(&_Mgz27.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz27 *Mgz27Caller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "chairperson")
	return *ret0, err
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz27 *Mgz27Session) Chairperson() (common.Address, error) {
	return _Mgz27.Contract.Chairperson(&_Mgz27.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz27 *Mgz27CallerSession) Chairperson() (common.Address, error) {
	return _Mgz27.Contract.Chairperson(&_Mgz27.CallOpts)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz27 *Mgz27Caller) CheckHashcode(opts *bind.CallOpts, hashcode [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "checkHashcode", hashcode)
	return *ret0, err
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz27 *Mgz27Session) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz27.Contract.CheckHashcode(&_Mgz27.CallOpts, hashcode)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz27 *Mgz27CallerSession) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz27.Contract.CheckHashcode(&_Mgz27.CallOpts, hashcode)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27Caller) Committed(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "committed", ballot, grp, position)
	return *ret0, err
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27Session) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz27.Contract.Committed(&_Mgz27.CallOpts, ballot, grp, position)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27CallerSession) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz27.Contract.Committed(&_Mgz27.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27Caller) CommittedStatistics(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "committedStatistics", ballot, grp, position)
	return *ret0, err
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27Session) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz27.Contract.CommittedStatistics(&_Mgz27.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz27 *Mgz27CallerSession) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz27.Contract.CommittedStatistics(&_Mgz27.CallOpts, ballot, grp, position)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) CurrentBallot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "currentBallot")
	return *ret0, err
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz27 *Mgz27Session) CurrentBallot() (*big.Int, error) {
	return _Mgz27.Contract.CurrentBallot(&_Mgz27.CallOpts)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) CurrentBallot() (*big.Int, error) {
	return _Mgz27.Contract.CurrentBallot(&_Mgz27.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz27 *Mgz27Caller) CurrentMessage(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "currentMessage")
	return *ret0, err
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz27 *Mgz27Session) CurrentMessage() ([32]byte, error) {
	return _Mgz27.Contract.CurrentMessage(&_Mgz27.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz27 *Mgz27CallerSession) CurrentMessage() ([32]byte, error) {
	return _Mgz27.Contract.CurrentMessage(&_Mgz27.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz27 *Mgz27Caller) DisableCandidateLink(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "disableCandidateLink")
	return *ret0, err
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz27 *Mgz27Session) DisableCandidateLink() (bool, error) {
	return _Mgz27.Contract.DisableCandidateLink(&_Mgz27.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz27 *Mgz27CallerSession) DisableCandidateLink() (bool, error) {
	return _Mgz27.Contract.DisableCandidateLink(&_Mgz27.CallOpts)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Caller) GetCampaignIpfsInfo(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "getCampaignIpfsInfo", pos)
	return *ret0, err
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Session) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GetCampaignIpfsInfo(&_Mgz27.CallOpts, pos)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27CallerSession) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GetCampaignIpfsInfo(&_Mgz27.CallOpts, pos)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Caller) GetCancellations(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "getCancellations", ballot, candidate)
	return *ret0, err
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Session) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.GetCancellations(&_Mgz27.CallOpts, ballot, candidate)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.GetCancellations(&_Mgz27.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz27 *Mgz27Caller) GetCandidate(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	ret := new(struct {
		Website [32]byte
		Count   *big.Int
	})
	out := ret
	err := _Mgz27.contract.Call(opts, out, "getCandidate", ballot, candidate)
	return *ret, err
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz27 *Mgz27Session) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz27.Contract.GetCandidate(&_Mgz27.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz27 *Mgz27CallerSession) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz27.Contract.GetCandidate(&_Mgz27.CallOpts, ballot, candidate)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[27], bytes32[27])
func (_Mgz27 *Mgz27Caller) GetGroupPubkeys(opts *bind.CallOpts, group *big.Int) ([27]*big.Int, [27][32]byte, error) {
	var (
		ret0 = new([27]*big.Int)
		ret1 = new([27][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz27.contract.Call(opts, out, "getGroupPubkeys", group)
	return *ret0, *ret1, err
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[27], bytes32[27])
func (_Mgz27 *Mgz27Session) GetGroupPubkeys(group *big.Int) ([27]*big.Int, [27][32]byte, error) {
	return _Mgz27.Contract.GetGroupPubkeys(&_Mgz27.CallOpts, group)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[27], bytes32[27])
func (_Mgz27 *Mgz27CallerSession) GetGroupPubkeys(group *big.Int) ([27]*big.Int, [27][32]byte, error) {
	return _Mgz27.Contract.GetGroupPubkeys(&_Mgz27.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[27])
func (_Mgz27 *Mgz27Caller) GetGroupVoters(opts *bind.CallOpts, group *big.Int) ([27]common.Address, error) {
	var (
		ret0 = new([27]common.Address)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "getGroupVoters", group)
	return *ret0, err
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[27])
func (_Mgz27 *Mgz27Session) GetGroupVoters(group *big.Int) ([27]common.Address, error) {
	return _Mgz27.Contract.GetGroupVoters(&_Mgz27.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[27])
func (_Mgz27 *Mgz27CallerSession) GetGroupVoters(group *big.Int) ([27]common.Address, error) {
	return _Mgz27.Contract.GetGroupVoters(&_Mgz27.CallOpts, group)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz27 *Mgz27Caller) GetPreVote(opts *bind.CallOpts, ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	ret := new(struct {
		Voter     [20]byte
		Candidate *big.Int
	})
	out := ret
	err := _Mgz27.contract.Call(opts, out, "getPreVote", ballot, group, fnumber)
	return *ret, err
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz27 *Mgz27Session) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz27.Contract.GetPreVote(&_Mgz27.CallOpts, ballot, group, fnumber)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz27 *Mgz27CallerSession) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz27.Contract.GetPreVote(&_Mgz27.CallOpts, ballot, group, fnumber)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Caller) GetTor(opts *bind.CallOpts, person common.Address, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "getTor", person, pos)
	return *ret0, err
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Session) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GetTor(&_Mgz27.CallOpts, person, pos)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27CallerSession) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GetTor(&_Mgz27.CallOpts, person, pos)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz27 *Mgz27Caller) GetVoter(opts *bind.CallOpts, voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	ret := new(struct {
		Pubkey   [32]byte
		Prefix   *big.Int
		Group    *big.Int
		HasGroup bool
	})
	out := ret
	err := _Mgz27.contract.Call(opts, out, "getVoter", voter)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz27 *Mgz27Session) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz27.Contract.GetVoter(&_Mgz27.CallOpts, voter)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz27 *Mgz27CallerSession) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz27.Contract.GetVoter(&_Mgz27.CallOpts, voter)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[27], uint256[27])
func (_Mgz27 *Mgz27Caller) GetVotes(opts *bind.CallOpts, ballot *big.Int, grp *big.Int) ([27][32]byte, [27]*big.Int, error) {
	var (
		ret0 = new([27][32]byte)
		ret1 = new([27]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz27.contract.Call(opts, out, "getVotes", ballot, grp)
	return *ret0, *ret1, err
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[27], uint256[27])
func (_Mgz27 *Mgz27Session) GetVotes(ballot *big.Int, grp *big.Int) ([27][32]byte, [27]*big.Int, error) {
	return _Mgz27.Contract.GetVotes(&_Mgz27.CallOpts, ballot, grp)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[27], uint256[27])
func (_Mgz27 *Mgz27CallerSession) GetVotes(ballot *big.Int, grp *big.Int) ([27][32]byte, [27]*big.Int, error) {
	return _Mgz27.Contract.GetVotes(&_Mgz27.CallOpts, ballot, grp)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Caller) GetVotesPerBallotCandidateCategory(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "getVotesPerBallotCandidateCategory", ballot, candidate, category)
	return *ret0, err
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Session) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.GetVotesPerBallotCandidateCategory(&_Mgz27.CallOpts, ballot, candidate, category)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.GetVotesPerBallotCandidateCategory(&_Mgz27.CallOpts, ballot, candidate, category)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Caller) GroupCategories(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "groupCategories", arg0)
	return *ret0, err
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27Session) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GroupCategories(&_Mgz27.CallOpts, arg0)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz27 *Mgz27CallerSession) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz27.Contract.GroupCategories(&_Mgz27.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz27 *Mgz27Caller) Groups(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	ret := new(struct {
		CPerson  common.Address
		Category *big.Int
		Size     *big.Int
	})
	out := ret
	err := _Mgz27.contract.Call(opts, out, "groups", arg0)
	return *ret, err
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz27 *Mgz27Session) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz27.Contract.Groups(&_Mgz27.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz27 *Mgz27CallerSession) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz27.Contract.Groups(&_Mgz27.CallOpts, arg0)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) HowManyBallots(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "howManyBallots")
	return *ret0, err
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz27 *Mgz27Session) HowManyBallots() (*big.Int, error) {
	return _Mgz27.Contract.HowManyBallots(&_Mgz27.CallOpts)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) HowManyBallots() (*big.Int, error) {
	return _Mgz27.Contract.HowManyBallots(&_Mgz27.CallOpts)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Caller) HowManyCandidatesInBallot(opts *bind.CallOpts, ballot *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "howManyCandidatesInBallot", ballot)
	return *ret0, err
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz27 *Mgz27Session) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.HowManyCandidatesInBallot(&_Mgz27.CallOpts, ballot)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz27.Contract.HowManyCandidatesInBallot(&_Mgz27.CallOpts, ballot)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) HowManyGroupCategories(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "howManyGroupCategories")
	return *ret0, err
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz27 *Mgz27Session) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz27.Contract.HowManyGroupCategories(&_Mgz27.CallOpts)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz27.Contract.HowManyGroupCategories(&_Mgz27.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) HowManyGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "howManyGroups")
	return *ret0, err
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz27 *Mgz27Session) HowManyGroups() (*big.Int, error) {
	return _Mgz27.Contract.HowManyGroups(&_Mgz27.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) HowManyGroups() (*big.Int, error) {
	return _Mgz27.Contract.HowManyGroups(&_Mgz27.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) Mgz(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "mgz")
	return *ret0, err
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz27 *Mgz27Session) Mgz() (*big.Int, error) {
	return _Mgz27.Contract.Mgz(&_Mgz27.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) Mgz() (*big.Int, error) {
	return _Mgz27.Contract.Mgz(&_Mgz27.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) RemainingRounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "remainingRounds")
	return *ret0, err
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz27 *Mgz27Session) RemainingRounds() (*big.Int, error) {
	return _Mgz27.Contract.RemainingRounds(&_Mgz27.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) RemainingRounds() (*big.Int, error) {
	return _Mgz27.Contract.RemainingRounds(&_Mgz27.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz27 *Mgz27Caller) Rounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "rounds")
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz27 *Mgz27Session) Rounds() (*big.Int, error) {
	return _Mgz27.Contract.Rounds(&_Mgz27.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz27 *Mgz27CallerSession) Rounds() (*big.Int, error) {
	return _Mgz27.Contract.Rounds(&_Mgz27.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz27 *Mgz27Caller) StoppingAccessionToGroups(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz27.contract.Call(opts, out, "stoppingAccessionToGroups")
	return *ret0, err
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz27 *Mgz27Session) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz27.Contract.StoppingAccessionToGroups(&_Mgz27.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz27 *Mgz27CallerSession) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz27.Contract.StoppingAccessionToGroups(&_Mgz27.CallOpts)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz27 *Mgz27Transactor) AddBallot(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addBallot", id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz27 *Mgz27Session) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddBallot(&_Mgz27.TransactOpts, id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddBallot(&_Mgz27.TransactOpts, id)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz27 *Mgz27Transactor) AddCandidateIntoBallot(opts *bind.TransactOpts, ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addCandidateIntoBallot", ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz27 *Mgz27Session) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddCandidateIntoBallot(&_Mgz27.TransactOpts, ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddCandidateIntoBallot(&_Mgz27.TransactOpts, ballot, position, website)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz27 *Mgz27Transactor) AddGroup(opts *bind.TransactOpts, cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addGroup", cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz27 *Mgz27Session) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddGroup(&_Mgz27.TransactOpts, cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddGroup(&_Mgz27.TransactOpts, cPerson, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz27 *Mgz27Transactor) AddGroupCategory(opts *bind.TransactOpts, category [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addGroupCategory", category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz27 *Mgz27Session) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddGroupCategory(&_Mgz27.TransactOpts, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.AddGroupCategory(&_Mgz27.TransactOpts, category)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz27 *Mgz27Transactor) AddPreVote(opts *bind.TransactOpts, ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addPreVote", ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz27 *Mgz27Session) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddPreVote(&_Mgz27.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddPreVote(&_Mgz27.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Transactor) AddVoterToGroup(opts *bind.TransactOpts, voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "addVoterToGroup", voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Session) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddVoterToGroup(&_Mgz27.TransactOpts, voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.AddVoterToGroup(&_Mgz27.TransactOpts, voter, grp, position)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Transactor) CloseBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "closeBallot", ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Session) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CloseBallot(&_Mgz27.TransactOpts, ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CloseBallot(&_Mgz27.TransactOpts, ballot)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Transactor) CommitVotationPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "commitVotationPerPosition", ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Session) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CommitVotationPerPosition(&_Mgz27.TransactOpts, ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CommitVotationPerPosition(&_Mgz27.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Transactor) CommitVotationStatisticsPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "commitVotationStatisticsPerPosition", ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27Session) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CommitVotationStatisticsPerPosition(&_Mgz27.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.CommitVotationStatisticsPerPosition(&_Mgz27.TransactOpts, ballot, grp, position)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27Transactor) DefineCampaignIpfsInfo(opts *bind.TransactOpts, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineCampaignIpfsInfo", pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27Session) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCampaignIpfsInfo(&_Mgz27.TransactOpts, pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCampaignIpfsInfo(&_Mgz27.TransactOpts, pos, value)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz27 *Mgz27Transactor) DefineCanCancel(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineCanCancel", b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz27 *Mgz27Session) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCanCancel(&_Mgz27.TransactOpts, b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCanCancel(&_Mgz27.TransactOpts, b)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Transactor) DefineCurrentBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineCurrentBallot", ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Session) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCurrentBallot(&_Mgz27.TransactOpts, ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCurrentBallot(&_Mgz27.TransactOpts, ballot)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz27 *Mgz27Transactor) DefineCurrentMessage(opts *bind.TransactOpts, message [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineCurrentMessage", message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz27 *Mgz27Session) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCurrentMessage(&_Mgz27.TransactOpts, message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineCurrentMessage(&_Mgz27.TransactOpts, message)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz27 *Mgz27Transactor) DefineDisableCandidateLink(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineDisableCandidateLink", b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz27 *Mgz27Session) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineDisableCandidateLink(&_Mgz27.TransactOpts, b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineDisableCandidateLink(&_Mgz27.TransactOpts, b)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz27 *Mgz27Transactor) DefineStoppingAccessionToGroups(opts *bind.TransactOpts, str [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineStoppingAccessionToGroups", str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz27 *Mgz27Session) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineStoppingAccessionToGroups(&_Mgz27.TransactOpts, str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineStoppingAccessionToGroups(&_Mgz27.TransactOpts, str)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27Transactor) DefineTor(opts *bind.TransactOpts, person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "defineTor", person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27Session) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineTor(&_Mgz27.TransactOpts, person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.DefineTor(&_Mgz27.TransactOpts, person, pos, value)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz27 *Mgz27Transactor) GiveRightToVote(opts *bind.TransactOpts, toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "giveRightToVote", toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz27 *Mgz27Session) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.GiveRightToVote(&_Mgz27.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz27 *Mgz27TransactorSession) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz27.Contract.GiveRightToVote(&_Mgz27.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz27 *Mgz27Transactor) IncrementCancellations(opts *bind.TransactOpts, ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "incrementCancellations", ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz27 *Mgz27Session) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.IncrementCancellations(&_Mgz27.TransactOpts, ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.IncrementCancellations(&_Mgz27.TransactOpts, ballot, candidate)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz27 *Mgz27Transactor) IterateCandidatesCounter(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "iterateCandidatesCounter", ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz27 *Mgz27Session) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.IterateCandidatesCounter(&_Mgz27.TransactOpts, ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.IterateCandidatesCounter(&_Mgz27.TransactOpts, ballot)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz27 *Mgz27Transactor) RemoveRightToVote(opts *bind.TransactOpts, toVoter common.Address) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "removeRightToVote", toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz27 *Mgz27Session) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz27.Contract.RemoveRightToVote(&_Mgz27.TransactOpts, toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz27 *Mgz27TransactorSession) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz27.Contract.RemoveRightToVote(&_Mgz27.TransactOpts, toVoter)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Transactor) StopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "stopBallot", ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Session) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.StopBallot(&_Mgz27.TransactOpts, ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.StopBallot(&_Mgz27.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Transactor) UnstopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "unstopBallot", ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27Session) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.UnstopBallot(&_Mgz27.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.UnstopBallot(&_Mgz27.TransactOpts, ballot)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz27 *Mgz27Transactor) Vote(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.contract.Transact(opts, "vote", ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz27 *Mgz27Session) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.Vote(&_Mgz27.TransactOpts, ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz27 *Mgz27TransactorSession) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz27.Contract.Vote(&_Mgz27.TransactOpts, ballot, grp, position, first_number, the_candidate)
}
