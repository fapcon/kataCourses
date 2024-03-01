package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 4, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	fmt.Println(averageOfSubtree(tree))
}

func averageOfSubtree(root *TreeNode) int {
	_, _, res := helper(root)
	return res
}

func helper(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	}
	sum1, count1, match1 := helper(node.Left)
	sum2, count2, match2 := helper(node.Right)
	sum := sum1 + sum2 + node.Val
	count := count1 + count2 + 1
	match := match1 + match2
	if node.Val == sum/count {
		match++
	}
	return sum, count, match
}
