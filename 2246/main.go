package main

import (
	"container/list"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPath(parent []int, s string) int {
	cnt := make([]int, len(parent))
	res := 1
	top1 := make([]int, len(parent))
	top2 := make([]int, len(parent))
	for i := 0; i < len(parent); i++ {
		if i != 0 {
			cnt[parent[i]]++
		}
		top1[i] = 1
		top2[i] = 1
	}
	var q = list.New()
	for i := 1; i < len(parent); i++ {
		if cnt[i] == 0 {
			q.PushBack(i)
		}
	}

	for q.Len() > 0 && q.Front().Value.(int) != 0 {
		i := q.Front().Value.(int)
		p := parent[i]
		q.Remove(q.Front())
		var len int = 1
		if s[i] != s[p] {
			len = top1[i] + 1
		}
		if top1[p] <= len {
			top2[p] = top1[p]
			top1[p] = len
		} else {
			top2[p] = max(top2[p], len)
		}

		cnt[p]--
		if cnt[p] == 0 {
			q.PushBack(p)
			res = max(res, top1[p]+top2[p]-1)
		}
	}

	return res
}

func main() {
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}
