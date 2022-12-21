package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	res, _ := dfs(root, 0)
	return res.Val
}

func dfs(root *TreeNode, level int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	left, leftLevel := dfs(root.Left, level+1)
	right, rightLevel := dfs(root.Right, level+1)
	if left == nil && right == nil {
		return root, level
	}
	if left == nil && right != nil {
		return right, rightLevel
	}
	if right == nil && left != nil {
		return left, leftLevel
	}
	if leftLevel >= rightLevel {
		return left, leftLevel
	} else {
		return right, rightLevel
	}
}
func main() {
	fmt.Println("vim-go")
}
