package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximumSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxEndHere = make([]int, len(nums))
	var maxStartHere = make([]int, len(nums))
	maxEndHere[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndHere[i] = max(nums[i], nums[i]+maxEndHere[i-1])
		res = max(res, maxEndHere[i])
	}

	maxStartHere[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		maxStartHere[i] = max(nums[i], nums[i]+maxStartHere[i+1])
	}

	for i := 1; i < len(nums)-1; i++ {
		res = max(res, maxEndHere[i-1]+maxStartHere[i+1])
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
