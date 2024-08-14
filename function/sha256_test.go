package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_Sha256(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "sha256(Darmaj4y4)",
			Expected: "1532a22b0d3525e94c7455945a30943a205ed15cde722532f4f4cda08ded0f88",
		},
		{
			Input:    "sha256(bangsat)",
			Expected: "4077902c11d06e1058060d1bd789e7c5a1db2bbecbc39b8d161c1d131c2ac1b7",
		},
		{
			Input:    "sha256(71947bea4c63-63ad-3489-a73a-2d73668356a460236703)",
			Expected: "4077902c11d06e1058060d1bd789e7c5a1db2bbecbc39b8d161c1d131c2ac1b7",
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
