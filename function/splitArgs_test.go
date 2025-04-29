package function

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SplitArgs(t *testing.T) {

	testCases := []TestCaseSplit{
		// {
		// 	Input: `"ntb",item.randomInt(12)`,
		// 	Expected: []string{
		// 		"ntb",
		// 		"item.randomInt(12)",
		// 	},
		// },
		// {
		// 	Input: "\"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \\\"Tag\\\" Tepat waktu:|Download PLN Mobile\", item.randomInt(12)",
		// 	Expected: []string{
		// 		"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \\\"Tag\\\" Tepat waktu:|Download PLN Mobile",
		// 		"item.randomInt(12)",
		// 	},
		// },
		// {
		// 	Input: "\"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \\\"Tag\\\" Tepat waktu:|Download PLN Mobile\", 0, 20",
		// 	Expected: []string{
		// 		"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \\\"Tag\\\" Tepat waktu:|Download PLN Mobile",
		// 		"0",
		// 		"20",
		// 	},
		// },
		// {
		// 	Input: "uuid(), \"-\"",
		// 	Expected: []string{
		// 		"uuid()",
		// 		"-",
		// 	},
		// },
		{
			Input: "\"\", 120",
			Expected: []string{
				"",
				"120",
			},
		},
	}

	for noTest, tc := range testCases {

		var err error

		result := splitArgs(tc.Input)

		errMsg := ""
		if len(tc.Expected) == len(result) {
			for i := 0; i < len(result); i++ {
				require.Equal(t, tc.Expected[i], result[i])
			}

		} else {
			errMsg = fmt.Sprintf("No Test.%v: %v expected (%v): %v result (%v): %v", (noTest + 1), "args total not same with expected total", len(tc.Expected), tc.Expected, len(result), result)
			err = errors.New(errMsg)
		}

		// require.Equal(t, len(tc.Expected), len(result))
		require.NoError(t, err, errMsg)
	}
}
