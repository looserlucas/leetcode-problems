package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxAncestorDiff(root *TreeNode) int {
	res = math.MinInt
	dfs(root)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return math.MinInt, math.MaxInt
	}
	maxLeft, minLeft := dfs(node.Left)
	maxRight, minRight := dfs(node.Right)
	maxTotal := max(maxLeft, maxRight)
	minTotal := min(minLeft, minRight)
	if maxTotal != math.MinInt {
		res = max(res, abs(maxTotal-node.Val))
	}
	if minTotal != math.MaxInt {
		res = max(res, abs(node.Val-minTotal))
	}
	maxTotal = max(node.Val, maxTotal)
	minTotal = min(node.Val, minTotal)
	return maxTotal, minTotal
}

func main() {
	fmt.Println("vim-go")
}
