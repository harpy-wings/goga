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
 */
// TODO add body
func OptionWithFitnessFunc(fn FitnessFunc) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("fitness function")
		}
		g.fitness = fn
		return nil
	}
}

// OptionWithPopulationFunc define operations such as cross over, mutation and replacement during each iteration of GA.
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
func OptionWithGeneratorFunc(fn func() Model) Option {
	return func(g *ga) error {
		if fn != nil {
			return ErrInvalidNilArgs("generator function")
		}
		g.genarator = fn
		return nil
	}
}

// OptionWithSelection make subset of the population is selected for the next generation based on their fitness scores.
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

// OptionWithReverse
// todo add comment
func OptionWithReverse(v bool) Option {
	return func(g *ga) error {
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
// specify the target cost to when comuptation continues
// todo add comment
func OptionWithTargetCost(v float64) Option {
	return func(g *ga) error {
		g.config.targetCost = v
		return nil
	}
}

// OptionWithInitialPopulation
/*
 */
//n number of the first population intitialization of randomly solutions
func OptionWithInitialPopulation(n uint64) Option {
	return func(g *ga) error {
		g.config.initialPopulation = n
		return nil
	}
}
