package main

import "fmt"

var inIndex map[int]int
var postIndex int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inIndex = make(map[int]int)
	for i, v := range inorder {
		inIndex[v] = i
	}
	postIndex = len(postorder) - 1

	return helper(postorder, 0, len(postorder)-1)
}

func helper(postorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	rootValue := postorder[postIndex]
	postIndex--
	root := &TreeNode{
		Val: rootValue,
	}
	root.Right = helper(postorder, inIndex[rootValue]+1, right)
	root.Left = helper(postorder, left, inIndex[rootValue]-1)
	return root
}

func main() {
	fmt.Println("vim-go")
}
