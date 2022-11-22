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
func numSquares(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(13))
}
