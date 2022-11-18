package main

import (
	"container/heap"
	"fmt"
)

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
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

func maximumImportance(n int, roads [][]int) int64 {
	h := CreateHeap()
	freq := make(map[int]int)
	for i := 0; i < len(roads); i++ {
		freq[roads[i][0]]++
		freq[roads[i][1]]++
	}

	for k, v := range freq {
		heap.Push(h, []int{v, k})
	}

	for h.Len() > 0 {
		x := heap.Pop(h).([]int)
		freq[x[1]] = n
		n--
	}
	var res int64 = 0
	for i := 0; i < len(roads); i++ {
		res += (int64(freq[roads[i][0]]) + int64(freq[roads[i][1]]))
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
