package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// dp[i][j] = max profit up until price[j] using at most i transaction
// dp[i][j] = max(dp[i][j-1], price[j] + max(dp[i-1][j1] - price[j1])) (j1 0 -> j-1)

func maxProfit(k int, prices []int) int {
	var dp [101][1001]int
	if len(prices) == 0 {
		return 0
	}

	for k1 := 1; k1 <= k; k1++ {
		// local max is used to keep track the max(dp[k1-1][i] - price[i])
		var localMax = dp[k1-1][0] - prices[0]
		for i := 1; i < len(prices); i++ {
			dp[k1][i] = max(dp[k1][i-1], prices[i]+localMax)
			// update local max as we travel
			localMax = max(localMax, dp[k1-1][i]-prices[i])
		}
	}

	for i := 0; i <= k; i++ {
		for j := 0; j < len(prices); j++ {
			fmt.Print(" ")
			fmt.Print(dp[i][j])
		}
		fmt.Println()
	}

	return dp[k][len(prices)-1]
}

func main() {
	k := 2
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(k, prices))
}
