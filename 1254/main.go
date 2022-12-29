package main

import "fmt"

var move = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var n, m int
var visited [][]bool

func closedIsland(a [][]int) int {
	n = len(a)
	m = len(a[0])
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		var tmp = make([]bool, m)
		visited[i] = tmp
	}

	for j := 0; j < m; j++ {
		if a[0][j] == 0 && !visited[0][j] {
			dfs(a, 0, j)
		}
		if a[n-1][j] == 0 && !visited[n-1][j] {
			dfs(a, n-1, j)
		}
	}
	for i := 0; i < n; i++ {
		if a[i][0] == 0 && !visited[i][0] {
			dfs(a, i, 0)
		}
		if a[i][m-1] == 0 && !visited[i][m-1] {
			dfs(a, i, m-1)
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 0 && !visited[i][j] {
				dfs(a, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(a [][]int, i, j int) {
	visited[i][j] = true
	for _, mo := range move {
		if i+mo[0] >= 0 && i+mo[0] < n && j+mo[1] >= 0 && j+mo[1] < m && !visited[i+mo[0]][j+mo[1]] && a[i+mo[0]][j+mo[1]] == 0 {
			dfs(a, i+mo[0], j+mo[1])
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
