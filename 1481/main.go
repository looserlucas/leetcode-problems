package main

import (
	"container/heap"
	"fmt"
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

func findLeastNumOfUniqueInts(arr []int, k int) int {
	var count = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	h := CreateHeap()
	for _, v := range count {
		heap.Push(h, v)
	}
	for h.Len() > 0 {
		if (*h)[0] <= k {
			k -= heap.Pop(h).(int)
		} else {
			break
		}
	}
	return h.Len()
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}
