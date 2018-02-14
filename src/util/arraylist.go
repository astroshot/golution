package util

// ArrayList implemented by slice of type interface{}
type ArrayList struct {
	items    []interface{}
	capacity int
	size     int
}

// New instance of arraylist
func New() *ArrayList {
	return &ArrayList{}
}

func (array *ArrayList) inRange(index int) bool {
	return index >= 0 && index < array.size
}

func (array *ArrayList) grow() {
	array.capacity = array.capacity << 1
	newItems := make([]interface{}, array.capacity)
	for i := 0; i < array.size; i++ {
		newItems[i] = array.items[i]
	}
	array.items = newItems
}

func (array *ArrayList) Add(values ...interface{}) {
	if array.size+len(values) >= array.capacity {
		array.grow()
	}
	for _, value := range values {
		array.items[array.size] = value
		array.size++
	}
}

// Push item in an ArrayList
func (array *ArrayList) Push(item interface{}) (size int) {
	if array.size+1 >= array.capacity {
		array.grow()
	}
	array.items[array.size] = item
	array.size++
	return array.size
}

// Pop item from an ArrayList
func (array *ArrayList) Pop() (item interface{}) {
	oldItem := array.items[array.size]
	array.size--
	return oldItem
}

// Get element by index
func (array *ArrayList) Get(index int) (interface{}, bool) {
	if !array.inRange(index) {
		return nil, false
	}
	return array.items[index], true
}

// IndexOf value in array
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

// Empty returns true if array has nothing
func (array *ArrayList) Empty() bool {
	return array.size == 0
}

// Clear removes all items in an array
func (array *ArrayList) Clear() {
	array.size = 0
	array.capacity = 0
	array.items = []interface{}{}
}
