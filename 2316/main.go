package main

import "fmt"

var adj [][]int
var visited map[int]bool

func countPairs(n int, edges [][]int) int64 {
	adj = make([][]int, n)
	visited = make(map[int]bool, n)
	sum := make([]int, n)

	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		var connected []int
		dfs(i, &connected)
		for j := 0; j < len(connected); j++ {
			sum[connected[j]] = len(connected)
		}
		fmt.Println(connected)
	}
	var res int64
	for i := 0; i < n; i++ {
		res += int64((n - sum[i]))
	}
	return res / 2
}

func dfs(index int, res *[]int) {
	if visited[index] {
		return
	}
	visited[index] = true
	*res = append(*res, index)
	for i := 0; i < len(adj[index]); i++ {
		dfs(adj[index][i], res)
	}
}

func main() {
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
}
