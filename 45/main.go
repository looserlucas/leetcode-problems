package main

import (
	"container/list"
	"fmt"
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}

	next := list.New()
	next.PushBack([]int{1, nums[0]})
	for i := 1; i < len(nums); i++ {
		right := next.Front().Value.([]int)
		back := next.Back().Value.([]int)
		//fmt.Println(right, back, i, nums[i])
		if i+nums[i] >= len(nums)-1 {
			return right[0] + 1
		}
		if right[1] == i {
			if i+nums[i] > back[1] {
				next.PushBack([]int{right[0] + 1, i + nums[i]})
			}
			next.Remove(next.Front())
		} else {
			if i+nums[i] > back[1] {
				next.PushBack([]int{right[0] + 1, i + nums[i]})
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(jump([]int{1, 1, 1, 1, 1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{6, 2, 6, 1, 7, 9, 3, 5, 3, 7, 2, 8, 9, 4, 7, 7, 2, 2, 8, 4, 6, 6, 1, 3}))
}
