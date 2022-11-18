package main

import "fmt"

func maximumTop(nums []int, k int) int {
	if k == 0 {
		return nums[0]
	}
	if len(nums) == 1 {
		if k%2 != 0 {
			return -1
		} else {
			return nums[0]
		}
	}
	var leftMax = make([]int, len(nums))
	var max = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		leftMax[i] = max
	}
	if k > len(nums) {
		return leftMax[len(nums)-1]
	} else {
		if k == len(nums) {
			return leftMax[k-2]
		}
		if k == 1 {
			return nums[1]
		}
		if leftMax[k-2] > nums[k] {
			return leftMax[k-2]
		} else {
			return nums[k]
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
