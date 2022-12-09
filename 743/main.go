package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
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

func networkDelayTime(times [][]int, n int, k int) int {
	// start sending from u
	// the signal will reach v for the first time at = shortest path (u->v)
	// result = max(shortest path(u -> v), v is other node in the network)
	var adj [][][]int = make([][][]int, n+1)
	for _, e := range times {
		adj[e[0]] = append(adj[e[0]], []int{e[1], e[2]})
	}
	var dist []int = make([]int, n+1)
	var h = CreateHeap()
	for i := 1; i <= n; i++ {
		if i != k {
			dist[i] = math.MaxInt
		} else {
			dist[i] = 0
			heap.Push(h, []int{i, 0})
		}
	}
	for h.Len() > 0 {
		top := heap.Pop(h).([]int)
		u := top[0]
		d := top[1]
		if dist[u] != d {
			continue
		}
		for _, e := range adj[u] {
			v := e[0]
			w := e[1]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(h, []int{v, dist[v]})
			}
		}
	}
	var res int = 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt {
			return -1
		} else if dist[i] > res {
			res = dist[i]
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
