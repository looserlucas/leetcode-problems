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

func furthestBuilding(a []int, bricks int, ladders int) int {
	var i int = 0
	h := CreateHeap()
	for i < len(a)-1 {
		if a[i] < a[i+1] {
			needed := a[i+1] - a[i]
			if bricks >= needed {
				bricks -= needed
				heap.Push(h, needed)
			} else {
				if ladders > 0 {
					if h.Len() > 0 && bricks+(*h)[0]-needed > bricks {
						x := heap.Pop(h).(int)
						bricks += x - needed
						heap.Push(h, needed)
					}
					ladders--
				} else {
					break
				}
			}
		}
		i++
	}
	return i
}

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}
