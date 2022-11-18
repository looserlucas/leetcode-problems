package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
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

// An IntHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
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

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  CreateMaxHeap(),
		right: CreateMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 {
		heap.Push(this.left, num)
		return
	}
	if this.right.Len() == 0 {
		if num >= (*this.left)[0] {
			heap.Push(this.right, num)
			return
		} else {
			tmp := heap.Pop(this.left)
			heap.Push(this.right, tmp)
			heap.Push(this.left, num)
			return
		}
	}

	if num > (*this.right)[0] {
		heap.Push(this.right, num)
	} else {
		heap.Push(this.left, num)
	}

	if this.right.Len() > this.left.Len() {
		tmp := heap.Pop(this.right)
		heap.Push(this.left, tmp)
		return
	}
	if this.left.Len() > this.right.Len()+1 {
		tmp := heap.Pop(this.left)
		heap.Push(this.right, tmp)
		return
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == 0 {
		return -1
	}
	if this.right.Len() == 0 {
		return float64((*this.left)[0])
	}
	if this.right.Len() == this.left.Len() {
		return (float64((*this.left)[0]) + float64((*this.right)[0])) / 2
	} else {
		return float64((*this.left)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func main() {
	fmt.Println("vim-go")
}
