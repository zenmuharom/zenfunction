package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestTrim(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "",
			Expected: "",
		},
		{
			Input:    "trim(000070000, 0)",
			Expected: "7",
		},
		{
			Input:    "trim(   hayo loooo   )",
			Expected: "hayo loooo",
		},
		{
			Input:    "trim(zzzzwoi taizzzzzz, z)",
			Expected: "woi tai",
		},
		{
			Input:    "trim()",
			Expected: "invalid parameter",
		},
		{
			Input:    "trim(007000, 0, 7)",
			Expected: "invalid parameter",
		},
	}

	for noTest, tc := range testCases {
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.NoError(t, err, errMsg)
		require.Equal(t, tc.Expected, result)
	}
}
