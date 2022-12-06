package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
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

// First, try to count the barcode[i] existences, then push all of it to the MAX heap (compare by count)
// Then, at every iteration, pop out the head of the heap, append that barcode to the res, if the count - 1 > 0,
// we save it to push it back to the heap in the next iteration
func rearrangeBarcodes(barcodes []int) []int {
	var count = make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		count[barcodes[i]]++
	}
	h := CreateHeap()
	for k, v := range count {
		heap.Push(h, []int{v, k})
	}
	var res []int
	var save []int
	for h.Len() > 0 {
		n1 := heap.Pop(h).([]int)
		res = append(res, n1[1])
		if save != nil {
			heap.Push(h, save)
		}
		if n1[0]-1 > 0 {
			save = []int{n1[0] - 1, n1[1]}
		} else {
			save = nil
		}
	}
	return res
}

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3}))
}
