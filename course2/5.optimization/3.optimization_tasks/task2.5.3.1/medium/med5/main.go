package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var i, l, res int
	node := head
	for node != nil {
		l++
		node = node.Next
	}
	sl := make([]int, l/2)
	node = head
	for node != nil {
		if i > (l/2 - 1) {
			res = max(res, sl[l-i-1]+node.Val)
		} else {
			sl[i] = node.Val
		}
		node = node.Next
		i++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	l := &ListNode{Val: 10, Next: &ListNode{Val: 200, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	fmt.Println(pairSum(l))
}
