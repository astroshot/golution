package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func toString(head *ListNode) string {
	if head == nil {
		fmt.Println("Empty ListNode")
	}
	str := ""
	values := make([]string, 0, 50) // only for test

	p := head
	for p != nil {
		values = append(values, fmt.Sprintf("%v", p.Val))
		p = p.Next
	}
	str += strings.Join(values, ", ")
	return str
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var p, q, r *ListNode
	p = head
	q = nil
	for p != nil {
		r = q
		q = p
		p = p.Next
		q.Next = r
	}
	return q
}

func buildListFromArray(arr []int) *ListNode {
	var p, head *ListNode
	p = nil
	for _, value := range arr {
		q := &ListNode{Val: value, Next: nil}
		if p != nil {
			p.Next = q
		} else {
			head = q
		}
		p = q
	}
	return head
}

func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		} else {
			// list contains even nodes

		}
	}
	tail := reverse(slow)
	p := head
	for p != tail {
		if p.Val != tail.Val {
			return false
		}
		p = p.Next
		tail = tail.Next
	}
	return true
}

func TestIsPalindromeList(t *testing.T) {
	x := []int{1, 2, 3}
	head := buildListFromArray(x)
	str := toString(head)
	assert.Equal(t, str, "1, 2, 3, 4")

	tail := reverse(head)
	str = toString(tail)
	assert.Equal(t, str, "4, 3, 2, 1")

}
