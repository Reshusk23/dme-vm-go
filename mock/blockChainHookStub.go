package mock

import (
	vmcommon "github.com/Reshusk23/dme-vm-common"
	"github.com/Reshusk23/dme-vm-common/data/dct"
)

var _ vmcommon.BlockchainHook = (*BlockchainHookStub)(nil)

type BlockchainHookStub struct {
	NewAddressCalled              func(creatorAddress []byte, creatorNonce uint64, vmType []byte) ([]byte, error)
	GetStorageDataCalled          func(accountsAddress []byte, index []byte) ([]byte, error)
	GetBlockHashCalled            func(nonce uint64) ([]byte, error)
	LastNonceCalled               func() uint64
	LastRoundCalled               func() uint64
	LastTimeStampCalled           func() uint64
	LastRandomSeedCalled          func() []byte
	LastEpochCalled               func() uint32
	GetStateRootHashCalled        func() []byte
	CurrentNonceCalled            func() uint64
	CurrentRoundCalled            func() uint64
	CurrentTimeStampCalled        func() uint64
	CurrentRandomSeedCalled       func() []byte
	CurrentEpochCalled            func() uint32
	ProcessBuiltInFunctionCalled  func(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error)
	GetBuiltinFunctionNamesCalled func() vmcommon.FunctionNames
	GetAllStateCalled             func(address []byte) (map[string][]byte, error)
	GetUserAccountCalled          func(address []byte) (vmcommon.UserAccountHandler, error)
	GetShardOfAddressCalled       func(address []byte) uint32
	IsSmartContractCalled         func(address []byte) bool
}

// ClearCompiledCodes implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) ClearCompiledCodes() {
	panic("unimplemented")
}

// GetCode implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) GetCode(vmcommon.UserAccountHandler) []byte {
	panic("unimplemented")
}

// GetCompiledCode implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) GetCompiledCode(codeHash []byte) (bool, []byte) {
	panic("unimplemented")
}

// GetDCTToken implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) GetDCTToken(address []byte, tokenID []byte, nonce uint64) (*dct.DCToken, error) {
	panic("unimplemented")
}

// GetSnapshot implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) GetSnapshot() int {
	panic("unimplemented")
}

// IsInterfaceNil implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) IsInterfaceNil() bool {
	panic("unimplemented")
}

// IsPayable implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) IsPayable(address []byte) (bool, error) {
	panic("unimplemented")
}

// RevertToSnapshot implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) RevertToSnapshot(snapshot int) error {
	panic("unimplemented")
}

// SaveCompiledCode implements vmcommon.BlockchainHook.
func (*BlockchainHookStub) SaveCompiledCode(codeHash []byte, code []byte) {
	panic("unimplemented")
}

func (b *BlockchainHookStub) NewAddress(creatorAddress []byte, creatorNonce uint64, vmType []byte) ([]byte, error) {
	if b.NewAddressCalled != nil {
		return b.NewAddressCalled(creatorAddress, creatorNonce, vmType)
	}
	return []byte("newAddress"), nil
}

func (b *BlockchainHookStub) GetStorageData(accountAddress []byte, index []byte) ([]byte, error) {
	if b.GetStorageDataCalled != nil {
		return b.GetStorageDataCalled(accountAddress, index)
	}
	return nil, nil
}

func (b *BlockchainHookStub) GetBlockhash(nonce uint64) ([]byte, error) {
	if b.GetBlockHashCalled != nil {
		return b.GetBlockHashCalled(nonce)
	}
	return []byte("roothash"), nil
}

func (b *BlockchainHookStub) LastNonce() uint64 {
	if b.LastNonceCalled != nil {
		return b.LastNonceCalled()
	}
	return 0
}

func (b *BlockchainHookStub) LastRound() uint64 {
	if b.LastRoundCalled != nil {
		return b.LastRoundCalled()
	}
	return 0
}

func (b *BlockchainHookStub) LastTimeStamp() uint64 {
	if b.LastTimeStampCalled != nil {
		return b.LastTimeStampCalled()
	}
	return 0
}

func (b *BlockchainHookStub) LastRandomSeed() []byte {
	if b.LastRandomSeedCalled != nil {
		return b.LastRandomSeedCalled()
	}
	return []byte("seed")
}

func (b *BlockchainHookStub) LastEpoch() uint32 {
	if b.LastEpochCalled != nil {
		return b.LastEpochCalled()
	}
	return 0
}

func (b *BlockchainHookStub) GetStateRootHash() []byte {
	if b.GetStateRootHashCalled != nil {
		return b.GetStateRootHashCalled()
	}
	return []byte("roothash")
}

func (b *BlockchainHookStub) CurrentNonce() uint64 {
	if b.CurrentNonceCalled != nil {
		return b.CurrentNonceCalled()
	}
	return 0
}

func (b *BlockchainHookStub) CurrentRound() uint64 {
	if b.CurrentRoundCalled != nil {
		return b.CurrentRoundCalled()
	}
	return 0
}

func (b *BlockchainHookStub) CurrentTimeStamp() uint64 {
	if b.CurrentTimeStampCalled != nil {
		return b.CurrentTimeStampCalled()
	}
	return 0
}

func (b *BlockchainHookStub) CurrentRandomSeed() []byte {
	if b.CurrentRandomSeedCalled != nil {
		return b.CurrentRandomSeedCalled()
	}
	return []byte("seed")
}

func (b *BlockchainHookStub) CurrentEpoch() uint32 {
	if b.CurrentEpochCalled != nil {
		return b.CurrentEpochCalled()
	}
	return 0
}

func (b *BlockchainHookStub) ProcessBuiltInFunction(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	if b.ProcessBuiltInFunctionCalled != nil {
		return b.ProcessBuiltInFunctionCalled(input)
	}
	return &vmcommon.VMOutput{}, nil
}

func (b *BlockchainHookStub) GetBuiltinFunctionNames() vmcommon.FunctionNames {
	if b.GetBuiltinFunctionNamesCalled != nil {
		return b.GetBuiltinFunctionNamesCalled()
	}
	return make(vmcommon.FunctionNames)
}

func (b *BlockchainHookStub) GetAllState(address []byte) (map[string][]byte, error) {
	if b.GetAllStateCalled != nil {
		return b.GetAllStateCalled(address)
	}
	return nil, nil
}

func (b *BlockchainHookStub) GetUserAccount(address []byte) (vmcommon.UserAccountHandler, error) {
	if b.GetUserAccountCalled != nil {
		return b.GetUserAccountCalled(address)
	}
	return nil, nil
}

func (b *BlockchainHookStub) GetShardOfAddress(address []byte) uint32 {
	if b.GetShardOfAddressCalled != nil {
		return b.GetShardOfAddressCalled(address)
	}
	return 0
}

func (b *BlockchainHookStub) IsSmartContract(address []byte) bool {
	if b.IsSmartContractCalled != nil {
		return b.IsSmartContractCalled(address)
	}
	return false
}
