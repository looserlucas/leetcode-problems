package main

import "fmt"

type stack []int

func (s *stack) Push(value int) {
	(*s) = append(*s, value)
}

func (s *stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func totalSteps(a []int) int {
	var s stack
	var res int
	var dp = make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		for len(s) > 0 && a[i] > a[s[len(s)-1]] {
			dp[i] = max(dp[i]+1, dp[s[len(s)-1]])
			res = max(res, dp[i])
			s.Pop()
		}
		s.Push(i)
	}
	return res
}

func main() {
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	//fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	//fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
}
