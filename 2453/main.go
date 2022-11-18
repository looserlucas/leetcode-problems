package main

import (
	"fmt"
	"sort"
)

func destroyTargets(nums []int, space int) int {
	var count = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]%space]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var now = make(map[int]int)
	var max = -1
	var res int
	for i := 0; i < len(nums); i++ {
		now[nums[i]%space]++
		if now[nums[i]%space] == count[nums[i]%space] {
			if count[nums[i]%space] >= max {
				max = count[nums[i]%space]
				res = nums[i]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2))
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2))
	fmt.Println(destroyTargets([]int{6, 2, 5}, 100))
}
