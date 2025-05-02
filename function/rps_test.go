package function

import (
	"fmt"
	"strconv"
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
		res, err := assigner.ReadCommandV2("string", tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}

		require.NoError(t, err, errMsg)
		require.Equal(t, tc.Expected, res)
	}
}
