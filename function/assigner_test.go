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

	billerResp := "FINNET - MUAMALAT\r\nSlamat thn baru 2025 - Byr Sbelum tgl 20 \"Tag\" Tepat waktu:"
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
			Input:    "substr(" + strconv.Quote(billerResp) + ", 0, 18)",
			Expected: "FINNET - MUAMALAT\r",
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

		res2, err2 := assigner.ReadCommandV2(variable.TYPE_STRING, tc.Input)
		if err2 != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.Equal(t, tc.Expected, res2)
		fmt.Println(errMsg)
	}
}
