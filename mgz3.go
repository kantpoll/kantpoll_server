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

// Mgz3ABI is the input ABI used to generate the binding from.
const Mgz3ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"checkHashcode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"howManyCandidatesInBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"incrementCancellations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"addGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getTor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"}],\"name\":\"getPreVote\",\"outputs\":[{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getCampaignIpfsInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCancellations\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"website\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stoppingAccessionToGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"iterateCandidatesCounter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disableCandidateLink\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ballots\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"closed\",\"type\":\"bool\"},{\"name\":\"stopped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"unstopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"howManyBallots\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"}],\"name\":\"removeRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"addVoterToGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"person\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineTor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toVoter\",\"type\":\"address\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"hashcode\",\"type\":\"bytes32\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"fnumber\",\"type\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"bytes20\"},{\"name\":\"candidate\",\"type\":\"uint256\"}],\"name\":\"addPreVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"first_number\",\"type\":\"bytes32\"},{\"name\":\"the_candidate\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[3]\"},{\"name\":\"\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[3]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groups\",\"outputs\":[{\"name\":\"cPerson\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"uint256\"}],\"name\":\"getGroupPubkeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"},{\"name\":\"\",\"type\":\"bytes32[3]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingRounds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"website\",\"type\":\"bytes32\"}],\"name\":\"addCandidateIntoBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineCanCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"stopBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"candidate\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"getVotesPerBallotCandidateCategory\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groupCategories\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"defineCurrentMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mgz\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"defineDisableCandidateLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"closeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canCancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"name\":\"prefix\",\"type\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\"},{\"name\":\"hasGroup\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"}],\"name\":\"defineCurrentBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"committedStatistics\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"str\",\"type\":\"bytes32\"}],\"name\":\"defineStoppingAccessionToGroups\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"category\",\"type\":\"bytes32\"}],\"name\":\"addGroupCategory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"defineCampaignIpfsInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballot\",\"type\":\"uint256\"},{\"name\":\"grp\",\"type\":\"uint256\"},{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"commitVotationStatisticsPerPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Mgz3 is an auto generated Go binding around an Ethereum contract.
type Mgz3 struct {
	Mgz3Caller     // Read-only binding to the contract
	Mgz3Transactor // Write-only binding to the contract
	Mgz3Filterer   // Log filterer for contract events
}

// Mgz3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mgz3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mgz3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mgz3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mgz3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mgz3Session struct {
	Contract     *Mgz3             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mgz3CallerSession struct {
	Contract *Mgz3Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Mgz3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mgz3TransactorSession struct {
	Contract     *Mgz3Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mgz3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mgz3Raw struct {
	Contract *Mgz3 // Generic contract binding to access the raw methods on
}

// Mgz3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mgz3CallerRaw struct {
	Contract *Mgz3Caller // Generic read-only contract binding to access the raw methods on
}

// Mgz3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mgz3TransactorRaw struct {
	Contract *Mgz3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMgz3 creates a new instance of Mgz3, bound to a specific deployed contract.
func NewMgz3(address common.Address, backend bind.ContractBackend) (*Mgz3, error) {
	contract, err := bindMgz3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mgz3{Mgz3Caller: Mgz3Caller{contract: contract}, Mgz3Transactor: Mgz3Transactor{contract: contract}, Mgz3Filterer: Mgz3Filterer{contract: contract}}, nil
}

// NewMgz3Caller creates a new read-only instance of Mgz3, bound to a specific deployed contract.
func NewMgz3Caller(address common.Address, caller bind.ContractCaller) (*Mgz3Caller, error) {
	contract, err := bindMgz3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz3Caller{contract: contract}, nil
}

// NewMgz3Transactor creates a new write-only instance of Mgz3, bound to a specific deployed contract.
func NewMgz3Transactor(address common.Address, transactor bind.ContractTransactor) (*Mgz3Transactor, error) {
	contract, err := bindMgz3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mgz3Transactor{contract: contract}, nil
}

// NewMgz3Filterer creates a new log filterer instance of Mgz3, bound to a specific deployed contract.
func NewMgz3Filterer(address common.Address, filterer bind.ContractFilterer) (*Mgz3Filterer, error) {
	contract, err := bindMgz3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mgz3Filterer{contract: contract}, nil
}

// bindMgz3 binds a generic wrapper to an already deployed contract.
func bindMgz3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Mgz3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz3 *Mgz3Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz3.Contract.Mgz3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz3 *Mgz3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz3.Contract.Mgz3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz3 *Mgz3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz3.Contract.Mgz3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mgz3 *Mgz3CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mgz3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mgz3 *Mgz3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mgz3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mgz3 *Mgz3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mgz3.Contract.contract.Transact(opts, method, params...)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz3 *Mgz3Caller) Ballots(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Mgz3.contract.Call(opts, out, "ballots", arg0)
	return *ret, err
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz3 *Mgz3Session) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz3.Contract.Ballots(&_Mgz3.CallOpts, arg0)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots( uint256) constant returns(id bytes32, closed bool, stopped bool)
func (_Mgz3 *Mgz3CallerSession) Ballots(arg0 *big.Int) (struct {
	Id      [32]byte
	Closed  bool
	Stopped bool
}, error) {
	return _Mgz3.Contract.Ballots(&_Mgz3.CallOpts, arg0)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz3 *Mgz3Caller) CanCancel(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "canCancel")
	return *ret0, err
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz3 *Mgz3Session) CanCancel() (bool, error) {
	return _Mgz3.Contract.CanCancel(&_Mgz3.CallOpts)
}

