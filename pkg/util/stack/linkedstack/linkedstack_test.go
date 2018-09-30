package linkedstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedStackPush(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push("str")
	assert.Equal(t, stack.Size(), 2)

	value, _ := stack.Pop()
	assert.Equal(t, value, "str")
	assert.Equal(t, stack.Size(), 1)

	value, _ = stack.Pop()
	assert.Equal(t, value, 1)
	assert.Equal(t, stack.Size(), 0)
}

func TestLinkedStackPeak(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push("str")
	value, _ := stack.Peak()
	assert.Equal(t, stack.Size(), 2)
	assert.Equal(t, value, "str")
	assert.Equal(t, stack.Size(), 2)
}
