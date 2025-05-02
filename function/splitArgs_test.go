package function

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SplitArgs(t *testing.T) {

	billerResp := "FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepat waktu:|Download PLN Mobile"
	testCases := []TestCaseSplit{
		{
			Input: "test",
			Expected: []string{
				"test",
			},
		},
		{
			Input: "\"ntb\", item.randomInt(12)",
			Expected: []string{
				"ntb",
				"item.randomInt(12)",
			},
		},
		{
			Input: strconv.Quote(billerResp) + ", item.randomInt(12)",
			Expected: []string{
				"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepat waktu:|Download PLN Mobile",
				"item.randomInt(12)",
			},
		},
		{
			Input: strconv.Quote(billerResp) + ", 0, 20",
			Expected: []string{
				"FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepat waktu:|Download PLN Mobile",
				"0",
				"20",
			},
		},
		{
			Input: "uuid(), \"-\"",
			Expected: []string{
				"uuid()",
				"-",
			},
		},
		{
			Input: "\"\", 120",
			Expected: []string{
				"",
				"120",
			},
		},
		{
			Input: "\"hallo \", \"anjing \", \"woi 1 \"",
			Expected: []string{
				"hallo ",
				"anjing ",
				"woi 1 ",
			},
		},
		{
			Input: strconv.Quote("FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:") + ", 0, 100",
			Expected: []string{
				"FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:",
				"0",
				"100",
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
			errMsg = fmt.Sprintf("No Test.%v: %v expected (%v): %v result (%#v): %v", (noTest + 1), "args total not same with expected total", len(tc.Expected), tc.Expected, len(result), result)
			err = errors.New(errMsg)
		}

		// require.Equal(t, len(tc.Expected), len(result))
		require.NoError(t, err, errMsg)
	}
}
