package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res, total int

const MODN = int(1e9 + 7)

func maxProduct(root *TreeNode) int {
	total = dfs(root)
	res = 0
	dfs(root)
	return res % MODN
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs(now *TreeNode) int {
	if now == nil {
		return 0
	}
	sum := now.Val + dfs(now.Left) + dfs(now.Right)
	res = max(res, sum*(total-sum))
	return sum
}

func main() {
	fmt.Println("vim-go")
}
