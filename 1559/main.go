package main

import (
	"fmt"
)

var visited [][]int
var cycle bool
var n, m int
var moves = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func containsCycle(a [][]byte) bool {
	n = len(a)
	m = len(a[0])

	cycle = false
	visited = make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp = make([]int, m)
		visited[i] = tmp
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 0 {
				dfs(a, i, j, -1, -1)
				//printMatrix(visited)
				//fmt.Println(i, j, cycle)
				if cycle {
					return true
				}
			}
		}
	}
	return false
}

func dfs(a [][]byte, i, j int, ip, jp int) {
	if visited[i][j] == 1 {
		cycle = true
		return
	}
	visited[i][j] = 1
	for _, k := range moves {
		x := i + k[0]
		y := j + k[1]
		/*
			if i == 0 && j == 1 {
				fmt.Println(i, j, x, y, string(a[x][y]), string(a[i][j]), n, m)
			}
		*/

		if x >= 0 && x < n && y >= 0 && y < m && (x != ip || y != jp) && visited[x][y] != 2 && a[x][y] == a[i][j] {
			dfs(a, x, y, i, j)
		}
	}
	visited[i][j] = 2
}

func printMatrix(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func main() {
	fmt.Println(containsCycle([][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}))
}
