package goga

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		GA, err := New(func(g *ga) error { return nil })
		require.NoError(t, err)
		require.NotNil(t, GA)
	})

	t.Run("option failure", func(t *testing.T) {
		GA, err := New(func(g *ga) error { return errors.New("any") })
		require.Error(t, err)
		require.Nil(t, GA)
	})
}

func TestRuntimeBestResult(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	require.NotNil(t, GA.RuntimeBestResult())
}

func TestResult(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	go func() {
		ga := GA.(*ga)
		ga.result <- nil
	}()
	Resp, err := GA.Result()
	require.NoError(t, err)
	require.NotNil(t, Resp)
}
