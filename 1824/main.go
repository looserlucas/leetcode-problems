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

var dp [][]int

func minSideJumps(obstacles []int) int {
	dp = nil
	dp = make([][]int, 3)
	for j := 0; j < len(obstacles); j++ {
		obstacles[j]--
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < len(obstacles); j++ {
			dp[i] = append(dp[i], -1)
		}
	}
	dfs(1, 0, obstacles)
	return dp[1][0]
}

func dfs(lane int, now int, o []int) int {
	if now == len(o)-1 {
		return 0
	}
	if dp[lane][now] != -1 {
		return dp[lane][now]
	}
	var res int = 0
	// if obstacle on current lane, switch
	if o[now+1] == lane {
		x := math.MaxInt
		if o[now] != (lane+1)%3 {
			x = dfs((lane+1)%3, now, o)
		}
		y := math.MaxInt
		if o[now] != (lane+2)%3 {
			y = dfs((lane+2)%3, now, o)
		}
		res = min(x, y) + 1
	} else {
		res = dfs(lane, now+1, o)
	}
	dp[lane][now] = res
	return res
}

func main() {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0}))
	fmt.Println(minSideJumps([]int{0, 1, 1, 3, 3, 0}))
	fmt.Println(minSideJumps([]int{0, 2, 1, 0, 3, 0}))
	fmt.Println(minSideJumps([]int{0, 3, 3, 0, 3, 2, 2, 0, 0, 3, 0}))
}
