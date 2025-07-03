package function

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_MD5(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		// {
		// 	Input:    "md5(\"Darmaj4y4\")",
		// 	Expected: "3c63e9db65160b08944286486f6f0672",
		// },
		// {
		// 	Input:    "md5(\"bangsat\")",
		// 	Expected: "528f980649c80a7269402447b51e815a",
		// },
		// {
		// 	Input:    "md5(\"210020250429111311224559FINNET501001000006106174468792427953000000000000000DwZ3cOwCt22CzOgaghCnKdw0DIGceg5WapkOIPcziqCIX\")",
		// 	Expected: "7676b5f7aecbb596f6935fcb6e0a6662",
		// },
		{
			Input:    "md5(\"nHDe5eonSjTYF86Aq8gCXYkjHTQXM1rpHG1751524492\")",
			Expected: "71f8171d8ac51e6351148e4033cbbad3",
		},
	}

	for noTest, tc := range testCases {
		var result any
		res, err := assigner.ReadCommand(tc.Input)
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
