package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestRandomInt(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "",
			Expected: "0",
		},
		{
			Input:    "randomInt(1, 2)",
			Expected: "2",
		},
		{
			Input:    "randomInt(1, 4)",
			Expected: "4",
		},
		{
			Input:    "randomInt()",
			Expected: "19",
		},
		{
			Input:    "randomInt(1, 2, 3)",
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
		if result != "invalid parameter" {
			require.Equal(t, tc.Expected, fmt.Sprintf("%v", len(fmt.Sprintf("%v", result))))
		} else {
			require.Equal(t, tc.Expected, result)
		}
	}
}
