package main

import (
	"fmt"
	"math"
)

var dp []int

func minPathCost(grid [][]int, moveCost [][]int) int {
	dp = nil
	for i := 0; i < len(moveCost); i++ {
		dp = append(dp, -1)
	}
	var res int = math.MaxInt
	for i := 0; i < len(grid[0]); i++ {
		res = min(res, dfs(grid, moveCost, 0, i))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func dfs(grid [][]int, moveCost [][]int, i, j int) int {
	if i == len(grid)-1 {
		return grid[i][j]
	}
	if dp[grid[i][j]] != -1 {
		return dp[grid[i][j]]
	}
	now := grid[i][j]
	res := math.MaxInt
	for k := 0; k < len(grid[i]); k++ {
		res = min(res, dfs(grid, moveCost, i+1, k)+now+moveCost[now][k])
	}
	dp[now] = res
	return dp[now]
}

func main() {
	fmt.Println(minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}))
}
