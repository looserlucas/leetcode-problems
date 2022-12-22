package main

import "fmt"

var total int
var count, res []int

func sumOfDistancesInTree(n int, edges [][]int) []int {
	total = n
	var adj [][]int = make([][]int, n)
	count = make([]int, n)
	res = make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	dfs1(adj, 0, -1)
	dfs2(adj, 0, -1)
	return res
}

func dfs1(adj [][]int, now, pre int) {
	for _, v := range adj[now] {
		if v == pre {
			continue
		}
		dfs1(adj, v, now)
		count[now] += count[v]
		res[now] += res[v] + count[v]
	}
	count[now]++
}

func dfs2(adj [][]int, now, pre int) {
	for _, v := range adj[now] {
		if v == pre {
			continue
		}
		res[v] = res[now] - count[v] + total - count[v]
		dfs2(adj, v, now)
	}
}

func main() {
	fmt.Println("vim-go")
}
