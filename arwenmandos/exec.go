package arwenmandos

import (
	worldhook "github.com/Reshusk23/dme-util/mock-hook-blockchain"
	cryptohook "github.com/Reshusk23/dme-util/mock-hook-crypto"
	mc "github.com/Reshusk23/dme-util/test-util/mandos/controller"
	mjparse "github.com/Reshusk23/dme-util/test-util/mandos/json/parse"
	vmcommon "github.com/Reshusk23/dme-vm-common"
	vmi "github.com/Reshusk23/dme-vm-common"
	arwen "github.com/Reshusk23/dme-vm/arwen"
	arwenHost "github.com/Reshusk23/dme-vm/arwen/host"
	"github.com/Reshusk23/dme-vm/config"
)

// TestVMType is the VM type argument we use in tests.
var TestVMType = []byte{0, 0}

// ArwenTestExecutor parses, interprets and executes both .test.json tests and .scen.json scenarios with Arwen.
type ArwenTestExecutor struct {
	fileResolver mjparse.FileResolver
	World        *worldhook.BlockchainHookMock
	vm           vmi.VMExecutionHandler
	checkGas     bool
}

var _ mc.TestExecutor = (*ArwenTestExecutor)(nil)
var _ mc.ScenarioExecutor = (*ArwenTestExecutor)(nil)

// NewArwenTestExecutor prepares a new ArwenTestExecutor instance.
func NewArwenTestExecutor() (*ArwenTestExecutor, error) {
	world := worldhook.NewMock()
	world.EnableMockAddressGeneration()

	blockGasLimit := uint64(10000000)
	gasSchedule := config.MakeGasMapForTests()
	vm, err := arwenHost.NewArwenVM(world, cryptohook.KryptoHookMockInstance, &arwen.VMHostParameters{
		VMType:                    TestVMType,
		BlockGasLimit:             blockGasLimit,
		GasSchedule:               gasSchedule,
		ProtocolBuiltinFunctions:  make(vmcommon.FunctionNames),
		ReshuskProtectedKeyPrefix: []byte(ReshuskProtectedKeyPrefix),
	})
	if err != nil {
		return nil, err
	}
	return &ArwenTestExecutor{
		fileResolver: nil,
		World:        world,
		vm:           vm,
		checkGas:     true,
	}, nil
}

// GetVM yields a reference to the VMExecutionHandler used.
func (ae *ArwenTestExecutor) GetVM() vmi.VMExecutionHandler {
	return ae.vm
}
