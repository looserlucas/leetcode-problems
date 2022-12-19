package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// the idea is to sort the array incrementally,
// then try to find the longest subarray [i,j] where sum(a[i]+...+a[j]) + k >= (j-i+1)*a[j]
func maxFrequency(a []int, k int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	j := 0
	sum := 0
	var res = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
		for j <= i && sum+k < (i-j+1)*a[i] {
			sum -= a[j]
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
