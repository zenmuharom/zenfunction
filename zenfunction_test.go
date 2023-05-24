package zenfunction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenlogger"
)

func TestNew(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	zf := New(logger)
	res, err := zf.ReadCommand("dateNow()")

	require.Equal(t, time.Now().Format(time.RFC3339), res)
	require.NoError(t, err)
}
