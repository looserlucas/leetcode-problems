package main

import (
	"fmt"
)

func brokenCalc(startValue int, target int) int {
	var res int
	for target > startValue {
		if target%2 == 0 {
			target = target / 2
		} else {
			target = target + 1
		}
		res++
	}
	res += startValue - target
	return res
}

func main() {
	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))
	fmt.Println(brokenCalc(3, 10))
}
