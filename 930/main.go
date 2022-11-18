package main

import "fmt"

func numSubarraysWithSumHashMap(nums []int, goal int) int {
	// using prefix sum, idea is with every j, find how many i (i<j) with sum(0:i) + goal = sum(0:j)
	// this means sum(i:j) = goal
	// keep track of count in a map
	var seen = make(map[int]int)
	var pre, res int
	seen[0]++
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		res += seen[pre-goal]
		seen[pre]++
	}

	return res
}

func numSubarraysWithSumSlidingWindow(nums []int, goal int) int {
	return atMost(nums, goal) - atMost(nums, goal-1)
}

// atMost function is to find how many subarrays are there with sum (at most) == t
// we use a sliding window with t as "window size". j is the end of the window, i is the start of the window
// at every j, we store res += j-i+1. This means every valid subarrays start with i and end with j is included in res
func atMost(nums []int, t int) int {
	if t < 0 {
		return 0
	}
	/*
		i := 0
		res := 0
		for j := 0; j < len(nums); j++ {
			t -= nums[j]
			for t < 0 && i < len(nums) {
				t += nums[i]
				i++
			}
			res += j - i + 1
		}
	*/
	if len(nums) == 0 {
		return 0
	}
	i := 0
	res := 0
	var sum = 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum > t && i < len(nums) {
			sum -= nums[i]
			i++
		}
		res += j - i + 1
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
