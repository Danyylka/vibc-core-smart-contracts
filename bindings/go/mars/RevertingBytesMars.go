// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mars

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AckPacket is an auto generated low-level Go binding around an user-defined struct.
type AckPacket struct {
	Success bool
	Data    []byte
}

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcEndpoint is an auto generated low-level Go binding around an user-defined struct.
type IbcEndpoint struct {
	PortId    string
	ChannelId [32]byte
}

// IbcPacket is an auto generated low-level Go binding around an user-defined struct.
type IbcPacket struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}

// RevertingBytesMarsMetaData contains all meta data concerning the RevertingBytesMars contract.
var RevertingBytesMarsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ackPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"connectedChannels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"greet\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"greetWithFee\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"skipAck\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvedPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedVersions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timeoutPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelClose\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelInitWithFee\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnRecvPacketRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnTimeoutPacket\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// RevertingBytesMarsABI is the input ABI used to generate the binding from.
// Deprecated: Use RevertingBytesMarsMetaData.ABI instead.
var RevertingBytesMarsABI = RevertingBytesMarsMetaData.ABI

// RevertingBytesMars is an auto generated Go binding around an Ethereum contract.
type RevertingBytesMars struct {
	RevertingBytesMarsCaller     // Read-only binding to the contract
	RevertingBytesMarsTransactor // Write-only binding to the contract
	RevertingBytesMarsFilterer   // Log filterer for contract events
}

// RevertingBytesMarsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevertingBytesMarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertingBytesMarsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevertingBytesMarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertingBytesMarsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevertingBytesMarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertingBytesMarsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevertingBytesMarsSession struct {
	Contract     *RevertingBytesMars // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RevertingBytesMarsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevertingBytesMarsCallerSession struct {
	Contract *RevertingBytesMarsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RevertingBytesMarsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevertingBytesMarsTransactorSession struct {
	Contract     *RevertingBytesMarsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RevertingBytesMarsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevertingBytesMarsRaw struct {
	Contract *RevertingBytesMars // Generic contract binding to access the raw methods on
}

// RevertingBytesMarsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevertingBytesMarsCallerRaw struct {
	Contract *RevertingBytesMarsCaller // Generic read-only contract binding to access the raw methods on
}

// RevertingBytesMarsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevertingBytesMarsTransactorRaw struct {
	Contract *RevertingBytesMarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevertingBytesMars creates a new instance of RevertingBytesMars, bound to a specific deployed contract.
func NewRevertingBytesMars(address common.Address, backend bind.ContractBackend) (*RevertingBytesMars, error) {
	contract, err := bindRevertingBytesMars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevertingBytesMars{RevertingBytesMarsCaller: RevertingBytesMarsCaller{contract: contract}, RevertingBytesMarsTransactor: RevertingBytesMarsTransactor{contract: contract}, RevertingBytesMarsFilterer: RevertingBytesMarsFilterer{contract: contract}}, nil
}

// NewRevertingBytesMarsCaller creates a new read-only instance of RevertingBytesMars, bound to a specific deployed contract.
func NewRevertingBytesMarsCaller(address common.Address, caller bind.ContractCaller) (*RevertingBytesMarsCaller, error) {
	contract, err := bindRevertingBytesMars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevertingBytesMarsCaller{contract: contract}, nil
}

// NewRevertingBytesMarsTransactor creates a new write-only instance of RevertingBytesMars, bound to a specific deployed contract.
func NewRevertingBytesMarsTransactor(address common.Address, transactor bind.ContractTransactor) (*RevertingBytesMarsTransactor, error) {
	contract, err := bindRevertingBytesMars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevertingBytesMarsTransactor{contract: contract}, nil
}

// NewRevertingBytesMarsFilterer creates a new log filterer instance of RevertingBytesMars, bound to a specific deployed contract.
func NewRevertingBytesMarsFilterer(address common.Address, filterer bind.ContractFilterer) (*RevertingBytesMarsFilterer, error) {
	contract, err := bindRevertingBytesMars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevertingBytesMarsFilterer{contract: contract}, nil
}

