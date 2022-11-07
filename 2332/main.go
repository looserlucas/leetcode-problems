package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func latestTimeCatchTheBus(b []int, p []int, c int) int {
	if len(b) == 0 {
		return 0
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	if len(p) == 0 {
		return b[len(b)-1]
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})
	var res int = min(b[len(b)-1], p[0]-1)
	j := 0
	for i := 0; i < len(b); i++ {
		var count int = c
		var lastExist bool = false
		for j < len(p) && count > 0 && p[j] <= b[i] {
			if j > 0 && p[j]-1 != p[j-1] {
				res = max(res, p[j]-1)
			}
			if p[j] == b[i] {
				lastExist = true
			}
			count--
			j++
		}
		if count > 0 && !lastExist {
			res = max(res, b[i])
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2)
	latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2)
	latestTimeCatchTheBus([]int{20, 30, 10}, []int{10, 29}, 2)
	latestTimeCatchTheBus([]int{20, 30, 10}, []int{7, 8, 9, 10, 11, 12, 13}, 2)
	latestTimeCatchTheBus([]int{2}, []int{4, 5}, 2)
	latestTimeCatchTheBus([]int{5, 15}, []int{11, 12, 13, 14, 15}, 2)
}
