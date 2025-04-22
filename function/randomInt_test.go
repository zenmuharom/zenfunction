package function

import (
	"fmt"
	"strings"
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
			Expected: "17",
		},
		{
			Input:    "randomInt(1, 2, 3)",
			Expected: "invalid parameter",
		},
		{
			Input:    "randomInt(20)",
			Expected: "20",
		},
	}

	for noTest, tc := range testCases {
		var result any
		res, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.NoError(t, err, errMsg)
		if result != "invalid parameter" {
			switch v := res.(type) {
			case string:

				if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
					// safe unwrap only outer quotes
					v = v[1 : len(v)-1]
				}
				result = v

				require.NoError(t, err, errMsg)

				if fmt.Sprintf("%v", result) != "invalid parameter" {
					require.Equal(t, tc.Expected, fmt.Sprintf("%v", len(result.(string))))
				}
			default:
				// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
				result = fmt.Sprintf("%v", v)
			}
		} else {
			require.Equal(t, tc.Expected, result)
		}
	}
}
