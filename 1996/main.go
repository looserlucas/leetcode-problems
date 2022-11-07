package main

import (
	"fmt"
	"sort"
)

type TwoDSlice [][]int

func numberOfWeakCharacters(a [][]int) int {
	sort.SliceStable(a, func(i, j int) bool {
		if a[i][0] != a[j][0] {
			return a[i][0] > a[j][0]
		} else {
			return a[i][1] < a[j][1]
		}
	})
	var min = -1
	var res int = 0
	for i := 0; i < len(a); i++ {
		if min > a[i][1] {
			res++
		} else {
			min = a[i][1]
		}
	}

	fmt.Println(a)
	fmt.Println(res)
	return res
}

func main() {
	/*
		numberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}})
		numberOfWeakCharacters([][]int{{2, 2}, {3, 3}})
		numberOfWeakCharacters([][]int{{1, 5}, {10, 4}, {4, 3}})
	*/
	numberOfWeakCharacters([][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}})
	/*
		numberOfWeakCharacters([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}})
	*/
}
