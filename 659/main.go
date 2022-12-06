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

func isPossible(nums []int) bool {
	last := make(map[int]*IntHeap)
	for i := 0; i < len(nums); i++ {
		if _, ok := last[nums[i]-1]; !ok {
			if _, ok1 := last[nums[i]]; !ok1 {
				h := CreateHeap()
				heap.Push(h, 1)
				last[nums[i]] = h
			} else {
				heap.Push(last[nums[i]], 1)
			}
		} else {
			x := heap.Pop(last[nums[i]-1]).(int)
			if _, ok1 := last[nums[i]]; !ok1 {
				h := CreateHeap()
				heap.Push(h, x+1)
				last[nums[i]] = h
			} else {
				heap.Push(last[nums[i]], x+1)
			}
			if last[nums[i]-1].Len() == 0 {
				delete(last, nums[i]-1)
			}
		}
	}
	for k, v := range last {
		for v.Len() > 0 {
			x := heap.Pop(v).(int)
			if x < 3 {
				return false
			}
		}
		delete(last, k)
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}
