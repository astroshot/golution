package linkedlist

import (
	"learning_go/src/util/list/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	list := linkedlist.New()
	list.Add(0, 1, 2, 3)
	assert.Equal(t, list.String(), "0, 1, 2, 3")
}

func TestLinkedListReverse(t *testing.T) {
	list := linkedlist.New()
	list.Add(0, 1, 2, 3)
	list.Reverse()
	assert.Equal(t, list.String(), "3, 2, 1, 0")
}

func TestLinkedListGet(t *testing.T) {
	list := linkedlist.New()
	list.Add(0, 1, 2, 3)
	item, ok := list.Get(2)
	assert.Equal(t, item, 2)
	assert.Equal(t, ok, true)

	item, ok = list.Get(5)
	assert.Equal(t, item, nil)
	assert.Equal(t, ok, false)
}

func TestLinkedListContains(t *testing.T) {
	list := linkedlist.New()
	list.Add(0, 1, 2, 3, 4, 5, 6)
	result := list.Contains(0, 6, 4)
	assert.Equal(t, result, true)
	result = list.Contains(1, 3, 9)
	assert.Equal(t, result, false)
}

func TestLinkedListInsert(t *testing.T) {
	list := linkedlist.New()
	list.Add(0, 1, 2, 3, 4, 5, 6)
	list.Insert(4, 7, 8, 9)
	assert.Equal(t, list.String(), "0, 1, 2, 3, 7, 8, 9, 4, 5, 6")

	list.Clear()
	list.Add(0, 1, 2, 3, 4, 5, 6)
	list.Insert(1, 7, 8, 9)
	assert.Equal(t, list.String(), "0, 7, 8, 9, 1, 2, 3, 4, 5, 6")
}
