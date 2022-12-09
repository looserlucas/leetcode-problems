package main

import (
	"fmt"
)

func mostCompetitive(nums []int, k int) []int {
	var s []int
	for i := 0; i < len(nums); i++ {
		for len(s) > 0 && nums[i] < nums[s[len(s)-1]] && len(nums)-i+len(s) > k {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, i)
		}
	}
	var res = make([]int, k)
	for i := k - 1; i >= 0; i-- {
		top := s[len(s)-1]
		res[i] = nums[top]
		s = s[:len(s)-1]
	}

	return res
}

func main() {
	fmt.Println(mostCompetitive([]int{3, 5, 2, 6}, 2))
	fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
	fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 1))
}
