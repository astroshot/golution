package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.
Solve it without converting the integer to a string.

Examples:

Input: 121
Output: true

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Idea:
Save every digit of input integer to an array.
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	NUM := 30
	arr := make([]int, NUM, NUM)
	n := 0
	for x != 0 {
		arr[n] = x % 10
		x = x / 10
		n++
	}

	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func isPalindromeV2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseNum := 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		x = x / 10
	}

	return x == reverseNum || x == reverseNum/10
}

func TestIsPalindrome(t *testing.T) {
	x := 12
	res := isPalindromeV2(x)
	assert.Equal(t, res, false)

	x = 2
	res = isPalindromeV2(x)
	assert.Equal(t, res, true)

	x = 1221
	res = isPalindromeV2(x)
	assert.Equal(t, res, true)

	x = -1221
	res = isPalindromeV2(x)
	assert.Equal(t, res, false)
}
