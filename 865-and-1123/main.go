package main

import "fmt"

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res, _ := dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) (*TreeNode, int) {
	if root == nil {
		return nil, -1
	}
	rootLeft, leftLevel := dfs(root.Left, level+1)
	rootRight, rightLevel := dfs(root.Right, level+1)
	if rootLeft == nil && rootRight != nil {
		return rootRight, rightLevel
	}
	if rootRight == nil && rootLeft != nil {
		return rootLeft, leftLevel
	}
	if rootRight == nil && rootLeft == nil {
		return root, level
	}
	if leftLevel == rightLevel {
		return root, leftLevel
	} else if leftLevel > rightLevel {
		return rootLeft, leftLevel
	} else {
		return rootRight, rightLevel
	}
}
func main() {
	fmt.Println("vim-go")
}
