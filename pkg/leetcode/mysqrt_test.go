package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// mySqrt calculates sqrt of x, using simple iteration. It can be imporved by Newton iteration.
func mySqrt(x int) int {
	var i int
	for i = 1; i*i < x; i++ {
	}
	if i*i == x {
		return i
	}
	return i - 1
}

func TestMySqrt(t *testing.T) {
	var x, res int
	x = 4
	res = mySqrt(x)
	assert.Equal(t, 2, res)

	x = 8
	res = mySqrt(x)
	assert.Equal(t, 2, res)
}
