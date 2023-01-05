package main

import (
	"fmt"
	"sort"
)

// sort by a[i][1]
// greedily shoot at a[i][1] position and see how many balloon we can take out with it
func findMinArrowShots(a [][]int) int {
	sort.SliceStable(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})
	curArr := a[0][1]
	res := 1
	for i := 0; i < len(a); i++ {
		if curArr >= a[i][0] {
			continue
		}
		curArr = a[i][1]
		res++
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
