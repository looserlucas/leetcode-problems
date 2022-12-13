package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var res = make([]int, len(nums))
	nums = append(nums, nums...)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	var s []int
	for i := 0; i < len(nums); i++ {
		for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
			j := s[len(s)-1]
			if j < n {
				res[j] = nums[i]
			}
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}

	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
