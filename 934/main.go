package main

import (
	"container/list"
	"fmt"
)

var dist [][]int
var n int
var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestBridge(a [][]int) int {
	n = len(a)
	dist = make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, -1)
		}
		dist[i] = tmp
	}
	for i := 0; i < n; i++ {
		ok := false
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				dfs(a, i, j)
				ok = true
				break
			}
		}
		if ok {
			break
		}
	}
	/*
		printMatrix(a)
		printMatrix(dist)
	*/
	q := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dist[i][j] == 0 {
				q.PushBack([]int{i, j})
			}
		}
	}
	for {
		now := q.Front()
		v := now.Value.([]int)
		for _, k := range moves {
			x := v[0] + k[0]
			y := v[1] + k[1]
			if x >= 0 && x < n && y >= 0 && y < n && dist[x][y] == -1 {
				if a[x][y] == 1 {
					return dist[v[0]][v[1]]
				} else {
					q.PushBack([]int{x, y})
					dist[x][y] = dist[v[0]][v[1]] + 1
				}
			}
		}
		q.Remove(now)
	}
}

func dfs(a [][]int, i, j int) {
	dist[i][j] = 0
	for _, k := range moves {
		x := i + k[0]
		y := j + k[1]
		if x >= 0 && x < n && y >= 0 && y < n && a[x][y] == 1 && dist[x][y] == -1 {
			dfs(a, x, y)
		}
	}
}

/*
func printMatrix(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
*/

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(shortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}))
}