// CanCancel is a free data retrieval call binding the contract method 0xd0f3725f.
//
// Solidity: function canCancel() constant returns(bool)
func (_Mgz3 *Mgz3CallerSession) CanCancel() (bool, error) {
	return _Mgz3.Contract.CanCancel(&_Mgz3.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz3 *Mgz3Caller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "chairperson")
	return *ret0, err
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz3 *Mgz3Session) Chairperson() (common.Address, error) {
	return _Mgz3.Contract.Chairperson(&_Mgz3.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Mgz3 *Mgz3CallerSession) Chairperson() (common.Address, error) {
	return _Mgz3.Contract.Chairperson(&_Mgz3.CallOpts)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz3 *Mgz3Caller) CheckHashcode(opts *bind.CallOpts, hashcode [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "checkHashcode", hashcode)
	return *ret0, err
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz3 *Mgz3Session) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz3.Contract.CheckHashcode(&_Mgz3.CallOpts, hashcode)
}

// CheckHashcode is a free data retrieval call binding the contract method 0x0fcf6143.
//
// Solidity: function checkHashcode(hashcode bytes32) constant returns(bool)
func (_Mgz3 *Mgz3CallerSession) CheckHashcode(hashcode [32]byte) (bool, error) {
	return _Mgz3.Contract.CheckHashcode(&_Mgz3.CallOpts, hashcode)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3Caller) Committed(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "committed", ballot, grp, position)
	return *ret0, err
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3Session) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz3.Contract.Committed(&_Mgz3.CallOpts, ballot, grp, position)
}

// Committed is a free data retrieval call binding the contract method 0x812d9dd4.
//
// Solidity: function committed(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3CallerSession) Committed(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz3.Contract.Committed(&_Mgz3.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3Caller) CommittedStatistics(opts *bind.CallOpts, ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "committedStatistics", ballot, grp, position)
	return *ret0, err
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3Session) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz3.Contract.CommittedStatistics(&_Mgz3.CallOpts, ballot, grp, position)
}

// CommittedStatistics is a free data retrieval call binding the contract method 0xe1975258.
//
// Solidity: function committedStatistics(ballot uint256, grp uint256, position uint256) constant returns(bool)
func (_Mgz3 *Mgz3CallerSession) CommittedStatistics(ballot *big.Int, grp *big.Int, position *big.Int) (bool, error) {
	return _Mgz3.Contract.CommittedStatistics(&_Mgz3.CallOpts, ballot, grp, position)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) CurrentBallot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "currentBallot")
	return *ret0, err
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz3 *Mgz3Session) CurrentBallot() (*big.Int, error) {
	return _Mgz3.Contract.CurrentBallot(&_Mgz3.CallOpts)
}

// CurrentBallot is a free data retrieval call binding the contract method 0xf80723a8.
//
// Solidity: function currentBallot() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) CurrentBallot() (*big.Int, error) {
	return _Mgz3.Contract.CurrentBallot(&_Mgz3.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz3 *Mgz3Caller) CurrentMessage(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "currentMessage")
	return *ret0, err
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz3 *Mgz3Session) CurrentMessage() ([32]byte, error) {
	return _Mgz3.Contract.CurrentMessage(&_Mgz3.CallOpts)
}

