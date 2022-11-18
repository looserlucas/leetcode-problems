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

func avoidFlood(r []int) []int {
	var last = make(map[int]int)
	var urgent = make([]int, len(r))
	var res []int = make([]int, len(r))
	var defaultValue int
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] != 0 {
			defaultValue = r[i]
		}
		v, ok := last[r[i]]
		if !ok {
			urgent[i] = math.MaxInt
			last[r[i]] = i
		} else {
			urgent[i] = v
			last[r[i]] = i
		}
	}
	h := CreateHeap()
	for i := 0; i < len(r); i++ {
		if r[i] != 0 {
			if h.Len() > 0 && i == (*h)[0] {
				return nil
			}
			if urgent[i] != math.MaxInt {
				heap.Push(h, urgent[i])
			}
			res[i] = -1
		} else {
			if h.Len() > 0 {
				j := heap.Pop(h).(int)
				res[i] = r[j]
			} else {
				res[i] = defaultValue
			}
		}
	}
	if h.Len() > 0 {
		return nil
	}
	return res
}

func main() {
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))
}
