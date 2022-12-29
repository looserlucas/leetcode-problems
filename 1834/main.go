package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	for i := 0; i < n; i++ {
		tasks[i] = append(tasks[i], i)
	}
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})

	h := CreateHeap()
	startTime := tasks[0][0]
	last := -1
	// Get first batch
	for i := 0; i < n; i++ {
		if tasks[i][0] == startTime {
			heap.Push(h, []int{tasks[i][1], tasks[i][2]})
			last = i
		}
	}
	var res []int
	for h.Len() > 0 {
		x := heap.Pop(h).([]int)
		res = append(res, x[1])
		next := startTime + x[0]
		y := sort.Search(len(tasks), func(i int) bool {
			return tasks[i][0] > next
		})
		for i := last + 1; i < y; i++ {
			heap.Push(h, []int{tasks[i][1], tasks[i][2]})
			last = i
		}
		startTime = next
		if y != len(tasks) && h.Len() == 0 {
			newStart := tasks[y][0]
			for i := y; i < n; i++ {
				if tasks[i][0] == newStart {
					heap.Push(h, []int{tasks[i][1], tasks[i][2]})
					last = i
				}
			}
			startTime = newStart
		}
	}
	return res
}

func main() {
	fmt.Println(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}))
	fmt.Println(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}))
	fmt.Println(getOrder([][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1}}))
}
