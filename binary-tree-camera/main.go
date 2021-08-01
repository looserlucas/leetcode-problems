package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	notCovered        int = 0
	coveredWithoutCam int = 1
	coveredWithCam    int = 2
)

var res int

func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := dfs(root)
	if r == notCovered {
		res++
	}
	return res
}

func dfs(n *TreeNode) int {
	if n == nil {
		return coveredWithoutCam
	}

	left := dfs(n.Left)
	right := dfs(n.Right)

	if left == coveredWithoutCam && right == coveredWithoutCam {
		return notCovered
	} else if left == notCovered || right == notCovered {
		res++
		return coveredWithCam
	} else {
		return coveredWithoutCam
	}
}

func main() {
	fmt.Println("vim-go")
}
