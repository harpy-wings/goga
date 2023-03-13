package bench

import (
	"crypto/rand"
	mrand "math/rand"
	"testing"

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

func Generator() goga.Model {
	m := make(model, len(target))
	rand.Read(m)
	return m
}
func BenchmarkBinary(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GA, err := goga.New(
			goga.OptionWithGeneratorFunc(Generator),
			goga.OptionWithTargetCost(0),
			goga.OptionWithSelection(1, 2, 5),
			goga.OptionWithInitialPopulation(600),
			// goga.OptionWithStepInterval(10*time.Millisecond),
			goga.OptionWithMaximumNumberOfSteps(100),
		)
		if err != nil {
			panic(err)
		}
		GA.Start()

		res, err := GA.Result()
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			C := GA.RuntimeBestResult()
			for range C {
			}
		}()
		_ = res
		b.ReportMetric(res.Cost(), "cost/op")

	}
}
