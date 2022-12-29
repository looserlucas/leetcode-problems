package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func minStoneSum(piles []int, k int) int {
	h := CreateHeap()
	for _, p := range piles {
		heap.Push(h, p)
	}
	for k > 0 {
		x := heap.Pop(h).(int)
		var newX int
		if x%2 == 0 {
			newX = x / 2
		} else {
			newX = x/2 + 1
		}
		heap.Push(h, newX)
		k--
	}
	var res int
	for h.Len() > 0 {
		res += heap.Pop(h).(int)
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
