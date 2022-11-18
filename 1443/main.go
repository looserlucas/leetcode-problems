package main

import "fmt"

var visited map[int]bool
var ha []bool
var adj [][]int

func minTime(n int, edges [][]int, hasApple []bool) int {
	adj = make([][]int, len(hasApple))
	ha = hasApple
	visited = make(map[int]bool)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	return dfs(0, 0)
}

func dfs(index int, cost int) int {
	if visited[index] {
		return 0
	}
	visited[index] = true
	childrenCost := 0
	for i := 0; i < len(adj[index]); i++ {
		childrenCost += dfs(adj[index][i], 2)
	}
	if childrenCost == 0 && !ha[index] {
		return 0
	}
	return cost + childrenCost
}

func main() {
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false}))
}
