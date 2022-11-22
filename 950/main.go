package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool {
		return deck[i] > deck[j]
	})

	var res []int
	for i := 0; i < len(deck); i++ {
		if len(res) == 0 {
			res = append(res, deck[i])
			continue
		}
		x := res[len(res)-1]
		res = res[:len(res)-1]
		res = append([]int{deck[i], x}, res...)
	}
	return res
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}
