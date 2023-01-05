package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func merge(a [][]int) [][]int {
	sort.SliceStable(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	g := 0
	for i := 0; i < len(a); i++ {
		g = max(g, a[i][1])
	}
	a = append(a, []int{g + 1, g + 1})
	var res [][]int
	end := a[0][1]
	start := a[0][0]
	for i := 1; i < len(a); i++ {
		if a[i][0] <= end {
			start = min(start, a[i][0])
			end = max(end, a[i][1])
		} else {
			res = append(res, []int{start, end})
			start = a[i][0]
			end = a[i][1]
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
