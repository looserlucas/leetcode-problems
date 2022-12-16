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
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

var dp [][]int

func minimizeTheDifference(mat [][]int, target int) int {
	n := len(mat)
	m := len(mat[0])
	var newMat [][]int
	for i := 0; i < n; i++ {
		var bucket = make([]bool, 71)
		for j := 0; j < m; j++ {
			bucket[mat[i][j]] = true
		}
		var tmp []int
		for j := 0; j <= 70; j++ {
			if bucket[j] {
				tmp = append(tmp, j)
			}
		}
		newMat = append(newMat, tmp)
	}

	dp = make([][]int, 71)
	for i := 0; i < len(dp); i++ {
		var tmp []int
		for j := 0; j <= 70*70; j++ {
			tmp = append(tmp, math.MaxInt)
		}
		dp[i] = tmp
	}

	return dfs(newMat, 0, 0, target)
}

func dfs(mat [][]int, i int, sum int, target int) int {
	if i == len(mat) {
		return abs(sum - target)
	}
	if dp[i][sum] == math.MaxInt {
		for j := 0; j < len(mat[i]); j++ {
			dp[i][sum] = min(dp[i][sum], dfs(mat, i+1, sum+mat[i][j], target))
			if dp[i][sum] == 0 || sum+mat[i][j] > target {
				break
			}
		}
	}
	return dp[i][sum]
}

func main() {
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
}
