package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func longestIdealString(s string, k int) int {
	var dp = make([]int, 26)
	for i := 0; i < len(s); i++ {
		c := int(s[i]) - 'a'
		var start, end int
		if c-k >= 0 {
			start = c - k
		} else {
			start = 0
		}
		if c+k >= 25 {
			end = 25
		} else {
			end = c + k
		}
		now := -1
		for j := start; j <= end; j++ {
			now = max(now, dp[j]+1)
		}
		dp[c] = now
	}
	res := -1
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}
func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
}
