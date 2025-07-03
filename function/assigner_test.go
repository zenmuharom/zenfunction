package function

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenfunction/variable"
	"github.com/zenmuharom/zenlogger"
)

func TestReadCommand(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	billerResp := "FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:"
	input := "substr($field, 0, 100)"
	testCases := []TestCase{
		{
			Input:    "",
			Expected: "",
		},
		{
			Input:    "test",
			Expected: "test",
		},
		{
			Input:    "dateNow",
			Expected: "dateNow",
		},
		{
			Input:    "trim(dateNow(), 2025)",
			Expected: strings.Trim(fmt.Sprintf("%v", time.Now().Format(time.RFC3339)), "2025"),
		},
		{
			Input:    "dateAdd(\"2006/01/02\", dateNow(), 30, day)",
			Expected: time.Now().AddDate(0, 0, 30).Format("2006/01/02"),
		},
		{
			Input:    "substr(dateAdd(2006, dateNow(2006), 1, year), 0, 2)",
			Expected: "20",
		},
		{
			Input:    "trim(substr(\"1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         \", 89, 30))",
			Expected: "IDHAM DHIYAULHAQ HABIBI",
		},
		{
			Input:    strings.ReplaceAll(input, "$field", strconv.Quote(billerResp)),
			Expected: "FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:",
		},
		{
			Input:    strings.ReplaceAll("rps($field, 200)", "$field", strconv.Quote(billerResp)),
			Expected: "FINNET - MUAMALAT\r\nHindari dend4 \"Tag+\", bayar sb3lum tgl 20 tiap bulannya ya:                                                                                                                          ",
		},
		{
			Input:    strings.ReplaceAll("substr($field, 0, 8)", "$field", strconv.Quote(billerResp)),
			Expected: "FINNET -",
		},
		{
			Input:    strings.ReplaceAll("substr(rps($field, 10), 0, 50)", "$field", strconv.Quote(billerResp)),
			Expected: billerResp[:50],
		},
	}

	for noTest, tc := range testCases {
		var err error
		var result interface{}
		res, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}

		require.NoError(t, err, errMsg)

		switch v := res.(type) {
		case string:

			unquote, errUnquote := strconv.Unquote(fmt.Sprintf("%v", res))
			if errUnquote != nil {
				result = fmt.Sprintf("%v", res)
			} else {
				result = fmt.Sprintf("%v", unquote)
			}
			require.Equal(t, tc.Expected, result)
		default:
			// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
			result = fmt.Sprintf("%v", v)
		}

		res2, err2 := assigner.ReadCommandV2(variable.TYPE_STRING, tc.Input)
		if err2 != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.Equal(t, tc.Expected, res2)
	}
}
