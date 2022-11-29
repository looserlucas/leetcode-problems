package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(now *TreeNode, a *[]int) {
	if now == nil {
		return
	}
	dfs(now.Left, a)
	(*a) = append(*a, now.Val)
	dfs(now.Right, a)
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var res [][]int
	var a []int
	dfs(root, &a)
	for i := 0; i < len(queries); i++ {
		t := queries[i]
		x := sort.Search(len(a), func(i int) bool {
			return a[i] >= t
		})
		if x == len(a) {
			res = append(res, []int{a[len(a)-1], -1})
		} else {
			if x == 0 && a[x] > t {
				res = append(res, []int{-1, a[0]})
			} else {
				if a[x] == t {
					res = append(res, []int{t, t})
				} else {
					res = append(res, []int{a[x-1], a[x]})
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
