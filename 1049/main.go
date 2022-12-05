package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func lastStoneWeightII(stones []int) int {
	dp := make([]bool, 1501)
	dp[0] = true
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
		for j := min(1500, sum); j >= stones[i]; j-- {
			dp[j] = dp[j] || dp[j-stones[i]]
		}
	}
	for i := sum / 2; i >= 0; i-- {
		if dp[i] {
			return sum - i - i
		}
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
}