// bindRevertingBytesMars binds a generic wrapper to an already deployed contract.
func bindRevertingBytesMars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevertingBytesMarsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertingBytesMars *RevertingBytesMarsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertingBytesMars.Contract.RevertingBytesMarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertingBytesMars *RevertingBytesMarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.RevertingBytesMarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertingBytesMars *RevertingBytesMarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.RevertingBytesMarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertingBytesMars *RevertingBytesMarsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertingBytesMars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertingBytesMars *RevertingBytesMarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertingBytesMars *RevertingBytesMarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.contract.Transact(opts, method, params...)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_RevertingBytesMars *RevertingBytesMarsCaller) AckPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "ackPackets", arg0)

	outstruct := new(struct {
		Success bool
		Data    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_RevertingBytesMars *RevertingBytesMarsSession) AckPackets(arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	return _RevertingBytesMars.Contract.AckPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) AckPackets(arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	return _RevertingBytesMars.Contract.AckPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_RevertingBytesMars *RevertingBytesMarsCaller) ConnectedChannels(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "connectedChannels", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_RevertingBytesMars *RevertingBytesMarsSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _RevertingBytesMars.Contract.ConnectedChannels(&_RevertingBytesMars.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _RevertingBytesMars.Contract.ConnectedChannels(&_RevertingBytesMars.CallOpts, arg0)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsSession) Dispatcher() (common.Address, error) {
	return _RevertingBytesMars.Contract.Dispatcher(&_RevertingBytesMars.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) Dispatcher() (common.Address, error) {
	return _RevertingBytesMars.Contract.Dispatcher(&_RevertingBytesMars.CallOpts)
}

// OnRecvPacket is a free data retrieval call binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns((bool,bytes) ack, bool skipAck)
func (_RevertingBytesMars *RevertingBytesMarsCaller) OnRecvPacket(opts *bind.CallOpts, arg0 IbcPacket) (struct {
	Ack     AckPacket
	SkipAck bool
}, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "onRecvPacket", arg0)

	outstruct := new(struct {
		Ack     AckPacket
		SkipAck bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ack = *abi.ConvertType(out[0], new(AckPacket)).(*AckPacket)
	outstruct.SkipAck = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// OnRecvPacket is a free data retrieval call binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns((bool,bytes) ack, bool skipAck)
func (_RevertingBytesMars *RevertingBytesMarsSession) OnRecvPacket(arg0 IbcPacket) (struct {
	Ack     AckPacket
	SkipAck bool
}, error) {
	return _RevertingBytesMars.Contract.OnRecvPacket(&_RevertingBytesMars.CallOpts, arg0)
}

// OnRecvPacket is a free data retrieval call binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns((bool,bytes) ack, bool skipAck)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) OnRecvPacket(arg0 IbcPacket) (struct {
	Ack     AckPacket
	SkipAck bool
}, error) {
	return _RevertingBytesMars.Contract.OnRecvPacket(&_RevertingBytesMars.CallOpts, arg0)
}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns()
func (_RevertingBytesMars *RevertingBytesMarsCaller) OnTimeoutPacket(opts *bind.CallOpts, arg0 IbcPacket) error {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "onTimeoutPacket", arg0)

	if err != nil {
		return err
	}

	return err

}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnTimeoutPacket(arg0 IbcPacket) error {
	return _RevertingBytesMars.Contract.OnTimeoutPacket(&_RevertingBytesMars.CallOpts, arg0)
}

// OnTimeoutPacket is a free data retrieval call binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) ) view returns()
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) OnTimeoutPacket(arg0 IbcPacket) error {
	return _RevertingBytesMars.Contract.OnTimeoutPacket(&_RevertingBytesMars.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsSession) Owner() (common.Address, error) {
	return _RevertingBytesMars.Contract.Owner(&_RevertingBytesMars.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) Owner() (common.Address, error) {
	return _RevertingBytesMars.Contract.Owner(&_RevertingBytesMars.CallOpts)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsCaller) RecvedPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "recvedPackets", arg0)

	outstruct := new(struct {
		Src              IbcEndpoint
		Dest             IbcEndpoint
		Sequence         uint64
		Data             []byte
		TimeoutHeight    Height
		TimeoutTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Src = *abi.ConvertType(out[0], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Dest = *abi.ConvertType(out[1], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Sequence = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.TimeoutHeight = *abi.ConvertType(out[4], new(Height)).(*Height)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsSession) RecvedPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _RevertingBytesMars.Contract.RecvedPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) RecvedPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _RevertingBytesMars.Contract.RecvedPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_RevertingBytesMars *RevertingBytesMarsCaller) SupportedVersions(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "supportedVersions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_RevertingBytesMars *RevertingBytesMarsSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _RevertingBytesMars.Contract.SupportedVersions(&_RevertingBytesMars.CallOpts, arg0)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _RevertingBytesMars.Contract.SupportedVersions(&_RevertingBytesMars.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsCaller) TimeoutPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	var out []interface{}
	err := _RevertingBytesMars.contract.Call(opts, &out, "timeoutPackets", arg0)

	outstruct := new(struct {
		Src              IbcEndpoint
		Dest             IbcEndpoint
		Sequence         uint64
		Data             []byte
		TimeoutHeight    Height
		TimeoutTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Src = *abi.ConvertType(out[0], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Dest = *abi.ConvertType(out[1], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Sequence = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.TimeoutHeight = *abi.ConvertType(out[4], new(Height)).(*Height)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsSession) TimeoutPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _RevertingBytesMars.Contract.TimeoutPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_RevertingBytesMars *RevertingBytesMarsCallerSession) TimeoutPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _RevertingBytesMars.Contract.TimeoutPackets(&_RevertingBytesMars.CallOpts, arg0)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsTransactor) Greet(opts *bind.TransactOpts, message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "greet", message, channelId, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsSession) Greet(message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.Greet(&_RevertingBytesMars.TransactOpts, message, channelId, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) Greet(message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.Greet(&_RevertingBytesMars.TransactOpts, message, channelId, timeoutTimestamp)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsTransactor) GreetWithFee(opts *bind.TransactOpts, message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "greetWithFee", message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsSession) GreetWithFee(message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.GreetWithFee(&_RevertingBytesMars.TransactOpts, message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) GreetWithFee(message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.GreetWithFee(&_RevertingBytesMars.TransactOpts, message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onAcknowledgementPacket", arg0, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnAcknowledgementPacket(arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnAcknowledgementPacket(&_RevertingBytesMars.TransactOpts, arg0, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnAcknowledgementPacket(arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnAcknowledgementPacket(&_RevertingBytesMars.TransactOpts, arg0, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onChanCloseConfirm", channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanCloseConfirm(&_RevertingBytesMars.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanCloseConfirm(&_RevertingBytesMars.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onChanCloseInit", channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanCloseInit(&_RevertingBytesMars.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanCloseInit(&_RevertingBytesMars.TransactOpts, channelId, arg1, arg2)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onChanOpenAck", channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenAck(&_RevertingBytesMars.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenAck(&_RevertingBytesMars.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenConfirm(&_RevertingBytesMars.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenConfirm(&_RevertingBytesMars.TransactOpts, channelId)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_RevertingBytesMars *RevertingBytesMarsTransactor) OnChanOpenTry(opts *bind.TransactOpts, arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "onChanOpenTry", arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_RevertingBytesMars *RevertingBytesMarsSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenTry(&_RevertingBytesMars.TransactOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.OnChanOpenTry(&_RevertingBytesMars.TransactOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) RenounceOwnership() (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.RenounceOwnership(&_RevertingBytesMars.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.RenounceOwnership(&_RevertingBytesMars.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TransferOwnership(&_RevertingBytesMars.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TransferOwnership(&_RevertingBytesMars.TransactOpts, newOwner)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) TriggerChannelClose(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "triggerChannelClose", channelId)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) TriggerChannelClose(channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelClose(&_RevertingBytesMars.TransactOpts, channelId)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) TriggerChannelClose(channelId [32]byte) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelClose(&_RevertingBytesMars.TransactOpts, channelId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) TriggerChannelInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "triggerChannelInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelInit(&_RevertingBytesMars.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelInit(&_RevertingBytesMars.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) TriggerChannelInitWithFee(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.Transact(opts, "triggerChannelInitWithFee", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) TriggerChannelInitWithFee(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelInitWithFee(&_RevertingBytesMars.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) TriggerChannelInitWithFee(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.TriggerChannelInitWithFee(&_RevertingBytesMars.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertingBytesMars.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertingBytesMars *RevertingBytesMarsSession) Receive() (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.Receive(&_RevertingBytesMars.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertingBytesMars *RevertingBytesMarsTransactorSession) Receive() (*types.Transaction, error) {
	return _RevertingBytesMars.Contract.Receive(&_RevertingBytesMars.TransactOpts)
}

// RevertingBytesMarsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RevertingBytesMars contract.
type RevertingBytesMarsOwnershipTransferredIterator struct {
	Event *RevertingBytesMarsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RevertingBytesMarsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevertingBytesMarsOwnershipTransferred)
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
		it.Event = new(RevertingBytesMarsOwnershipTransferred)
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
func (it *RevertingBytesMarsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevertingBytesMarsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevertingBytesMarsOwnershipTransferred represents a OwnershipTransferred event raised by the RevertingBytesMars contract.
type RevertingBytesMarsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RevertingBytesMars *RevertingBytesMarsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RevertingBytesMarsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RevertingBytesMars.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RevertingBytesMarsOwnershipTransferredIterator{contract: _RevertingBytesMars.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RevertingBytesMars *RevertingBytesMarsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RevertingBytesMarsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RevertingBytesMars.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevertingBytesMarsOwnershipTransferred)
				if err := _RevertingBytesMars.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RevertingBytesMars *RevertingBytesMarsFilterer) ParseOwnershipTransferred(log types.Log) (*RevertingBytesMarsOwnershipTransferred, error) {
	event := new(RevertingBytesMarsOwnershipTransferred)
	if err := _RevertingBytesMars.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
