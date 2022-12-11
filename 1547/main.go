package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findLengthOfShortestSubarray(a []int) int {
	res := 1
	for i := 1; i < len(a); i++ {
		if a[i] >= a[i-1] {
			res = max(res, i+1)
		} else {
			break
		}
	}

	var s []int
	s = append(s, len(a)-1)
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] <= a[i+1] {
			res = max(res, len(a)-i)
			s = append(s, i)
		} else {
			break
		}
	}

	last := 0
	for i := 0; i < len(a); i++ {
		if a[i] >= last && len(s) > 0 && i < s[len(s)-1] {
			for len(s) > 0 && a[i] > a[s[len(s)-1]] {
				s = s[0 : len(s)-1]
			}
			res = max(res, i+1+len(s))
			fmt.Println(res)
			last = a[i]
		} else {
			break
		}
	}

	return len(a) - res
}

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 3, 10, 1, 3, 3, 5}))
}
