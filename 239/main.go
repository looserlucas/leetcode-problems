package main

import (
	"container/list"
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	q := list.New()
	for i := 0; i < len(nums); i++ {
		if q.Len() > 0 && q.Front().Value.(int) < i-k+1 {
			q.Remove(q.Front())
		}
		for q.Len() > 0 && nums[q.Back().Value.(int)] < nums[i] {
			q.Remove(q.Back())
		}
		q.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[q.Front().Value.(int)])

		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
