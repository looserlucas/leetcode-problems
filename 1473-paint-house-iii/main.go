package main

import "fmt"

const MAXINT int = 9999999

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	var dp [101][101][21]int
	res := dfs(dp, houses, cost, 0, target, 0)

	if res == MAXINT {
		return -1
	} else {
		return res
	}
}

func dfs(dp [101][101][21]int, house []int, cost [][]int, index int, blocks int, lastColor int) int {
	if index == len(house) || blocks < 0 || len(house)-index < blocks {
		if blocks == 0 && index == len(house) {
			return 0
		} else {
			return MAXINT
		}
	}
	if dp[index][blocks][lastColor] != 0 {
		return dp[index][blocks][lastColor]
	}

	dp[index][blocks][lastColor] = MAXINT // set init value == maxint
	// if this house is paint-able
	if house[index] == 0 {
		// try to paint the house with every type of color
		for i := 1; i <= len(cost[index]); i++ {
			if i != lastColor {
				dp[index][blocks][lastColor] = min(dp[index][blocks][lastColor], dfs(dp, house, cost, index+1, blocks-1, i)+cost[index][i-1])
			} else {
				dp[index][blocks][lastColor] = min(dp[index][blocks][lastColor], dfs(dp, house, cost, index+1, blocks, i)+cost[index][i-1])
			}
		}
	} else {
		// just move on
		if lastColor != house[index] {
			dp[index][blocks][lastColor] = dfs(dp, house, cost, index+1, blocks-1, house[index])
		} else {
			dp[index][blocks][lastColor] = dfs(dp, house, cost, index+1, blocks, house[index])
		}
	}

	return dp[index][blocks][lastColor]
}
func main() {
	houses := []int{0, 2, 1, 2, 0}
	cost := [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}
	m := 5
	n := 2
	target := 3
	fmt.Println(minCost(houses, cost, m, n, target))
}
