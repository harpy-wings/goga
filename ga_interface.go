package goga

// todo add comment
type GA interface {
	Start() error
	Result() (Model, error)
	RuntimeBestResult() chan RunTimeResult
}

// todo add comment
type RunTimeResult struct {
	Model Model
	Step  int64
	Cost  float64
}
