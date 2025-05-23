package function

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitWithEscapedCommas(t *testing.T) {
	testCases := []TestCase{
		{
			Input:    "1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         , 89, 30",
			Expected: "1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         ",
		},
	}

	for _, tc := range testCases {
		stringArr := splitWithEscapedCommas(tc.Input)

		require.Equal(t, 3, len(stringArr))
		require.Equal(t, tc.Expected, stringArr[0])
	}
}

func TestSplitWithEscapedCommas2(t *testing.T) {
	testCases := []TestCase{
		{
			Input:    "\"1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         \", 89, 30",
			Expected: "\"1267345625003090001303GAYCGKDPS 7502208061803GAYCGKDPS 7502208061803GAYCGKDPS 75022080618IDHAM DHIYAULHAQ HABIBI       ABC123         \"",
		},
	}

	for _, tc := range testCases {
		stringArr := splitWithEscapedCommas(tc.Input)

		require.Equal(t, 3, len(stringArr))
		require.Equal(t, tc.Expected, stringArr[0])
	}
}
