package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxEndHere = make([]int, len(nums))
	maxEndHere[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndHere[i] = max(nums[i], nums[i]+maxEndHere[i-1])
		res = max(res, maxEndHere[i])
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
