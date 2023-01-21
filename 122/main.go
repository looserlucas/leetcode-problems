package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxProfit(a []int) int {
	s1 := make([]int, len(a))
	s2 := make([]int, len(a))
	s1[0] = 0
	s2[0] = -a[0]
	for i := 1; i < len(a); i++ {
		s1[i] = max(s1[i-1], s2[i-1]+a[i])
		s2[i] = max(s2[i-1], s1[i-1]-a[i])
	}
	fmt.Println(s1, s2)
	res := 0
	for i := 0; i < len(s1); i++ {
		res = max(res, s1[i])
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
