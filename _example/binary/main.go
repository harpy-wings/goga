package main

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/harpy-wings/goga"
)

var target = []byte("Hello World!")

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

func (s model) Encode() ([]byte, error) {
	return s, nil
}
func (s model) Decode(v []byte) error {
	s = v
	return nil
}

func Generator() goga.Model {
	m := make(model, len(target))
	rand.Read(m)
	return m
}

func main() {
	GA, err := goga.New(
		goga.OptionWithGeneratorFunc(Generator),
		goga.OptionWithTargetCost(0),
		goga.OptionWithSelection(2, 3, 2),
		goga.OptionWithInitialPopulation(600),
		// goga.OptionWithStepInterval(10*time.Millisecond),
		goga.OptionWithMaximumNumberOfSteps(1000000000),
	)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	GA.Start()
	C := GA.GetRunChan()
	go func() {
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
	M, err := GA.Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", time.Since(t))
	fmt.Println(M)
	fmt.Println(string(M.(model)))
}
