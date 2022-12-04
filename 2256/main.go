package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func minimumAverageDifference(nums []int) int {
	var pre = make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}
	res := 0
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		k1 := 0
		if len(nums)-1-i != 0 {
			k1 = (pre[len(nums)-1] - pre[i]) / (len(nums) - i - 1)
		}
		if k := abs(pre[i]/(i+1) - k1); k < min {
			min = k
			res = i
		}
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
