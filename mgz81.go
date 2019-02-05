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

// Mgz81ABI is the input ABI used to generate the binding from.
const Mgz81ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"checkHashcode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"howManyCandidatesInBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"incrementCancellations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"addGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getTor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"}],\"name\":\"getPreVote\",\"outputs\":[{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getCampaignIpfsInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCancellations\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"website\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stoppingAccessionToGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"iterateCandidatesCounter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disableCandidateLink\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ballots\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"closed\",\"type\":\"bool\"},{\"name\":\"stopped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"unstopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyBallots\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"}],\"name\":\"removeRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"addVoterToGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineTor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"addPreVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"first_number\",\"type\":\"bytes32\"},{\"name\":\"the_candidate\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[81]\"},{\"name\":\"\",\"type\":\"uint256[81]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[81]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groups\",\"outputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupPubkeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[81]\"},{\"name\":\"\",\"type\":\"bytes32[81]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingRounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"website\",\"type\":\"bytes32\"}],\"name\":\"addCandidateIntoBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineCanCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"stopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"getVotesPerBallotCandidateCategory\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"defineCurrentMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mgz\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineDisableCandidateLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"closeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canCancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"hasGroup\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"defineCurrentBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committedStatistics\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"str\",\"type\":\"bytes32\"}],\"name\":\"defineStoppingAccessionToGroups\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"category\",\"type\":\"bytes32\"}],\"name\":\"addGroupCategory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineCampaignIpfsInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationStatisticsPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Mgz81 is an auto generated Go binding around an Ethereum contract.
type Mgz81 struct {
	Mgz81Caller     // Read-only binding to the contract
	Mgz81Transactor // Write-only binding to the contract
	Mgz81Filterer   // Log filterer for contract events
}

