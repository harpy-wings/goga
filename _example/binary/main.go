package main

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/harpy-wings/goga"
)

var target = []byte("Hello World!")

// define the model which implements the goga.Model interface.
type model []byte

func (s model) Cost() float64 {
	var cost float64 = 0
	for i := range s {
		if s[i] == target[i] {
			continue
		}
		cost++
	}

	return cost
}

func (s model) Mutation(m goga.Model) goga.Model {
	bm := m.(model)
	res := make(model, len(s))
	n := mrand.Intn(len(s) - 1)
	for i := 0; i < len(s); i++ {
		if i <= n {
			res[i] = s[i]
		} else {
			res[i] = bm[i]
		}
	}

	return res
}

// define the Generator function to generate Random model.
func Generator() goga.Model {
	m := make(model, len(target))
	rand.Read(m)
	return m
}

func main() {
	//initializing the GA algorithm
	GA, err := goga.New(
		goga.OptionWithGeneratorFunc(Generator),   // pass the generator function
		goga.OptionWithTargetCost(0),              // target cost
		goga.OptionWithSelection(1, 2, 5),         // Gens selection configuration
		goga.OptionWithInitialPopulation(600),     // initial population
		goga.OptionWithMaximumNumberOfSteps(1000), // maximum number of steps.
	)
	if err != nil {
		panic(err)
	}

	//define the current time for measuring the GA duration, Optional you can remove it if you want.
	t := time.Now()

	// start will start the algorithm in background and will not lock the current thread.
	GA.Start()
	// you can get the runtime best result in each step from this channel.
	C := GA.RuntimeBestResult()
	go func() {
		// log the best Result in each step just to see. this is Optional you can remove this part.
		m := ""
		for v := range C {
			strVal := string(v.Model.(model))
			if m != strVal && v.Cost < 4 {
				fmt.Printf("\n")
				m = strVal
			}
			fmt.Printf("\rStep: %d\tCost: %0.1f\tVal: %s", v.Step, v.Cost, strVal)
		}
	}()

	// get the final result once the target cost find or the maximum number of steps reached. this will lock the current thread until the GA find the result.
	M, err := GA.Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", time.Since(t))
	fmt.Println(string(M.(model)))
}
