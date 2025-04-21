package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestArrayLength(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "lengthArray([{\"alamat\": \"JL.TEBET TIMUR DALAM III NO.15\", \"amount\": \"329400\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"329400\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030804300010\", \"nama\": \"PT. JCO DONUTS   COFFE\", \"no_skpd\": \"\", \"nop\": \"311508030804300010\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL. TEBET TIMUR DALAM III/15\", \"amount\": \"292800\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"292800\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030805260090\", \"nama\": \"JCO DONUTS   COFFEE\", \"no_skpd\": \"\", \"nop\": \"311508030805260090\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL.TEBET TIMUR DALAM III NO.15\", \"amount\": \"512400\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"512400\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030804330016\", \"nama\": \"PT. JCO DONUTS   COFFEE\", \"no_skpd\": \"\", \"nop\": \"311508030804330016\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL. TAMAN RATU INDAH BLOK A-1/15 JAKARTA 1151\", \"amount\": \"118950\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"118950\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311507310804270022\", \"nama\": \"TALKINDO SELAKSA ANUGRAH, PT\", \"no_skpd\": \"\", \"nop\": \"311507310804270022\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL. TEBET TIMUR DALAM III NO. 15\", \"amount\": \"291657\", \"amount_bunga\": \"0\", \"amount_denda\": \"5719\", \"amount_pokok\": \"285938\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030805170070\", \"nama\": \"PT. J.CO DONUTS COFFEE\", \"no_skpd\": \"\", \"nop\": \"311508030805170070\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JLN. CASABLANCA KAV. 88 MALL KOTA KASABLANKA\", \"amount\": \"714488\", \"amount_bunga\": \"0\", \"amount_denda\": \"98550\", \"amount_pokok\": \"615938\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311509020810490002\", \"nama\": \"ROYAL PANCA PERSADA ANUGERAH,\", \"no_skpd\": \"\", \"nop\": \"311509020810490002\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL. TAMAN RATU INDAH BLOK A-1/15 JAKARTA 1151\", \"amount\": \"1363350\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"1363350\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311507310804250018\", \"nama\": \"TALKINDO SELAKSA ANUGRAH, PT\", \"no_skpd\": \"\", \"nop\": \"311507310804250018\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JLN. SULTAN ISKANDAR MUDA /\", \"amount\": \"446063\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"446063\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508310811230002\", \"nama\": \"JOHNNY ANDREAN\", \"no_skpd\": \"\", \"nop\": \"311508310811230002\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL.TEBET TIMUR DALAM III NO.15\", \"amount\": \"146400\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"146400\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030804320014\", \"nama\": \"PT. JCO DONUTS   COFFE\", \"no_skpd\": \"\", \"nop\": \"311508030804320014\", \"tahun_pajak\": \"\" }, { \"alamat\": \"JL. TEBET TIMUR DALAM III/15\", \"amount\": \"576450\", \"amount_bunga\": \"0\", \"amount_denda\": \"0\", \"amount_pokok\": \"576450\", \"amount_sanksi\": \"0\", \"bulan_pajak\": \"\", \"jns_pajak\": \"REKLAME\", \"kd_rek_bunga\": \"4149707001\", \"kd_rek_denda\": \"4149707002\", \"kd_rek_pokok\": \"4110801\", \"kd_rek_sanksi\": \"0\", \"kode_billing\": \"311508030805220082\", \"nama\": \"JCO DONUTS   COFFEE\", \"no_skpd\": \"\", \"nop\": \"311508030805220082\", \"tahun_pajak\": \"\"}])",
			Expected: "10",
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
