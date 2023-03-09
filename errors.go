package goga

import "fmt"

var (
	ErrInvalidNilArgs   = func(args ...interface{}) error { return fmt.Errorf("nil args %v", args) }
	ErrInvalidSelection = func(args ...interface{}) error { return fmt.Errorf("invalid selection, %v", args) }
)
