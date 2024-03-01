package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}}}}
	newTree := bstToGst(tree)
	//print(tree)
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

func bstToGst(root *TreeNode) *TreeNode {
	h := 0
	return helper(root, &h)
}

func helper(root *TreeNode, lastVal *int) *TreeNode {
	if root == nil {
		return nil
	}

	right := helper(root.Right, lastVal)
	newRoot := &TreeNode{}
	newRoot.Val = root.Val + *lastVal
	newRoot.Right = right

	*lastVal = newRoot.Val

	left := helper(root.Left, lastVal)
	newRoot.Left = left

	return newRoot

}
