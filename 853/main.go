package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	var cow [][]float64 = make([][]float64, len(position))
	for i := 0; i < len(position); i++ {
		cow[i] = []float64{float64(position[i]), float64(target-position[i]) / float64(speed[i])}
	}
	sort.SliceStable(cow, func(i, j int) bool {
		return cow[i][0] > cow[j][0]
	})
	var res = 0
	var curMax float64 = 0
	for i := 0; i < len(cow); i++ {
		if cow[i][1] > curMax {
			curMax = cow[i][1]
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
