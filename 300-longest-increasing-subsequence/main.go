package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	res := 1
	var dd1 = make([]int, len(nums))
	dd1[0] = 1
	// can turn this O(n^2) look up to O(nlogn) look up with Balanced Binary Search Tree (AVL tree, red-and-black tree, ..)
	for i := 1; i < len(nums); i++ {
		dd1[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dd1[i] = max(dd1[i], dd1[j]+1)
				res = max(res, dd1[i])
			}
		}
	}
	return res
}
func main() {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums))
}
