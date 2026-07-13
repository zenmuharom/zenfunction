package function

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenfunction/variable"
	"github.com/zenmuharom/zenlogger"
)

func TestToBoolViaReadCommandAndReadCommandV2(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []struct {
		name            string
		input           string
		expectedRead     string
		expectedReadV2   string
	}{
		{
			name:          "string true",
			input:         `toBool("true")`,
			expectedRead:   "true",
			expectedReadV2: "true",
		},
		{
			name:          "numeric one",
			input:         `toBool("1")`,
			expectedRead:   "true",
			expectedReadV2: "true",
		},
		{
			name:          "invalid bool",
			input:         `toBool("hello")`,
			expectedRead:   `"invalid parameter"`,
			expectedReadV2: "invalid parameter",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := assigner.ReadCommand(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expectedRead, res)

			resV2, err := assigner.ReadCommandV2(variable.TYPE_STRING, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expectedReadV2, resV2)
		})
	}
}