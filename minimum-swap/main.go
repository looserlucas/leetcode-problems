package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/*
 stay[i]: minimum swaps till i to make a[0]..a[i] strictly incremental while NOT swapping a[i], b[i]
 swap[i]: minimum swaps till i to make a[0]..a[i] strictly incremental while swapping a[i],b[i]

 this means stay[0] = 0, swap[0] = 1 (cuz we chose to swap a[i] b[i])

 the formula is as follow:
 - if nums1[i-1] >= nums2[i] || nums2[i-1] >= nums1[i] then the manipulation at index i should be the same as the one at index i-1 (if we swap a[i-1], b[i-1] then we should swap a[i], b[i] too, or the other way around)
 -> this means swap[i] = swap[i-1] + 1, stay[i-1] = stay[i]

 - if nums1[i-1] >= nums1[i] || nums2[i-1] >= nums2[i] then the opposite case of the former case should happen, if we swap a[i-1] b[i-1], we should keep a[i] b[i] untouched.

 - if else, get the minimum of stay[i-1] and swap[i-1], then stay[i] = min, swap[i] = min + 1
*/
func minSwap(nums1 []int, nums2 []int) int {
	stay := 0
	swap := 1
	for i := 1; i < len(nums1); i++ {
		if nums1[i-1] >= nums2[i] || nums2[i-1] >= nums1[i] {
			swap++
		} else if nums1[i-1] >= nums1[i] || nums2[i-1] >= nums2[i] {
			temp := stay
			stay = swap
			swap = temp + 1
		} else {
			temp := min(stay, swap)
			stay = temp
			swap = temp + 1
		}
	}

	return min(stay, swap)
}

func main() {
	nums1 := []int{1, 3, 5, 4}
	nums2 := []int{1, 2, 3, 7}
	fmt.Println(minSwap(nums1, nums2))
}
