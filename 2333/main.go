package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// This solution is insipred by Bucket Sort.
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	var diff []int
	max := -1
	for i := 0; i < len(nums1); i++ {
		diff = append(diff, abs(nums1[i]-nums2[i]))
		if diff[i] > max {
			max = diff[i]
		}
	}

	bucket := make([]int, max+1)
	// Put all the diff in its bucket
	for i := 0; i < len(diff); i++ {
		bucket[diff[i]]++
	}

	k := k1 + k2
	for i := max; i > 0; i-- {
		if k == 0 {
			break
		}
		if bucket[i] > 0 {
			// we will try to decrease the value of diff in the current bucket by 1
			// this mean we also move the diff to bucket[i-1]
			minus := min(bucket[i], k)
			bucket[i] -= minus
			bucket[i-1] += minus
			k -= minus
		}
	}

	var res int64
	for i := 1; i <= max; i++ {
		res += int64(bucket[i]) * int64(i) * int64(i)
	}
	return res
}

func main() {
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
}
