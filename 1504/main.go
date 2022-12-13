package main

import (
	"fmt"
)

func numSubmat(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	var res int
	var c []int = make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				c[j] = 0
			} else {
				c[j] += 1
			}
		}
		res += helper(c)
	}
	return res
}

func helper(a []int) int {
	var s []int
	var sum []int = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		for len(s) > 0 && a[s[len(s)-1]] >= a[i] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			preIndex := s[len(s)-1]
			sum[i] = sum[preIndex]
			sum[i] += a[i] * (i - preIndex)
		} else {
			sum[i] = a[i] * (i + 1)
		}
		s = append(s, i)
	}

	var res int
	for _, s := range sum {
		res += s
	}

	return res
}
func main() {
	fmt.Println(numSubmat([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))
}
