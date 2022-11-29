package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func maxSubarraySumCircular(nums []int) int {
	curMax := 0
	curMin := 0
	maxSum := nums[0]
	minSum := nums[0]
	total := 0
	for i := 0; i < len(nums); i++ {
		curMax = max(curMax+nums[i], nums[i])
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+nums[i], nums[i])
		minSum = min(minSum, curMin)
		total += nums[i]
	}
	if maxSum > 0 {
		return max(maxSum, total-minSum)
	} else {
		return maxSum
	}
}
func main() {
	fmt.Println("vim-go")
}
