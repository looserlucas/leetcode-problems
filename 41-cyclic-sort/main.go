package main

import (
	"fmt"
)

func firstMissingPositive(a []int) int {
	a = append(a, 0)
	n := len(a)
	i := 0
	for i < n {
		if a[i] > 0 {
			if a[i]-1 < n && a[a[i]-1] != a[i] {
				a[a[i]-1], a[i] = a[i], a[a[i]-1]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	fmt.Println(a)
	var res = 1
	for i := 0; i < n; i++ {
		if a[i]-1 != i {
			res = i + 1
			break
		}
	}
	return res
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
