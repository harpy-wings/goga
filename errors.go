package goga

import "fmt"

var (
	// ErrInvalidNilArgs in cases which an required argument is nil.
	ErrInvalidNilArgs = func(args ...interface{}) error { return fmt.Errorf("nil args %v", args) }

	// ErrInvalidSelection in cases which selection of genome is not valid
	ErrInvalidSelection = func(args ...interface{}) error { return fmt.Errorf("invalid selection, %v", args) }

	// ErrExecutionFailed any runtime failure that can have many reasons.
	ErrExecutionFailed = func(args ...interface{}) error { return fmt.Errorf("execution failed, %v", args) }
)
