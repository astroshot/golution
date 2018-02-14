package util

import (
	util "learning_go/src/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayListPush(t *testing.T) {
	num := 20
	array := util.NewArrayList(num)
	array.Push(num)
	tmp := array.Pop()
	assert.Equal(t, tmp, num)
}
