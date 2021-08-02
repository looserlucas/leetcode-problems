package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func minimumMountainRemovals(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dd1 = make([]int, len(nums))
	var dd2 = make([]int, len(nums))
	dd1[0] = 1
	dd2[len(nums)-1] = 1
	// can turn this O(n^2) look up to O(nlogn) look up with Balanced Binary Search Tree (AVL tree, red-and-black tree, ..)
	for i := 1; i < len(nums); i++ {
		dd1[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dd1[i] = max(dd1[i], dd1[j]+1)
			}
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		dd2[i] = 1
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[i] {
				dd2[i] = max(dd2[i], dd2[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		if dd1[i] == 1 || dd2[i] == 1 {
			continue
		} else {
			res = max(res, dd1[i]+dd2[i]-1)
		}
	}

	return len(nums) - res
}

func main() {
	nums := []int{4, 3, 2, 1, 1, 2, 3, 1}
	fmt.Println(minimumMountainRemovals(nums))
}
