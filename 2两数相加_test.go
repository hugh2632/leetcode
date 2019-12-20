package leetcode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var parent ListNode
	var res = &parent
	var plusRes = 0
	for plusRes != 0 || l1 != nil || l2 != nil {
		var child ListNode
		if l1 != nil {
			plusRes += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			plusRes += l2.Val
			l2 = l2.Next
		}
		if plusRes > 9 {
			child.Val = plusRes - 10
			plusRes = 1
		} else {
			child.Val = plusRes
			plusRes = 0
		}
		res.Next = &child
		res = res.Next
	}
	return parent.Next
}

func TestAddTwo(t *testing.T) {
	var l1 = ListNode{
		Val: 9,
	}
	var l2 = ListNode{
		Val: 9,
	}
	var res = addTwoNumbers(&l1, &l2)
	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val)
}

