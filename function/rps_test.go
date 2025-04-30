package function

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_Rps(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "rps(zaw, 10)",
			Expected: "zaw       ",
		},
		{
			Input:    "rps(\"zeni\", 10)",
			Expected: "zeni      ",
		},
		{
			Input:    "rps(\"zeni\", 2)",
			Expected: "zeni",
		},
		{
			Input:    "rps(\"\", 120)",
			Expected: "                                                                                                                        ",
		},
		{
			Input:    "rps(" + strconv.Quote("FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:") + ", 120)",
			Expected: "FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:                                          ",
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
