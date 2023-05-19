package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestSubstr(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "",
			Expected: "",
		},
		{
			Input:    "substr(test woi, 0, 4)",
			Expected: "test",
		},
		{
			Input:    "substr(tets lah, 5)",
			Expected: "lah",
		},
		{
			Input:    "substr(test lah)",
			Expected: "invalid parameter",
		},
		{
			Input:    "substr()",
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
