package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minOperations(nums []int, x int) int {
	var pre = make([]int, len(nums)+1)
	var seen = make(map[int]int)
	pre[0] = 0
	seen[0] = 0
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
		seen[pre[i+1]] = i + 1
	}
	var res int = math.MaxInt
	r := pre[len(nums)]
	for i := len(nums); i >= 0; i-- {
		v, ok := seen[x-(r-pre[i])]
		if !ok {
			continue
		} else {
			if v <= i {
				res = min(res, len(nums)-i+v)
			}
		}
	}
	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

func main() {
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}
