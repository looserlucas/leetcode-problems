package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minPairSum(nums []int) int {
	var res = 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := 0
	j := len(nums) - 1
	for i < j {
		res = max(res, nums[i]+nums[j])
		i++
		j--
	}

	return res
}
func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
}
