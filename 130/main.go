package main

import "fmt"

var move = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var n, m int
var visited [][]bool

func solve(board [][]byte) {
	n = len(board)
	m = len(board[0])
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		var tmp = make([]bool, m)
		visited[i] = tmp
	}
	for j := 0; j < m; j++ {
		if board[0][j] == 'O' && !visited[0][j] {
			dfs(board, 0, j)
		}
		if board[n-1][j] == 'O' && !visited[n-1][j] {
			dfs(board, n-1, j)
		}
	}
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' && !visited[i][0] {
			dfs(board, i, 0)
		}
		if board[i][m-1] == 'O' && !visited[i][m-1] {
			dfs(board, i, m-1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
	fmt.Println(board)
}

func dfs(board [][]byte, i, j int) {
	visited[i][j] = true
	for _, mo := range move {
		if i+mo[0] >= 0 && i+mo[0] < n && j+mo[1] >= 0 && j+mo[1] < m && !visited[i+mo[0]][j+mo[1]] && board[i+mo[0]][j+mo[1]] == 'O' {
			dfs(board, i+mo[0], j+mo[1])
		}
	}
}

func main() {
	solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}
