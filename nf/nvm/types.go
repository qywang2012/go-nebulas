package nvm

import (
	"errors"

	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/core/pb"
	"github.com/nebulasio/go-nebulas/core/state"
	"github.com/nebulasio/go-nebulas/storage"
	"github.com/nebulasio/go-nebulas/util"
	"github.com/nebulasio/go-nebulas/util/byteutils"
)

// Error Types
var (
	ErrEngineRepeatedStart      = errors.New("engine repeated start")
	ErrEngineNotStart           = errors.New("engine not start")
	ErrContextConstructArrEmpty = errors.New("context construct err by args empty")
	ErrEngineNotFound           = errors.New("Failed to get engine")

	ErrDisallowCallPrivateFunction     = errors.New("disallow call private function")
	ErrExecutionTimeout                = errors.New("execution timeout")
	ErrInsufficientGas                 = errors.New("insufficient gas")
	ErrExceedMemoryLimits              = errors.New("exceed memory limits")
	ErrInjectTracingInstructionFailed  = errors.New("inject tracing instructions failed")
	ErrTranspileTypeScriptFailed       = errors.New("transpile TypeScript failed")
	ErrUnsupportedSourceType           = errors.New("unsupported source type")
	ErrArgumentsFormat                 = errors.New("arguments format error")
	ErrLimitHasEmpty                   = errors.New("limit args has empty")
	ErrSetMemorySmall                  = errors.New("set memory small than v8 limit")
	ErrDisallowCallNotStandardFunction = errors.New("disallow call not standard function")

	ErrNvmNumLimit          = errors.New("out of limit nvm count")
	ErrInnerTransferFailed  = errors.New("inner transfer failed")
	ErrInnerInsufficientGas = errors.New("preparation inner nvm insufficient gas")
	ErrInnerInsufficientMem = errors.New("preparation inner nvm insufficient mem")

	ErrOutOfNvmMaxGasLimit = errors.New("out of nvm max gas limit")
)

//define
var (
	EventNameSpaceContract    = "chain.contract" //ToRefine: move to core
	InnerTransactionErrPrefix = "inner transation err ["
	InnerTransactionResult    = "] result ["
	InnerTransactionErrEnding = "] engine index:%v"
)

//common err
var (
	ErrKeyNotFound = storage.ErrKeyNotFound
)

//transfer err code enum
const (
	TransferFuncSuccess = iota
	TransferSuccess
	TransferGetEngineErr
	TransferAddressParseErr
	TransferGetAccountErr
	TransferStringToBigIntErr
	TransferSubBalance
	TransferAddBalance
	TransferRecordEventFailed
	TransferAddressFailed
)

//the max recent block number can query
const (
	maxQueryBlockInfoValidTime = 30
	maxBlockOffset             = maxQueryBlockInfoValidTime * 24 * 3600 * 1000 / 15000 //TODO:dpos.BlockIntervalInMs
)

// define gas consume
const (
	// crypto
	CryptoSha256GasBase         = 20000
	CryptoSha3256GasBase        = 20000
	CryptoRipemd160GasBase      = 20000
	CryptoRecoverAddressGasBase = 100000
	CryptoMd5GasBase            = 6000
	CryptoBase64GasBase         = 3000

	//In blockChain
	GetTxByHashGasBase     = 1000
	GetAccountStateGasBase = 2000
	TransferGasBase        = 2000
	VerifyAddressGasBase   = 100
	GetPreBlockHashGasBase = 2000
	GetPreBlockSeedGasBase = 2000
)

//MultiV8error err info, err only in InnerContractFunc .so not to deine #
type MultiV8error struct {
	errCode uint32
	index   uint32
	errStr  string
}

//nvm args define //TODO: 确定所有值的大小
var (
	MultiNvmMax               = 3
	GetTxByHashFuncCost       = 1000
	GetAccountStateFuncCost   = 1000
	TransferFuncCost          = 2000
	VerifyAddressFuncCost     = 100
	GetContractSourceFuncCost = 5000
	InnerContractFuncCost     = 30000
	GetTxRandomFuncCost       = 1000
)

// Block interface breaks cycle import dependency and hides unused services.
type Block interface {
	Hash() byteutils.Hash
	Height() uint64 // ToAdd: timestamp interface
	Timestamp() int64
	RandomSeed() string
	RandomAvailable() bool
	DateAvailable() bool
}

// Transaction interface breaks cycle import dependency and hides unused services.
type Transaction interface {
	ChainID() uint32
	Hash() byteutils.Hash
	From() *core.Address
	To() *core.Address
	Value() *util.Uint128
	Nonce() uint64
	Timestamp() int64
	GasPrice() *util.Uint128
	GasLimit() *util.Uint128
}

// Account interface breaks cycle import dependency and hides unused services.
type Account interface {
	Address() byteutils.Hash
	Balance() *util.Uint128
	Nonce() uint64
	AddBalance(value *util.Uint128) error
	SubBalance(value *util.Uint128) error
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Del(key []byte) error
	ContractMeta() *corepb.ContractMeta
}

// WorldState interface breaks cycle import dependency and hides unused services.
type WorldState interface {
	GetOrCreateUserAccount(addr byteutils.Hash) (state.Account, error)
	GetTx(txHash byteutils.Hash) ([]byte, error)
	RecordEvent(txHash byteutils.Hash, event *state.Event)
	GetBlockHashByHeight(height uint64) ([]byte, error)
	GetBlock(txHash byteutils.Hash) ([]byte, error)
	CreateContractAccount(owner byteutils.Hash, birthPlace byteutils.Hash, contractMeta *corepb.ContractMeta) (state.Account, error)
	Dynasty() ([]byteutils.Hash, error)
	DynastyRoot() byteutils.Hash
	FetchEvents(byteutils.Hash) ([]*state.Event, error)
	GetContractAccount(addr byteutils.Hash) (state.Account, error)
	PutTx(txHash byteutils.Hash, txBytes []byte) error
	RecordGas(from string, gas *util.Uint128) error
	Reset(addr byteutils.Hash) error //Need to consider risk
}
