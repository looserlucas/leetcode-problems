package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res int
	var cur = math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < cur {
				cur = abs(sum - target)
				res = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
