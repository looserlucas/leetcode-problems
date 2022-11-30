package main

import "fmt"

var flag bool

func canFinish(n int, edges [][]int) bool {
	flag = false
	adj := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	var visit = make([]int, n)
	var q []int = nil
	for i := 0; i < n; i++ {
		if visit[i] == 0 {
			dfsTopo(adj, i, visit, &q)
			if flag {
				return false
			}
		}
	}
	return len(q) == n
}

func dfsTopo(adj [][]int, i int, visit []int, q *[]int) {
	visit[i] = 1
	for j := 0; j < len(adj[i]); j++ {
		if visit[adj[i][j]] == 1 {
			flag = true
			return
		}
		if visit[adj[i][j]] == 0 {
			dfsTopo(adj, adj[i][j], visit, q)
		}
	}
	(*q) = append(*q, i)
	visit[i] = 2
}
func main() {
	fmt.Println("vim-go")
}
