package main

import (
	"container/heap"
	"fmt"
	"math"
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

func maxResult(nums []int, k int) int {
	h := CreateHeap()
	var dp []int
	for i := 0; i < len(nums); i++ {
		dp = append(dp, math.MinInt)
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			heap.Push(h, []int{dp[i], i})
			continue
		}
		if i < k+1 {
			dp[i] = nums[i] + (*h)[0][0]
			heap.Push(h, []int{dp[i], i})
			continue
		} else {
			pop := i - k - 1
			for h.Len() > 0 && (*h)[0][1] <= pop {
				heap.Pop(h)
			}
			if h.Len() > 0 {
				dp[i] = nums[i] + (*h)[0][0]
			} else {
				dp[i] = nums[i]
			}
			heap.Push(h, []int{dp[i], i})
		}
	}
	return dp[len(nums)-1]
}
func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
	fmt.Println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3))
	fmt.Println(maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2))
}
