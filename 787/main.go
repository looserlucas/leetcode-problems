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

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var adj = make([][][]int, n)
	var stop = make([]int, n)
	for i := 0; i < len(flights); i++ {
		adj[flights[i][0]] = append(adj[flights[i][0]], []int{flights[i][1], flights[i][2]})
	}
	for i := 0; i < n; i++ {
		stop[i] = math.MaxInt
	}

	h := CreateHeap()
	heap.Push(h, []int{0, src, 0})
	stop[src] = 0
	for h.Len() > 0 {
		x := heap.Pop(h).([]int)
		d := x[0]
		u := x[1]
		st := x[2]
		if st > stop[u] || st > k+1 {
			continue
		}
		stop[u] = st
		if u == dst {
			return d
		}
		for _, v := range adj[u] {
			heap.Push(h, []int{d + v[1], v[0], st + 1})
		}
	}

	return -1
}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
}
