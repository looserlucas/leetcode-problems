package main

import (
	"fmt"
	"math"
)

var visited []bool

func minScore(n int, roads [][]int) int {
	var adj [][][]int = make([][][]int, n+1)
	for i := 0; i < len(roads); i++ {
		adj[roads[i][0]] = append(adj[roads[i][0]], []int{roads[i][1], roads[i][2]})
		adj[roads[i][1]] = append(adj[roads[i][1]], []int{roads[i][0], roads[i][2]})
	}

	visited = make([]bool, n+1)
	return dfs(adj, 1)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(adj [][][]int, i int) int {
	res := math.MaxInt
	visited[i] = true
	for _, edge := range adj[i] {
		if !visited[edge[0]] {
			res = min(res, dfs(adj, edge[0]))
		}
		res = min(res, edge[1])
	}
	return res
}

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
}
