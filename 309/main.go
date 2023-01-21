package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 0 -> i
//
func maxProfit(a []int) int {
	var s0 []int = make([]int, len(a))
	var s1 []int = make([]int, len(a))
	var s2 []int = make([]int, len(a))
	s0[0] = 0
	s1[0] = -a[0]
	s2[0] = math.MinInt
	for i := 1; i < len(a); i++ {
		s0[i] = max(s2[i-1], s0[i-1])
		s1[i] = max(s0[i-1]-a[i], s1[i-1])
		s2[i] = s1[i-1] + a[i]
	}
	var res int
	for i := 0; i < len(a); i++ {
		res = max(res, s0[i])
		res = max(res, s2[i])
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
