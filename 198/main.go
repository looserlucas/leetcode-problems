package main

import "fmt"

var dp [][]int

func rob(nums []int) int {
	dp = nil
	for i := 0; i < len(nums); i++ {
		var tmp []int
		for j := 0; j <= 1; j++ {
			tmp = append(tmp, -1)
		}
		dp = append(dp, tmp)
	}
	return dfs(nums, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs(nums []int, i int, rob int) int {
	if i == len(nums) {
		return 0
	}
	if dp[i][rob] != -1 {
		return dp[i][rob]
	}
	if rob == 1 {
		dp[i][rob] = dfs(nums, i+1, 0)
	} else {
		dp[i][rob] = max(dfs(nums, i+1, 1)+nums[i], dfs(nums, i+1, 0))
	}
	return dp[i][rob]
}

func main() {
	fmt.Println("vim-go")
}
