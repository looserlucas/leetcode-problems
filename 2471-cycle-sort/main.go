package main

import (
	"container/list"
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var idxMap map[int]int

func cycleSort(a []int) int {
	var swap = 0
	var b []int = make([]int, len(a))
	copy(b, a)
	n := len(a)
	for i := 0; i < n; i++ {
		idxMap[a[i]] = i
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			idx1 := idxMap[b[i]]
			idx2 := idxMap[a[i]]
			idxMap[a[i]], idxMap[b[i]] = idx1, idx2
			a[idx1] = a[i]
			a[idx2] = b[i]
			swap++
		}
	}
	return swap
}

func minimumOperations(root *TreeNode) int {
	idxMap = make(map[int]int)
	q := list.New()
	var res int
	q.PushBack(root)
	for q.Len() > 0 {
		l := q.Len()
		var a []int
		for i := 0; i < l; i++ {
			now := q.Front()
			v := now.Value.(*TreeNode)
			a = append(a, v.Val)
			if v.Left != nil {
				q.PushBack(v.Left)
			}
			if v.Right != nil {
				q.PushBack(v.Right)
			}
			q.Remove(now)
		}
		res += cycleSort(a)
	}
	return res
}

func main() {
	fmt.Println(cycleSort([]int{3, 2}))
	//fmt.Println("vim-go")
}
