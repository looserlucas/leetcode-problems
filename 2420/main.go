package main

import "fmt"

func goodIndices(nums []int, k int) []int {
	var leftDown = make([]int, len(nums))
	var rightDown = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftDown[i] = 1
			continue
		}
		if nums[i] <= nums[i-1] {
			leftDown[i] = leftDown[i-1] + 1
		} else {
			leftDown[i] = 1
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightDown[i] = 1
			continue
		}
		if nums[i] <= nums[i+1] {
			rightDown[i] = rightDown[i+1] + 1
		} else {
			rightDown[i] = 1
		}
	}
	var res []int
	for i := k; i < len(nums)-k; i++ {
		if rightDown[i+1] >= k && leftDown[i-1] >= k {
			res = append(res, i)
		}
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
