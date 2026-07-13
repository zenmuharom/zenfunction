package function

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isArgumentOrganic(t *testing.T) {
	arg := "ltrim(15700.00,0)"
	result := isArgumentOrganic(arg)

	require.Equal(t, false, result, "OK")
}

func Test_convertToString(t *testing.T) {
	require.Equal(t, "1000000", convertToString(float64(1000000)))
	require.Equal(t, "true", convertToString(true))
	require.Equal(t, "12345", convertToString(json.Number("12345")))
	require.Equal(t, "", convertToString(nil))
}

func Test_convertToInt64(t *testing.T) {
	v, err := convertToInt64("1000000")
	require.NoError(t, err)
	require.EqualValues(t, 1000000, v)

	v, err = convertToInt64(json.Number("42"))
	require.NoError(t, err)
	require.EqualValues(t, 42, v)

	_, err = convertToInt64("10.5")
	require.Error(t, err)
}

func Test_convertToFloat64(t *testing.T) {
	v, err := convertToFloat64("1000.25")
	require.NoError(t, err)
	require.EqualValues(t, 1000.25, v)

	v, err = convertToFloat64(json.Number("123"))
	require.NoError(t, err)
	require.EqualValues(t, 123, v)
}

func Test_convertToBool(t *testing.T) {
	v, err := convertToBool("true")
	require.NoError(t, err)
	require.True(t, v)

	v, err = convertToBool("0")
	require.NoError(t, err)
	require.False(t, v)

	_, err = convertToBool("not-bool")
	require.Error(t, err)
}
