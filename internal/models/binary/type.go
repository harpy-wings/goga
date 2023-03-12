// binary package is binary internal model for testing purpose which designed to gets empty as the best Case.
package binary

import (
	mrand "math/rand"

	"github.com/harpy-wings/goga"
)

// target is to be empty

// Model is a binary internal model for testing purposes.
type Model []byte

func (s Model) Cost() float64 {
	return float64(len(s))
}

func (s Model) Mutation(m goga.Model) goga.Model {
	bm := m.(Model)
	res := make(Model, len(s))
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
