package main

import (
	"fmt"
	"math"
)

var mid, left, right []int

func findKthLargest(nums []int, k int) int {
	mid, left, right = nil, nil, nil
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	if len(nums) == 0 {
		return math.MinInt
	}
	pivot := nums[len(nums)-1]

	mid, left, right = nil, nil, nil
	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			left = append(left, nums[i])
		} else if nums[i] < pivot {
			right = append(right, nums[i])
		} else if nums[i] == pivot {
			mid = append(mid, nums[i])
		}
	}
	if len(left) >= k {
		return quickSelect(left, k)
	} else if len(left)+len(mid) < k {
		return quickSelect(right, k-len(left)-len(mid))
	} else {
		return pivot
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{2, 1}, 1))
}
