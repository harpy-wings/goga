package goga

import "time"

type ga struct {

	// generator generate random Models
	generator func() Model

	// fitness is weight of each model to be selected for the Mutation for the next generation.
	fitness FitnessFunc

	// population calculate the next generation population based on the step, and best cost value.
	population PopulationFunc

	config struct {

		// initialPopulation is the population of the first generation.
		initialPopulation uint64

		//maxNumOfSteps maximum number of steps
		// -1 for infinity run
		maxNumOfSteps int64

		// zero for infinity run.
		//
		// ga stops when  error < targetCost
		targetCost float64

		// the interval between two step
		// useful for infinity mood.
		stepsInterval time.Duration

		// selection is the property for distribution to the next generation.
		//
		// top + mutation + random = 1
		selection struct {
			top      float64
			mutation float64
			random   float64
		}
	}
}
