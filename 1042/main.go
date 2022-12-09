package main

import (
	"fmt"
)

var res []int

func gardenNoAdj(n int, paths [][]int) []int {
	var res = make([]int, n+1)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	var adj [][]int = make([][]int, n+1)
	for _, e := range paths {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	for i := 1; i <= n; i++ {
		color := make([]bool, 5)
		for _, v := range adj[i] {
			color[res[v]] = true
		}

		for c := 4; c >= 1; c-- {
			if !color[c] {
				res[i] = c
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
