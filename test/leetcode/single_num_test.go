package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
func singleNumber(nums []int) int {
	var res int
	for i, val := range nums {
		if i == 0 {
			res = val
		} else {
			res = res ^ val
		}
	}
	return res
}

func TestSingleNum(t *testing.T) {
	input := []int{1, 1, 2, 4, 4}
	res := singleNumber(input)
	assert.Equal(t, res, 2)
}

func TestNor(t *testing.T) {
	assert.Equal(t, 2^2, 0)
	assert.Equal(t, 1^0, 1)
}
