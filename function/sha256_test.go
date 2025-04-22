package function

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func Test_Sha256(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "sha256(\"Darmaj4y4\")",
			Expected: "1532a22b0d3525e94c7455945a30943a205ed15cde722532f4f4cda08ded0f88",
		},
		{
			Input:    "sha256(\"bangsat\")",
			Expected: "4077902c11d06e1058060d1bd789e7c5a1db2bbecbc39b8d161c1d131c2ac1b7",
		},
		{
			Input:    "sha256(\"71947bea4c63-63ad-3489-a73a-2d73668356a460236703\")",
			Expected: "16afeefe52a2fc5fae12023b56b990fac59ffd0674c26aafd2e2e061f0629504",
		},
		{
			Input:    "sha256(\"9a698a48-30f9-4801-b39e-6fa43f3f2d94|2025-02-25T13:52:44+07:00\", \"-----BEGIN PRIVATE KEY----- MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCPMAI+Z/30MZnL 8EZEXmCB1bxz/bIbAOFrIkaLPjMESGPGJc4cI4B/90MQoL5WRbE7uc2fFIaWJzco 7CO4H8zPo9iOQ5rfXEyO/49S5DHBw0Wj8p+PvhDQ/XyFCW40gzL9B1UMu4driANV 7wSBz4XB3NhVUcqQI1oXQhk9eLJf5tqryZ3OuvqYheY+st4IEn86foVL3t1KxdKV EpYNn0TlDIBlgDCcHKqO6v+Sec8Xa3FMfxY/RpUr9oZ+wWKN+uzXHW8DMglN+w4k FgThIaookUjC9qfsm5LqU6/X4uByL4Gv+0PLFLIdvawbzjaYhQ6AV2s2kjW7qUui J7XCcJEpAgMBAAECggEAGQKCsdHKKrfrLBLowxJU+viRNRVP4aWSufq/pZyUnp/i RX7e9E1DKZbPsPzSohEENvFqN+oG4/0vhkdQqd1sEayFq7ulNnRRrvx2VT0kb7qi 1FMRibLiDw7ni/kstaFg6483TMUIj0CvjHf2KpJLw6PGaYiiv/Bc0QkrjXAif9t4 11Lr48KYsog6xZbF6M9Xk40twPMiHxz7s+v4NPOssaDaiGlW6u7/CWoKpwB6J/zr 5YVKuxJ4XnLwlt3NcnSzkV0juKS5ZAGXlXLwFErLhjtw/9OS4HZHYlYpkBpZpwpj 7O0alEB+0Srm4TbzD3izui40wz6m1ZXuzkZrUR99HQKBgQDFCFE1zHVqZhLm1gLf QBu1Bia/8utu7xy0zBhqy9geE6imEWPzO5aSUxlNuDVrY5jWIH3BItARPtrmcnpD 3J/P7tRIZcbbUsegtbl3hyWqhfyi6MGS60jliALBlVV80zXVjhMtsqMenIGMPTQy 1+hExh7zFfxdlHcMmhU6nLWlkwKBgQC6ClmKDcz3eIoXxG89kOX8ZDrCDp2ees+t BzkwC2nyWGlInx8vNu3GSHP3SqpX1tIMtCcTwW/uM5HcyvGpB8RP9niVNZ4fuBZL 4l+9Brnq0EsWcuU/2opYNVo2x8D8+/DzgtTMG4zytzSQfeGj5CG2/TO6t44VETID ie8+nlIj0wKBgAv+46LY3dUqfcAcC3S4HHe69iT9jyPj3uWK/3mRC4lZPQ1PRbyL RjGGaaX3rxjoqWdv9vgJPI2wO/eHxLXY+snYCoiV2bOEqK66IZ6LVdm56pWoghCF zpxa2YAbrWa6HS7xRW2k0JWOhbyaBVGLH5MAVOYL0p+H6G+V+fDllZGNAoGBALGv pOloWPWbmTkuEpkYxbCUAlLKJtzwq121Yndyz1P6AUStRdmQevVAyhHMrHmM4b3k atZBkKhPdOcOplUs5+D/pRfNyCK/bfw4T/x4aiXNn4nnXvHnxu6MtodPrhFyiCXs NVZkkfBX7sp6kII6J8FggIG7Qub4L26V1X1XNVilAoGAHNzBTVeY+l1eE65Wwlxm AIBPZygptzB2Jqs6x5G6kj1+qwdyKhDWXDQr2W5LPl6/iwuwlJS9zyBGYg00/luC JzanZ/glRiJVqo23PnVVsnlpKmjYTqXAH+T4xh2imUTKZIE3O1tvji4gL4mFPAZL omfZxLVtIe8vpxCJsXIWWr0= -----END PRIVATE KEY-----\")",
			Expected: "a607776069236c9f94d3c67543d39f42ad228b7358f5f3f3a2070f57e9ae6ca1",
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