// CurrentMessage is a free data retrieval call binding the contract method 0xd9203db6.
//
// Solidity: function currentMessage() constant returns(bytes32)
func (_Mgz3 *Mgz3CallerSession) CurrentMessage() ([32]byte, error) {
	return _Mgz3.Contract.CurrentMessage(&_Mgz3.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz3 *Mgz3Caller) DisableCandidateLink(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "disableCandidateLink")
	return *ret0, err
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz3 *Mgz3Session) DisableCandidateLink() (bool, error) {
	return _Mgz3.Contract.DisableCandidateLink(&_Mgz3.CallOpts)
}

// DisableCandidateLink is a free data retrieval call binding the contract method 0x5715a6e0.
//
// Solidity: function disableCandidateLink() constant returns(bool)
func (_Mgz3 *Mgz3CallerSession) DisableCandidateLink() (bool, error) {
	return _Mgz3.Contract.DisableCandidateLink(&_Mgz3.CallOpts)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Caller) GetCampaignIpfsInfo(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "getCampaignIpfsInfo", pos)
	return *ret0, err
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Session) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GetCampaignIpfsInfo(&_Mgz3.CallOpts, pos)
}

// GetCampaignIpfsInfo is a free data retrieval call binding the contract method 0x2e8112fd.
//
// Solidity: function getCampaignIpfsInfo(pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3CallerSession) GetCampaignIpfsInfo(pos *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GetCampaignIpfsInfo(&_Mgz3.CallOpts, pos)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Caller) GetCancellations(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "getCancellations", ballot, candidate)
	return *ret0, err
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Session) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.GetCancellations(&_Mgz3.CallOpts, ballot, candidate)
}

