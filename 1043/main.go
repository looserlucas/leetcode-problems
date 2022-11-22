package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxSumAfterPartitioning(arr []int, k int) int {
	var dp = make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		var best int = -1
		var curMax int = -1
		for j := 1; j <= k && i-j >= 0; j++ {
			curMax = max(curMax, arr[i-j])
			best = max(best, dp[i-j]+j*curMax)
		}
		dp[i] = best
	}
	return dp[len(arr)]
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
}
