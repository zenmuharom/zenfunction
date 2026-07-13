package function

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenfunction/variable"
	"github.com/zenmuharom/zenlogger"
)

func TestToStringViaReadCommandAndReadCommandV2(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "integer literal",
			input:    `toString(1000000)`,
			expected: "1000000",
		},
		{
			name:     "boolean literal",
			input:    `toString(true)`,
			expected: "true",
		},
		{
			name:     "string literal",
			input:    `toString("nominal")`,
			expected: "nominal",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := assigner.ReadCommand(tc.input)
			require.NoError(t, err)
			require.Equal(t, `"`+tc.expected+`"`, res)

			resV2, err := assigner.ReadCommandV2(variable.TYPE_STRING, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, resV2)
		})
	}
}