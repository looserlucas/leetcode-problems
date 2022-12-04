package main

import (
	"fmt"
	"math"
)

func minimumCardPickup(cards []int) int {
	var seen = make(map[int]bool)
	j := 0
	min := math.MaxInt
	for i := 0; i < len(cards); i++ {
		seen[cards[i]] = false
	}
	for i := 0; i < len(cards); i++ {
		if seen[cards[i]] {
			for j < i && seen[cards[i]] {
				if i-j+1 < min {
					min = i - j + 1
				}
				seen[cards[j]] = false
				j++
			}
		}
		seen[cards[i]] = true
	}
	if min == math.MaxInt {
		return -1
	} else {
		return min
	}
}
func main() {
	fmt.Println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
}
