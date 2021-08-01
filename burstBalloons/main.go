package main

import "fmt"

// https://leetcode.com/problems/burst-balloons
// dp[i,j] is the maximum money can get from poping all baloon between i and j (except for i and j)
// => dp[i,j] = max(balloon[i] * balloon[k] * balloon[j] + dp[i,k] + dp[k,j]) (with k from i+1 -> j-1)
func maxCoins(iNums []int) int {
	nums := make([]int, len(iNums)+2)
	var n = 1
	for _, v := range iNums {
		if n > 0 {
			nums[n] = v
			n++
		}
	}

	nums[0] = 1
	nums[n] = 1
	n++

	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for k := 2; k < n; k++ {
		for left := 0; left < n-k; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				dp[left][right] = max(dp[left][right], nums[left]*nums[i]*nums[right]+dp[left][i]+dp[i][right])
			}
		}
	}
	return dp[0][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{1, 5}
	fmt.Println(maxCoins(nums))
}
