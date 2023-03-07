package goga

type Model interface {
	// the cost function of the model to be Optimized.
	Cost() float64

	// Mutation is concatation of the model with another model.
	//
	// Note: A.Mutate(B) != B.Mutate(A)
	Mutation(Model) Model

	// Encode must encode the Model to array of bytes
	//
	// * required if you want to resote the GA from disk.
	Encode() ([]byte, error)

	// Decode must restore the model from array of bytes
	//
	// * required if you want to restore the model from Disk.
	Decode([]byte) error
}
