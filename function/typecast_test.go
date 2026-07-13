package function

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenmuharom/zenfunction/variable"
	"github.com/zenmuharom/zenlogger"
)

func TestTypeCastFunctions(t *testing.T) {
	logger := zenlogger.NewZenlogger()
	assigner := NewAssigner(logger)

	t.Run("toString converts integer string safely", func(t *testing.T) {
		res, err := assigner.ReadCommand(`toString(1000000)`)
		require.NoError(t, err)
		require.Equal(t, `"1000000"`, res)

		resV2, err := assigner.ReadCommandV2(variable.TYPE_STRING, `toString(1000000)`)
		require.NoError(t, err)
		require.Equal(t, "1000000", resV2)
	})

	t.Run("toInt converts string to integer", func(t *testing.T) {
		res, err := assigner.ReadCommand(`toInt("1000000")`)
		require.NoError(t, err)
		require.Equal(t, "1000000", res)
	})

	t.Run("toFloat converts string to float", func(t *testing.T) {
		res, err := assigner.ReadCommand(`toFloat("1000.25")`)
		require.NoError(t, err)
		require.Equal(t, "1000.25", res)
	})

	t.Run("toBool converts numeric string", func(t *testing.T) {
		res, err := assigner.ReadCommand(`toBool("1")`)
		require.NoError(t, err)
		require.Equal(t, "true", res)
	})

	t.Run("toInt rejects non integer", func(t *testing.T) {
		res, err := assigner.ReadCommand(`toInt("10.5")`)
		require.NoError(t, err)
		require.Equal(t, `"invalid parameter"`, res)
	})
}
