package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* https://leetcode.com/problems/valid-parentheses/description/
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := -1
	dictLeft := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	dictRight := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, item := range s {
		_, ok := dictLeft[string(item)]
		if ok {
			top++
			stack[top] = item
		}
		_, ok = dictRight[string(item)]
		if ok {
			if top == -1 {
				return false
			}
			if string(stack[top]) == dictRight[string(item)] {
				top--
			} else {
				top++
				stack[top] = item
			}
		}
	}
	if top == -1 {
		return true
	}
	return false
}

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	result := isValid(s)
	assert.Equal(t, result, true)

	s = "]"
	result = isValid(s)
	assert.Equal(t, result, false)

	s = "(]"
	result = isValid(s)
	assert.Equal(t, result, false)
}
