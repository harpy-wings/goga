package goga

import "time"

var (
	defaultSelectionTop      = 0.2
	defaultSelectionMutation = 0.6

	defaultMaxNumOfSteps = int64(1000)

	defaultTargetCost = float64(0.05)

	defaultStepsInterval = time.Duration(0)

	// defaultFitness default fitness function, equal chance.
	defaultFitness = func(int64, float64) float64 { return 1 }

	defaultInitialPopulation = uint64(1000)
	defaultPopulation        = func(currentPopulation int64, step int64, bestCost, worstCost float64) int64 { return currentPopulation }
)
