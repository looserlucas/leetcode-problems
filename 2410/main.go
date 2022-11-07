package main

import (
	"fmt"
	"sort"
)

func matchPlayersAndTrainers(p []int, t []int) int {
	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})
	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})
	var res int = 0
	i := 0
	j := 0
	for i < len(p) && j < len(t) {
		if p[i] > t[j] {
			i++
			continue
		}
		res++
		j++
		i++
	}
	return res
}
func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
	fmt.Println(matchPlayersAndTrainers([]int{1, 1, 1}, []int{10}))
	fmt.Println(matchPlayersAndTrainers([]int{2}, []int{1}))
}
