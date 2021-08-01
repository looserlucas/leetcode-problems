package main

import (
	"fmt"
)

func compare(a, b int, min bool) int {
	if a < b {
		if min {
			return a
		} else {
			return b
		}
	} else {
		if min {
			return b
		} else {
			return a
		}
	}
}
func countDigitOne(n int) int {
	if n == 1 {
		return 1
	}
	var total int
	for i := 1; i < n; i *= 10 {
		div := i * 10
		total += (n/div)*i + compare(compare(n%div-i+1, 0, false), i, true)
	}
	return total
}

func main() {
	fmt.Println(countDigitOne(13))
}
