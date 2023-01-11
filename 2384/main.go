package main

import (
	"fmt"
	"strconv"
	"strings"
)

func largestPalindromic(num string) string {
	var count = make([]int, 10)
	for i := 0; i < len(num); i++ {
		count[int(num[i])-int('0')]++
	}
	fmt.Println(count)
	var res []int
	for i := 9; i >= 0; i-- {
		if len(res) == 0 && i == 0 {
			continue
		} else if count[i] > 1 {
			for j := 0; j < count[i]/2; j++ {
				res = append(res, i)
			}
			count[i] -= (count[i] / 2 * 2)
		}
	}
	fmt.Println(count)
	mid := false
	for i := 9; i >= 0; i-- {
		if count[i] >= 1 {
			res = append(res, i)
			mid = true
			break
		}
	}
	var r strings.Builder
	for i := 0; i < len(res); i++ {
		r.WriteString(strconv.Itoa(res[i]))
	}

	if mid {
		for i := len(res) - 2; i >= 0; i-- {
			r.WriteString(strconv.Itoa(res[i]))
		}
	} else {
		for i := len(res) - 1; i >= 0; i-- {
			r.WriteString(strconv.Itoa(res[i]))
		}
	}

	return r.String()
}
func main() {
	fmt.Println(largestPalindromic("444947137"))
	fmt.Println(largestPalindromic("0000"))
}
