package function

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestJsonDecode(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "json_decode({\"NIK\":\"7302046706800003\",\"NAMA_LENGKAP\":\"HASANAINY ANAS\",\"MASA_AKTIF\":{\"TGL_GRACE\":\"21-09-2024 00:00:00\",\"TGL_EXPIRED\":\"21-06-2024 00:00:00\",\"TGL_EFEKTIF\":\"22-06-2023 06:45:51\",\"TGL_AKTIF\":\"22-06-2023 06:45:51\"},\"DATA_IURAN\":{\"BLTH\":\"06-2023\",\"PROG_JKM\":\"81600\",\"PROG_JKK\":\"120000\",\"UMP\":\"0\",\"KODE_IURAN\":\"923065279115\",\"BIAYA_TRANSAKSI\":\"0\",\"BIAYA_REGISTRASI\":\"0\",\"TOTAL\":\"201600\",\"BLN_PROGRAM\":\"12\",\"UPAH\":\"1000000\",\"DASAR_UPAH\":\"1000000\"},\"HITUNG_IURAN\":\"\",\"PROGRAM\":\"JKK,JKM\",\"KODE_KANTOR\":\"W15\",\"NAMA_KANTOR\":\"BULUKUMBA SAM RATULANGI(KCP)\",\"REQID\":\"IDM516127770935064919064915\"})",
			Expected: "map[DATA_IURAN:map[BIAYA_REGISTRASI:0 BIAYA_TRANSAKSI:0 BLN_PROGRAM:12 BLTH:06-2023 DASAR_UPAH:1000000 KODE_IURAN:923065279115 PROG_JKK:120000 PROG_JKM:81600 TOTAL:201600 UMP:0 UPAH:1000000] HITUNG_IURAN: KODE_KANTOR:W15 MASA_AKTIF:map[TGL_AKTIF:22-06-2023 06:45:51 TGL_EFEKTIF:22-06-2023 06:45:51 TGL_EXPIRED:21-06-2024 00:00:00 TGL_GRACE:21-09-2024 00:00:00] NAMA_KANTOR:BULUKUMBA SAM RATULANGI(KCP) NAMA_LENGKAP:HASANAINY ANAS NIK:7302046706800003 PROGRAM:JKK,JKM REQID:IDM516127770935064919064915]",
		},
	}

	for noTest, tc := range testCases {
		result, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		require.NoError(t, err, errMsg)
		resValueOf := reflect.ValueOf(result)
		require.Equal(t, "map", resValueOf.Kind().String())
	}
}

func TestJsonUnpack(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	t.Run("json_unpack via ReadCommand with no keys", func(t *testing.T) {
		input := `json_unpack({"a": {"b": {"c": 123}}})`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)

		resultMap, ok := result.(map[string]interface{})
		require.True(t, ok)
		require.Contains(t, resultMap, "a")
	})

	t.Run("json_unpack via ReadCommand with single key", func(t *testing.T) {
		input := `json_unpack({"a": {"b": {"c": 123}}}, "a")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)

		resultMap, ok := result.(map[string]interface{})
		require.True(t, ok)
		require.Contains(t, resultMap, "b")
	})

	t.Run("json_unpack via ReadCommand with multiple keys", func(t *testing.T) {
		input := `json_unpack({"a": {"b": {"c": 123}}}, "a", "b", "c")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, float64(123), result)
	})

	t.Run("json_unpack via ReadCommand extracting string value", func(t *testing.T) {
		input := `json_unpack({"user": {"name": "John Doe"}}, "user", "name")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "John Doe", result)
	})

	t.Run("json_unpack via ReadCommand with nested object extraction", func(t *testing.T) {
		input := `json_unpack({"a": {"b": {"c": 123}}}, "a", "b")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)

		resultMap, ok := result.(map[string]interface{})
		require.True(t, ok)
		require.Equal(t, float64(123), resultMap["c"])
	})

	t.Run("json_unpack via ReadCommand with non-existent key returns original string", func(t *testing.T) {
		input := `json_unpack({"a": {"b": {"c": 123}}}, "a", "x")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		// When key is not found, json_unpack returns nil, but coreReadCommand returns the original string
		// This is the expected behavior of the system
		require.Equal(t, input, result)
	})

	t.Run("json_unpack via ReadCommand with real world example", func(t *testing.T) {
		input := `json_unpack({"NIK":"7302046706800003","MASA_AKTIF":{"TGL_GRACE":"21-09-2024 00:00:00"}}, "MASA_AKTIF", "TGL_GRACE")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "21-09-2024 00:00:00", result)
	})

	t.Run("json_unpack via ReadCommand with array extraction", func(t *testing.T) {
		input := `json_unpack({"data": {"items": [1, 2, 3]}}, "data", "items")`
		result, err := assigner.ReadCommand(input)

		require.NoError(t, err)
		require.NotNil(t, result)

		resultArr, ok := result.([]interface{})
		require.True(t, ok)
		require.Len(t, resultArr, 3)
		require.Equal(t, float64(1), resultArr[0])
	})

	t.Run("json_unpack via ReadCommandV2 with string type", func(t *testing.T) {
		input := `json_unpack({"user": {"name": "Alice"}}, "user", "name")`
		result, err := assigner.ReadCommandV2("string", input)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "Alice", result)
	})

	t.Run("json_unpack via ReadCommandV2 with default type", func(t *testing.T) {
		input := `json_unpack({"price": 99.99}, "price")`
		result, err := assigner.ReadCommandV2("", input)

		require.NoError(t, err)
		require.NotNil(t, result)
		// ReadCommandV2 converts to string for default type
		require.Equal(t, "99.99", result)
	})
}
