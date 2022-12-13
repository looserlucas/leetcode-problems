package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CreateMaxHeap() *MaxHeap {
	h := &MaxHeap{}
	heap.Init(h)
	return h
}

// An IntHeap is a min-heap of ints.
type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CreateMinHeap() *MinHeap {
	h := &MinHeap{}
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
func longestSubarray(nums []int, limit int) int {
	minH := CreateMinHeap()
	maxH := CreateMaxHeap()
	j := 0
	var res int
	for i := 0; i < len(nums); i++ {
		heap.Push(minH, []int{nums[i], i})
		heap.Push(maxH, []int{nums[i], i})
		for j < i && (*maxH)[0][0]-(*minH)[0][0] > limit {
			j++
			for (*maxH)[0][1] < j {
				heap.Pop(maxH)
			}
			for (*minH)[0][1] < j {
				heap.Pop(minH)
			}
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}
