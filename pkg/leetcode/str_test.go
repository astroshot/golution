package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// NaiveStrMatch returns index of `needle` in `haystack` using basic string iteration
func NaiveStrMatch(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	i, j := 0, 0
	if ln < 1 {
		return 0
	}
	if lh < ln {
		return -1
	}
	for i < lh {
		for haystack[i] != needle[j] {
			i++
			if i >= lh {
				return -1
			}
		}
		index := i
		for haystack[i] == needle[j] {
			i++
			j++
			if i == lh && j < ln-1 {
				return -1
			}
			if j == ln {
				return index
			}
		}
		i = index + 1
		j = 0
	}
	return -1
}

func KMP(haystack string, needle string) int {
	return -1
}

func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}

	return -1
}

func TestNaiveStrMatch(t *testing.T) {
	var (
		haystack string
		needle   string
		index    int
	)

	haystack = "hello"
	needle = "ll"
	index = NaiveStrMatch(haystack, needle)
	assert.Equal(t, 2, index)

	haystack = "hello"
	needle = "lo"
	index = NaiveStrMatch(haystack, needle)
	assert.Equal(t, 3, index)

	needle = "to"
	index = NaiveStrMatch(haystack, needle)
	assert.Equal(t, -1, index)

	haystack = "aaaa"
	needle = "aaa"
	index = NaiveStrMatch(haystack, needle)
	assert.Equal(t, 0, index)
}
