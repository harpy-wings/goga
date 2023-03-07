package goga

import "fmt"

var (
	ErrInvalidNilArgs = func(args ...interface{}) error { return fmt.Errorf("nil args %v", args) }
)
