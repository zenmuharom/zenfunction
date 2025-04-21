package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestConcat(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	TestCase := []TestCase{
		{
			Input:    "concat(\"hallo \", \"anjing \", \"woi 1 \")",
			Expected: "hallo anjing woi 1 ",
		},
		{
			Input:    "concat(\"hallo\", \"anjing\", \"woi2\")",
			Expected: "halloanjingwoi2",
		},
		{
			Input:    "concat(\"hallo, anjing\", woi3)",
			Expected: "hallo, anjingwoi3",
		},
	}

	for noTest, tc := range TestCase {
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.NoError(t, err, errMsg)
		require.Equal(t, tc.Expected, result)
	}
}
