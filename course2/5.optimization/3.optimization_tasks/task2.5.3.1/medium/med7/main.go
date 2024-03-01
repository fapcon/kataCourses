package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	newTree := balanceBST(tree)
	print(tree)
	fmt.Println()
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

func balanceBST(root *TreeNode) *TreeNode {
	var arr []int
	InOrder(root, &arr)
	result := createBalancedBST(arr, 0, len(arr)-1)
	return result
}

func InOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	InOrder(root.Right, arr)

}

func createBalancedBST(A []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := (l + r) / 2
	root := &TreeNode{Val: A[m]}
	root.Left = createBalancedBST(A, l, m-1)
	root.Right = createBalancedBST(A, m+1, r)
	return root
}
