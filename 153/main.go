package main

import "fmt"

func findMin(a []int) int {
	l := 0
	r := len(a) - 1
	for l < r {
		m := (l + r) / 2
		if a[m] < a[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return a[l]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
