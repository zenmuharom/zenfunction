package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestReadCommand(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		// {
		// 	Input:    "",
		// 	Expected: "",
		// },
		// {
		// 	Input:    "test",
		// 	Expected: "test",
		// },
		// {
		// 	Input:    "dateNow",
		// 	Expected: "dateNow",
		// },
		// // {
		// // 	Input:    "trim(dateNow(), 2023)",
		// // 	Expected: strings.Trim(fmt.Sprintf("%v", time.Now().Format(time.RFC3339)), "2023"),
		// // },
		// // {
		// // 	Input:    "dateAdd(2006/01/02, dateNow(), 30, day)",
		// // 	Expected: "2023/06/22",
		// // },
		// {
		// 	Input:    "substr(dateAdd(2006, dateNow(2006), 1, year), 0, 2)",
		// 	Expected: "20",
		// },
		{
			Input:    "trim(substr(1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         , 89, 30))",
			Expected: "IDHAM DHIYAULHAQ HABIBI",
		},
	}

	for noTest, tc := range testCases {
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.NoError(t, err, errMsg)
		require.Equal(t, tc.Expected, result)
	}
}
