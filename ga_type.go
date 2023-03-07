package goga

import "time"

type ga struct {

	// genarator generate random Models
	genarator func() Model

	// fitness is weght of each model to be selected for the Mutation for the next generation.
	fitness func(rank int64, cost float64) float64

	// population calculate the next generation popoulation based on the step, and best cost value.
	population func(step int64, bestCost float64) int64

	config struct {

		//maxNumOfSteps maximum number of steps
		// -1 for infinity run
		maxNumOfSteps int64

		// zero for infinity run.
		//
		// ga stops when  error < targetCost
		targetCost float64

		// the inverval between two step
		// usefull for infinity mood.
		stepsInterval time.Duration

		// reverse reverse the process to maximize the cost.
		reverse bool

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
