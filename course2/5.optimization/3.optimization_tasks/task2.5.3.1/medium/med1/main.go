package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}}}}
	fmt.Println(deepestLeavesSum(tree))
}

func deepestLeavesSum(root *TreeNode) int {
	md := maxDepth(root)

	return sumMaxLevelRec(root, md)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)),
		float64(maxDepth(root.Right))))
}

func sumMaxLevelRec(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	if max == 1 {
		return root.Val
	}
	return sumMaxLevelRec(root.Left, max-1) + sumMaxLevelRec(root.Right, max-1)
}
