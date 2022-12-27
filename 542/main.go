package main

import (
	"container/list"
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	n := len(mat)
	m := len(mat[0])
	q := list.New()
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				tmp[j] = math.MaxInt
			} else {
				tmp[j] = 0
				q.PushBack([]int{i, j, 0})
			}
		}
		dist[i] = tmp
	}
	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for q.Len() > 0 {
		now := q.Front()
		value := now.Value.([]int)
		for _, k := range moves {
			if value[0]+k[0] < n && value[0]+k[0] >= 0 && value[1]+k[1] < m && value[1]+k[1] >= 0 && dist[value[0]+k[0]][value[1]+k[1]] == math.MaxInt {
				dist[value[0]+k[0]][value[1]+k[1]] = value[2] + 1
				q.PushBack([]int{value[0] + k[0], value[1] + k[1], value[2] + 1})
			}
		}
		q.Remove(now)
	}
	return dist
}

func main() {
	fmt.Println(updateMatrix([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
}
