package goga

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOptionWithFitnessFunc(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := &ga{}
		err := OptionWithWeightFunc(func(rank int, cost float64) int { return 1 })(g)
		require.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		// nil func
		g := &ga{}
		err := OptionWithWeightFunc(nil)(g)
		require.Error(t, err)
	})
}

func TestOptionWithPopulationFunc(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := &ga{}
		err := OptionWithPopulationFunc(func(currentPopulation int, step int64, bestCost, worstCost float64) int { return 1 })(g)
		require.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		// nil func
		g := &ga{}
		err := OptionWithPopulationFunc(nil)(g)
		require.Error(t, err)
	})
}

func TestOptionWithGeneratorFunc(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := &ga{}
		err := OptionWithGeneratorFunc(func() Model { return nil })(g)
		require.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		// nil func
		g := &ga{}
		err := OptionWithGeneratorFunc(nil)(g)
		require.Error(t, err)
	})
}

func TestOptionWithSelection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := &ga{}
		err := OptionWithSelection(0.1, 0.2, 0.3)(g)
		require.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		// nil func
		g := &ga{}
		err := OptionWithSelection(0, 0, 0)(g)
		require.Error(t, err)
	})
}

func TestOptionWithStepInterval(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stepsInterval := 1 * time.Second
		g := &ga{}
		err := OptionWithStepInterval(stepsInterval)(g)
		require.NoError(t, err)
		require.Equal(t, stepsInterval, g.config.stepsInterval)
	})
}
func TestOptionWithMaximumNumberOfSteps(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		maxStep := int64(1000)
		g := &ga{}
		err := OptionWithMaximumNumberOfSteps(maxStep)(g)
		require.NoError(t, err)
		require.Equal(t, maxStep, g.config.maxNumOfSteps)
	})
}
func TestOptionWithTargetCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		targetCost := float64(0.2)
		g := &ga{}
		err := OptionWithTargetCost(targetCost)(g)
		require.NoError(t, err)
		require.Equal(t, targetCost, g.config.targetCost)
	})
}
func TestOptionWithInitialPopulation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		initialPopulation := int(2000)
		g := &ga{}
		err := OptionWithInitialPopulation(initialPopulation)(g)
		require.NoError(t, err)
		require.Equal(t, initialPopulation, g.config.initialPopulation)
	})
}

func TestOptionWithNumberOfThreads(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		numberOfThreads := int(20)
		g := &ga{}
		err := OptionWithNumberOfThreads(numberOfThreads)(g)
		require.NoError(t, err)
		require.Equal(t, numberOfThreads, g.config.numberOfThreads)
	})
}
