package main

import (
	"fmt"
	"sort"
)

// The idea is try with every i, get the closest a[j] (i<j<len(a)-1) where:
// a[j] > a[i]
// a[j] is as small as possible
// and then swap the two
func prevPermOpt1(a []int) []int {
	n := len(a) - 1
	var b [][]int
	b = append(b, []int{a[n], n})
	for i := n - 1; i >= 0; i-- {
		t := a[i]
		x := sort.Search(len(b), func(i int) bool {
			return b[i][0] < t
		})
		if x != len(b) {
			if b[x][0] < t {
				a[i], a[b[x][1]] = a[b[x][1]], a[i]
				break
			}
		}

		if b[len(b)-1][0] == a[i] {
			b = b[:len(b)-1]
		}
		b = append(b, []int{a[i], i})
	}
	return a
}

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
	fmt.Println(prevPermOpt1([]int{1, 9, 5, 4, 7}))
	fmt.Println(prevPermOpt1([]int{3, 1, 1, 3}))
}
