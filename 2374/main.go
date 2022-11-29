package main

import "fmt"

func edgeScore(edges []int) int {
	var count = make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		count[edges[i]] += i
	}
	var curMax, res = 0, 0
	for i := 0; i < len(count); i++ {
		if count[i] > curMax {
			curMax = count[i]
			res = i
		}
	}
	return res
}

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5}))
}
