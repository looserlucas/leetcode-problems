package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2))
	}
	var res = 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if i-1 < 0 || j-1 < 0 || nums1[i-1] != nums2[j-1] {
					dp[i][j] = 1
				} else {
					dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
				}
				res = max(dp[i][j], res)
			}
		}
	}
	fmt.Println(dp)
	return res
}

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
