package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestDateAdd(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "dateAdd(2006/01/02 15:04:05, 2023/05/17 10:33:01, 5, second)",
			Expected: "2023/05/17 10:33:06",
		},
		{
			Input:    "dateAdd(2006/01/02 15:04:05, 2023/05/17 10:33:01, 2, minute)",
			Expected: "2023/05/17 10:35:01",
		},
		{
			Input:    "dateAdd(2006/01/02 15:04:05, 2023/05/17 10:33:01, 20, hour)",
			Expected: "2023/05/18 06:33:01",
		},
		{
			Input:    "dateAdd(2006/01/02 15:04:05, 2023/05/17 10:33:01, 30, day)",
			Expected: "2023/06/16 10:33:01",
		},
		{
			Input:    "dateAdd(2006-01-02 15:04:05, 2023/05/17, 2, month)",
			Expected: "2023-07-17 00:00:00",
		},
		{
			Input:    "dateAdd()",
			Expected: "invalid parameter",
		},
		{
			Input:    "dateAdd(2006)",
			Expected: "invalid parameter",
		},
		{
			Input:    "dateAdd(2006-01-02, 2023/05/17)",
			Expected: "invalid parameter",
		},
		{
			Input:    "dateAdd(2006-01-02 15:04:05, 2023/05/17, 2)",
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
