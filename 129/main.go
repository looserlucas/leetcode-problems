package main

import "fmt"

var num []int

func sumNumbers(root *TreeNode) int {
	num = nil
	dfs(root, 1)
	var res int
	for _, n := range num {
		res += n
	}
	return res
}

func dfs(node *TreeNode, now int) {
	if node.Left == nil && node.Right == nil {
		num = append(num, now*10+node.Val)
		return
	}

	if node.Left != nil {
		dfs(node.Left, now*10+node.Val)
	}
	if node.Right != nil {
		dfs(node.Right, now*10+node.Val)
	}
}
func main() {
	fmt.Println("vim-go")
}
