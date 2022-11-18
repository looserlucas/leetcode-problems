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

func movesToMakeZigzag(nums []int) int {
	res := make([]int, 2)
	n := len(nums)
	left, right := 0, 0
	for i := 0; i < n; i++ {
		if i == 0 {
			left = 1001
		} else {
			left = nums[i-1]
		}
		if i == n-1 {
			right = 1001
		} else {
			right = nums[i+1]
		}
		res[i%2] += max(0, nums[i]-min(left, right)+1)
	}
	return min(res[0], res[1])
}
func main() {
	fmt.Println("vim-go")
}
