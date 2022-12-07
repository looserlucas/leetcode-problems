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

func leastInterval(tasks []byte, n int) int {
	coolDown := make(map[int][]int)
	count := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		count[tasks[i]]++
	}

	h := CreateHeap()
	for k, v := range count {
		heap.Push(h, []int{v, int(k) - int('A')})
	}

	res := 0
	for h.Len() > 0 || len(coolDown) != 0 {
		for k, v := range coolDown {
			if v[0] == 0 {
				heap.Push(h, []int{v[1], k})
				delete(coolDown, k)
			}
		}
		// decrease cooldown
		for k, _ := range coolDown {
			coolDown[k][0]--
		}
		if h.Len() > 0 {
			x := heap.Pop(h).([]int)
			if x[0]-1 > 0 {
				// add new cooldown
				coolDown[x[1]] = []int{n, x[0] - 1}
			}
		}
		res++
	}
	return res
}
func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0))
}
