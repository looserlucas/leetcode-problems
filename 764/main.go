package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	up := make([][]int, n)
	down := make([][]int, n)
	left := make([][]int, n)
	right := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp1 := make([]int, n)
		tmp2 := make([]int, n)
		tmp3 := make([]int, n)
		tmp4 := make([]int, n)
		up[i] = tmp1
		down[i] = tmp2
		left[i] = tmp3
		right[i] = tmp4
	}
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, 1)
		}
		a[i] = tmp
	}
	for _, m := range mines {
		a[m[0]][m[1]] = 0
	}
	for i := 0; i < n; i++ {
		if a[i][0] == 1 {
			left[i][0] = 1
		}
		for j := 1; j < n; j++ {
			if a[i][j] == 1 {
				left[i][j] = left[i][j-1] + 1
			}
		}
		if a[i][n-1] == 1 {
			right[i][n-1] = 1
		}
		for j := n - 2; j >= 0; j-- {
			if a[i][j] == 1 {
				right[i][j] = right[i][j+1] + 1
			}
		}
	}
	for j := 0; j < n; j++ {
		if a[0][j] == 1 {
			up[0][j] = 1
		}
		for i := 1; i < n; i++ {
			if a[i][j] == 1 {
				up[i][j] = up[i-1][j] + 1
			}
		}

		if a[n-1][j] == 1 {
			down[n-1][j] = 1
		}
		for i := n - 2; i >= 0; i-- {
			if a[i][j] == 1 {
				down[i][j] = down[i+1][j] + 1
			}
		}
	}
	var res = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				res = max(res, min(min(up[i][j]-1, down[i][j]-1), min(right[i][j]-1, left[i][j]-1))+1)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
}
