package main

import "fmt"

const MIN_INT = -int64(1e9)

func abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxPoints(points [][]int) int64 {
	n := len(points)
	m := len(points[0])
	var dp [][]int64
	for i := 0; i < n; i++ {
		if i == 0 {
			var tmp []int64
			for j := 0; j < m; j++ {
				tmp = append(tmp, int64(points[i][j]))
			}
			dp = append(dp, tmp)
		} else {
			var tmp []int64
			for j := 0; j < m; j++ {
				tmp = append(tmp, MIN_INT)
			}
			dp = append(dp, tmp)
		}
	}

	for i := 1; i < n; i++ {
		var left = make([]int64, m)
		left[0] = dp[i-1][0]
		for j := 1; j < m; j++ {
			left[j] = max(left[j-1]-1, dp[i-1][j])
		}

		var right = make([]int64, m)
		right[m-1] = dp[i-1][m-1]
		for j := m - 2; j >= 0; j-- {
			right[j] = max(right[j+1]-1, dp[i-1][j])
		}

		for j := 0; j < m; j++ {
			dp[i][j] = max(right[j], left[j]) + int64(points[i][j])
		}
	}

	var res int64
	for i := 0; i < m; i++ {
		res = max(res, dp[n-1][i])
	}
	return res
}
func main() {
	points := [][]int{{1, 5}, {2, 3}, {4, 2}}
	fmt.Println(maxPoints(points))
}
