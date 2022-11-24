package main

import "fmt"

var dp [][27]int
var cal []int
var isCyclic bool
var ans int

func largestPathValue(colors string, edges [][]int) int {
	dp = make([][27]int, len(colors))
	adj := make([][]int, len(colors))
	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
	}
	isCyclic = false
	ans = 0
	cal = make([]int, len(colors))
	for i := 0; i < len(cal); i++ {
		cal[i] = -1
	}
	for i := 0; i < len(colors); i++ {
		if cal[i] == -1 {
			dfs(adj, colors, i)
			if isCyclic {
				return -1
			}
			for j := 0; j < len(dp[i]); j++ {
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
}

func dfs(adj [][]int, colors string, now int) {
	// if cycle -> return -1
	if cal[now] == 0 {
		isCyclic = true
		return
	}
	// if dp[i] != nil -> return dp[i]
	if cal[now] == 1 {
		return
	}

	cal[now] = 0
	// for children -> calculate dp[i] by chilldren outcome
	for _, v := range adj[now] {
		dfs(adj, colors, v)
		for i := 0; i < len(dp[v]); i++ {
			if dp[v][i] > dp[now][i] {
				dp[now][i] = dp[v][i]
			}
		}
	}

	// return dp[i]
	dp[now][int(colors[now])-int('a')]++
	cal[now] = 1
}

func main() {
	fmt.Println(largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
	fmt.Println(largestPathValue("a", [][]int{{0, 0}}))
}
