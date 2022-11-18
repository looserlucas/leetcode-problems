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

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := CreateHeap()
	rightHeap := CreateHeap()
	var res int64 = 0
	i := 0
	j := len(costs) - 1
	for k > 0 {
		for leftHeap.Len() < candidates && i <= j {
			heap.Push(leftHeap, costs[i])
			i++
		}
		for rightHeap.Len() < candidates && j >= i {
			heap.Push(rightHeap, costs[j])
			j--
		}
		var a, b int = math.MaxInt, math.MaxInt
		if leftHeap.Len() > 0 {
			a = (*leftHeap)[0]
		}
		if rightHeap.Len() > 0 {
			b = (*rightHeap)[0]
		}
		if a <= b {
			res += int64(a)
			heap.Pop(leftHeap)
		} else {
			res += int64(b)
			heap.Pop(rightHeap)
		}
		k--
	}
	return res
}

func main() {
	/*
		fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
		fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
	*/
	fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
	//fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
}
