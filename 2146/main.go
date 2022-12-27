package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	for k := 0; k < len(h[i]); k++ {
		if h[i][k] < h[j][k] {
			return true
		} else if h[i][k] > h[j][k] {
			return false
		}
	}
	return false
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CreateHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	if grid[start[0]][start[1]] == 0 {
		return nil
	}
	n := len(grid)
	m := len(grid[0])
	q := list.New()
	var rank = make([]*IntHeap, n+m+1)
	for i := 0; i <= n+m; i++ {
		tmp := CreateHeap()
		rank[i] = tmp
	}

	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	q.PushBack([]int{start[0], start[1], 0, grid[start[0]][start[1]]})
	for q.Len() > 0 {
		now := q.Front()
		value := now.Value.([]int)
		if value[3] >= pricing[0] && value[3] <= pricing[1] {
			heap.Push(rank[value[2]], []int{value[3], value[0], value[1]})
		}
		for _, k := range moves {
			if value[0]+k[0] < n && value[0]+k[0] >= 0 && value[1]+k[1] < m && value[1]+k[1] >= 0 && grid[value[0]+k[0]][value[1]+k[1]] != 0 {
				q.PushBack([]int{value[0] + k[0], value[1] + k[1], value[2] + 1, grid[value[0]+k[0]][value[1]+k[1]]})
				grid[value[0]+k[0]][value[1]+k[1]] = 0
			}
		}
		q.Remove(now)
	}
	var res [][]int
	for i := 0; i < len(rank); i++ {
		for rank[i].Len() > 0 && len(res) < k {
			x := heap.Pop(rank[i]).([]int)
			res = append(res, []int{x[1], x[2]})
		}
		if len(res) == k {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(highestRankedKItems([][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}, []int{2, 5}, []int{0, 0}, 3))
}
