package main

import "fmt"

func subArrayRanges(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums); i++ {
		curMax := nums[i]
		curMin := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > curMax {
				curMax = nums[j]
			}
			if nums[j] < curMin {
				curMin = nums[j]
			}
			res += int64(curMax) - int64(curMin)
		}
	}
	return res
}

func main() {
	fmt.Println(subArrayRanges([]int{1, 2, 3}))
	fmt.Println(subArrayRanges([]int{1, 3, 3}))
	fmt.Println(subArrayRanges([]int{4, -2, -3, 4, 1}))
}
