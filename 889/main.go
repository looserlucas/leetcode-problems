package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preIndex int
var postIndex int

func constructFromPrePost(pre []int, post []int) *TreeNode {
	preIndex = 0
	postIndex = 0
	return helper(pre, post)
}

func helper(pre []int, post []int) *TreeNode {
	root := &TreeNode{
		Val: pre[preIndex],
	}
	preIndex++
	if root.Val != post[postIndex] {
		root.Left = helper(pre, post)
	}
	if root.Val != post[postIndex] {
		root.Right = helper(pre, post)
	}
	postIndex++
	return root
}

func main() {
	//fmt.Println(constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
	fmt.Println(constructFromPrePost([]int{1, 2, 4}, []int{2, 4, 1}))
}
