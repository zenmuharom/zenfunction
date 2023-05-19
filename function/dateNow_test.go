package function

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestDateNow(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "dateNow()",
			Expected: time.Now().Format(time.RFC3339),
		},
		{
			Input:    "dateNow(2006)",
			Expected: time.Now().Format("2006"),
		},
		{
			Input:    "dateNow(200601)",
			Expected: time.Now().Format("200601"),
		},
		{
			Input:    "dateNow(2006-01)",
			Expected: time.Now().Format("2006-01"),
		},
		{
			Input:    "dateNow(2006/01)",
			Expected: time.Now().Format("2006/01"),
		},
		{
			Input:    "dateNow(20060102)",
			Expected: time.Now().Format("20060102"),
		},
		{
			Input:    "dateNow(2006-01-02)",
			Expected: time.Now().Format("2006-01-02"),
		},
		{
			Input:    "dateNow(2006/01/02)",
			Expected: time.Now().Format("2006/01/02"),
		},
		{
			Input:    "dateNow(20060102150405)",
			Expected: time.Now().Format("20060102150405"),
		},
		{
			Input:    "dateNow(2006-01-02 15:04:05)",
			Expected: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			Input:    "dateNow(2006/01/02 15:04:05)",
			Expected: time.Now().Format("2006/01/02 15:04:05"),
		},
		{
			Input:    "dateNow(15:04:05)",
			Expected: time.Now().Format("15:04:05"),
		},
		{
			Input:    "dateNow(15/04/05)",
			Expected: time.Now().Format("15/04/05"),
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
