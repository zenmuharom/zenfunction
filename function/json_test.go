package function

import (
	"fmt"
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
		require.Equal(t, tc.Expected, result)
	}
}
