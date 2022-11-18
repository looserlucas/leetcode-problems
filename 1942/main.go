package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
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

func smallestChair(times [][]int, targetFriend int) int {
	arrivedAt := times[targetFriend][0]
	sort.SliceStable(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	//fmt.Println(arrivedAt)
	//fmt.Println(times)
	h := CreateHeap()
	chair := CreateHeap()
	now := 0
	for i := 0; i < len(times); i++ {
		var seat int = -1
		for h.Len() > 0 && times[i][0] >= (*h)[0][0] {
			lastSeat := heap.Pop(h).([]int)
			heap.Push(chair, []int{lastSeat[1], 0})
		}
		if chair.Len() > 0 {
			seat = heap.Pop(chair).([]int)[0]
		} else {
			seat = now
			now++
		}
		heap.Push(h, []int{times[i][1], seat})

		if arrivedAt == times[i][0] {
			return seat
		}
	}
	return 0
}

func main() {
	fmt.Println(smallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	fmt.Println(smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
	fmt.Println(smallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6))
}
