package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	var move [][]int = [][]int{{1, -1}, {1, 0}, {1, 1}}
	var dp [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			tmp[j] = math.MaxInt
		}
		dp[i] = tmp
	}

	for i := 0; i < m; i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			for _, k := range move {
				if j+k[1] >= 0 && j+k[1] < m {
					dp[i+k[0]][j+k[1]] = min(dp[i+k[0]][j+k[1]], matrix[i+k[0]][j+k[1]]+dp[i][j])
				}
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < m; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
