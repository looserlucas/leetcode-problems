package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count map[string]int
var res []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = nil
	count = make(map[string]int)
	PostOrderTraverse(root)
	return res
}

func PostOrderTraverse(now *TreeNode) string {
	if now == nil {
		return "#"
	}
	serial := strconv.Itoa(now.Val) + "," + PostOrderTraverse(now.Left) + "," + PostOrderTraverse(now.Right)
	count[serial]++
	if count[serial] == 2 {
		res = append(res, now)
	}
	return serial
}

func main() {
	fmt.Println("vim-go")
}
