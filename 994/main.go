package main

import (
	"container/list"
	"fmt"
)

func orangesRotting(a [][]int) int {
	n := len(a)
	m := len(a[0])
	q := list.New()
	count1 := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 2 {
				q.PushBack([]int{i, j, 0})
			} else if a[i][j] == 1 {
				count1++
			}
		}
	}
	var move = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	res := 0
	for q.Len() > 0 {
		now := q.Front()
		v := now.Value.([]int)
		res = v[2]
		for _, k := range move {
			x := v[0] + k[0]
			y := v[1] + k[1]
			if x >= 0 && x < n && y >= 0 && y < m && a[x][y] == 1 {
				q.PushBack([]int{x, y, v[2] + 1})
				a[x][y] = 2
				count1--
			}
		}
		q.Remove(now)
	}
	if count1 == 0 {
		return res
	} else {
		return -1
	}
}

func main() {
	fmt.Println("vim-go")
}
