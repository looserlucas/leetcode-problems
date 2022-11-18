package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func minGroups(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return false
		}
	})
	q := CreateHeap()
	var res = 1
	heap.Push(q, intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		if (*q)[0] < intervals[i][0] {
			heap.Pop(q)
			heap.Push(q, intervals[i][1])
		} else {
			res++
			heap.Push(q, intervals[i][1])
		}
	}
	return res
}

func main() {
	/*
		fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))
		fmt.Println(minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))
	*/
	fmt.Println(minGroups([][]int{{441459, 446342}, {801308, 840640}, {871890, 963447}, {228525, 336985}, {807945, 946787}, {479815, 507766}, {693292, 944029}, {751962, 821744}}))
}
