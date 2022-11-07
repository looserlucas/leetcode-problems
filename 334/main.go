package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	var leftMin = make([]int, len(nums))
	var rightMax = make([]int, len(nums))
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		leftMin[i] = min
	}

	max := -math.MaxInt
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > max {
			max = nums[i]
		}
		rightMax[i] = max
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
