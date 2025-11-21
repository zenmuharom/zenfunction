package function

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestReplace(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "replace(\"hello world hello universe\",\"hello\",\"hi\",1)",
			Expected: "hi world hello universe",
		},
		{
			Input:    "replace(\"foo bar foo baz foo\",\"foo\",\"test\",2)",
			Expected: "test bar test baz foo",
		},
		{
			Input:    "replace(\"hello world hello universe\",\"hello\",\"hi\",-1)",
			Expected: "hi world hi universe",
		},
		{
			Input:    "replace(\"hello world\",\"hello\",\"hi\",0)",
			Expected: "hello world",
		},
		{
			Input:    "replace(\"hello world hello\",\"hello\",\"\",1)",
			Expected: " world hello",
		},
		{
			Input:    "replace(\"hello world\",\"goodbye\",\"farewell\",1)",
			Expected: "hello world",
		},
		{
			Input:    "replace(\"aaa bbb aaa ccc aaa\",\"aaa\",\"xxx\",2)",
			Expected: "xxx bbb xxx ccc aaa",
		},
		{
			Input:    "replace(\"a-b-c-d-e\",\"-\",\"_\",3)",
			Expected: "a_b_c_d-e",
		},
		{
			Input:    "replace(\"\",\"hello\",\"hi\",1)",
			Expected: "",
		},
		{
			Input:    "replace(\"hello world\",\"hello\",\"hi\",5)",
			Expected: "hi world",
		},
	}

	for noTest, tc := range testCases {
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}

		require.NoError(t, err, errMsg)

		// Unquote the result if it's a string
		switch v := result.(type) {
		case string:
			if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
				// safe unwrap only outer quotes
				result = v[1 : len(v)-1]
			}
		default:
			result = fmt.Sprintf("%v", v)
		}

		require.Equal(t, tc.Expected, result, fmt.Sprintf("Test #%d failed", noTest))
	}
}
