package main

import (
	"container/list"
	"fmt"
)

func shortestPathBinaryMatrix(a [][]int) int {
	if a[0][0] == 1 {
		return -1
	}
	n := len(a)
	m := len(a[0])
	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var dist = make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			tmp[j] = -1
		}
		dist[i] = tmp
	}
	q := list.New()
	q.PushBack([]int{0, 0, 1})
	dist[0][0] = 1
	for q.Len() > 0 {
		now := q.Front()
		v := now.Value.([]int)
		for _, k := range moves {
			x := v[0] + k[0]
			y := v[1] + k[1]
			if x >= 0 && x < n && y >= 0 && y < m && a[x][y] == 0 && dist[x][y] == -1 {
				q.PushBack([]int{x, y, v[2] + 1})
				dist[x][y] = v[2] + 1
			}
		}
		q.Remove(now)
	}
	return dist[n-1][m-1]
}

func main() {
	fmt.Println("vim-go")
}
