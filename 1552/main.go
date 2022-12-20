package main

import (
	"fmt"
	"sort"
)

func check(a []int, m, dist int) bool {
	last := a[0]
	m--
	for i := 1; i < len(a); i++ {
		if a[i]-last >= dist {
			last = a[i]
			m--
			if m == 0 {
				break
			}
		}
	}
	if m == 0 {
		return true
	} else {
		return false
	}
}

func maxDistance(a []int, m int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	lo := 0
	hi := a[len(a)-1]
	for lo < hi {
		mid := (lo + hi) / 2
		if check(a, m, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
	fmt.Println(maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4))
}
