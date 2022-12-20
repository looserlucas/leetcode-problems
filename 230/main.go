package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int
var res int

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	inorder(root, k)
	return res
}

func inorder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inorder(root.Left, k)
	if count < k {
		res = root.Val
		count++
	}
	inorder(root.Right, k)
}
func main() {
	fmt.Println("vim-go")
}
