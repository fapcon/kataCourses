package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(mergeNodes(&ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: nil}}}}))
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Val == 0 && head.Next == nil {
		return nil
	}
	curr := head
	mod := &ListNode{}
	for curr.Next != nil {
		if curr.Val == 0 {
			mod = curr
			curr = curr.Next
		}
		if curr.Val != 0 {
			mod.Val += curr.Val
			*curr = *curr.Next
		}
	}
	mod.Next = nil
	return head
}
