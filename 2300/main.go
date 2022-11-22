package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, p []int, success int64) []int {
	var si [][]int
	for i := 0; i < len(spells); i++ {
		si = append(si, []int{spells[i], i})
	}
	sort.SliceStable(si, func(i, j int) bool {
		return si[i][0] < si[j][0]
	})

	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})

	var res = make([]int, len(spells))
	j := 0
	for i := 0; i < len(si); i++ {
		for j < len(p) && int64(si[i][0])*int64(p[j]) >= success {
			j++
		}
		res[si[i][1]] = j
	}
	return res
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 1))
}
