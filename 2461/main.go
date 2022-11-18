package main

import "fmt"

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximumSubarraySum(nums []int, k int) int64 {
	seen := make(map[int]int)
	var sum, res int64 = 0, 0
	i := 0
	for i < k && i < len(nums) {
		seen[nums[i]]++
		sum += int64(nums[i])
		i++
	}
	if len(seen) == k {
		res = sum
	}
	for i < len(nums) {
		seen[nums[i]]++
		seen[nums[i-k]]--
		if seen[nums[i-k]] == 0 {
			delete(seen, nums[i-k])
		}

		sum += int64(nums[i])
		sum -= int64(nums[i-k])
		if len(seen) == k {
			res = max(res, sum)
		}
		i++
	}

	return res
}

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	fmt.Println(maximumSubarraySum([]int{9, 9, 9}, 3))
}
