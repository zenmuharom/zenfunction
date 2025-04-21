package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestLtrim(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "ltrim(\"          halooooo       \")",
			Expected: "halooooo       ",
		},
		{
			Input:    "ltrim(00000000055000, 0)",
			Expected: "55000",
		},
		{
			Input:    "ltrim(zzzz125000, z)",
			Expected: "125000",
		},
		{
			Input:    "ltrim(zzzz35000, t)",
			Expected: "zzzz35000",
		},
		{
			Input:    "ltrim()",
			Expected: "invalid parameter",
		},
		{
			Input:    "ltrim(0002500, 0, 3)",
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
