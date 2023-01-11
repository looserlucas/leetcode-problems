package main

import "fmt"

func minimizeArrayValue(nums []int) int {
	l := 0
	r := int(1e9)
	for l < r {
		m := (l + r) / 2
		if ok(nums, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func ok(a []int, m int) bool {
	space := 0
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			needed := a[i] - m
			if needed > space {
				return false
			}
			space -= needed
		} else {
			space += (m - a[i])
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
