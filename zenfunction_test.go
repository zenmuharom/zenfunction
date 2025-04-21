package zenfunction

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestNew(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	zf := New(logger)
	res, err := zf.ReadCommand("dateNow()")
	var result any

	switch v := res.(type) {
	case string:

		if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) && len(v) >= 2 {
			// safe unwrap only outer quotes
			v = v[1 : len(v)-1]
		}
		result = v

		require.Equal(t, time.Now().Format(time.RFC3339), result)
	default:
		// for numbers, arrays, objects: convert to string (optional, sesuai kebutuhan)
		result = fmt.Sprintf("%v", v)
	}

	require.NoError(t, err)
}
