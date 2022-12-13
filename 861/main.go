package main

import "fmt"

var pow2 []int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func matrixScore(a [][]int) int {
	n := len(a)
	m := len(a[0])
	pow2 = make([]int, m)
	pow2[0] = 1
	for i := 1; i < m; i++ {
		pow2[i] = pow2[i-1] * 2
	}
	count1 := 0
	for i := 0; i < n; i++ {
		if a[i][0] == 1 {
			count1++
		}
	}
	if count1 == n {
		return maxmize(a)
	} else {
		var b [][]int = make([][]int, n)
		for i := 0; i < n; i++ {
			tmp := make([]int, m)
			b[i] = tmp
			if a[i][0] == 1 {
				b[i][0] = 0
			} else {
				b[i][0] = 1
			}
			for j := 1; j < m; j++ {
				b[i][j] = a[i][j]
			}
		}

		for i := 0; i < n; i++ {
			if a[i][0] == 0 {
				for j := 0; j < m; j++ {
					if a[i][j] == 0 {
						a[i][j] = 1
					} else {
						a[i][j] = 0
					}
				}
			}
		}
		return max(maxmize(a), maxmize(b))
	}
}

func maxmize(a [][]int) int {
	n := len(a)
	m := len(a[0])
	for i := 0; i < n; i++ {
		if a[i][0] == 0 {
			for j := 0; j < m; j++ {
				if a[i][j] == 0 {
					a[i][j] = 1
				} else {
					a[i][j] = 0
				}
			}
		}
	}
	res := pow2[m-1] * n

	for j := 1; j < m; j++ {
		var count1, count0 int = 0, 0
		for i := 0; i < n; i++ {
			if a[i][j] == 0 {
				count0++
			} else {
				count1++
			}
		}
		res += (max(count1, count0) * pow2[m-1-j])
	}

	return res
}

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
}
