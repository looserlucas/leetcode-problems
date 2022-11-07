package main

import (
	"fmt"
	"sort"
)

func IsGreater(a, b string) bool {
	if len(a) > len(b) {
		return true
	}
	if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			return true
		} else {
			return false
		}
	}

	return false
}

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		return IsGreater(nums[i], nums[j])
	})
	return nums[k-1]
}

func main() {
	fmt.Println(kthLargestNumber([]string{"3", "6", "7", "10"}, 4))
	fmt.Println(kthLargestNumber([]string{"2", "21", "12", "1"}, 3))
	fmt.Println(kthLargestNumber([]string{"0", "0"}, 2))
}
