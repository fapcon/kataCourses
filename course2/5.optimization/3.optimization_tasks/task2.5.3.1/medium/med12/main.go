package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}}
	list2 := ListNode{Val: 777, Next: &ListNode{Val: 777, Next: &ListNode{Val: 777, Next: &ListNode{Val: 888, Next: nil}}}}
	//fmt.Println(mergeInBetween(&list1, 3, 4, &list2))
	mergeInBetween(&list1, 3, 4, &list2)
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node := list1
	var start, end *ListNode
	for i := 0; node != nil; i++ {
		if i == a-1 {
			start = node
		}
		if i == b {
			end = node
		}
		node = node.Next
	}
	start.Next = list2
	node = list2
	for node.Next != nil {
		node = node.Next
	}
	node.Next = end.Next
	res := list1
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	return list1
}
