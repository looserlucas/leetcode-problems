package main

import (
	"fmt"
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {
	var nums2WithIndex [][]int = make([][]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		nums2WithIndex[i] = []int{nums2[i], i}
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.SliceStable(nums2WithIndex, func(i, j int) bool {
		return nums2WithIndex[i][0] < nums2WithIndex[j][0]
	})
	var res = make([]int, len(nums1))
	i := 0
	j := len(nums1) - 1
	k := len(nums2) - 1
	for i <= j {
		if nums1[j] > nums2WithIndex[k][0] {
			idx := nums2WithIndex[k][1]
			res[idx] = nums1[j]
			j--
			k--
		} else {
			idx := nums2WithIndex[k][1]
			res[idx] = nums1[i]
			i++
			k--
		}
	}
	return res
}

func main() {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
