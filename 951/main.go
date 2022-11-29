package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// other solution
/*
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
        return false
    }
    if root1.Val != root2.Val {
        return false
    }

    return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) ||
        flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
}
*/

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var parent1 = make([]int, 100)
	var parent2 = make([]int, 100)
	for i := 0; i < len(parent1); i++ {
		parent1[i] = -2
		parent2[i] = -2
	}
	dfs(root1, -1, &parent1)
	dfs(root2, -1, &parent2)
	for i := 0; i < len(parent1); i++ {
		if parent1[i] != parent2[i] {
			return false
		}
	}
	return true
}

func dfs(now *TreeNode, parentValue int, parent *[]int) {
	if now == nil {
		return
	}
	(*parent)[now.Val] = parentValue
	dfs(now.Left, now.Val, parent)
	dfs(now.Right, now.Val, parent)
}

func main() {
	fmt.Println("vim-go")
}
