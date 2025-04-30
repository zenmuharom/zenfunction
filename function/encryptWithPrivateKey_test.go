package function

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_EncryptWithPrivateKey(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	privKey := `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCPMAI+Z/30MZnL
8EZEXmCB1bxz/bIbAOFrIkaLPjMESGPGJc4cI4B/90MQoL5WRbE7uc2fFIaWJzco
7CO4H8zPo9iOQ5rfXEyO/49S5DHBw0Wj8p+PvhDQ/XyFCW40gzL9B1UMu4driANV
7wSBz4XB3NhVUcqQI1oXQhk9eLJf5tqryZ3OuvqYheY+st4IEn86foVL3t1KxdKV
EpYNn0TlDIBlgDCcHKqO6v+Sec8Xa3FMfxY/RpUr9oZ+wWKN+uzXHW8DMglN+w4k
FgThIaookUjC9qfsm5LqU6/X4uByL4Gv+0PLFLIdvawbzjaYhQ6AV2s2kjW7qUui
J7XCcJEpAgMBAAECggEAGQKCsdHKKrfrLBLowxJU+viRNRVP4aWSufq/pZyUnp/i
RX7e9E1DKZbPsPzSohEENvFqN+oG4/0vhkdQqd1sEayFq7ulNnRRrvx2VT0kb7qi
1FMRibLiDw7ni/kstaFg6483TMUIj0CvjHf2KpJLw6PGaYiiv/Bc0QkrjXAif9t4
11Lr48KYsog6xZbF6M9Xk40twPMiHxz7s+v4NPOssaDaiGlW6u7/CWoKpwB6J/zr
5YVKuxJ4XnLwlt3NcnSzkV0juKS5ZAGXlXLwFErLhjtw/9OS4HZHYlYpkBpZpwpj
7O0alEB+0Srm4TbzD3izui40wz6m1ZXuzkZrUR99HQKBgQDFCFE1zHVqZhLm1gLf
QBu1Bia/8utu7xy0zBhqy9geE6imEWPzO5aSUxlNuDVrY5jWIH3BItARPtrmcnpD
3J/P7tRIZcbbUsegtbl3hyWqhfyi6MGS60jliALBlVV80zXVjhMtsqMenIGMPTQy
1+hExh7zFfxdlHcMmhU6nLWlkwKBgQC6ClmKDcz3eIoXxG89kOX8ZDrCDp2ees+t
BzkwC2nyWGlInx8vNu3GSHP3SqpX1tIMtCcTwW/uM5HcyvGpB8RP9niVNZ4fuBZL
4l+9Brnq0EsWcuU/2opYNVo2x8D8+/DzgtTMG4zytzSQfeGj5CG2/TO6t44VETID
ie8+nlIj0wKBgAv+46LY3dUqfcAcC3S4HHe69iT9jyPj3uWK/3mRC4lZPQ1PRbyL
RjGGaaX3rxjoqWdv9vgJPI2wO/eHxLXY+snYCoiV2bOEqK66IZ6LVdm56pWoghCF
zpxa2YAbrWa6HS7xRW2k0JWOhbyaBVGLH5MAVOYL0p+H6G+V+fDllZGNAoGBALGv
pOloWPWbmTkuEpkYxbCUAlLKJtzwq121Yndyz1P6AUStRdmQevVAyhHMrHmM4b3k
atZBkKhPdOcOplUs5+D/pRfNyCK/bfw4T/x4aiXNn4nnXvHnxu6MtodPrhFyiCXs
NVZkkfBX7sp6kII6J8FggIG7Qub4L26V1X1XNVilAoGAHNzBTVeY+l1eE65Wwlxm
AIBPZygptzB2Jqs6x5G6kj1+qwdyKhDWXDQr2W5LPl6/iwuwlJS9zyBGYg00/luC
JzanZ/glRiJVqo23PnVVsnlpKmjYTqXAH+T4xh2imUTKZIE3O1tvji4gL4mFPAZL
omfZxLVtIe8vpxCJsXIWWr0=
-----END PRIVATE KEY-----	
	`

	testCases := []TestCase{
		{
			Input:    "encryptWithPrivateKey(" + strconv.Quote("9a698a48-30f9-4801-b39e-6fa43f3f2d94|2025-02-25T13:52:44+07:00") + ", " + privKey + ")",
			Expected: "Wi6dGrkoi2f1GvyPtQgt2+wXakSMfb8mDt1mhKKHiTsK4i/67lAtzbzV/iIOzC6/VFrXdBdGnC7ZX6iicOIq7i/AIxg7hArvTLuq30UxmhFLMkTwxkrTmc8FuRKeDumQAE51RGvm/4LlPAZsX+fY4PvbKYwXSTfkyXuk18pRQb50ZPrXc2ntjr24B9vEXK8uaxEz03kLMV/v9oprCcSSxI9Kl0Yf57mTp643xg0s6jEHE0dPaioV2iY0dXnREXwAbFrXPkQjp1RuhwTbhz8aNhtipq0HuzLVtv6DiAEz5R6iFK356KOjHy/luteC7H3xv7jWytWeUG53EpGY6LWkag==",
		},
	}

	for noTest, tc := range testCases {
		var result any
		res, err := assigner.ReadCommand(tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}

		switch v := res.(type) {
		case string:

			if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
				// safe unwrap only outer quotes
				v = v[1 : len(v)-1]
			}
			result = v

			require.NoError(t, err, errMsg)
			require.Equal(t, tc.Expected, result)
		default:
			// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
			result = fmt.Sprintf("%v", v)
		}
	}
}
