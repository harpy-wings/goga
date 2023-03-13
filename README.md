# goga
[![Go Reference](https://pkg.go.dev/badge/github.com/harpy-wings/goga.svg)](https://pkg.go.dev/github.com/harpy-wings/goga)
[![Go Report Card](https://goreportcard.com/badge/github.com/harpy-wings/goga)](https://goreportcard.com/report/github.com/harpy-wings/goga)
[![codecov](https://codecov.io/gh/harpy-wings/goga/branch/main/graph/badge.svg?token=R4Z4JTD87I)](https://codecov.io/gh/harpy-wings/goga)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=harpy-wings_goga&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=harpy-wings_goga)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=harpy-wings_goga&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=harpy-wings_goga)

---

Genetic Algorithm implementation in go.

## Installation

```bash
go get -u github.com/harpy-wings/goga
```

## Usage

You need to define a type that implements the `goga.Model` interface based on your requirement. It can be a simple string matcher or a CNN neural network. Also, you need to pass a generator function to the *GA* algorithm, The *GA* uses the generator function to produce the new random genes or model.

```go
 // GeneratorFunc generates random Models.
 type GeneratorFunc  func() Model

 //Model is an interface which is about to being optimized.
 type Model interface {
    // the cost function of the model to be Optimized.
    Cost() float64
    // Mutation is concatenation of the model with another model.
    //
    // Note: A.Mutate(B) != B.Mutate(A)
    Mutation(Model) Model
}
```

> You can check the `_example` for more detailed examples.

### string matcher example

```go
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
		goga.OptionWithGeneratorFunc(Generator),        // pass the generator function
		goga.OptionWithTargetCost(0),                   // target cost
		goga.OptionWithSelection(1, 2, 5),              // Gens selection configuration
		goga.OptionWithInitialPopulation(600),          // initial population
		goga.OptionWithMaximumNumberOfSteps(1000),      // maximum number of steps.
	)
	if err != nil {
		panic(err)
	}
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

    // log the result and duration.
	fmt.Printf("\n%s\n", time.Since(t))
	fmt.Println(string(M.(model)))
}

```

## Parameters


## Benchmark

### Binary String Match

Use the GoGa package to find the `Hello World!`.
You can find the benchmark code at `test/bench/binary_test.go`.

*Result* :
```bs
goos: windows
goarch: amd64
pkg: github.com/harpy-wings/goga/test/bench
cpu: Intel(R) Core(TM) i9-10850K CPU @ 3.60GHz
BenchmarkBinary-20            88          12442376 ns/op                 0 cost/op       3502057 B/op      96357 allocs/op
PASS
ok      github.com/harpy-wings/goga/test/bench  2.046s

```

- target string: `Hello World!`
- cost func:

 ```go
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
 ```

- selection config:
    - top: 0.125
    - mutation: 0.25
    - random: 0.625
- maximum number of step: 100
- population: constant 600
- target cost: 0 (all cases match)

> You may get different results based on the random value generated or your system configuration. Also, the Genetic Algorithm may be locked to find the exact value. as observed it should find the exact result within the `25~35` steps, or else it is locked. Therefore we chose the `100` as the maximum number of steps for this test.
