package main

import "fmt"

func findDuplicates(nums []int) []int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	var res []int
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
