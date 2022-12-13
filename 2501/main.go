package main

import (
	"fmt"
	"sort"
)

func longestSquareStreak(nums []int) int {
	var dp = make(map[uint64]int)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res int = 0
	for i := 0; i < len(nums); i++ {
		_, ok := dp[uint64(nums[i])]
		if !ok {
			dp[uint64(nums[i])] = 1
		}
		if dp[uint64(nums[i])] > res {
			res = dp[uint64(nums[i])]
		}
		dp[uint64(nums[i])*uint64(nums[i])] = dp[uint64(nums[i])] + 1
	}
	if res == 1 {
		return -1
	} else {
		return res
	}
}

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2}))
}
