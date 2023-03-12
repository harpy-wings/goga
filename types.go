package goga

type (
	WeightFunc     func(rank int, cost float64) int
	PopulationFunc func(currentPopulation int, step int64, bestCost, worstCost float64) int
)
