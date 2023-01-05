package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(a [][]int) int {
	sort.SliceStable(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})
	end := a[0][1]
	count := 1
	for i := 1; i < len(a); i++ {
		if a[i][0] >= end {
			end = a[i][1]
			count++
		}
	}

	return len(a) - count
}

func main() {
	fmt.Println("vim-go")
}
