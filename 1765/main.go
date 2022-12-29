package main

import (
	"container/list"
	"fmt"
)

func highestPeak(isWater [][]int) [][]int {
	n := len(isWater)
	m := len(isWater[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < m; j++ {
			tmp = append(tmp, -1)
		}
		res[i] = tmp
	}

	q := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = 0
				q.PushBack([]int{i, j})
			}
		}
	}
	move := [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	for q.Len() > 0 {
		now := q.Front()
		v := now.Value.([]int)
		for _, k := range move {
			x := v[0] + k[0]
			y := v[1] + k[1]
			if x >= 0 && x < n && y >= 0 && y < m && res[x][y] == -1 {
				q.PushBack([]int{x, y})
				res[x][y] = res[v[0]][v[1]] + 1
			}
		}
		q.Remove(now)
	}
	return res
}

func main() {
	fmt.Println(highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}))
}
