package main

import (
	"fmt"
	"sort"
)

func helper(a []int, x, k int) bool {
	count := 1
	last := a[0]
	for i := 0; i < len(a); i++ {
		if count == k {
			break
		}
		if a[i]-last >= x {
			last = a[i]
			count++
		}
	}
	return count == k
}

func maximumTastiness(a []int, k int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	hi := int(1e9)
	lo := 0
	for lo < hi {
		mid := (hi + lo) / 2
		if helper(a, mid, k) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

func main() {
	fmt.Println("vim-go")
}
