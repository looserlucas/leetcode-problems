package main

import "fmt"

var visited []bool

func minReorder(n int, connections [][]int) int {
	var adj [][]int = make([][]int, n)
	visited = make([]bool, n)
	for i := 0; i < len(connections); i++ {
		adj[connections[i][0]] = append(adj[connections[i][0]], connections[i][1])
		adj[connections[i][1]] = append(adj[connections[i][1]], -connections[i][0])
	}
	return dfs(adj, 0)
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func dfs(adj [][]int, now int) int {
	visited[now] = true
	var pos int
	for _, v := range adj[now] {
		if !visited[abs(v)] {
			if v > 0 {
				pos++
			}
			pos += dfs(adj, abs(v))
		}
	}
	return pos
}

func main() {
	fmt.Println()
}
