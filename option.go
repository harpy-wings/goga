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
func OptionWithDefaultGenerator(t interface{}) Option {
	return func(g *ga) error {

		return nil
	}
}

// OptionWithFitnessFunc
// todo add comment
func OptionWithFitnessFunc(fn FitnessFunc) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("fitness function")
		}
		g.fitness = fn
		return nil
	}
}

// OptionWithPopulationFunc
// todo add comment
func OptionWithPopulationFunc(fn PopulationFunc) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("population function")
		}
		g.population = fn
		return nil
	}
}

// OptionWithGeneratorFunc
// todo add comment
func OptionWithGeneratorFunc(fn func() Model) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("generator function")
		}
		g.genarator = fn
		return nil
	}
}

// OptionWithSelection
// todo add comment
func OptionWithSelection(top, mutaion, random float64) Option {
	return func(g *ga) error {
		sum := top + mutaion + random
		if sum == 0 {
			return ErrInvalidSlection(top, mutaion, random, "sum must not be zero")
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

// OptionWithStepInterval
// todo add comment
func OptionWithStepInterval(d time.Duration) Option {
	return func(g *ga) error {
		g.config.stepsInterval = d
		return nil
	}
}

// OptionWithMaximumNumberOfSteps
// todo add comment
func OptionWithMaximumNumberOfSteps(n int64) Option {
	return func(g *ga) error {
		g.config.maxNumOfSteps = n
		return nil
	}
}

// OptionWithTargetCost
// todo add comment
func OptionWithTargetCost(v float64) Option {
	return func(g *ga) error {
		g.config.targetCost = v
		return nil
	}
}

// OptionWithInitialPopulation
// todo add comment
func OptionWithInitialPopulation(n uint64) Option {
	return func(g *ga) error {
		g.config.initialPopulation = n
		return nil
	}
}
