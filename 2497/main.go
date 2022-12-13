package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
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

func maxStarSum(vals []int, edges [][]int, k int) int {
	type verticeH struct {
		curSum int
		h      *IntHeap
	}
	var adj []verticeH = make([]verticeH, len(vals))
	for i := 0; i < len(adj); i++ {
		tmp := verticeH{
			curSum: 0,
			h:      CreateHeap(),
		}
		adj[i] = tmp
	}
	for _, e := range edges {
		v := e[0]
		u := e[1]

		hu := adj[u].h
		if hu.Len() == k {
			if hu.Len() > 0 && vals[v] > (*hu)[0] {
				x := heap.Pop(adj[u].h).(int)
				adj[u].curSum -= x
				adj[u].curSum += vals[v]
				heap.Push(adj[u].h, vals[v])
			}
		} else {
			if adj[u].curSum+vals[v] > adj[u].curSum {
				adj[u].curSum += vals[v]
				heap.Push(adj[u].h, vals[v])
			}
		}

		hv := adj[v].h
		if hv.Len() == k {
			if hv.Len() > 0 && vals[u] > (*hv)[0] {
				x := heap.Pop(adj[v].h).(int)
				adj[v].curSum -= x
				adj[v].curSum += vals[u]
				heap.Push(adj[v].h, vals[u])
			}
		} else {
			if adj[v].curSum+vals[u] > adj[v].curSum {
				adj[v].curSum += vals[u]
				heap.Push(adj[v].h, vals[u])
			}
		}

		//fmt.Println(u, adj[u].h, v, adj[v].h)
	}
	var res int = math.MinInt
	for i := 0; i < len(adj); i++ {
		if vals[i]+adj[i].curSum > res {
			res = max(vals[i], vals[i]+adj[i].curSum)
		}
	}

	return res
}

func main() {
	fmt.Println(maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2))
	fmt.Println(maxStarSum([]int{-5}, [][]int{}, 0))
	fmt.Println(maxStarSum([]int{1, -8, 0}, [][]int{{1, 0}, {2, 1}}, 2))
}
