package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	lose := make(map[int]int)
	var res [][]int
	for i := 0; i < len(matches); i++ {
		lose[matches[i][1]]++
		_, ok := lose[matches[i][0]]
		if !ok {
			lose[matches[i][0]] = 0
		}
	}
	var winner, second []int
	for k, v := range lose {
		if v == 0 {
			winner = append(winner, k)
		} else if v == 1 {
			second = append(second, k)
		}
	}
	sort.Slice(winner, func(i, j int) bool {
		return winner[i] < winner[j]
	})
	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})
	res = append(res, winner)
	res = append(res, second)
	return res
}

func main() {
	fmt.Println("vim-go")
}
