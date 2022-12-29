package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func minimumEffortPath(a [][]int) int {
	n := len(a)
	m := len(a[0])
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dist = make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < m; j++ {
			tmp = append(tmp, math.MaxInt)
		}
		dist[i] = tmp
	}
	dist[0][0] = 0
	h := CreateHeap()
	heap.Push(h, []int{0, 0, 0})
	for h.Len() > 0 {
		now := heap.Pop(h).([]int)
		cost := now[0]
		i := now[1]
		j := now[2]
		if cost != dist[i][j] {
			continue
		}
		for _, k := range moves {
			x := i + k[0]
			y := j + k[1]
			if x >= 0 && x < n && y >= 0 && y < m {
				if r := max(cost, abs(a[i][j]-a[x][y])); r < dist[x][y] {
					dist[x][y] = r
					heap.Push(h, []int{r, x, y})
				}
			}
		}
	}
	return dist[n-1][m-1]
}

func main() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
}
