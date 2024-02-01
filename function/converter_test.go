package function

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isArgumentOrganic(t *testing.T) {
	arg := "ltrim(15700.00,0)"
	result := isArgumentOrganic(arg)

	require.Equal(t, false, result, "OK")
}
