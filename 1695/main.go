package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maximumUniqueSubarray(nums []int) int {
	var res, sum int = 0, 0
	var j = 0
	var count = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		sum += nums[i]
		for j <= i && count[nums[i]] > 1 {
			count[nums[j]]--
			sum -= nums[j]
			j++
		}
		res = max(res, sum)
	}
	return res
}
func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}
