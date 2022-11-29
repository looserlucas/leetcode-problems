package main

import (
	"fmt"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// jBad is the index i where nums[i] > maxK || nums[i] < minK
// jMin is the index i where nums[i] = minK
// jMax is the index i where nums[i] = maxK
// Subarrays end at i are only valid to the condition when
// j...jBad...jMin...Jmax...i or
// j...jBad...jMax...Jmin...i
// the idea that we can pick the start of subarray from [jBad+1, min(jMin, jMax)], end of the array
// is i, so we have min(jMin, jMax) - jBad subarrays
func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	var jMin, jMax, jBad = -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			jBad = i
		}
		if nums[i] == minK {
			jMin = i
		}
		if nums[i] == maxK {
			jMax = i
		}
		res += max(0, int64(min(jMax, jMin))-int64(jBad))
	}

	return res
}

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
}
