package linkedlist

import (
	"fmt"
	"strings"
)

// LinkedList holds items in a double link
type LinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	prev  *node
	next  *node
}

// New instance of a new empty LinkedList
func New() *LinkedList {
	return &LinkedList{}
}

// Add one or more items to a LinkedList
func (list *LinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &node{value: value, prev: list.tail, next: nil}
		if list.size == 0 {
			list.head = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		list.size++
	}
}

// Reverse reverses a LinkedList
func (list *LinkedList) Reverse() {
	if list.size < 2 {
		return
	}
	newHead := list.tail
	newTail := list.head

	p := list.head
	r := p.prev

	for p != nil {
		p.prev = p.next
		p.next = r
		r = p
		p = p.prev
	}
	list.head = newHead
	list.tail = newTail
}

// String returns string consists of values of LinkedList
func (list *LinkedList) String() string {
	str := ""
	values := make([]string, 0, list.size)
	cur := list.head
	for cur != nil {
		values = append(values, fmt.Sprintf("%v", cur.value))
		cur = cur.next
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *LinkedList) inRange(index int) bool {
	return index >= 0 && index < list.size
}

// Get returns item at the given index. O(n)
func (list *LinkedList) Get(index int) (interface{}, bool) {
	if !list.inRange(index) {
		return nil, false
	}
	i := 0
	p := list.head
	for i < index {
		p = p.next
		i++
	}
	return p.value, true
}

// Size returns number of items in LinkedList.
func (list *LinkedList) Size() int {
	return list.size
}

// Clear clears all items in LinkedList.
func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Values returns slice consists of all items in LinkedList.
func (list *LinkedList) Values() []interface{} {
	values := make([]interface{}, list.size)
	for i, item := 0, list.head; item != nil; i, item = i+1, item.next {
		values[i] = item.value
	}
	return values
}

// IsEmpty returns true when LinkedList contains no items.
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// Remove removes item at the given index.
func (list *LinkedList) Remove(index int) interface{} {
	if !list.inRange(index) {
		return nil
	}

	var item *node

	if list.size == 1 {
		item = list.head
		list.Clear()
		return item.value
	}

	if list.size-index < index {
		item = list.tail
		for i := list.size - 1; i != index; i, item = i-1, item.prev {
		}
	} else {
		item = list.head
		for i := 0; i != index; i, item = i+1, item.next {
		}
	}

	if item == list.head {
		list.head = item.next
	}
	if item == list.tail {
		list.tail = item.prev
	}
	if item.prev != nil {
		item.prev.next = item.next
	}
	if item.next != nil {
		item.next.prev = item.prev
	}
	list.size--
	return item.value
}

// Contains returns true when all values are in LinkedList.
func (list *LinkedList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for item := list.head; item != nil; item = item.next {
			if item.value == value {
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

// Insert add values at the given index.
func (list *LinkedList) Insert(index int, values ...interface{}) {
	if !list.inRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	valuesList := New()
	valuesList.Add(values...)
	var indexItem, prevItem *node
	if list.size-index < index {
		indexItem = list.tail
		for i := list.size - 1; i > index; i, indexItem = i-1, indexItem.prev {
		}
	} else {
		indexItem = list.head
		for i := 0; i < index; i, indexItem = i+1, indexItem.next {
		}
	}
	prevItem = indexItem.prev
	prevItem.next = valuesList.head
	valuesList.head.prev = prevItem

	valuesList.tail.next = indexItem
	indexItem.prev = valuesList.tail
	list.size += len(values)
}

// Swap swaps values of two items at the given indexes.
func (list *LinkedList) Swap(i, j int) {
	if list.inRange(i) && list.inRange(j) && i != j {
		var itemI, itemJ *node
		for i, item := 0, list.head; itemI == nil || itemJ == nil; i, item = i+1, item.next {
			switch i {
			case i:
				itemI = item
			case j:
				itemJ = item
			}
		}
		itemI.value, itemJ.value = itemJ.value, itemI.value
	}
}
