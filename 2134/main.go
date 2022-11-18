package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func minSwaps(nums []int) int {
	var count1 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count1++
		}
	}

	a := append(nums, nums...)
	var onesInWindow, tmp = 0, 0
	for i := 0; i < len(a); i++ {
		if i >= count1 && a[i-count1] == 1 {
			tmp--
		}
		if a[i] == 1 {
			tmp++
		}
		onesInWindow = max(onesInWindow, tmp)
	}
	return count1 - onesInWindow
}
func main() {
	fmt.Println(minSwaps([]int{1, 2, 3}))
}
