package main

import (
	"container/list"
	"fmt"
)

func findFarmland(land [][]int) [][]int {
	n := len(land)
	m := len(land[0])
	/*
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			tmp := make([]bool, m)
			visited[i] = tmp
		}
	*/

	var move [][]int = [][]int{{0, 1}, {1, 0}}
	var res [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 1 {
				q := list.New()
				q.PushBack([]int{i, j})
				var last = make([]int, 2)
				for q.Len() > 0 {
					now := q.Front()
					v := now.Value.([]int)
					last[0] = v[0]
					last[1] = v[1]
					for _, k := range move {
						if v[0]+k[0] < n && v[1]+k[1] < m && land[v[0]+k[0]][v[1]+k[1]] == 1 {
							q.PushBack([]int{v[0] + k[0], v[1] + k[1]})
							land[v[0]+k[0]][v[1]+k[1]] = 0
						}
					}
					q.Remove(now)
				}
				res = append(res, []int{i, j, last[0], last[1]})
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
