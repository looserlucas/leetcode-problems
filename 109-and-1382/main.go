package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var a []int
	inorder(root, &a)
	return buildTree(a, 0, len(a)-1)
}

func inorder(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, a)
	(*a) = append(*a, root.Val)
	inorder(root.Right, a)
}

func buildTree(a []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := &TreeNode{
		Val: a[mid],
	}
	root.Left = buildTree(a, l, mid-1)
	root.Right = buildTree(a, mid+1, r)
	return root
}

func main() {
	fmt.Println("vim-go")
}