// GetCancellations is a free data retrieval call binding the contract method 0x419cc25f.
//
// Solidity: function getCancellations(ballot uint256, candidate uint256) constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) GetCancellations(ballot *big.Int, candidate *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.GetCancellations(&_Mgz3.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz3 *Mgz3Caller) GetCandidate(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	ret := new(struct {
		Website [32]byte
		Count   *big.Int
	})
	out := ret
	err := _Mgz3.contract.Call(opts, out, "getCandidate", ballot, candidate)
	return *ret, err
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz3 *Mgz3Session) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz3.Contract.GetCandidate(&_Mgz3.CallOpts, ballot, candidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(ballot uint256, candidate uint256) constant returns(website bytes32, count uint256)
func (_Mgz3 *Mgz3CallerSession) GetCandidate(ballot *big.Int, candidate *big.Int) (struct {
	Website [32]byte
	Count   *big.Int
}, error) {
	return _Mgz3.Contract.GetCandidate(&_Mgz3.CallOpts, ballot, candidate)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[3], bytes32[3])
func (_Mgz3 *Mgz3Caller) GetGroupPubkeys(opts *bind.CallOpts, group *big.Int) ([3]*big.Int, [3][32]byte, error) {
	var (
		ret0 = new([3]*big.Int)
		ret1 = new([3][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz3.contract.Call(opts, out, "getGroupPubkeys", group)
	return *ret0, *ret1, err
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[3], bytes32[3])
func (_Mgz3 *Mgz3Session) GetGroupPubkeys(group *big.Int) ([3]*big.Int, [3][32]byte, error) {
	return _Mgz3.Contract.GetGroupPubkeys(&_Mgz3.CallOpts, group)
}

// GetGroupPubkeys is a free data retrieval call binding the contract method 0x99733838.
//
// Solidity: function getGroupPubkeys(group uint256) constant returns(uint256[3], bytes32[3])
func (_Mgz3 *Mgz3CallerSession) GetGroupPubkeys(group *big.Int) ([3]*big.Int, [3][32]byte, error) {
	return _Mgz3.Contract.GetGroupPubkeys(&_Mgz3.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[3])
func (_Mgz3 *Mgz3Caller) GetGroupVoters(opts *bind.CallOpts, group *big.Int) ([3]common.Address, error) {
	var (
		ret0 = new([3]common.Address)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "getGroupVoters", group)
	return *ret0, err
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[3])
func (_Mgz3 *Mgz3Session) GetGroupVoters(group *big.Int) ([3]common.Address, error) {
	return _Mgz3.Contract.GetGroupVoters(&_Mgz3.CallOpts, group)
}

// GetGroupVoters is a free data retrieval call binding the contract method 0x939cf16d.
//
// Solidity: function getGroupVoters(group uint256) constant returns(address[3])
func (_Mgz3 *Mgz3CallerSession) GetGroupVoters(group *big.Int) ([3]common.Address, error) {
	return _Mgz3.Contract.GetGroupVoters(&_Mgz3.CallOpts, group)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz3 *Mgz3Caller) GetPreVote(opts *bind.CallOpts, ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	ret := new(struct {
		Voter     [20]byte
		Candidate *big.Int
	})
	out := ret
	err := _Mgz3.contract.Call(opts, out, "getPreVote", ballot, group, fnumber)
	return *ret, err
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz3 *Mgz3Session) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz3.Contract.GetPreVote(&_Mgz3.CallOpts, ballot, group, fnumber)
}

// GetPreVote is a free data retrieval call binding the contract method 0x2b9127d4.
//
// Solidity: function getPreVote(ballot uint256, group uint256, fnumber bytes32) constant returns(voter bytes20, candidate uint256)
func (_Mgz3 *Mgz3CallerSession) GetPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte) (struct {
	Voter     [20]byte
	Candidate *big.Int
}, error) {
	return _Mgz3.Contract.GetPreVote(&_Mgz3.CallOpts, ballot, group, fnumber)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Caller) GetTor(opts *bind.CallOpts, person common.Address, pos *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "getTor", person, pos)
	return *ret0, err
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Session) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GetTor(&_Mgz3.CallOpts, person, pos)
}

// GetTor is a free data retrieval call binding the contract method 0x1dcd1e17.
//
// Solidity: function getTor(person address, pos uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3CallerSession) GetTor(person common.Address, pos *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GetTor(&_Mgz3.CallOpts, person, pos)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz3 *Mgz3Caller) GetVoter(opts *bind.CallOpts, voter common.Address) (struct {
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
	err := _Mgz3.contract.Call(opts, out, "getVoter", voter)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz3 *Mgz3Session) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz3.Contract.GetVoter(&_Mgz3.CallOpts, voter)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(voter address) constant returns(pubkey bytes32, prefix uint256, group uint256, hasGroup bool)
func (_Mgz3 *Mgz3CallerSession) GetVoter(voter common.Address) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {
	return _Mgz3.Contract.GetVoter(&_Mgz3.CallOpts, voter)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[3], uint256[3])
func (_Mgz3 *Mgz3Caller) GetVotes(opts *bind.CallOpts, ballot *big.Int, grp *big.Int) ([3][32]byte, [3]*big.Int, error) {
	var (
		ret0 = new([3][32]byte)
		ret1 = new([3]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Mgz3.contract.Call(opts, out, "getVotes", ballot, grp)
	return *ret0, *ret1, err
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[3], uint256[3])
func (_Mgz3 *Mgz3Session) GetVotes(ballot *big.Int, grp *big.Int) ([3][32]byte, [3]*big.Int, error) {
	return _Mgz3.Contract.GetVotes(&_Mgz3.CallOpts, ballot, grp)
}

// GetVotes is a free data retrieval call binding the contract method 0x91f36633.
//
// Solidity: function getVotes(ballot uint256, grp uint256) constant returns(bytes32[3], uint256[3])
func (_Mgz3 *Mgz3CallerSession) GetVotes(ballot *big.Int, grp *big.Int) ([3][32]byte, [3]*big.Int, error) {
	return _Mgz3.Contract.GetVotes(&_Mgz3.CallOpts, ballot, grp)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Caller) GetVotesPerBallotCandidateCategory(opts *bind.CallOpts, ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "getVotesPerBallotCandidateCategory", ballot, candidate, category)
	return *ret0, err
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Session) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.GetVotesPerBallotCandidateCategory(&_Mgz3.CallOpts, ballot, candidate, category)
}

