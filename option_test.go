package goga

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionWithFitnessFunc(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := &ga{}
		err := OptionWithFitnessFunc(func(rank int64, cost float64) float64 { return 1 })(g)
		require.NoError(t, err)
	})
}
