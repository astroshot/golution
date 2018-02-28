package arraylist

import (
	arraylist "go_tools/src/util/list/arraylist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayListGet(t *testing.T) {
	array, _ := arraylist.New(16)
	array.Add(0, 1, 2, 3, "test", 5)
	item, _ := array.Get(4)
	assert.Equal(t, item, "test")
}

func TestArrayListAdd(t *testing.T) {
	num := 2
	array, _ := arraylist.New(num)
	array.Add(1)
	array.Add("str", 3, 4)
	assert.Equal(t, array.Size(), 4)
	assert.Equal(t, array.String(), "1, str, 3, 4")
}

func TestArrayListRemove(t *testing.T) {
	capacity := 10
	array, _ := arraylist.New(capacity)
	array.Add(0, 1, 2, 3, 4, 5, 6, 7)
	removedItem := array.Remove(5)
	assert.Equal(t, removedItem, 5)
	item, _ := array.Get(5)
	assert.Equal(t, item, 6)
	assert.Equal(t, array.Size(), 7)
}

func TestArrayListContains(t *testing.T) {
	capacity := 10
	array, _ := arraylist.New(capacity)
	array.Add(0, 1, 2, 3, 4, 5, 6, 7)
	exist := array.Contains(1, 3, 6)
	assert.Equal(t, exist, true)
	exist = array.Contains(2, 6, 9)
	assert.Equal(t, exist, false)
}

func TestArrayListToString(t *testing.T) {
	capacity := 10
	array, _ := arraylist.New(capacity)
	array.Add(0, 1, 2, 3, 4, 5, 6)
	assert.Equal(t, array.String(), "0, 1, 2, 3, 4, 5, 6")
}

func TestArrayListInsert(t *testing.T) {
	capacity := 10
	array, _ := arraylist.New(capacity)
	array.Add(0, 1, 2, 3, 4, 5, 6)
	array.Insert(3, 7, 8, 9)
	assert.Equal(t, array.String(), "0, 1, 2, 7, 8, 9, 3, 4, 5, 6")
}

func TestArrayListSwap(t *testing.T) {
	capacity := 10
	array, _ := arraylist.New(capacity)
	array.Add(0, 1, 2, 3, 4, 5, 6)
	array.Swap(0, 4)
	assert.Equal(t, array.String(), "4, 1, 2, 3, 0, 5, 6")
}
