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

> You can check the [`_example`](https://github.com/harpy-wings/goga/tree/main/_example/) for more detailed examples.

## Parameters

- ### Generator Function

  - Required
  - `OptionWithGeneratorFunc(fn func() Model)`

  - The Generator function is used by *GA* to generate Random model.

- ### Weight Function

  - Optional
  - Default: equal weight
  - `OptionWithWeightFunc(fn WeightFunc)`
    - `type WeightFunc func(rank int, cost float64) int`
  - The Weight Function can be use to select the next generation biased on their weights. you can make the design based it the values which are given you as arguments and return a weight as int.
    - > note you can use Decorator Pattern to access some other values from your code.

- ### Population Function

  - Optional
  - Default: the initial population as a constant Population.
  - `OptionWithPopulationFunc(fn PopulationFunc)`
    - `PopulationFunc func(currentPopulation int, step int64, bestCost, worstCost float64) int`
  - The Population function can be use to manage the population in the process, based on the current population, step, bestCost and worst cost.

- ### Initial Population

  - Optional
  - Default: 1000
  - `OptionWithInitialPopulation(n int)`
  - The initial Population of *GA*.

- ### Maximum Number of Steps

  - Optional
  - Default: 1000
  - `OptionWithMaximumNumberOfSteps(n int64)`
  - The maximum number of steps if *GA* could not reach the `targetCost`.

- ### Target Cost

  - Required
  - Default: 0.05
  - `OptionWithTargetCost(v float64)`
  - The target cost specifies the point at which the process will stop once it reaches that cost.

- ### Steps Interval

  - Optional
  - Default: 0
  - `OptionWithStepInterval(d time.Duration)`
  - In Always run mode, once the *GA* reach the CostTarget, it will sleep that much time between two steps. It is useful in situation which you want to continuously improve the model.

- ### Selection Configuration

  - Optional
  - Default: Top: 20%, Mutation: 40%, Random: 40%
  - `OptionWithSelection(top, mutation, random float64) Option`
  - The subset of the population is selected for the next generation.

- ### Number of Threads

  - Optional
  - Default: 12
  - `OptionWithNumberOfThreads(n int)`
  - Specify the number of thread in Cost calculation step.

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
- maximum number of steps: 100
- population: constant 600
- target cost: 0 (all cases match)

> You may get different results based on the random value generated or your system configuration. Also, the Genetic Algorithm may be locked to find the exact value. as observed it should find the exact result within the `25~35` steps, or else it is locked. Therefore we chose the `100` as the maximum number of steps for this test.
