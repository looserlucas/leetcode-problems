package main

import (
	"fmt"
	"strconv"
)

func largestMultipleOfThree(digits []int) string {
	var memo [3][]string
	for i := 0; i < 3; i++ {
		tmp := make([]string, len(digits)+1)
		memo[i] = tmp
	}

	fmt.Println(dfsMemo(digits, 0, digits[0]%3, memo))
	for i := 0; i < 3; i++ {
		for j := 0; j < len(digits); j++ {
			fmt.Print(memo[i][j], " ")
		}
		fmt.Println()
	}
	return ""
}

func max(a, b string) string {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfsMemo(digits []int, index int, mod int, memo [3][]string) string {
	if index == len(digits)-1 {
		return strconv.Itoa(digits[index])
	}

	if memo[mod][index] != "" {
		return memo[mod][index]
	}

	c := strconv.Itoa(digits[index])

	memo[mod][index] = max(dfsMemo(digits, index+1, mod, memo), c+dfsMemo(digits, index+1, (mod+digits[index]%3)%3, memo))

	return memo[mod][index]
}

func main() {
	digits := []int{8, 1, 9}
	largestMultipleOfThree(digits)
}
