package main

import "fmt"

type stack []int

func (s *stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	x := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return x
}

func (s *stack) Push(x int) {
	(*s) = append(*s, x)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestValidParentheses(s string) int {
	var res = 0
	var st stack
	var dp = make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		fmt.Println(dp)
		if s[i] == '(' {
			st.Push(i)
		} else {
			if len(st) == 0 {
				continue
			}
			x := st.Pop()
			if x-1 >= 0 {
				dp[i] = dp[i-1] + dp[x-1] + 2
			} else {
				dp[i] = dp[i-1] + 2
			}
			res = max(res, dp[i])
		}
	}
	return res
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()()(())"))
	fmt.Println(longestValidParentheses("()(())(("))
	fmt.Println(longestValidParentheses(""))
}