// GetVotesPerBallotCandidateCategory is a free data retrieval call binding the contract method 0xb7b066c4.
//
// Solidity: function getVotesPerBallotCandidateCategory(ballot uint256, candidate uint256, category uint256) constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) GetVotesPerBallotCandidateCategory(ballot *big.Int, candidate *big.Int, category *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.GetVotesPerBallotCandidateCategory(&_Mgz3.CallOpts, ballot, candidate, category)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Caller) GroupCategories(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "groupCategories", arg0)
	return *ret0, err
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3Session) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GroupCategories(&_Mgz3.CallOpts, arg0)
}

// GroupCategories is a free data retrieval call binding the contract method 0xbfc3f2d5.
//
// Solidity: function groupCategories( uint256) constant returns(bytes32)
func (_Mgz3 *Mgz3CallerSession) GroupCategories(arg0 *big.Int) ([32]byte, error) {
	return _Mgz3.Contract.GroupCategories(&_Mgz3.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz3 *Mgz3Caller) Groups(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Mgz3.contract.Call(opts, out, "groups", arg0)
	return *ret, err
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz3 *Mgz3Session) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz3.Contract.Groups(&_Mgz3.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups( uint256) constant returns(cPerson address, category uint256, size uint256)
func (_Mgz3 *Mgz3CallerSession) Groups(arg0 *big.Int) (struct {
	CPerson  common.Address
	Category *big.Int
	Size     *big.Int
}, error) {
	return _Mgz3.Contract.Groups(&_Mgz3.CallOpts, arg0)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) HowManyBallots(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "howManyBallots")
	return *ret0, err
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz3 *Mgz3Session) HowManyBallots() (*big.Int, error) {
	return _Mgz3.Contract.HowManyBallots(&_Mgz3.CallOpts)
}

// HowManyBallots is a free data retrieval call binding the contract method 0x656e2a37.
//
// Solidity: function howManyBallots() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) HowManyBallots() (*big.Int, error) {
	return _Mgz3.Contract.HowManyBallots(&_Mgz3.CallOpts)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Caller) HowManyCandidatesInBallot(opts *bind.CallOpts, ballot *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "howManyCandidatesInBallot", ballot)
	return *ret0, err
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz3 *Mgz3Session) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.HowManyCandidatesInBallot(&_Mgz3.CallOpts, ballot)
}

// HowManyCandidatesInBallot is a free data retrieval call binding the contract method 0x103eb2f7.
//
// Solidity: function howManyCandidatesInBallot(ballot uint256) constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) HowManyCandidatesInBallot(ballot *big.Int) (*big.Int, error) {
	return _Mgz3.Contract.HowManyCandidatesInBallot(&_Mgz3.CallOpts, ballot)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) HowManyGroupCategories(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "howManyGroupCategories")
	return *ret0, err
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz3 *Mgz3Session) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz3.Contract.HowManyGroupCategories(&_Mgz3.CallOpts)
}

// HowManyGroupCategories is a free data retrieval call binding the contract method 0x0a7c172b.
//
// Solidity: function howManyGroupCategories() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) HowManyGroupCategories() (*big.Int, error) {
	return _Mgz3.Contract.HowManyGroupCategories(&_Mgz3.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) HowManyGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "howManyGroups")
	return *ret0, err
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz3 *Mgz3Session) HowManyGroups() (*big.Int, error) {
	return _Mgz3.Contract.HowManyGroups(&_Mgz3.CallOpts)
}

// HowManyGroups is a free data retrieval call binding the contract method 0x422f675b.
//
// Solidity: function howManyGroups() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) HowManyGroups() (*big.Int, error) {
	return _Mgz3.Contract.HowManyGroups(&_Mgz3.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) Mgz(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "mgz")
	return *ret0, err
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz3 *Mgz3Session) Mgz() (*big.Int, error) {
	return _Mgz3.Contract.Mgz(&_Mgz3.CallOpts)
}

// Mgz is a free data retrieval call binding the contract method 0xc4c1ce45.
//
// Solidity: function mgz() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) Mgz() (*big.Int, error) {
	return _Mgz3.Contract.Mgz(&_Mgz3.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) RemainingRounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "remainingRounds")
	return *ret0, err
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz3 *Mgz3Session) RemainingRounds() (*big.Int, error) {
	return _Mgz3.Contract.RemainingRounds(&_Mgz3.CallOpts)
}

