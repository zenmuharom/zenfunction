package function

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestSubstr(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	middleware_response_value := "512233350072 0700010220230614150549721391424021249469                                Plg.,De'mo512233350072   R3  00002400051223          123            202107  2007202105072021000000358200000000000000000000035820000000000000000000284936000000284981000000358200202108  2008202105082021000000310440000000000000000000031044000000000000000000284981000000285020000000310440                        000000000000000000000000000000000000000000000000000000000000000000000000000000000000                        000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	middleware_response_value = strings.ReplaceAll(middleware_response_value, ",", `\,`)
	fmt.Println(middleware_response_value)
	specialCase := "trim(substr($middleware_response_id, 0, 13))"
	// specialCase := "trim(substr($middleware_response_id, 85, 20), P)"
	// specialCase := "$middleware_response_id"
	specialCase = strings.ReplaceAll(specialCase, "$middleware_response_id", middleware_response_value)

	testCases := []TestCase{
		{
			Input:    "",
			Expected: "",
		},
		{
			Input:    "substr(\"test woi\", 0, 4)",
			Expected: "test",
		},
		{
			Input:    "substr(\"test woi\", 4, 4)",
			Expected: " woi",
		},
		{
			Input:    "substr(tets lah, 5)",
			Expected: "lah",
		},
		{
			Input:    "substr(test lah)",
			Expected: "invalid parameter",
		},
		{
			Input:    "substr()",
			Expected: "invalid parameter",
		},
		{
			Input:    specialCase,
			Expected: "512233350072",
		},
		{
			Input:    "substr(" + strconv.Quote("FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepat waktu:|Download PLN Mobile") + ", 0, 70)",
			Expected: "FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepa",
		},
		{
			Input:    "substr(\"halo\", 6, 8)",
			Expected: "halo",
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
