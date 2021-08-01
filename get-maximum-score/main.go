package main

import "fmt"

const MOD = 1e9 + 7

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// use 2 pointer, travel both on nums1 and nums2, when it s switchable (nums[i] = nums[j]), get the max point between then 2 value, then reset
// NOTE: care for not finishing one out of 2 array
func maxSum(nums1 []int, nums2 []int) int {
	i := 0
	j := 0
	sb := 0
	sa := 0
	n := len(nums1)
	m := len(nums2)
	var res int
	for i < n && j < m {
		// if can travel on nums1, keep on goin
		if nums1[i] < nums2[j] {
			sa += nums1[i]
			i++
			// if can travel on nums2, keep on goin
		} else if nums1[i] > nums2[j] {
			sb += nums2[j]
			j++
			// switch point, get max value
		} else if nums1[i] == nums2[j] {
			res += max(sa, sb) + nums1[i]
			sa = 0
			sb = 0
			i++
			j++
		}
	}

	// if one of the array is not finished
	for i < n {
		sa += nums1[i]
		i++
	}
	for j < m {
		sb += nums2[j]
		j++
	}
	res += max(sa, sb)

	return res % MOD
}

func main() {
	nums1 := []int{2, 4, 5, 8, 10}
	nums2 := []int{4, 6, 8, 9}
	fmt.Println(maxSum(nums1, nums2))
}
