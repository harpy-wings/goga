package goga

type Model interface {
	// the cost function of the model to be Optimized.
	Cost() float64

	// Mutation is concatenation of the model with another model.
	//
	// Note: A.Mutate(B) != B.Mutate(A)
	Mutation(Model) Model
}
