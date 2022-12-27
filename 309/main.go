package main

import "fmt"

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
	var dp []int = make([]int, len(a))
	for i := 1; i < len(a); i++ {
		for j := 0; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]+a[i]-a[j+1])
		}
	}
	return dp[len(a)-1]
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
