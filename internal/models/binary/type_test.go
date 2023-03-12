package binary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCost(t *testing.T) {
	m := Model{1, 2, 3}
	require.Equal(t, float64(3), m.Cost())
}

func TestMutation(t *testing.T) {
	A := Model{1, 2, 3}
	B := Model{4, 5, 6}
	C := A.Mutation(B)
	D := B.Mutation(A)
	require.NotEqual(t, C, A)
	require.NotEqual(t, C, D)
	require.NotNil(t, C)
}