// Mgz81Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mgz81Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz81Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mgz81Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz81Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mgz81Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz81Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mgz81Session struct {
	Contract     *Mgz81            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz81CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mgz81CallerSession struct {
	Contract *Mgz81Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Mgz81TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mgz81TransactorSession struct {
	Contract     *Mgz81Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz81Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mgz81Raw struct {
	Contract *Mgz81 // Generic contract binding to access the raw methods on
}

// Mgz81CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mgz81CallerRaw struct {
	Contract *Mgz81Caller // Generic read-only contract binding to access the raw methods on
}

// Mgz81TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mgz81TransactorRaw struct {
	Contract *Mgz81Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMgz81 creates a new instance of Mgz81, bound to a specific deployed contract.
func NewMgz81(address common.Address, backend bind.ContractBackend) (*Mgz81, error) {
	contract, err := bindMgz81(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mgz81{Mgz81Caller: Mgz81Caller{contract: contract}, Mgz81Transactor: Mgz81Transactor{contract: contract}, Mgz81Filterer: Mgz81Filterer{contract: contract}}, nil
}

// NewMgz81Caller creates a new read-only instance of Mgz81, bound to a specific deployed contract.
func NewMgz81Caller(address common.Address, caller bind.ContractCaller) (*Mgz81Caller, error) {
	contract, err := bindMgz81(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz81Caller{contract: contract}, nil
}

// NewMgz81Transactor creates a new write-only instance of Mgz81, bound to a specific deployed contract.
func NewMgz81Transactor(address common.Address, transactor bind.ContractTransactor) (*Mgz81Transactor, error) {
	contract, err := bindMgz81(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz81Transactor{contract: contract}, nil
}

// NewMgz81Filterer creates a new log filterer instance of Mgz81, bound to a specific deployed contract.
func NewMgz81Filterer(address common.Address, filterer bind.ContractFilterer) (*Mgz81Filterer, error) {
	contract, err := bindMgz81(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mgz81Filterer{contract: contract}, nil
}

// bindMgz81 binds a generic wrapper to an already deployed contract.
func bindMgz81(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Mgz81ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz81 *Mgz81Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz81.Contract.Mgz81Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz81 *Mgz81Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz81.Contract.Mgz81Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz81 *Mgz81Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz81.Contract.Mgz81Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz81 *Mgz81CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz81.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz81 *Mgz81TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz81.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz81 *Mgz81TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz81.Contract.contract.Transact(opts, method, params...)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz81 *Mgz81Caller) Ballots(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Mgz81.contract.Call(opts, out, "ballots", arg0)
	return *ret, err
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz81 *Mgz81Session) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz81.Contract.Ballots(&_Mgz81.CallOpts, arg0)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz81 *Mgz81CallerSession) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz81.Contract.Ballots(&_Mgz81.CallOpts, arg0)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz81 *Mgz81Caller) CanCancel(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "canCancel")
	return *ret0, err
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz81 *Mgz81Session) CanCancel() (bool, error) {
	return _Mgz81.Contract.CanCancel(&_Mgz81.CallOpts)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz81 *Mgz81CallerSession) CanCancel() (bool, error) {
	return _Mgz81.Contract.CanCancel(&_Mgz81.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz81 *Mgz81Caller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "chairperson")
	return *ret0, err
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz81 *Mgz81Session) Chairperson() (common.Address, error) {
	return _Mgz81.Contract.Chairperson(&_Mgz81.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz81 *Mgz81CallerSession) Chairperson() (common.Address, error) {
	return _Mgz81.Contract.Chairperson(&_Mgz81.CallOpts)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz81 *Mgz81Caller) CheckHashcode(opts *bind.CallOpts, hashcode [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "checkHashcode", hashcode)
	return *ret0, err
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz81 *Mgz81Session) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz81.Contract.CheckHashcode(&_Mgz81.CallOpts, hashcode)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz81 *Mgz81CallerSession) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz81.Contract.CheckHashcode(&_Mgz81.CallOpts, hashcode)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81Caller) Committed(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "committed", ballot, grp, position)
	return *ret0, err
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81Session) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz81.Contract.Committed(&_Mgz81.CallOpts, ballot, grp, position)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81CallerSession) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz81.Contract.Committed(&_Mgz81.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81Caller) CommittedStatistics(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "committedStatistics", ballot, grp, position)
	return *ret0, err
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81Session) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz81.Contract.CommittedStatistics(&_Mgz81.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz81 *Mgz81CallerSession) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz81.Contract.CommittedStatistics(&_Mgz81.CallOpts, ballot, grp, position)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) CurrentBallot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "currentBallot")
	return *ret0, err
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz81 *Mgz81Session) CurrentBallot() (*big.Int, error) {
	return _Mgz81.Contract.CurrentBallot(&_Mgz81.CallOpts)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) CurrentBallot() (*big.Int, error) {
	return _Mgz81.Contract.CurrentBallot(&_Mgz81.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz81 *Mgz81Caller) CurrentMessage(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "currentMessage")
	return *ret0, err
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz81 *Mgz81Session) CurrentMessage() ([32]byte, error) {
	return _Mgz81.Contract.CurrentMessage(&_Mgz81.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz81 *Mgz81CallerSession) CurrentMessage() ([32]byte, error) {
	return _Mgz81.Contract.CurrentMessage(&_Mgz81.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz81 *Mgz81Caller) DisableCandidateLink(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "disableCandidateLink")
	return *ret0, err
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz81 *Mgz81Session) DisableCandidateLink() (bool, error) {
	return _Mgz81.Contract.DisableCandidateLink(&_Mgz81.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz81 *Mgz81CallerSession) DisableCandidateLink() (bool, error) {
	return _Mgz81.Contract.DisableCandidateLink(&_Mgz81.CallOpts)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Caller) GetCampaignIpfsInfo(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "getCampaignIpfsInfo", pos)
	return *ret0, err
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Session) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GetCampaignIpfsInfo(&_Mgz81.CallOpts, pos)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81CallerSession) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GetCampaignIpfsInfo(&_Mgz81.CallOpts, pos)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Caller) GetCancellations(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "getCancellations", ballot, candidate)
	return *ret0, err
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Session) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.GetCancellations(&_Mgz81.CallOpts, ballot, candidate)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.GetCancellations(&_Mgz81.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz81 *Mgz81Caller) GetCandidate(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	ret := new(struct {
		Website [32]byte
		Count   *big.Int
	})
	out := ret
	err := _Mgz81.contract.Call(opts, out, "getCandidate", ballot, candidate)
	return *ret, err
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz81 *Mgz81Session) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz81.Contract.GetCandidate(&_Mgz81.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz81 *Mgz81CallerSession) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz81.Contract.GetCandidate(&_Mgz81.CallOpts, ballot, candidate)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[81], bytes32[81])
func (_Mgz81 *Mgz81Caller) GetGroupPubkeys(opts *bind.CallOpts, group *big.Int) ([81]*big.Int, [81][32]byte, error) {
	var (
		ret0 = new([81]*big.Int)
		ret1 = new([81][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz81.contract.Call(opts, out, "getGroupPubkeys", group)
	return *ret0, *ret1, err
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[81], bytes32[81])
func (_Mgz81 *Mgz81Session) GetGroupPubkeys(group *big.Int) ([81]*big.Int, [81][32]byte, error) {
	return _Mgz81.Contract.GetGroupPubkeys(&_Mgz81.CallOpts, group)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[81], bytes32[81])
func (_Mgz81 *Mgz81CallerSession) GetGroupPubkeys(group *big.Int) ([81]*big.Int, [81][32]byte, error) {
	return _Mgz81.Contract.GetGroupPubkeys(&_Mgz81.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[81])
func (_Mgz81 *Mgz81Caller) GetGroupVoters(opts *bind.CallOpts, group *big.Int) ([81]common.Address, error) {
	var (
		ret0 = new([81]common.Address)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "getGroupVoters", group)
	return *ret0, err
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[81])
func (_Mgz81 *Mgz81Session) GetGroupVoters(group *big.Int) ([81]common.Address, error) {
	return _Mgz81.Contract.GetGroupVoters(&_Mgz81.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[81])
func (_Mgz81 *Mgz81CallerSession) GetGroupVoters(group *big.Int) ([81]common.Address, error) {
	return _Mgz81.Contract.GetGroupVoters(&_Mgz81.CallOpts, group)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz81 *Mgz81Caller) GetPreVote(opts *bind.CallOpts, ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	ret := new(struct {
		Voter     [20]byte
		Candidate *big.Int
	})
	out := ret
	err := _Mgz81.contract.Call(opts, out, "getPreVote", ballot, group, fnumber)
	return *ret, err
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz81 *Mgz81Session) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz81.Contract.GetPreVote(&_Mgz81.CallOpts, ballot, group, fnumber)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz81 *Mgz81CallerSession) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz81.Contract.GetPreVote(&_Mgz81.CallOpts, ballot, group, fnumber)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Caller) GetTor(opts *bind.CallOpts, person common.Address, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "getTor", person, pos)
	return *ret0, err
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Session) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GetTor(&_Mgz81.CallOpts, person, pos)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81CallerSession) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GetTor(&_Mgz81.CallOpts, person, pos)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz81 *Mgz81Caller) GetVoter(opts *bind.CallOpts, voter common.Address) (struct {
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
	err := _Mgz81.contract.Call(opts, out, "getVoter", voter)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz81 *Mgz81Session) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz81.Contract.GetVoter(&_Mgz81.CallOpts, voter)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz81 *Mgz81CallerSession) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz81.Contract.GetVoter(&_Mgz81.CallOpts, voter)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[81], uint256[81])
