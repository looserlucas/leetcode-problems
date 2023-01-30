package main

import "fmt"

func countGood(nums []int, k int) int64 {
	var t int64 = 0
	var total int64 = 0
	seen := make(map[int]int)
	n := int64(len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
		t += int64(seen[nums[i]] - 1)
		for j < i && t > int64(k-1) {
			t -= int64(seen[nums[j]] - 1)
			seen[nums[j]]--
			j++
		}
		total += int64(i - j + 1)
	}
	return n*(n+1)/2 - total
}

func main() {
	fmt.Println(countGood([]int{1, 1, 1, 1, 1}, 10))
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
}
