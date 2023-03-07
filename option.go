package goga

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

// OptionWithFitnessFunc
// todo add comment
func OptionWithFitnessFunc(fn FitnessFunc) Option {
	return func(g *ga) error {
		if fn == nil {
			return ErrInvalidNilArgs("fitness function")
		}
		g.fitness = fn
		return nil
	}
}

// OptionWithPopulationFunc
// todo add comment
func OptionWithPopulationFunc(fn PopulationFunc) Option {
	return func(g *ga) error {
		if fn != nil {
			return ErrInvalidNilArgs("fitness function")
		}
		g.population = fn
		return nil
	}
}
