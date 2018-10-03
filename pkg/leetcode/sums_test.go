package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) [][]int {
	var res [][]int
	numMap := make(map[int]int)
	skipIndex := make(map[int]int)

	for i, val := range nums {
		numMap[val] = i
	}

	for i, val := range nums {
		if _, ok := skipIndex[i]; ok {
			continue
		}
		key := target - val
		if index, ok := numMap[key]; ok {
			row := []int{val, nums[index]}
			res = append(res, row)
			skipIndex[index] = val
		}
	}
	return res
}

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 4, 9}
	res := twoSum(nums, 5)
	expect := [][]int{
		{1, 4},
	}
	assert.Equal(t, res, expect)
}

func TestTwoSumNotFound(t *testing.T) {
	nums := []int{1, 2, 4, 9}
	res := twoSum(nums, 15)
	expect := [][]int(nil) // What is the difference between [][]int{} and [][]int{nil}
	assert.Equal(t, expect, res)
}

// func threeSum(nums []int, target int) [][]int {
// var res [][]int
// for i, val := range nums {
// resTwo := twoSum(nums, target-val)
// }
// return res
// }
