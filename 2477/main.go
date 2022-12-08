package main

import "fmt"

var visited []bool
var finalRes int64

func minimumFuelCost(roads [][]int, seats int) int64 {
	visited = make([]bool, len(roads)+1)
	finalRes = 0
	var adj = make([][]int, len(roads)+1)
	for i := 0; i < len(roads); i++ {
		adj[roads[i][0]] = append(adj[roads[i][0]], roads[i][1])
		adj[roads[i][1]] = append(adj[roads[i][1]], roads[i][0])
	}
	dfs(adj, 0, seats)
	return finalRes
}

func dfs(adj [][]int, now int, seats int) int64 {
	visited[now] = true
	var people int64 = 1
	for _, v := range adj[now] {
		if !visited[v] {
			people += dfs(adj, v, seats)
		}
	}
	if now != 0 {
		finalRes += (people + int64(seats) - 1) / int64(seats)
	}
	return people
}

func main() {
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 5))
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
}
