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

func reductionOperations(nums []int) int {
	var bucket = make(map[int]int)
	h := CreateHeap()
	for i := 0; i < len(nums); i++ {
		_, ok := bucket[nums[i]]
		if !ok {
			bucket[nums[i]] = 1
			heap.Push(h, nums[i])
		} else {
			bucket[nums[i]]++
		}
	}

	res := 0
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		res += bucket[x]
		bucket[(*h)[0]] += bucket[x]
	}
	return res
}

func main() {
	fmt.Println(reductionOperations([]int{5, 1, 3}))
	fmt.Println(reductionOperations([]int{1, 1, 1}))
	fmt.Println(reductionOperations([]int{1, 1, 2, 2, 3}))
}
