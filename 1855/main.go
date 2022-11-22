package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxDistance(nums1 []int, nums2 []int) int {
	i := 0
	var res = 0
	for j := 0; j < len(nums2); j++ {
		for i < j && nums1[i] > nums2[j] {
			i++
		}
		if i == len(nums1) {
			return res
		}
		res = max(res, j-i)
	}
	return res
}

func main() {
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{55, 30, 5, 4, 2}))
	fmt.Println(maxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	fmt.Println(maxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}))
}
