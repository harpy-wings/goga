package goga

type (
	FitnessFunc    func(rank int64, cost float64) float64
	PopulationFunc func(currentPopulation int64, step int64, bestCost, worstCost float64) int64
)
