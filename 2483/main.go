package main

import (
	"fmt"
	"math"
)

func bestClosingTime(customers string) int {
	var countY, countN int
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			countY++
		} else {
			countN++
		}
	}
	var res int = -1
	var min int = math.MaxInt
	var nBefore, yBefore int = 0, 0
	for i := 0; i < len(customers); i++ {
		tmp := nBefore + countY - yBefore
		if tmp < min {
			min = tmp
			res = i
		}
		if customers[i] == 'N' {
			nBefore++
		} else {
			yBefore++
		}
	}

	if countN < min {
		return len(customers)
	}
	return res
}

func main() {
	fmt.Println(bestClosingTime("YYNY"))
	fmt.Println(bestClosingTime("NNNNN"))
	fmt.Println(bestClosingTime("YYYY"))
}
