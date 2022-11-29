package main

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})
	i := 0
	j := len(tokens) - 1
	var res int
	for i < j {
		fmt.Println(i, j, power)
		if power >= tokens[i] {
			power -= tokens[i]
			res++
			i++
		} else {
			if res == 0 {
				return res
			}
			power += tokens[j]
			res--
			j--
		}
	}
	if power >= tokens[i] {
		res++
	}
	return res
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
