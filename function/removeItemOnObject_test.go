package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestRemoveItemOnObject(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	testCases := []TestCase{
		{
			Input:    "removeItemOnObject({\"name\":\"ZeniMuharom\",\"gender\":\"male\",\"age\":30,\"address\":{\"street\":\"Jl.BasukiRachmatNo.1A\",\"city\":\"EastJakarta\",\"nation\":\"Indonesia\"}}, \"name\", \"gender\")",
			Expected: "{\"age\":30,\"address\":{\"street\":\"Jl.BasukiRachmatNo.1A\",\"city\":\"EastJakarta\",\"nation\":\"Indonesia\"}",
		},
	}

	for noTest, tc := range testCases {
		result, err := assigner.ReadCommandV2("object", tc.Input)
		errMsg := ""
		if err != nil {
			errMsg = fmt.Sprintf("No Test.%v: %v", noTest, err.Error())
		}
		// fmt.Println(fmt.Sprintf("%#v", result))
		require.NoError(t, err, errMsg)

		if result != "invalid parameter" {
			// fmt.Println("VALID")
			var res any
			e := json.Unmarshal([]byte(fmt.Sprintf("%v", result)), &res)
			if e != nil {
				// fmt.Println("e != nil")
				require.Error(t, e, e.Error())
			} else {
				// fmt.Println("else")
				switch v := res.(type) {
				case map[string]interface{}:
					// fmt.Println("map string")
					require.Equal(t, "tai", v)
				default:
					// fmt.Println("default")
					err = errors.New("not object")
				}
			}
		} else {
			// fmt.Println("invalid parameter")
		}
		// require.Equal(t, "asdf", "msbdf")
	}
}
