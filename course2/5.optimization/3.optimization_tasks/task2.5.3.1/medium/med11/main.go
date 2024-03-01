package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}}}
	newTree := removeLeafNodes(tree, 2)
	print(newTree)
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

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = removeLeafNodes(root.Right, target)
	root.Left = removeLeafNodes(root.Left, target)
	if root.Right == nil && root.Left == nil && root.Val == target {
		return nil
	}
	return root
}
