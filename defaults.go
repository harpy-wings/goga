package goga

import "time"

var (
	defaultSelectionTop      = 0.2
	defaultSelectionMutation = 0.6
	defaultNumberOfThreads   = 12

	defaultMaxNumOfSteps = int64(1000)

	defaultTargetCost = float64(0.05)

	defaultStepsInterval = time.Duration(0)

	// defaultWeightFunc default fitness function, equal chance.
	defaultWeightFunc = func(int, float64) int { return 1 }

	defaultInitialPopulation = int(1000)

	defaultPopulation = func(currentPopulation int, step int64, bestCost, worstCost float64) int { return currentPopulation }
)
