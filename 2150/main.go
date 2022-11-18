package main

import "fmt"

func findLonely(nums []int) []int {
	var count = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if count[nums[i]] == 1 && count[nums[i]-1] != 0 && count[nums[i]+1] != 0 {
			res = append(res, nums[i])
		}
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
