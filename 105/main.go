package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inIndex map[int]int
var preIndex int

func buildTree(preorder []int, inorder []int) *TreeNode {
	inIndex = make(map[int]int)
	for i, v := range inorder {
		inIndex[v] = i
	}
	preIndex = 0

	return helper(preorder, 0, len(preorder)-1)
}

func helper(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	rootValue := preorder[preIndex]
	preIndex++
	root := &TreeNode{
		Val: rootValue,
	}
	root.Left = helper(preorder, left, inIndex[rootValue]-1)
	root.Right = helper(preorder, inIndex[rootValue]+1, right)
	return root
}

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
