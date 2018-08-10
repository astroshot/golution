package arraylist

import (
	"fmt"
	"strings"
)

// ArrayList implemented by slice of type interface{}
type ArrayList struct {
	items    []interface{}
	capacity int // capacity of items
	size     int // actually used of items
}

const (
	minCap = 16
)

// New instance of arraylist, error happens when capacity <= 0.
func New(capacity int) (*ArrayList, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("Initial capacity of ArrayList should be positive")
	}
	items := make([]interface{}, capacity)
	return &ArrayList{capacity: capacity, size: 0, items: items}, nil
}

func (array *ArrayList) inRange(index int) bool {
	return index >= 0 && index < array.size
}

func (array *ArrayList) resize(newCapacity int) {
	newItems := make([]interface{}, newCapacity)
	copy(newItems, array.items)
	array.items = newItems
	array.capacity = newCapacity
}

func (array *ArrayList) grow(n int) error {
	if array.size+n <= array.capacity {
		return nil
	}

	newCapacity := array.capacity << 1
	if newCapacity < 0 {
		return fmt.Errorf("Capacity of ArrayList overflows: %d", newCapacity)
	}
	array.resize(newCapacity)
	return nil
}

func (array *ArrayList) shrink() {
	newCap := array.capacity >> 1
	if newCap <= minCap {
		return
	}
	if array.size <= newCap {
		array.resize(newCap)
	}
}

// Add values to a given ArrayList
func (array *ArrayList) Add(values ...interface{}) error {
	err := array.grow(len(values))
	if err != nil {
		return fmt.Errorf("Add failed")
	}

	for _, value := range values {
		array.items[array.size] = value
		array.size++
	}
	return nil
}

// Get element by index
func (array *ArrayList) Get(index int) (interface{}, bool) {
	if !array.inRange(index) {
		return nil, false
	}
	return array.items[index], true
}

// IndexOf value first appeared in ArrayList
func (array *ArrayList) IndexOf(value interface{}) int {
	if array.size == 0 {
		return -1
	}
	for index, item := range array.items {
		if item == value {
			return index
		}
	}
	return -1
}

// Size returns number of items in an array
func (array *ArrayList) Size() int {
	return array.size
}

// IsEmpty returns true if array has nothing
func (array *ArrayList) IsEmpty() bool {
	return array.size == 0
}

// Clear removes all items in an ArrayList
func (array *ArrayList) Clear() {
	array.size = 0
	array.capacity = 0
	array.items = []interface{}{}
}

// Swap swaps two values at the given index i and j.
func (array *ArrayList) Swap(i, j int) {
	if array.inRange(i) && array.inRange(j) && i != j {
		array.items[i], array.items[j] = array.items[j], array.items[i]
	}
}

// Remove the given index item and return the removed item
func (array *ArrayList) Remove(index int) interface{} {
	if !array.inRange(index) {
		return nil
	}
	item := array.items[index]
	array.items[index] = nil
	copy(array.items[index:], array.items[index+1:array.size])
	array.size--
	array.shrink()
	return item
}

// Contains returns true when all values are in the ArrayList.
func (array *ArrayList) Contains(values ...interface{}) bool {
	for _, target := range values {
		found := false
		for _, item := range array.items {
			if item == target {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Insert values at the given index.
func (array *ArrayList) Insert(index int, values ...interface{}) error {
	if !array.inRange(index) {
		return fmt.Errorf("Index out of range")
	}

	num := len(values)
	array.grow(num)
	copy(array.items[index+num:array.size+num], array.items[index:array.size])
	copy(array.items[index:index+num], values[:])
	array.size += num
	return nil
}

// String returns a string filled with values of ArrayList
func (array *ArrayList) String() string {
	str := ""
	values := make([]string, 0, array.size)
	for _, value := range array.items[:array.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Values returns all items in ArrayList.
func (array *ArrayList) Values() []interface{} {
	results := make([]interface{}, array.size)
	copy(results, array.items[:array.size])
	return results
}
