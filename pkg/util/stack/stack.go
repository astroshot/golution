package stack

import "golution/pkg/util/container"

// Stack interface
type Stack interface {
	Push(value interface{})
	// Pop returns the top value and remove it from stack
	Pop() (value interface{}, ok bool)
	// Peak returns the top value without removing it from stack
	Peak() (value interface{}, ok bool)
	container.Container
}
