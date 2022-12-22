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

var sum map[int]int

func maxLevelSum(root *TreeNode) int {
	res := 0
	maxSum := math.MinInt
	sum = make(map[int]int)
	dfs(root, 0)
	for k, v := range sum {
		if v > maxSum {
			maxSum = v
			res = k
		} else if v == maxSum {
			if k < res {
				res = k
			}
		}
	}
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	sum[level] += root.Val
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func main() {
	fmt.Println("vim-go")
}