// RemainingRounds is a free data retrieval call binding the contract method 0xa7e22a3e.
//
// Solidity: function remainingRounds() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) RemainingRounds() (*big.Int, error) {
	return _Mgz3.Contract.RemainingRounds(&_Mgz3.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz3 *Mgz3Caller) Rounds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "rounds")
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz3 *Mgz3Session) Rounds() (*big.Int, error) {
	return _Mgz3.Contract.Rounds(&_Mgz3.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xa2e800ad.
//
// Solidity: function rounds() constant returns(uint256)
func (_Mgz3 *Mgz3CallerSession) Rounds() (*big.Int, error) {
	return _Mgz3.Contract.Rounds(&_Mgz3.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz3 *Mgz3Caller) StoppingAccessionToGroups(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mgz3.contract.Call(opts, out, "stoppingAccessionToGroups")
	return *ret0, err
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz3 *Mgz3Session) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz3.Contract.StoppingAccessionToGroups(&_Mgz3.CallOpts)
}

// StoppingAccessionToGroups is a free data retrieval call binding the contract method 0x4f2dc68d.
//
// Solidity: function stoppingAccessionToGroups() constant returns(bytes32)
func (_Mgz3 *Mgz3CallerSession) StoppingAccessionToGroups() ([32]byte, error) {
	return _Mgz3.Contract.StoppingAccessionToGroups(&_Mgz3.CallOpts)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz3 *Mgz3Transactor) AddBallot(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addBallot", id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz3 *Mgz3Session) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddBallot(&_Mgz3.TransactOpts, id)
}

// AddBallot is a paid mutator transaction binding the contract method 0xb8d97473.
//
// Solidity: function addBallot(id bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) AddBallot(id [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddBallot(&_Mgz3.TransactOpts, id)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz3 *Mgz3Transactor) AddCandidateIntoBallot(opts *bind.TransactOpts, ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addCandidateIntoBallot", ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz3 *Mgz3Session) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddCandidateIntoBallot(&_Mgz3.TransactOpts, ballot, position, website)
}

// AddCandidateIntoBallot is a paid mutator transaction binding the contract method 0xa88a74af.
//
// Solidity: function addCandidateIntoBallot(ballot uint256, position uint256, website bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) AddCandidateIntoBallot(ballot *big.Int, position *big.Int, website [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddCandidateIntoBallot(&_Mgz3.TransactOpts, ballot, position, website)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz3 *Mgz3Transactor) AddGroup(opts *bind.TransactOpts, cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addGroup", cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz3 *Mgz3Session) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddGroup(&_Mgz3.TransactOpts, cPerson, category)
}

// AddGroup is a paid mutator transaction binding the contract method 0x1a409c24.
//
// Solidity: function addGroup(cPerson address, category uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) AddGroup(cPerson common.Address, category *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddGroup(&_Mgz3.TransactOpts, cPerson, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz3 *Mgz3Transactor) AddGroupCategory(opts *bind.TransactOpts, category [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addGroupCategory", category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz3 *Mgz3Session) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddGroupCategory(&_Mgz3.TransactOpts, category)
}

// AddGroupCategory is a paid mutator transaction binding the contract method 0xe4f97e5d.
//
// Solidity: function addGroupCategory(category bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) AddGroupCategory(category [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.AddGroupCategory(&_Mgz3.TransactOpts, category)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz3 *Mgz3Transactor) AddPreVote(opts *bind.TransactOpts, ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addPreVote", ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz3 *Mgz3Session) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddPreVote(&_Mgz3.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddPreVote is a paid mutator transaction binding the contract method 0x81779e38.
//
// Solidity: function addPreVote(ballot uint256, group uint256, fnumber bytes32, voter bytes20, candidate uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) AddPreVote(ballot *big.Int, group *big.Int, fnumber [32]byte, voter [20]byte, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddPreVote(&_Mgz3.TransactOpts, ballot, group, fnumber, voter, candidate)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Transactor) AddVoterToGroup(opts *bind.TransactOpts, voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "addVoterToGroup", voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Session) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddVoterToGroup(&_Mgz3.TransactOpts, voter, grp, position)
}

