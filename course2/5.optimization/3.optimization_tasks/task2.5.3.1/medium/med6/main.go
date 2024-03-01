package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	tree := constructMaximumBinaryTree(nums)
	print(tree)
}

func print(n *TreeNode) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Val)
		print(n.Left)
		print(n.Right)
	}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			j = i
		}
	}
	sliceLeft := nums[:j]
	sliceRight := nums[j+1:]

	return &TreeNode{Val: max, Left: constructMaximumBinaryTree(sliceLeft), Right: constructMaximumBinaryTree(sliceRight)}
}
