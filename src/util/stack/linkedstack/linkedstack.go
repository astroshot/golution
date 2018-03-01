package linkedstack

import (
	"modules/src/util/list/linkedlist"
)

// LinkedStack provides stack interface with a double linked list.
type LinkedStack struct {
	list *linkedlist.LinkedList
}

// New returns LinkedStack.
func New() *LinkedStack {
	return &LinkedStack{list: &linkedlist.LinkedList{}}
}

// Push add items to a LinkedStack.
func (stack *LinkedStack) Push(value interface{}) {
	stack.list.Add(value)
}

// Pop returns the top value and remove it from stack.
func (stack *LinkedStack) Pop() (value interface{}, ok bool) {
	top := stack.list.Size() - 1
	value, ok = stack.list.Get(top)
	stack.list.Remove(top)
	return value, ok
}

// Peak returns the top value without removing it from stack
func (stack *LinkedStack) Peak() (value interface{}, ok bool) {
	top := stack.list.Size() - 1
	value, ok = stack.list.Get(top)
	return value, ok
}

// Size returns number of items in LinkedStack
func (stack *LinkedStack) Size() int {
	return stack.list.Size()
}

// IsEmpty returns true when LinkedStack contains no item
func (stack *LinkedStack) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// Clear clears LinkedStack
func (stack *LinkedStack) Clear() {
	stack.list.Clear()
}

// Values returns slice consists of all items in LinkedStack.
func (stack *LinkedStack) Values() []interface{} {
	return stack.list.Values()
}
