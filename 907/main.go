package main

import (
	"fmt"
)

var MODN = int(1e9 + 7)

func sumSubarrayMins(arr []int) int {
	var res int = 0
	var s []int
	for i := 0; i <= len(arr); i++ {
		for len(s) > 0 && (i == len(arr) || arr[s[len(s)-1]] >= arr[i]) {
			x := s[len(s)-1]
			s = s[:len(s)-1]
			y := -1
			if len(s) > 0 {
				y = s[len(s)-1]
			}
			res = (res + arr[x]*(x-y)*(i-x)) % MODN
		}

		s = append(s, i)
	}

	return res
}

func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}
