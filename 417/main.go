package main

import (
	"container/list"
	"fmt"
)

func pacificAtlantic(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	var pState = make([][]bool, n)
	var aState = make([][]bool, n)
	for i := 0; i < n; i++ {
		tmp := make([]bool, m)
		tmp1 := make([]bool, m)
		pState[i] = tmp
		aState[i] = tmp1
	}

	q := list.New()
	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < n; i++ {
		q.PushBack([]int{i, 0})
		pState[i][0] = true
	}
	for j := 0; j < m; j++ {
		q.PushBack([]int{0, j})
		pState[0][j] = true
	}
	for q.Len() > 0 {
		now := q.Front()
		value := now.Value.([]int)
		for _, k := range moves {
			if value[0]+k[0] < n && value[0]+k[0] >= 0 && value[1]+k[1] < m && value[1]+k[1] >= 0 &&
				!pState[value[0]+k[0]][value[1]+k[1]] && a[value[0]+k[0]][value[1]+k[1]] >= a[value[0]][value[1]] {
				q.PushBack([]int{value[0] + k[0], value[1] + k[1]})
				pState[value[0]+k[0]][value[1]+k[1]] = true
			}
		}
		q.Remove(now)
	}

	for i := 0; i < n; i++ {
		q.PushBack([]int{i, m - 1})
		aState[i][m-1] = true
	}
	for j := 0; j < m; j++ {
		q.PushBack([]int{n - 1, j})
		aState[n-1][j] = true
	}
	for q.Len() > 0 {
		now := q.Front()
		value := now.Value.([]int)
		for _, k := range moves {
			if value[0]+k[0] < n && value[0]+k[0] >= 0 && value[1]+k[1] < m && value[1]+k[1] >= 0 &&
				!aState[value[0]+k[0]][value[1]+k[1]] && a[value[0]+k[0]][value[1]+k[1]] >= a[value[0]][value[1]] {
				q.PushBack([]int{value[0] + k[0], value[1] + k[1]})
				aState[value[0]+k[0]][value[1]+k[1]] = true
			}
		}
		q.Remove(now)
	}

	fmt.Println(aState)
	fmt.Println(pState)
	var res [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if aState[i][j] && pState[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(pacificAtlantic([][]int{{1}}))
}