// AddVoterToGroup is a paid mutator transaction binding the contract method 0x6e85f432.
//
// Solidity: function addVoterToGroup(voter address, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) AddVoterToGroup(voter common.Address, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.AddVoterToGroup(&_Mgz3.TransactOpts, voter, grp, position)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Transactor) CloseBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "closeBallot", ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Session) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CloseBallot(&_Mgz3.TransactOpts, ballot)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xd045f1e4.
//
// Solidity: function closeBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) CloseBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CloseBallot(&_Mgz3.TransactOpts, ballot)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Transactor) CommitVotationPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "commitVotationPerPosition", ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Session) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CommitVotationPerPosition(&_Mgz3.TransactOpts, ballot, grp, position)
}

// CommitVotationPerPosition is a paid mutator transaction binding the contract method 0x6f3cb6d6.
//
// Solidity: function commitVotationPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) CommitVotationPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CommitVotationPerPosition(&_Mgz3.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Transactor) CommitVotationStatisticsPerPosition(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "commitVotationStatisticsPerPosition", ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3Session) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CommitVotationStatisticsPerPosition(&_Mgz3.TransactOpts, ballot, grp, position)
}

// CommitVotationStatisticsPerPosition is a paid mutator transaction binding the contract method 0xee931581.
//
// Solidity: function commitVotationStatisticsPerPosition(ballot uint256, grp uint256, position uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) CommitVotationStatisticsPerPosition(ballot *big.Int, grp *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.CommitVotationStatisticsPerPosition(&_Mgz3.TransactOpts, ballot, grp, position)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3Transactor) DefineCampaignIpfsInfo(opts *bind.TransactOpts, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineCampaignIpfsInfo", pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3Session) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCampaignIpfsInfo(&_Mgz3.TransactOpts, pos, value)
}

// DefineCampaignIpfsInfo is a paid mutator transaction binding the contract method 0xe9068033.
//
// Solidity: function defineCampaignIpfsInfo(pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineCampaignIpfsInfo(pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCampaignIpfsInfo(&_Mgz3.TransactOpts, pos, value)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz3 *Mgz3Transactor) DefineCanCancel(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineCanCancel", b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz3 *Mgz3Session) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCanCancel(&_Mgz3.TransactOpts, b)
}

// DefineCanCancel is a paid mutator transaction binding the contract method 0xadedcf79.
//
// Solidity: function defineCanCancel(b bool) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineCanCancel(b bool) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCanCancel(&_Mgz3.TransactOpts, b)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Transactor) DefineCurrentBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineCurrentBallot", ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Session) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCurrentBallot(&_Mgz3.TransactOpts, ballot)
}

// DefineCurrentBallot is a paid mutator transaction binding the contract method 0xd508d799.
//
// Solidity: function defineCurrentBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineCurrentBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCurrentBallot(&_Mgz3.TransactOpts, ballot)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz3 *Mgz3Transactor) DefineCurrentMessage(opts *bind.TransactOpts, message [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineCurrentMessage", message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz3 *Mgz3Session) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCurrentMessage(&_Mgz3.TransactOpts, message)
}

// DefineCurrentMessage is a paid mutator transaction binding the contract method 0xc4471494.
//
// Solidity: function defineCurrentMessage(message bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineCurrentMessage(message [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineCurrentMessage(&_Mgz3.TransactOpts, message)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz3 *Mgz3Transactor) DefineDisableCandidateLink(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineDisableCandidateLink", b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz3 *Mgz3Session) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineDisableCandidateLink(&_Mgz3.TransactOpts, b)
}

// DefineDisableCandidateLink is a paid mutator transaction binding the contract method 0xc6483da0.
//
// Solidity: function defineDisableCandidateLink(b bool) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineDisableCandidateLink(b bool) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineDisableCandidateLink(&_Mgz3.TransactOpts, b)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz3 *Mgz3Transactor) DefineStoppingAccessionToGroups(opts *bind.TransactOpts, str [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineStoppingAccessionToGroups", str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz3 *Mgz3Session) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineStoppingAccessionToGroups(&_Mgz3.TransactOpts, str)
}

// DefineStoppingAccessionToGroups is a paid mutator transaction binding the contract method 0xe4ed8665.
//
// Solidity: function defineStoppingAccessionToGroups(str bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineStoppingAccessionToGroups(str [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineStoppingAccessionToGroups(&_Mgz3.TransactOpts, str)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3Transactor) DefineTor(opts *bind.TransactOpts, person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "defineTor", person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3Session) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineTor(&_Mgz3.TransactOpts, person, pos, value)
}