func (_Mgz81 *Mgz81Caller) GetVotes(opts *bind.CallOpts, ballot *big.Int, grp *big.Int) ([81][32]byte, [81]*big.Int, error) {
	var (
		ret0 = new([81][32]byte)
		ret1 = new([81]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz81.contract.Call(opts, out, "getVotes", ballot, grp)
	return *ret0, *ret1, err
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[81], uint256[81])
func (_Mgz81 *Mgz81Session) GetVotes(ballot *big.Int, grp *big.Int) ([81][32]byte, [81]*big.Int, error) {
	return _Mgz81.Contract.GetVotes(&_Mgz81.CallOpts, ballot, grp)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[81], uint256[81])
func (_Mgz81 *Mgz81CallerSession) GetVotes(ballot *big.Int, grp *big.Int) ([81][32]byte, [81]*big.Int, error) {
	return _Mgz81.Contract.GetVotes(&_Mgz81.CallOpts, ballot, grp)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Caller) GetVotesPerBallotCandidateCategory(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "getVotesPerBallotCandidateCategory", ballot, candidate, category)
	return *ret0, err
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Session) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.GetVotesPerBallotCandidateCategory(&_Mgz81.CallOpts, ballot, candidate, category)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.GetVotesPerBallotCandidateCategory(&_Mgz81.CallOpts, ballot, candidate, category)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Caller) GroupCategories(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "groupCategories", arg0)
	return *ret0, err
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81Session) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GroupCategories(&_Mgz81.CallOpts, arg0)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz81 *Mgz81CallerSession) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz81.Contract.GroupCategories(&_Mgz81.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz81 *Mgz81Caller) Groups(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Mgz81.contract.Call(opts, out, "groups", arg0)
	return *ret, err
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz81 *Mgz81Session) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz81.Contract.Groups(&_Mgz81.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz81 *Mgz81CallerSession) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz81.Contract.Groups(&_Mgz81.CallOpts, arg0)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) HowManyBallots(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "howManyBallots")
	return *ret0, err
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz81 *Mgz81Session) HowManyBallots() (*big.Int, error) {
	return _Mgz81.Contract.HowManyBallots(&_Mgz81.CallOpts)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) HowManyBallots() (*big.Int, error) {
	return _Mgz81.Contract.HowManyBallots(&_Mgz81.CallOpts)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Caller) HowManyCandidatesInBallot(opts *bind.CallOpts, ballot *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "howManyCandidatesInBallot", ballot)
	return *ret0, err
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz81 *Mgz81Session) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.HowManyCandidatesInBallot(&_Mgz81.CallOpts, ballot)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz81.Contract.HowManyCandidatesInBallot(&_Mgz81.CallOpts, ballot)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) HowManyGroupCategories(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "howManyGroupCategories")
	return *ret0, err
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz81 *Mgz81Session) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz81.Contract.HowManyGroupCategories(&_Mgz81.CallOpts)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz81.Contract.HowManyGroupCategories(&_Mgz81.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) HowManyGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "howManyGroups")
	return *ret0, err
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz81 *Mgz81Session) HowManyGroups() (*big.Int, error) {
	return _Mgz81.Contract.HowManyGroups(&_Mgz81.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) HowManyGroups() (*big.Int, error) {
	return _Mgz81.Contract.HowManyGroups(&_Mgz81.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) Mgz(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "mgz")
	return *ret0, err
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz81 *Mgz81Session) Mgz() (*big.Int, error) {
	return _Mgz81.Contract.Mgz(&_Mgz81.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) Mgz() (*big.Int, error) {
	return _Mgz81.Contract.Mgz(&_Mgz81.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) RemainingRounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "remainingRounds")
	return *ret0, err
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz81 *Mgz81Session) RemainingRounds() (*big.Int, error) {
	return _Mgz81.Contract.RemainingRounds(&_Mgz81.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) RemainingRounds() (*big.Int, error) {
	return _Mgz81.Contract.RemainingRounds(&_Mgz81.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz81 *Mgz81Caller) Rounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "rounds")
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz81 *Mgz81Session) Rounds() (*big.Int, error) {
	return _Mgz81.Contract.Rounds(&_Mgz81.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz81 *Mgz81CallerSession) Rounds() (*big.Int, error) {
	return _Mgz81.Contract.Rounds(&_Mgz81.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz81 *Mgz81Caller) StoppingAccessionToGroups(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz81.contract.Call(opts, out, "stoppingAccessionToGroups")
	return *ret0, err
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz81 *Mgz81Session) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz81.Contract.StoppingAccessionToGroups(&_Mgz81.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz81 *Mgz81CallerSession) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz81.Contract.StoppingAccessionToGroups(&_Mgz81.CallOpts)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz81 *Mgz81Transactor) AddBallot(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addBallot", id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz81 *Mgz81Session) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddBallot(&_Mgz81.TransactOpts, id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddBallot(&_Mgz81.TransactOpts, id)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz81 *Mgz81Transactor) AddCandidateIntoBallot(opts *bind.TransactOpts, ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addCandidateIntoBallot", ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz81 *Mgz81Session) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddCandidateIntoBallot(&_Mgz81.TransactOpts, ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddCandidateIntoBallot(&_Mgz81.TransactOpts, ballot, position, website)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz81 *Mgz81Transactor) AddGroup(opts *bind.TransactOpts, cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addGroup", cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz81 *Mgz81Session) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddGroup(&_Mgz81.TransactOpts, cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddGroup(&_Mgz81.TransactOpts, cPerson, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz81 *Mgz81Transactor) AddGroupCategory(opts *bind.TransactOpts, category [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addGroupCategory", category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz81 *Mgz81Session) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddGroupCategory(&_Mgz81.TransactOpts, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.AddGroupCategory(&_Mgz81.TransactOpts, category)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz81 *Mgz81Transactor) AddPreVote(opts *bind.TransactOpts, ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addPreVote", ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz81 *Mgz81Session) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddPreVote(&_Mgz81.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddPreVote(&_Mgz81.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Transactor) AddVoterToGroup(opts *bind.TransactOpts, voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "addVoterToGroup", voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Session) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddVoterToGroup(&_Mgz81.TransactOpts, voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.AddVoterToGroup(&_Mgz81.TransactOpts, voter, grp, position)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Transactor) CloseBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "closeBallot", ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Session) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CloseBallot(&_Mgz81.TransactOpts, ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CloseBallot(&_Mgz81.TransactOpts, ballot)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Transactor) CommitVotationPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "commitVotationPerPosition", ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Session) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CommitVotationPerPosition(&_Mgz81.TransactOpts, ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CommitVotationPerPosition(&_Mgz81.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Transactor) CommitVotationStatisticsPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "commitVotationStatisticsPerPosition", ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81Session) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CommitVotationStatisticsPerPosition(&_Mgz81.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.CommitVotationStatisticsPerPosition(&_Mgz81.TransactOpts, ballot, grp, position)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81Transactor) DefineCampaignIpfsInfo(opts *bind.TransactOpts, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineCampaignIpfsInfo", pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81Session) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCampaignIpfsInfo(&_Mgz81.TransactOpts, pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCampaignIpfsInfo(&_Mgz81.TransactOpts, pos, value)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz81 *Mgz81Transactor) DefineCanCancel(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineCanCancel", b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz81 *Mgz81Session) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCanCancel(&_Mgz81.TransactOpts, b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCanCancel(&_Mgz81.TransactOpts, b)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Transactor) DefineCurrentBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineCurrentBallot", ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Session) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCurrentBallot(&_Mgz81.TransactOpts, ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCurrentBallot(&_Mgz81.TransactOpts, ballot)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz81 *Mgz81Transactor) DefineCurrentMessage(opts *bind.TransactOpts, message [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineCurrentMessage", message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz81 *Mgz81Session) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCurrentMessage(&_Mgz81.TransactOpts, message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineCurrentMessage(&_Mgz81.TransactOpts, message)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz81 *Mgz81Transactor) DefineDisableCandidateLink(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineDisableCandidateLink", b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz81 *Mgz81Session) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineDisableCandidateLink(&_Mgz81.TransactOpts, b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineDisableCandidateLink(&_Mgz81.TransactOpts, b)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz81 *Mgz81Transactor) DefineStoppingAccessionToGroups(opts *bind.TransactOpts, str [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineStoppingAccessionToGroups", str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz81 *Mgz81Session) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineStoppingAccessionToGroups(&_Mgz81.TransactOpts, str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineStoppingAccessionToGroups(&_Mgz81.TransactOpts, str)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81Transactor) DefineTor(opts *bind.TransactOpts, person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "defineTor", person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81Session) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineTor(&_Mgz81.TransactOpts, person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.DefineTor(&_Mgz81.TransactOpts, person, pos, value)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz81 *Mgz81Transactor) GiveRightToVote(opts *bind.TransactOpts, toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "giveRightToVote", toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz81 *Mgz81Session) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.GiveRightToVote(&_Mgz81.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz81 *Mgz81TransactorSession) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz81.Contract.GiveRightToVote(&_Mgz81.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz81 *Mgz81Transactor) IncrementCancellations(opts *bind.TransactOpts, ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "incrementCancellations", ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz81 *Mgz81Session) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.IncrementCancellations(&_Mgz81.TransactOpts, ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.IncrementCancellations(&_Mgz81.TransactOpts, ballot, candidate)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz81 *Mgz81Transactor) IterateCandidatesCounter(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "iterateCandidatesCounter", ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz81 *Mgz81Session) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.IterateCandidatesCounter(&_Mgz81.TransactOpts, ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.IterateCandidatesCounter(&_Mgz81.TransactOpts, ballot)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz81 *Mgz81Transactor) RemoveRightToVote(opts *bind.TransactOpts, toVoter common.Address) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "removeRightToVote", toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz81 *Mgz81Session) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz81.Contract.RemoveRightToVote(&_Mgz81.TransactOpts, toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz81 *Mgz81TransactorSession) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz81.Contract.RemoveRightToVote(&_Mgz81.TransactOpts, toVoter)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Transactor) StopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "stopBallot", ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Session) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.StopBallot(&_Mgz81.TransactOpts, ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.StopBallot(&_Mgz81.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Transactor) UnstopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "unstopBallot", ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81Session) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.UnstopBallot(&_Mgz81.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.UnstopBallot(&_Mgz81.TransactOpts, ballot)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz81 *Mgz81Transactor) Vote(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.contract.Transact(opts, "vote", ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz81 *Mgz81Session) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.Vote(&_Mgz81.TransactOpts, ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz81 *Mgz81TransactorSession) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz81.Contract.Vote(&_Mgz81.TransactOpts, ballot, grp, position, first_number, the_candidate)
}
