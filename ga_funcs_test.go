package goga

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockedModel struct{}

func (*mockedModel) Cost() float64 {
	return float64(rand.Intn(10)) + 0.01 // to not be empty
}
func (s *mockedModel) Mutation(Model) Model { return s }

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

func TestStart(t *testing.T) {
	t.Run("reach target cost", func(t *testing.T) {
		GA, err := New()
		require.NoError(t, err)
		require.NotNil(t, GA)
		population := 10
		ga := GA.(*ga)
		ga.config.numberOfThreads = 3
		ga.generator = func() Model { return &mockedModel{} }
		ga.config.targetCost = 0.01
		ga.config.maxNumOfSteps = 10000
		ga.config.initialPopulation = population
		require.NoError(t, err)
		err = ga.Start()
		require.NoError(t, err)

		res, err := ga.Result()
		t.Log(ga.step)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, 0.01, ga.curetGeneration[0].cost)
	})

	t.Run("reach maximum number of steps", func(t *testing.T) {
		GA, err := New()
		require.NoError(t, err)
		require.NotNil(t, GA)
		population := 10
		ga := GA.(*ga)
		ga.config.numberOfThreads = 3
		ga.generator = func() Model { return &mockedModel{} }
		ga.config.targetCost = 0.001
		ga.config.maxNumOfSteps = 100
		ga.config.initialPopulation = population
		require.NoError(t, err)
		err = ga.Start()
		require.NoError(t, err)

		res, err := ga.Result()
		t.Log(ga.step)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, ga.config.maxNumOfSteps, ga.step)
	})

}

func TestResult(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		GA, err := New()
		require.NoError(t, err)
		require.NotNil(t, GA)
		go func() {
			ga := GA.(*ga)
			ga.result <- nil
		}()
		_, err = GA.Result()
		require.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {

		t.Run("closed channel", func(t *testing.T) {
			GA, err := New()
			require.NoError(t, err)
			require.NotNil(t, GA)
			go func() {
				ga := GA.(*ga)
				close(ga.result)
			}()
			Res, err := GA.Result()
			require.Error(t, err)
			require.Nil(t, Res)
		})

		t.Run("runtime error", func(t *testing.T) {
			GA, err := New()
			require.NoError(t, err)
			require.NotNil(t, GA)
			go func() {
				ga := GA.(*ga)
				ga.runtimeError <- errors.New("any")
			}()
			Res, err := GA.Result()
			require.Error(t, err)
			require.Nil(t, Res)
		})

	})
}

func TestGenerateRemindingGeneration(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	ga := GA.(*ga)
	population := 10
	ga.generator = func() Model { return &mockedModel{} }
	ga.curetGeneration = modelSortedList{{}, {}, {}}
	err = ga.generateRemindingGeneration(population)
	require.NoError(t, err)
	require.Equal(t, population, len(ga.curetGeneration))
}

func TestCalculateCosts(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	ga := GA.(*ga)
	population := 10
	ga.config.numberOfThreads = 3
	ga.generator = func() Model { return &mockedModel{} }
	err = ga.generateRemindingGeneration(population)
	require.NoError(t, err)

	ga.calculateCosts()
	for _, v := range ga.curetGeneration {
		require.NotEmpty(t, v.cost)
	}

}
func TestSort(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	ga := GA.(*ga)
	population := 10
	ga.config.numberOfThreads = 3
	ga.generator = func() Model { return &mockedModel{} }
	err = ga.generateRemindingGeneration(population)
	require.NoError(t, err)

	ga.calculateCosts()
	ga.sort()
	require.Less(t, ga.curetGeneration[0].cost, ga.curetGeneration[population-1].cost)

}
func TestGetTop(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	population := 10
	ga := GA.(*ga)
	ga.config.numberOfThreads = 3
	ga.generator = func() Model { return &mockedModel{} }
	err = ga.generateRemindingGeneration(population)
	require.NoError(t, err)

	ga.calculateCosts()
	ga.sort()
	ga.config.selection.top = 0.2
	res := ga.getTop(population)
	require.Equal(t, int(ga.config.selection.top*float64(population)), len(res))

}

func TestGetMutation(t *testing.T) {
	GA, err := New()
	require.NoError(t, err)
	require.NotNil(t, GA)
	population := 10
	ga := GA.(*ga)
	ga.config.numberOfThreads = 3
	ga.generator = func() Model { return &mockedModel{} }
	err = ga.generateRemindingGeneration(population)
	require.NoError(t, err)

	ga.calculateCosts()
	ga.sort()
	ga.config.selection.top = 0.1
	ga.config.selection.mutation = 0.2 // total 30%
	res := ga.getMutation(population)
	require.Equal(t, int(ga.config.selection.mutation*float64(population)), len(res))

}
