package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	//return atMost(nums, k) - atMost(nums, k-1)
	return atMost(nums, k)
}

func atMost(nums []int, t int) int {
	if t < 0 {
		return 0
	}
	var seen = make(map[int]int)
	i := 0
	res := 0
	for j := 0; j < len(nums); j++ {
		seen[nums[j]]++
		for i < len(nums) && len(seen) >= t {
			seen[nums[i]]--
			if seen[nums[i]] == 0 {
				delete(seen, nums[i])
			}
			i++
		}
		fmt.Println(i, j)
		res += i
	}
	return res
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	//fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}
