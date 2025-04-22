package function

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_BasicAuth(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "basicAuth(admin:test123)",
			Expected: "Basic YWRtaW46dGVzdDEyMw==",
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
