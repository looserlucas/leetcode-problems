package main

import "fmt"

var n, m, res int
var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var visited [][]bool

func numEnclaves(grid [][]int) int {
	n = len(grid)
	m = len(grid[0])
	res = 0
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		tmp := make([]bool, m)
		visited[i] = tmp
	}
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && !visited[i][0] {
			dfs(grid, i, 0, false)
		}
		if grid[i][m-1] == 1 && !visited[i][m-1] {
			dfs(grid, i, m-1, false)
		}
	}
	for j := 0; j < m; j++ {
		if grid[0][j] == 1 && !visited[0][j] {
			dfs(grid, 0, j, false)
		}
		if grid[n-1][j] == 1 && !visited[n-1][j] {
			dfs(grid, n-1, j, false)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				dfs(grid, i, j, true)
			}
		}
	}
	return res
}

func dfs(a [][]int, i, j int, count bool) {
	visited[i][j] = true
	if count {
		res++
	}
	for _, k := range moves {
		x := i + k[0]
		y := j + k[1]
		if x >= 0 && x < n && y >= 0 && y < m && a[x][y] == 1 && !visited[x][y] {
			dfs(a, x, y, count)
		}
	}
}
func main() {
	fmt.Println("vim-go")
}
