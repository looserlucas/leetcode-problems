package main

import (
	"fmt"
	"math"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var dp []int

func mincostTickets(days []int, costs []int) int {
	dp = make([]int, len(days)+1)
	for i := 0; i < len(days); i++ {
		dp[i] = math.MaxInt
	}
	dp[len(days)] = 0
	sort.Slice(days, func(i, j int) bool {
		return days[i] < days[j]
	})
	dfs(days, costs, 0)
	return dp[0]
}

func dfs(days, costs []int, i int) {
	if i == len(days) {
		return
	}
	if dp[i] != math.MaxInt {
		return
	}
	x := sort.Search(len(days), func(j int) bool {
		return days[j] > days[i]
	})
	y := sort.Search(len(days), func(j int) bool {
		return days[j] > days[i]+6
	})
	z := sort.Search(len(days), func(j int) bool {
		return days[j] > days[i]+29
	})
	dfs(days, costs, x)
	dfs(days, costs, y)
	dfs(days, costs, z)
	dp[i] = min(dp[x]+costs[0], min(dp[y]+costs[1], dp[z]+costs[2]))
}

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}
