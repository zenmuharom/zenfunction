package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_Sha1(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "sha1(Darmaj4y4)",
			Expected: "2cff38c4dd478b1c80f4612bbed67af8db0865a9",
		},
		{
			Input:    "sha1(bangsat)",
			Expected: "93bccfd866e61053c3f769435b40574b94352a69",
		},
		{
			Input:    "sha1(71947bea4c63-63ad-3489-a73a-2d73668356a460236703)",
			Expected: "a95d947b6c5696c0713b6d6427bc50f6c16175cd",
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