// DefineTor is a paid mutator transaction binding the contract method 0x755bd7a5.
//
// Solidity: function defineTor(person address, pos uint256, value bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) DefineTor(person common.Address, pos *big.Int, value [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.DefineTor(&_Mgz3.TransactOpts, person, pos, value)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz3 *Mgz3Transactor) GiveRightToVote(opts *bind.TransactOpts, toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "giveRightToVote", toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz3 *Mgz3Session) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.GiveRightToVote(&_Mgz3.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x7db6a029.
//
// Solidity: function giveRightToVote(toVoter address, prefix uint256, pubkey bytes32, hashcode bytes32) returns()
func (_Mgz3 *Mgz3TransactorSession) GiveRightToVote(toVoter common.Address, prefix *big.Int, pubkey [32]byte, hashcode [32]byte) (*types.Transaction, error) {
	return _Mgz3.Contract.GiveRightToVote(&_Mgz3.TransactOpts, toVoter, prefix, pubkey, hashcode)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz3 *Mgz3Transactor) IncrementCancellations(opts *bind.TransactOpts, ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "incrementCancellations", ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz3 *Mgz3Session) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.IncrementCancellations(&_Mgz3.TransactOpts, ballot, candidate)
}

// IncrementCancellations is a paid mutator transaction binding the contract method 0x19f1e816.
//
// Solidity: function incrementCancellations(ballot uint256, candidate uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) IncrementCancellations(ballot *big.Int, candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.IncrementCancellations(&_Mgz3.TransactOpts, ballot, candidate)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz3 *Mgz3Transactor) IterateCandidatesCounter(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "iterateCandidatesCounter", ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz3 *Mgz3Session) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.IterateCandidatesCounter(&_Mgz3.TransactOpts, ballot)
}

// IterateCandidatesCounter is a paid mutator transaction binding the contract method 0x5182d3f1.
//
// Solidity: function iterateCandidatesCounter(ballot uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) IterateCandidatesCounter(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.IterateCandidatesCounter(&_Mgz3.TransactOpts, ballot)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz3 *Mgz3Transactor) RemoveRightToVote(opts *bind.TransactOpts, toVoter common.Address) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "removeRightToVote", toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz3 *Mgz3Session) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz3.Contract.RemoveRightToVote(&_Mgz3.TransactOpts, toVoter)
}

// RemoveRightToVote is a paid mutator transaction binding the contract method 0x685a9dd8.
//
// Solidity: function removeRightToVote(toVoter address) returns()
func (_Mgz3 *Mgz3TransactorSession) RemoveRightToVote(toVoter common.Address) (*types.Transaction, error) {
	return _Mgz3.Contract.RemoveRightToVote(&_Mgz3.TransactOpts, toVoter)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Transactor) StopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "stopBallot", ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Session) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.StopBallot(&_Mgz3.TransactOpts, ballot)
}

// StopBallot is a paid mutator transaction binding the contract method 0xb1bb8fdc.
//
// Solidity: function stopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) StopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.StopBallot(&_Mgz3.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Transactor) UnstopBallot(opts *bind.TransactOpts, ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "unstopBallot", ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3Session) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.UnstopBallot(&_Mgz3.TransactOpts, ballot)
}

// UnstopBallot is a paid mutator transaction binding the contract method 0x5e59d00a.
//
// Solidity: function unstopBallot(ballot uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) UnstopBallot(ballot *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.UnstopBallot(&_Mgz3.TransactOpts, ballot)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz3 *Mgz3Transactor) Vote(opts *bind.TransactOpts, ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.contract.Transact(opts, "vote", ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz3 *Mgz3Session) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.Vote(&_Mgz3.TransactOpts, ballot, grp, position, first_number, the_candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x8fbd82ef.
//
// Solidity: function vote(ballot uint256, grp uint256, position uint256, first_number bytes32, the_candidate uint256) returns()
func (_Mgz3 *Mgz3TransactorSession) Vote(ballot *big.Int, grp *big.Int, position *big.Int, first_number [32]byte, the_candidate *big.Int) (*types.Transaction, error) {
	return _Mgz3.Contract.Vote(&_Mgz3.TransactOpts, ballot, grp, position, first_number, the_candidate)
}
