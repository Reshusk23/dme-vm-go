package contexts

import (
	"testing"

	vmcommon "github.com/Reshusk23/dme-vm-common"
	"github.com/Reshusk23/dme-vm/arwen"
	"github.com/stretchr/testify/require"
)

func TestReservedFunctions_IsFunctionReserved(t *testing.T) {
	scAPINames := vmcommon.FunctionNames{
		"rockets": {},
	}

	fromProtocol := vmcommon.FunctionNames{
		"protocolFunctionFoo": {},
		"protocolFunctionBar": {},
	}

	reserved := NewReservedFunctions(scAPINames, fromProtocol)

	require.False(t, reserved.IsReserved("foo"))
	require.True(t, reserved.IsReserved("rockets"))
	require.True(t, reserved.IsReserved("protocolFunctionFoo"))
	require.True(t, reserved.IsReserved("protocolFunctionBar"))
	require.True(t, reserved.IsReserved(arwen.UpgradeFunctionName))
}
