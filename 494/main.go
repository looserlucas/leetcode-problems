package main

import (
	"fmt"
	"math"
)

var dp [][]int

func findTargetSumWays(nums []int, target int) int {
	var total = 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	dp = nil
	dp = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var tmp = make([]int, total*2+1)
		for j := 0; j < total*2+1; j++ {
			tmp[j] = math.MinInt
		}
		dp[i] = tmp
	}
	return dfs(nums, total, 0, 0, target)
}

func dfs(nums []int, total int, i int, sum, target int) int {
	if i == len(nums) {
		if sum == target {
			return 1
		} else {
			return 0
		}
	}
	if dp[i][sum+total] != math.MinInt {
		return dp[i][sum+total]
	}
	dp[i][sum+total] = dfs(nums, total, i+1, sum+nums[i], target) + dfs(nums, total, i+1, sum-nums[i], target)
	return dp[i][sum+total]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
