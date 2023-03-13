# goga
[![Go Reference](https://pkg.go.dev/badge/github.com/harpy-wings/goga.svg)](https://pkg.go.dev/github.com/harpy-wings/goga)
[![Go Report Card](https://goreportcard.com/badge/github.com/harpy-wings/goga)](https://goreportcard.com/report/github.com/harpy-wings/goga)
[![codecov](https://codecov.io/gh/harpy-wings/goga/branch/main/graph/badge.svg?token=R4Z4JTD87I)](https://codecov.io/gh/harpy-wings/goga)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=harpy-wings_goga&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=harpy-wings_goga)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=harpy-wings_goga&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=harpy-wings_goga)

---

Genetic Algorithm implementation in go.

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
