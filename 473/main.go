package main

import (
	"fmt"
	"sort"
)

func makesquare(a []int) bool {
	if len(a) == 0 && len(a) < 4 {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	t := 0
	for i := 0; i < len(a); i++ {
		t += a[i]
	}
	if t%4 != 0 {
		return false
	}
	var e []int = make([]int, 4)
	return dfs(a, e, 0, t/4)
}

func dfs(a []int, e []int, index int, target int) bool {
	if index == len(a) {
		return e[0] == target && e[1] == target && e[2] == target && e[3] == target
	}

	for i := 0; i < 4; i++ {
		if e[i]+a[index] > target {
			continue
		}
		e[i] += a[index]
		if dfs(a, e, index+1, target) {
			return true
		}
		e[i] -= a[index]
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
