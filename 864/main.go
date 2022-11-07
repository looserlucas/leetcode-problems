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

// Broken down into 2 thoughts:
// - If we know the first element of the group, for example: [1,2,3], and we know that this will start with 1,
// we need to have a way to check whether there are 2 and 3 in hands or not with O(1) time complexity -> use hashmap (count)
// - To find the first element of next the group, we need to get min value of in array every time we done creating group.
// -> use min heap
func isNStraightHand(hand []int, groupSize int) bool {
	h := CreateHeap()
	count := make(map[int]int)
	for i := 0; i < len(hand); i++ {
		if _, ok := count[hand[i]]; !ok {
			count[hand[i]] = 1
		} else {
			count[hand[i]]++
		}
	}
	for k, _ := range count {
		heap.Push(h, k)
	}
	for h.Len() > 0 {
		now := (*h)[0]
		if count[now] == 0 {
			heap.Pop(h)
			continue
		}
		for i := 0; i < groupSize; i++ {
			v, ok := count[now+i]
			if !ok || v == 0 {
				return false
			}
			count[now+i]--
		}
	}
	return true
}

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}
