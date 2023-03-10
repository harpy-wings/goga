package goga

import "time"

type Option func(*ga) error

// OptionWithDefaultGenerator use the reflect to generate random values for the properties of provided type.
/*
 type model struct {
		age 	int64		`ga:min:0,max:100`
		height 	float64 	`ga:"min:0,max:1"`
 }
 // ...
 OptionWithDefaultGenerator(model{})
*/
// TODO add body
func OptionWithDefaultGenerator(t any) Option {
	return func(g *ga) error {
		return nil
	}
}

// OptionWithFitnessFunc takes a function to calculate chance of the solution to be selected during crossover.
// Cost and score will be defined separately (Generic solution) based on how the individual is calculated and gets score
/*
	TODO add example
*/
func OptionWithFitnessFunc(fn WeightFunc) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("fitness function")
		}
		g.weightFunc = fn
		return nil
	}
}

// OptionWithPopulationFunc define operations such as cross over, mutation and replacement during each iteration of GA.
/*
	TODO add example
*/
func OptionWithPopulationFunc(fn PopulationFunc) Option {
	return func(g *ga) error {
		if fn != nil {
			return ErrInvalidNilArgs("population function")
		}
		g.population = fn
		return nil
	}
}

// OptionWithGeneratorFunc generator to make solutions randomly
/*
	TODO add example
*/
func OptionWithGeneratorFunc(fn func() Model) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("generator function")
		}
		g.generator = fn
		return nil
	}
}

// OptionWithSelection make subset of the population is selected for the next generation based on their fitness scores.
//
// default: 0.2,0.6,0.2
//
// ex: the following options works the same.
//
//	OptionWithSelection(0.2,0.6,0.2) 	// 20% top, 60% crossover, 20% random
//	OptionWithSelection(20,60,20) 		// 20% top, 60% crossover, 20% random
//	OptionWithSelection(1,3,1) 		// 20% top, 60% crossover, 20% random
func OptionWithSelection(top, mutaion, random float64) Option {
	return func(g *ga) error {
		sum := top + mutaion + random
		if sum == 0 {
			return ErrInvalidSelection(top, mutaion, random, "sum must not be zero")
		}
		top = top / sum
		mutaion = top / mutaion
		random = top / random
		g.config.selection.top = top
		g.config.selection.mutation = mutaion
		g.config.selection.random = random
		return nil
	}
}

// OptionWithStepInterval is a function that returns an Option which sets the interval between two generations.
// It is useful for runtime processes where you want to continuously improve the parameters, especially in cases where the cost function is not constant and changes over time.
//
// default: 0 (no wait)
//
// ex:
//
//	OptionWithStepInterval(30 * time.Second)
func OptionWithStepInterval(d time.Duration) Option {
	return func(g *ga) error {
		g.config.stepsInterval = d
		return nil
	}
}

// OptionWithMaximumNumberOfSteps is a function that returns an Option which sets the maximum number of steps for a genetic algorithm.
//
// default: 1000
//
// ex:
//
//	OptionWithMaximumNumberOfSteps(1000)
func OptionWithMaximumNumberOfSteps(n int64) Option {
	return func(g *ga) error {
		g.config.maxNumOfSteps = n
		return nil
	}
}

// OptionWithTargetCost is a function that returns an Option which sets the target cost for a genetic algorithm.
// The target cost specifies the point at which the process will stop once it reaches that cost.
//
// default: 0.05
//
// ex:
//
//	OptionWithTargetCost(0.2)
func OptionWithTargetCost(v float64) Option {
	return func(g *ga) error {
		g.config.targetCost = v
		return nil
	}
}

// OptionWithInitialPopulation is a function that returns an Option that sets the initial population of the genetic algorithm.
//
// default: 1000
//
// ex:
//
//	OptionWithInitialPopulation(10000)
func OptionWithInitialPopulation(n uint64) Option {
	return func(g *ga) error {
		g.config.initialPopulation = n
		return nil
	}
}
