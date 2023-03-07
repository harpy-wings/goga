package goga

import "fmt"

var (
	ErrInvalidNilArgs  = func(args ...interface{}) error { return fmt.Errorf("nil args %v", args) }
	ErrInvalidSlection = func(args ...interface{}) error { return fmt.Errorf("invalid selection, %v", args) }
)
