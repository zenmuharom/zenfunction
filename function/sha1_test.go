package function

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_Sha1(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "sha1(\"Darmaj4y4\")",
			Expected: "2cff38c4dd478b1c80f4612bbed67af8db0865a9",
		},
		{
			Input:    "sha1(\"bangsat\")",
			Expected: "93bccfd866e61053c3f769435b40574b94352a69",
		},
		{
			Input:    "sha1(\"71947bea4c63-63ad-3489-a73a-2d73668356a460236703\")",
			Expected: "a95d947b6c5696c0713b6d6427bc50f6c16175cd",
		},
		{
			Input:    "sha1(\"210020250429111311224559FINNET501001000006106174468792427953000000000000000DwZ3cOwCt22CzOgaghCnKdw0DIGceg5WapkOIPcziqCIX\")",
			Expected: "9feaa8545018006f0315eba52e6bfaeeb68aff9a",
		},
	}

	for noTest, tc := range testCases {
		var res any
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}

		switch v := res.(type) {
		case string:

			if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
				// safe unwrap only outer quotes
				v = v[1 : len(v)-1]
			}
			result = v

			require.NoError(t, err, errMsg)

			require.Equal(t, tc.Expected, result)
		default:
			// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
			result = fmt.Sprintf("%v", v)
		}
	}
}
